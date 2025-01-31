package mapreduce

import (
	"reflect"
	"testing"
)

type WordCount struct{}

func (wc WordCount) Map(input string) []KeyValue[string, int] {
	return []KeyValue[string, int]{{Key: input, Value: 1}}
}

func (wc WordCount) Reduce(key string, values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func TestMapReduce(t *testing.T) {
	engine := MapReduceEngine[string, int, string, int]{}
	inputs := []string{"hello", "world", "hello"}
	expected := map[string]int{"hello": 2, "world": 1}

	result := engine.Run(WordCount{}, inputs)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
