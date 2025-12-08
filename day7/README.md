# Day 7: Laboratories

Now we're starting to get into the whole `O(log N)` problems with a vengeance. Part one seemed innocent enough. Just aim the beam downward, follow how many times the beam splits, if it splits and re-joins with another beam then two beams become one. Regardless though, just process the beam and count how many times it splits given the puzzle input. Barring a few off-by-one issues, I managed to get this relatively quickly. I did manage to get this done later in the afternoon on Sunday, though I did try and get a handle on it at midnight so I could tackle it once I was done with the errands I had to run today (as evidenced by the annotations to my code), but that wasn't happening as midnight is very quickly becoming past my bedtime. I started a little bit and then powered down and went to bed. Next day I ran my errands and came back in the late afternoon to get it done. Not too much of a problem.

Part 2 looked exhausting at first. You could just tell that this was going to result in some astronomical number. Knowing this challenge I knew there was going to be two ways to solve this problem:

1. Power through it, let it run, come back in about 10,000,000 years and enter the result.
2. The _right_ way.

And I immediately remembered [this problem from 2024](https://adventofcode.com/2024/day/21). Controlling the robot that controlled the robot that controlled the robot...(x25) -- and the answer to figuring out this one was _memoization!_

I learned allllll about that last year, and now I realized that this is yet another perfect reason to do this again! If this is a cascading effect of splitters, I can instead work from the bottom up. If I know that this particular splitter will have 2 possible outcomes, then the splitter up the chain will take that into account when computing _its_ possible outcomes, all the way up the chain until it reaches the top.

And I _just so happened to have created a utility for exactly this purpose last year!_ `utils.cache`!

I printed out the map for debugging, so running it with printing its output runs in `649ms`, but if I comment out the map printing functions this whole thing completes in only `2.0482ms`! Really proud of this one, and I'm all caught up so far!