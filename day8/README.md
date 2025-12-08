# Day 8: Playground

Eight days in and I tripped over myself. I will leave my over-complicated ridiculous code filled with attempts and failures and just incredibly complicated methods to accomplish something in as a testament to my ineptitude. If I don't stumble, I never truly learn to walk.

Part one just had me scratching my head for a while. Sadly it was me just trying to understand what the puzzle even wanted. I was hung up over the notion that every single junction box needed to have a connection with some other junction box, and the connection was whatever was the next closest point based on [Euclidean Distance](https://en.wikipedia.org/wiki/Euclidean_distance). So that to me made me think that every single junction needed to be connected, which would mean that I would have no less than `10` circuits with `20` junction boxes (as per the example data), and there wouldn't be any single junction circuits...but the example text showed that there would be single-junction circuits! After reading and re-reading, I believe it wanted me to create a list of pairs of points that could potentially be connected (I referred to them as _edges_), then choose `10` for the example data. From those `10`, define the `3` largest circuits.

This concept had me scratching my head for a while. Suffice it to say, I was at a complete loss for how best to accomplish this. I had to look to people far smarter than me to get an idea as to how to approach this, so I checked some of the other various Youtubers who had completed the challenge already. Unfortunately it doesn't look like my favorite go-to [HyperNeutrino](https://www.youtube.com/@hyper-neutrino) had made a video about this one already, so I found [Programming Live with Larry](https://www.youtube.com/watch?v=-GyDPsfVunA) who lead me down the right path...Union Find!

Or, more specifically, a [Disjoint-set/Union-find Forest](https://en.wikipedia.org/wiki/Disjoint-set_data_structure). The best way I can explain it was the following steps were taken:

1. Determine the distances of every possible pair of points in the dataset  (referred to as edges), store in a slice.
2. Sort the slice by distance.
3. Create the Union-Find structure of both `parent` and `rank`, where `parent` is a list of every single junction box, and `rank` is a list of every single junction box initialized to a score of `0`.
4. Define the `find()` anonymous function, which, given a junction box, finds the "root" junction box (aka the "beginning" junction box in a circuit). If it just points to itself, it's a circuit of `1`:

```go
var find func(*Junction) *Junction
find = func(j *Junction) *Junction {
    if parent[j] != j {
        parent[j] = find(parent[j])
    }
    return parent[j]
}
```

5. Define the `union()` anonymous function, which will connect two junctions by searching for the root of that circuit (if it belongs to one), and appending it to the same circuit. This is how things can be joined! This is basically lifted almost right out of the Wikipedia article I showed above:

```go
union := func(a, b *Junction) {
    rootA := find(a)
    rootB := find(b)
    if rootA != rootB {
        if rank[rootA] < rank[rootB] {
            parent[rootA] = rootB
        } else if rank[rootA] > rank[rootB] {
            parent[rootB] = rootA
        } else {
            parent[rootB] = rootA
            rank[rootA]++
        }
    }
}
```

6. Find the top `N` circuits given the challenge, take the top 3, multiply them together. This is your answer.

This was the meat & potatoes of the challenge right here. After a while I finally figured out the logic and got Part 1 answered.

Part 2 was (mercifully) easier than Part 1. Basically keep doing what I've been doing, but keep creating unions until there is only one giant circuit. Then return the last two Junctions that created that final circuit, multiply the X-axes together, return the answer. Managed to get part two in about 30 minutes after part 1.

Yikes, glad that's over with. This code finishes part 1 and 2 in about `215ms`.