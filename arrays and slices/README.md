# Arrays and Slices

## Arrays having fixed capacity
This is... hard to get over...

In python an array can be any length of size. This is difficult for me to accept, but I guess that's sort of the power of Go?...


Per Quii's documentation, there are two ways to setup an array.
- [N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}
- [...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5}

Also printing out the array in the test is useful. `%v` which seems to be the default variable for printing an array.


## Ranges and "Blank Identifiers"

oof. This is a learning concept... `_` is a variable... But like a variable that's temporary. I see it's uses but It's hard when it comes to reading your code. I guess that's where all those programming humor meme's come from.

[Effective Go; Blank Identifiers](https://go.dev/doc/effective_go#blank)
> The blank identifier can be assigned or declared with any value of any type, with the value discarded harmlessly. It's a bit like writing to the Unix /dev/null file: it represents a write-only value to be used as a place-holder where a variable is needed but the actual value is irrelevant. It has uses beyond those we've seen already. 

> You may be thinking it's quite cumbersome that arrays have a fixed length, and most of the time you probably won't be using them! Go has slices which do not encode the size of the collection and instead can have any size.

AHH there it is... I knew there wasn't a requirement to hard code arrays!... I guess let's learn about slices...

## Slices
Well, they are easy. Instead of `[#]` just do `[]` with an array and you have a slice!

I honestly thought that was going to be a whole thing, and I'm really glad it didn't turn out to be one!

## reflect.DeepEqual

It's not "type safe" and will "work" regardless. Which is important to note. 

## After Thoughts

Slices seem to be a lot more usable unless I am being very particular in what I want as far as capacity goes. So far this has been quite a bit of fun, though I think Structs Methods and Interfaces are going to be quite a new concept.