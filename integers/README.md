# Integers

## Packages

> The name of the package should match the name of the directory in which its source is located. If you package is called logger, then its source files may be located in
```
$GOPATH/src/github.com/yourname/logger
```
Package names _should be lowercase_. 

Only one package per directory, so apparently you need to make sure your stuff is organized appropriately. Only one package per directory.
[5 Suggestions for setting up a go project](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)

> The `go` commands; `go build`, `go install`, `go test`, `go get`, all work with packages, not individual files. 
Neat I assumed it worked directly with each individual file.


## Variables:
It seems `%q` is for strings, and `%d` is for digits, I should google this at some point...

## More learning with Tests...

Well french toast... 

```
func testAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected '%d', got '%d'", expected, sum)
	}
}
```

This gives you an interesting error stating that this particular test isn't used. The reason behind this?  `testAdder` should be `TestAdder`. This is important...

## One last thought...

I figured the `+` piece when I wrote the code instead of following the actual tutorial, I went back and read and was like wtf.. Why are they hard coding `0` or `4` like... that's not how that works. I guess it's just to teach you how you can "cheat" your way to passing a test, something you shouldn't do to begin