# Day 6: Trash Compactor

Homework for the cephalopods...okay!

Part one took me longer than I'd like to admit. I was trying to wrap my head around the proper way to loop through this. I parsed everything as fields, ignoring all the whitespace, and just cast each field with non-whitespace to either a number or an operator. I tried to loop as logically as possible through the homework problems, taking the operator at the last row and applying that operation to each of the numbers above it. Typing it out now, it seems incredibly easy. But I got so hung up on the rows and cols that I kept doing it wrong and setting my time back. Eventually I got it.

Part two was a bit difficult to figure out how to process things, because now whitespace actually mattered. Read from right to left, and from top to bottom. Each individual character's row now mattered, so as per the example:

```text
64
23
314
+
```

Would translate to `4 + 431 + 623`. This was the first time I ever had to completely re-write the parser function I used here. And after some careful thinking, I decided that I could kinda retro-fit my trusty `utils.RuneMap` object for this!

This time, it took every single character and made the runemap from it. Now after careful scrutiny, I used the operators at the last row to denote the columns where I should parse the data. So in the example provided...

```text
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
```

I noticed that the operator always started at the left-most column of the homework problem. What I needed was to build a sub-map of the numbers for each problem. So I would turn the above into 4 separate boxes. Once I chopped it up into separate problems, I wrote a `Solve()` function for them, which would start at the right-most column and build a number by reading from the top down, skipping it if it was a whitespace and not counting the last row. Once I built the operands properly, I applied the operator to them and returned the result. Then I just added it all up and that was the result. Took `310ms`, and that was mostly because I printed everything to the screen.

Managed to finish this just before lunch. I'm starving!