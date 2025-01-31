package mapreduce

import "sync"

type MapReduce[Key comparable, Value any, Input any, Output any] interface {
	Map(input Input) []KeyValue[Key, Value]
	Reduce(key Key, values []Value) Output
}

type KeyValue[Key comparable, Value any] struct {
	Key   Key
	Value Value
}

type MapReduceEngine[Key comparable, Value any, Input any, Output any] struct{}

func (engine MapReduceEngine[Key, Value, Input, Output]) Run(job MapReduce[Key, Value, Input, Output], inputs []Input) map[Key]Output {
	var wg sync.WaitGroup
	intermediate := make(chan KeyValue[Key, Value], len(inputs)*10)

	for _, input := range inputs {
		wg.Add(1)
		go func(data Input) {
			defer wg.Done()
			for _, kv := range job.Map(data) {
				intermediate <- kv
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(intermediate)
	}()

	grouped := make(map[Key][]Value)
	var mu sync.Mutex

	for kv := range intermediate {
		mu.Lock()
		grouped[kv.Key] = append(grouped[kv.Key], kv.Value)
		mu.Unlock()
	}

	output := make(map[Key]Output)
	var wgReduce sync.WaitGroup
	for key, values := range grouped {
		wgReduce.Add(1)
		go func(k Key, v []Value) {
			defer wgReduce.Done()
			output[k] = job.Reduce(k, v)
		}(key, values)
	}

	wgReduce.Wait()
	return output
}
