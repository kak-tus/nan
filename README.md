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
type pointerSmall struct {
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
type valueSmall struct {
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
BenchmarkPointersSmall-8   	1000000000	         0.296 ns/op	       0 B/op	       0 allocs/op
BenchmarkValueSmall-8      	184559210	         6.50 ns/op	       0 B/op	       0 allocs/op
```

Let's do some function calls with structs. Structs with Pointers still processed with zero allocations and with better time, then structs with Values.

```
BenchmarkPointerSmallChain-8   	1000000000	         0.305 ns/op	       0 B/op	       0 allocs/op
BenchmarkValuesSmallChain-8    	56732866	        20.5 ns/op	       0 B/op	       0 allocs/op
```

We can use JSONs for input and output in our service. Try to do marshalling and unmarshalling with jsoniter. Here situation changed. Struct with Values got lower allocations, memory usage, lower processing time.

```
BenchmarkPointersSmallJSON-8   	   58467	     20060 ns/op	   21635 B/op	      36 allocs/op
BenchmarkValueSmallJSON-8      	   60211	     19375 ns/op	   21539 B/op	      23 allocs/op
```

Let's try to improve json speed with easyjson. Better all results for both structs.

```
BenchmarkPointersSmallEasyJSON-8   	   64222	     17797 ns/op	   14590 B/op	      21 allocs/op
BenchmarkValueSmallEasyJSON-8      	   67212	     17365 ns/op	   14442 B/op	      14 allocs/op
```

**Intermediate conclusion**: if your code processing pipeline is to produce some value, chains it thru some function call and got some result - sometimes better to use struct with Pointers. But if your processing is to do some conversions (marshalling, unmarshalling) with structs - prefer to use struct with Values.

Go next. Sometimes struct groth bigger.

```
type pointerBig struct {
	Field000 *string
	...
	Field999 *string
}
```

```
type valueBig struct {
	Field000 string
	...
	Field999 string
}
```

Try to do this struct. Here we see, that struct with Values left zero allocations and got bigger processing time (this is normal, because struct is bigger). And struct with Pointers lost advantage - non zero allocations, much worst processing time and memory usage.

```
BenchmarkPointersBig-8         	   37318	     31707 ns/op	   24192 B/op	    1001 allocs/op
BenchmarkValuesBig-8           	  751268	      1618 ns/op	       0 B/op	       0 allocs/op
```

Try function calls chain. Not changed anything in struct with Pointers and slightly added processing time for struct with Values (still lower then struct with Pointers).

```
BenchmarkPointerBigChain-8     	   39541	     31800 ns/op	   24192 B/op	    1001 allocs/op
BenchmarkValuesBigChain-8      	  321734	      3310 ns/op	       0 B/op	       0 allocs/op
```

Try do marshal and unmarshal. Struct with Values better in all, except of memory usage per operation.

```
BenchmarkPointerBigJSON-8   	     303	   3530988 ns/op	 3777091 B/op	    4068 allocs/op
BenchmarkValueBigJSON-8     	     314	   3504458 ns/op	 3877041 B/op	    2065 allocs/op
```

Try to fix this with easyjson. Fixed and also got better results.

```
BenchmarkPointerBigEasyJSON-8   	     364	   3239433 ns/op	 2354781 B/op	    3066 allocs/op
BenchmarkValueBigEasyJSON-8     	     378	   3123912 ns/op	 2308829 B/op	    1064 allocs/op
```

**Final conclusion**: Don't do optimisations at first code stage - prefer to use struct with Values, then struct with Pointers. And only when perfomance tuning needed - review your processing pipeline and try swith to struct with Pointers on hot places. Prefer to use codegens (easyjson and others) then native runtime processing - they get better results in mostly cases.
