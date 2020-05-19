# nel

# Benchmarks

## Usual (small) values vs pointers

```
BenchmarkPointersSmall-8       	1000000000	         0.604 ns/op	       0 B/op	       0 allocs/op
BenchmarkValuesSmall-8         	183266614			        6.53 ns/op		     0 B/op	       0 allocs/op
```

## Big values vs pointers

```
BenchmarkPointersBig-8         	   37318	     31707 ns/op	   24192 B/op	    1001 allocs/op
BenchmarkValuesBig-8           	  751268	      1618 ns/op	       0 B/op	       0 allocs/op
```

## Small values vs pointers in function calls chain

```
BenchmarkPointerSmallChain-8   	1000000000	         0.305 ns/op	       0 B/op	       0 allocs/op
BenchmarkValuesSmallChain-8    	56732866			        20.5 ns/op	       0 B/op	       0 allocs/op
```

## Big values vs pointers in function calls chain

```
BenchmarkPointerBigChain-8     	   39541	     31800 ns/op	   24192 B/op	    1001 allocs/op
BenchmarkValuesBigChain-8      	  321734	      3310 ns/op	       0 B/op	       0 allocs/op
```

## Small and big values vs pointers in marshalling

```
BenchmarkPointersSmallJSON-8   	 3215271	       363 ns/op	     154 B/op	       9 allocs/op
BenchmarkValuesSmallJSON-8     	 6429516	       185 ns/op	      98 B/op	       2 allocs/op

BenchmarkPointersBigJSON-8     	   38011	     31940 ns/op	   24212 B/op	    1003 allocs/op
BenchmarkValuesBigJSON-8       	  195124	      7522 ns/op	   16393 B/op	       2 allocs/op
```
