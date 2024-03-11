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

With 50% of `None` and 50% of `Some` the benchmark is:
| Name                | Iterations | Time per operation | Memory per operation | Allocations per operation |
|---------------------|------------|--------------------|----------------------|---------------------------|
| BenchmarkPointer-10 | 7597224    | 161.7 ns/op        | 304 B/op             | 6 allocs/op               |                    |                      |                           |
| **BenchmarkOption-10**  | 5809479    | 212.6 ns/op        | 336 B/op             | 8 allocs/op               |                                                               |                      |                           |
| BenchmarkValue-10   | 7430028    | 162.2 ns/op        | 304 B/op             | 6 allocs/op               |                                                               |                      |                           |

