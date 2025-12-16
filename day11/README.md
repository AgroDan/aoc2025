# Day 11: Reactor

What a difference between days. Day 10 had me stumped for days, but Day 11 I finished right after lunch. Not to complain that this one was easy per se, it certainly wasn't. But it just _reeked_ of a way of introducing me to a new algorithmic paradigm.

## Part 1

I boldly approached part 1 with a sense of hesitation. The question was straightforward, determine a path from `you` to `out`. Or rather, determine exactly how many paths existed that went from `you` to `out`. I knew ahead of time that this was going to be one of those situations where a simple Depth-First Search algorithm wouldn't cut it. Didn't matter, I pressed on anyway. I created that DFS algorithm for the first part of traversal knowing full well what to expect for part 2. It searched through the maze blazingly fast and spat out a number. Part 1 was easy. I knew not to be so confident. Part 2's rug-pull was bound to be some ridiculous `O(2^n)` spread that would completely invalidate DFS. I braced for impact and continued onwards.

## Part 2

Sure enough, DFS was not the proper means of accomplishing this task. Now instead of going from `you` to `out`, I had to go from `svr`, through `dac` and `fft` (order didn't matter), and land through `out`. Count how many paths allow that. My initial instinct (again, knowing it was wrong) was to perform that DFS algorithm again and write a list of pathways for each successful traversal, then eliminate any pathways that didn't have `fft` and `dac` through it. Since the standard DFS algorithm I wrote didn't even complete with the `svr -> out` pathway amounts, I knew that there must have been some sort of different means. Didn't take too much googling to learn about [Digital Dynamic Programming](https://en.wikipedia.org/wiki/Dynamic_programming).

Truth be told, there isn't a lot of "magic" to this one. I'm not even sure why this is considered something new because all it's doing is applying some neat constraint tricks to the traversal. In this particular example, I'm basically traversing down the path, noting every single node as one particular potential path to the end, in a sense kinda working backwards from the endpoint by way of recursion. But the real muscle to this algorithm? A combination of memoization and that handy "DAC/FFT Mask" trick where I just use bitwise math to determine whether or not I passed through `DAC` or `FFT`. Mostly memoization though. Don't bother going down this path fully because we already know how many paths exist through here, but if you happened to pass through `DAC` and `FFT` when you're checking this path then you know it goes to the end, so count this path as valid then.

This flew through part 1 and part 2 in `1.0216ms`.

> **NOTE**: if you are running this code yourself, it will fail if you try this against either `testinput` or `testinput2`. This is because both test data inputs had different values specific to the puzzle it was placed under. `svr` and `you` are not in both of those input files!

> **ADDITIONAL NOTE**: I lost some time in this because I kept wanting to look for `srv` and not `svr`. I was debugging why this was bombing out so much when I realized I read the puzzle input too quickly.