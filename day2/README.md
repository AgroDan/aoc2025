# Day 2: Gift Shop

Plagued yet again with reading too fast and not truly understanding the specific instructions, which caused me some downtime and backing up a bit to re-write things. You'll see some rewriting of pattern-finding in the `utils.go` file as I read and re-read the instructions. Turns out that I was overcomplicating it yet again, and looking for _all_ patterns repetitions, when really for part 1 all I needed to do was just split the number and check if it's a repeat. I was creating huge slices of potential patterns to check for in gigantic numbers and getting it all wrong. After I re-read it for the third time did I realize that I was checking for way too many patterns.

I personally felt my method for taking on Part 2 was clever, if I do say so myself. Since the rules stipulate now there should be a repeated pattern _through the entirety of the number_, then an invalid number is considered invalid if a number group repeats twice or more. So by that rule, the following numbers are considered **invalid**, which is to say the numbers that I'm specifically looking for:
    - `2222`
    - `1212`
    - `123123`
    - `9999599995`

While the following numbers are considered **valid**, which is to say numbers that I should ignore:
    - `1231231`
    - `3334333`
    - `9`
    - `10110`

At first this tripped me up because I had to build patterns and make sure the pattern existed everywhere, but then I discovered that I could just _create_ a number based on specific patterns, and if it's _the same as the number I'm checking then it's considered invalid!_

For example, let's say I have the number `123123123`, which is considered invalid and I should count it towards the challenge results. I'd get the length of the number, in this case it's `9`. I'll start at the beginning, take the first number and re-create it 9 times. so I'd have `111111111`. Does that equal the number I'm checking? No, so now build a new pattern using the first two numbers and repeat it until it's the exact same size as the number I'm checking. Problem is the length of the number is `9`, and I have an even number to check with, so I can't even re-create the number with just `12`. So toss that out before I even work with it.

Now moving to add the next number as well, and repeat it as many times as to re-create the size of the number. Since it's just 3 numbers, `123`, I can recreate the size of the number cleanly so let's check on that. I can repeat it 3 times to build a number of the same size: `123123123`. Is that the number I'm checking? YES! So this is considered an _invalid number_ and I will add it to the numbers I will count towards the challenge.

I just did this for every single number in all the ranges and it finished part 1 and part 2 in about `511ms`.