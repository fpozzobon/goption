# goption
**EXPERIMENTAL** Option pattern for Golang - generic way to handle optional values, as seen in benchmark it has overhead.
Strongly inspired from [Scala option pattern](https://www.scala-lang.org/api/2.13.3/scala/Option.html)

## Usage

```go
    myOptionalStruct := goption.Some(myStruct{
		Name: "test",
	})
```

## Benchmark

| Name                | Iterations | Time per operation | Memory per operation | Allocations per operation |
|---------------------|------------|--------------------|----------------------|---------------------------|
| BenchmarkPointer-10 | 13428400   | 89.83 ns/op        | 160 B/op             | 3 allocs/op               |
| **BenchmarkOption-10**  | 6814846    | 174.3 ns/op        | 272 B/op             | 5 allocs/op               |
| BenchmarkValue-10   | 13407333   | 90.60 ns/op        | 160 B/op             | 3 allocs/op               |


