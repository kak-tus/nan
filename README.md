# nan - **N**o more **A**llocations **N**ever

Zero allocation Nullable structures in one library with handy conversion functions, marshallers and unmarshallers.

Features:
- short name "nan"

Supported types:
- time.Time
- string
- int

Supported marshallers:
- JSON
- Scylla and Cassandra (gocql, gocqlx)

# Benchmarks

In Go code you usually use structures with some Pointer fields like this

```
type pointersSmall struct {
	Field000 *string
	Field001 *string
	Field002 *string
	Field003 *string
	Field004 *string
	Field005 *string
}
```

Let's try compare this struct benchmark with benchmark for struct with simple Values

```
type valuesSmall struct {
	Field000 string
	Field001 string
	Field002 string
	Field003 string
	Field004 string
	Field005 string
}
```

They processed with zero allocation, as fields with Values. And we see here better processing time for struct with Pointers.

```
BenchmarkPointersSmall-8       	1000000000	         0.604 ns/op	       0 B/op	       0 allocs/op
BenchmarkValuesSmall-8         	183266614	         6.53 ns/op	       0 B/op	       0 allocs/op
```

Let's do some function calls with structs. Structs with Pointers still processed with zero allocations and with better time, then structs with Values.

```
BenchmarkPointerSmallChain-8   	1000000000	         0.305 ns/op	       0 B/op	       0 allocs/op
BenchmarkValuesSmallChain-8    	56732866	        20.5 ns/op	       0 B/op	       0 allocs/op
```

We can use JSONs for input and output in our service. Try to do marshalling and unmarshalling with jsoniter. Here situation changed. Struct with Values got lower allocations, memory usage and better processing time.

```
BenchmarkPointersSmallJSON-8       	   95233	     12222 ns/op	   14816 B/op	      25 allocs/op
BenchmarkValuesSmallJSON-8         	   42004	     28349 ns/op	   28397 B/op	      26 allocs/op
```

Let's try to improve json speed with easyjson. Better all results for struct with Values, but (strange) worse processing time results (in comparison with jsoniter) for struct with Pointers.

```
BenchmarkPointersSmallEasyJSON-8   	   64302	     17107 ns/op	   14589 B/op	      21 allocs/op
BenchmarkValuesSmallEasyJSON-8     	   67329	     18270 ns/op	   14450 B/op	      14 allocs/op
```

**Intermediate conclusion**: if your code processing pipeline is to produce some value, chains it thru some function call and got some result - sometimes better to use struct with Pointers. But if your processing is to do some conversions (marshalling, unmarshalling) with structs - prefer to use struct with Values.

Go next. Sometimes struct groth bigger.

```
type pointersBig struct {
	Field000 *string
	...
	Field999 *string
}
```

```
type valuesBig struct {
	Field000 string
	...
	Field999 string
}
```

Try to do this struct. Here we see, that struct with Values left zero allocations and got bigger processing time (this is normal, because struct is bigger). And struct with Pointers lost advantage - non zero allocations and much worst processing time and memory usage.

```
BenchmarkPointersBig-8         	   37318	     31707 ns/op	   24192 B/op	    1001 allocs/op
BenchmarkValuesBig-8           	  751268	      1618 ns/op	       0 B/op	       0 allocs/op
```

Try function calls chain. Not changed anything in struct with Pointers and slightly added processing time for struct with Values (still lower then struct with Pointers).

```
BenchmarkPointerBigChain-8     	   39541	     31800 ns/op	   24192 B/op	    1001 allocs/op
BenchmarkValuesBigChain-8      	  321734	      3310 ns/op	       0 B/op	       0 allocs/op
```

Try do marshal and unmarshal. So sad result for struct with Values. Twice processing, time, memory usage and allocations.

```
BenchmarkPointersBigJSON-8   	     429	   2691710 ns/op	 2001118 B/op	    1020 allocs/op
BenchmarkValuesBigJSON-8     	     273	   4316146 ns/op	 4005063 B/op	    2014 allocs/op
```

And finally try easyjson. Sad results for struct with Pointers - all worse, then jsoniter. But better results for struct with Values (worse then jsoniter struct with Pointers).

```
BenchmarkPointersBigEasyJSON-8   	     362	   3151776 ns/op	 2357775 B/op	    3066 allocs/op
BenchmarkValuesBigEasyJSON-8     	     382	   3061267 ns/op	 2307679 B/op	    1063 allocs/op
```

**Final conclusion**: Don't do optimisations at first code stage - prefer to use struct with Values, then struct with Pointers. And only when perfomance tuning needed - review your processing pipeline and try swith to struct with Pointers on hot places. Prefer to use codegens (easyjson and others) then native runtime processing - they get better results in mostly cases.
