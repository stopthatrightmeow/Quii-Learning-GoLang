# Hello World

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


