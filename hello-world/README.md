# Hello World

[Hello World](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)

## Rune Literals

OK So the first thing I learned, `'` is not a thing when printing generic items, so when I typed out:

```
    fmt.Println('Hello World')
```
I got:
```
# command-line-arguments
hello-world/hello.go:6:14: more than one character in rune literal
```

There is a good article that discusses what rune literals are, and explains how they work in GoLang: [GeeksForGeeks - Rune In GoLang](https://www.geeksforgeeks.org/rune-in-golang/).

So in my hello world script, the rune literal represents a constant where an integer value recognizes a Unicode code point. So basically it's a magical translation for an [ASCII table](https://www.asciitable.com/). 

So if we were to play around with them a bit more: 
```
func main() {
	fmt.Println("Hello World")
	fmt.Println("\u0009SomeInformationHere")
	fmt.Println('\u0009')
}
```

In the case of the `fmt.Println("\u0009SomeInformationHere")` line, the `\u00095` represents a tab, so when that section runs the output will be:
```
    SomeInformationHere
```

Where the rune literal with the single quotes, we end up with the integer `9`. If we check that out on the previously linked ASCII table, we get a TAB character. Pretty neat!

## Writing my first test

Ok so _technically_ it's not my first test, I've written some basic ones before I actually knew what I was doing within Python BUT with Go it is actually my first test. 


## Setting variables

So it turns out you can set variables in two separate ways, the first was is:

```
	got := Hello()
```

Another way is to use:

```
	var got = Hello()
```

I think the specifying of the var is easier to read, and write VS the `:=` which feels a little weird to type out, but it's good to know either way!

> Side note after playing with it a bit more, while the `var` style worked well, I feel like the `:=` isn't _terrible_... but still learning maybe I will change my mind later...

## Running Tests

Requirements for building your tests:

- They need to be in a file with a name like XYZ_test.go
- Test function must start with `Test`
- Test function takes one argument only `t *testing.T`
- In order to use the `*testing.T` type you need to `import "testing"`

To test, you need ot build a module by running `go mod init hello`

Then to run your tests you can simply run `go test`

## Go Documentation
This is a pretty neat feature, you can run `godoc -http :8000` and it will allow you access to your godocs on localhost:8000! Pretty handy, though it's not quite like stack overflow where one can just copy pasta.

Also localhost:8000/pkg will show you all of your installed packages and loclahost:8000/pkg/testing/ is also quite useful. 


## Adding more languages

~~OK so this is where I ran into some problems, the first of which happened to be that when we added the language variable, I was unable to run my tests due to needing a second parameter. Makes sense, honestly, and I solved it by adding blank strings to be passed within my tests.~~

I just need to learn to read better... It literally tells you that you should add blank strings to fill in for the language.

## Final thoughts on Hello World

Honestly I didn't think hello world could get som complex, but learning test driven development was interesting, and quite a bit of fun! Go is starting to make a little bit more sense as well, I'm sure I just need to practice it quite a bit more.

One key thing I learned was writing a function for your testing function through the `assertCorrectMessage` which was pretty neat. Saves you from having to write a bunch of "If got != want` statements.
