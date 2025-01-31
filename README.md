# mapreduce

![Go CI](https://github.com/Palladium02/mapreduce/actions/workflows/ci.yml/badge.svg)


As the name implies this repo contains a small but neat implementation of map reduce [\[1\]](https://static.googleusercontent.com/media/research.google.com/en//archive/mapreduce-osdi04.pdf) as a Go library.

## Installation

```
go get https://github.com/Palladium02/mapreduce
```

## Usage

Simply import the library like so:

```go
import "github.com/Palladium02/mapreduce"
```

After that you have to implement `MapReduce` interface which you can then pass to the `MapReduceEngine` together with the workload.
An example of that can be found in `examples/wordcount.go`.
