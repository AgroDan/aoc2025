# Day 5: Cafeteria

We're 5 days in (out of 12 this year) and we are now dealing with one of those `O(log N)` issues. I was naive in thinking that I could just brute-force this method. Of course this is a fantastic opportunity to dive into previously-solved computational equations! I do have a soft spot for these because I always learn something new, and frankly the elegance involved in this algorithm is all wrapped up in simplicity. I'll explain in a bit, but first, my methods.

I started this by simply writing a `struct` object for the Ingredient Ranges. I created an integer slice of all the ingredients that are being checked and simply went through each ingredient range and determined if that particular ingredient was inside of one of the ranges. It was deceptively simple, and ran relatively quickly.

Then came part two. Get the amount of unique IDs that are listed in the "Fresh" ranges, knowing full well that there is a decent amount of overlap involved. I first just decided to create a gigantic finite set of individual IDs, which works in principle!...on smaller ranges. I didn't really get a good look at the challenge data until afterwards when I just ran it...and it just started hanging on the first range! Whoops.

So yeah, need to find a better way of handling this. I started trying to work out complex math functions in my head, but kept butting up against weird outliers. Eventually I decided that this seems like one of those complex problems that have already been solved before by people much, much smarter than me. Sure enough, I was right. Enter the **Sweeping Line Algorithm**!

The real muscle behind this algorithm has to do with properly sorting things here. The issue is _what_ should you sort? And the answer is a specific type of object denoted by the ranges. For that, I created a localized struct that stored the **position** and the **type**:

```go
type event struct {
    position int
    typ      int
}
```
> Note that I had to use `typ` instead of `type` because `type` is a reserved word in Go, as I'm sure most would understand

The logic here is that I'm going to loop over every single range, and create a new `event{}` object, taking the lower-bound number of the range and assigning it the `1` type, denoting that it's the start of a range. Then taking the upper-bound number, adding `1` to it, and assigning it the `-1` type, denoting the end of the range. Adding one to the end of the range just makes the math work, essentially.

Then you sort! Instead of doing some `MergeSort()` function that I'm so fond of, I just relied on the Go standard library and made my own kind of sorting method, given that I'm sorting a "complex" data type. That is, the slice of `event{}` struct objects. So I leaned on the `cmp` and `slices` libraries to do the comparison.

```go
import (
    "cmp"
    "slices"
)

// ... snip ...

slices.SortFunc(events, func(a, b event) int {
    if a.position == b.position {
        return cmp.Compare(a.typ, b.typ)
    }
    return cmp.Compare(a.position, b.position)
})
```

Took a little bit to really grasp how the `SortFunc()` function works, but basically it takes a function as a parameter in which you tell it how you want to sort something, so in this case I will sort by the `position` in the event slice, and if both are equal then sort by the `typ`. This is the key here!

Now what I do is basically loop through the newly-sorted event slice and record three different variables which will be updated for each iteration. Namely:

```go
activeRanges := 0
prevPosition := events[0].position
totalFresh := 0
```

Where `activeRanges` is the "status" indicator, letting us know if we should count things based on whether or not we're inside a range. If `activeRanges` is ever `0`, then don't add anything to the `totalFresh` variable, which is the counter. Then of course `prevPosition` is just letting us know what the previous position was, obviously.

Loop through every single [sorted] event, then just run this calculation for each iteration:

```go
if activeRanges > 0 {
    totalFresh += e.position - prevPosition
}
activeRanges += e.typ
prevPosition = e.position
```

So as long as `activeRanges` is greater trhan `0`, then you'll be adding the difference between the current position and the previous position!

Honestly, the simplicity of this was so cool. Definitely one of those things where once you go over it on paper (or in my case, a whiteboard), it all just clicks, and you wonder why you never came close to anything that simple. At least that's how I felt.