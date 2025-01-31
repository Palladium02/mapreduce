package examples

import (
	"fmt"
	"strings"

	"github.com/Palladium02/mapreduce"
)

type WordCount struct{}

func (wc WordCount) Map(input string) []mapreduce.KeyValue[string, int] {
	words := strings.Fields(strings.ToLower(input))
	var result []mapreduce.KeyValue[string, int]
	for _, word := range words {
		result = append(result, mapreduce.KeyValue[string, int]{Key: word, Value: 1})
	}
	return result
}

func (wc WordCount) Reduce(key string, values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func main() {
	inputs := []string{
		"hello world",
		"hello Go",
		"world of Go",
	}

	engine := mapreduce.MapReduceEngine[string, int, string, int]{}
	result := engine.Run(WordCount{}, inputs)

	fmt.Println(result)
}
