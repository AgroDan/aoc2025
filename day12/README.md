# Day 12: Christmas Tree Farm

I...can't believe my solution worked. I just had a hunch that it would. It builds itself up to be this really difficult polygonical flood-fill problem that may or may not involve flipping and rotating shapes. I figured I'd just get all the busy-work stuff out of the way like parsing the data, creating struct objects and potential functions that I might have to call, when I figured...alright, first and foremost, I can't overlap shapes. So by that logic, if I get the _volume_ of the shape, multiply it by how many shapes are required of that particular region, and compare it against the _area_ of the region I'd have to fit these shapes, then in theory if any of the intended volumes exceed the area then it's impossible to fill, so they can be immediately eliminated from the running. So I added a quick little `Area()` and `Volume()` function that calculated those numbers, then multiplied the volume with all the amount of shapes I've been tasked adding to the region and comparing them. I ran it against the test data first and it was wrong. It said that all three regions would fit their required presents. I knew that was wrong.

But I figured...eh, let's see if that number is the same in the puzzle input. And it _wasn't_. I figured alright, what the hell, let's just see if this number works.

And it _did_.

Mercifully so, as is true Advent of Code tradition, the final puzzle was _much_ simpler than the preceeding puzzles, and it only consisted of one part.

Brothers and Sisters, I did it. _Before_ Christmas, too! Granted this year's AOC is half the size, but I no longer have this challenge breathing down my neck.

The Advent of Code is done! Another year in the books!

So long [Eric Wastl](https://was.tl/), and thanks for all the puzzles!