# Day 9: Movie Theater

This one made me doubt my capabilities. Part 1 was pretty easy. Find the biggest rectangle. I just created as many pairs of coordinates as I could, drew rectangles with both of the points, calculated the areas of each, sorted in ascending order and chose the one with the largest area by just picking the one at the top of the slice. That took no time at all.

Part 2 is where things got excessive. I'm day 9 of 12 in and I should have known better than to think I could brute-force my way through this one. I figured I would just draw a border around every single point, but I'd create a finite set of points specifically to determine the border. Initially I thought, "OK, once I create the border, I'll just flood-fill and I'll have a finite set of every known point that's considered _inside_ the polygon! I didn't even think to consider that the area of this polygon would be astro-freakin-nomical. So it looks like I'll have to find some efficiencies here.

I remembered back to a previous advent of code challenge where I had to determine if a point existed inside another polygon, and I was brought back to the concept of ray-casting and employing the [Even-Odd Rule](https://en.wikipedia.org/wiki/Even%E2%80%93odd_rule) to determine "inside-ness." Generally speaking, draw a line from a point out to infinity in a specific direction and count how many times it passes a border. If it crosses a border a zero or even amount of times, then the point is _outside the polygon_. If it crosses it an odd number of times, then the point exists _inside the polygon_. So I figure, "OK, I'll generate a rectangle and check every point in that rectangle to determine if it's inside the polygon."

Again, this was naive. First of all, the rectangles will also be tremendous in size as well. So each one will take forever. Second, why bother with the inside? Just use the perimeter.

This method seemed pretty good...but the perimeters were gigantic too. Not as gigantic as the area, but gigantic enough to still be a pain.

I admit I had to look elsewhere for guidance here. I checked out 0xdf's video on it and he mentioned something I really should have thought about doing from the get-go: caching!

I figured I'd use my caching algorithm I created for this, but I figured I could just use a finite set and make that check one of the first things a function does, check to see if the point we're checking has already been checked and confirmed inside already before doing any ray-casting. By marking down every point I've already confirmed is inside already, I managed to skip quite a few checks and it scrolled through the rectangle checks a whole lot faster this time. This time I could see how many rectangles my code was invalidating was going faster and faster, increasing in speed as it went along and built up its cache.

Still though, I stumbled quite a bit to get to this point. I wrote lots of code that seemed to make sense at the time. In the end though, I barely used much of it. At least it managed to find the answer though...`52m2s` later...

Glad to get this one out of the way though. Three more challenges...oof.