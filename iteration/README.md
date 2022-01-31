# Iteration

## No `while`, `do`, `until`
Only `for` loops. 

## Repeating character 5 times...
Just trying to fully understand this code below...

- We set `i` equal to `0` with `i := 0` and we separate the statements with a `;`
- Then we say if `i < 5`; 
- Add `1` to `i`?...


```
	// for number; if i < 5; add number?...
	for i := 0; i < 5; i++ {
		repeated += character
	}
```


If you write out a function you can see what it prints. I wonder why it ends with a `%`?...
```
01234%                                                                                                                                                               
```

## Benchmarks
This is neat, you don't specify a particular number of times your program will be run but instead Go will select a number and run it an appropriate number of times which is pretty neat. 
```
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```


## Some of the challenges
At the end there were some interesting challenges. I went ahead and wrote a portion that also adds in the ability to print all uppercase A's, VS just lowers. Also wrote a benchmark for it. Pretty interesting. I don't really find the `ExampleRepeat` very useful as a way to document one's code, but maybe I'm just not seeing the entire point of it? Just seems like it's easier to write in a description in comments VS writing out some code that just adds more un-needed lines to your code. I guess at some point I will need to read more into it. 

One super useful things was the documentation though! The [strings](https://golang.org/pkg/strings) documentation was super easy to browse through and pickup on!