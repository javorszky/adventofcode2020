# advent of code 2020

This is the first year I managed to complete everything. I wrote tests for all, but the most trivial tasks. You can see for yourself if you clone the repository, and start the tests by running the following from the folder.
```bash
$ go mod tidy
$ go test ./...
```

I only had issues with 3 days, going to detail them here.

## day 13

This had me stuck for a few days. I came close when I found common frequencies, but could not put it together. Had to reach out to a friend of mine who guided me on how to solve it.

Essentially: given two buses with a start, frequency, and remainder for each, I needed to combine them into one overlapping bus whose start, frequency, and remainder would always produce a valid result for both of the contained buses.

For example a bus that starts at minute 12, with frequency 17, with remainder 0, and another bus at start 9, frequency 23, remainder 8 would have an overlapping bus of start 369, frequency 391, offset 0. From here it's a classic reduce implementation between the new 369/391/0 bus and the next one.

## day 20

This was fairly easy, though doing it in a performant manner was the hardest. Task 1 was easy, but few numbers of steps made any and all implementation complete in a reasonable time.

Task 2 made me think hard on performance. The initial map based implementation took 190ms per step.

The second, ring linked list based implementation, while faster, was still taking 120ms or so.

I had to go to the subreddit to get a hint. Someone suggested that the cup labels and the ones to their right map nicely to a slice of int keys and their values, at which point the lightbulb came on: I only need to change a few entries in the 1 million long ring, and I don't even need to traverse it first, like I had to with a linked list.

## day 23

I got stuck. The issue was bad variable naming. I was comparing sides of tiles that were above and below each other, and top/bottom for tiles where they were next to each other horizontally. This made task 1 pass by accident (the layout of the tiles still worked), but once I had to create the image, and manipulate that, everything fell apart.

Once I wrote enough tests to surface the issue, and fixed the problem, everything fell into place.
