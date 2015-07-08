# topological-sort
A simple implementation of the topological sort algorithm.

## Overview

Suppose you have a directed graph. To sort it topologically is to generate a linear ordering of its vertices such that 
for every edge *UV* from vertex *U* to vertex *V*, *U* comes before *V* in the ordering. It is important to notice that 
a topological sort can only be done if - and only if - the graph has no directed cycles, usually called a directed 
acyclic graph, or DAG.

A directed graph may have more than one topological ordering.

## Example

![Example of a directed acyclic graph](dag.png)

All the following orderings are valid topological sort of the example graph above.

- **7, 5, 3, 11, 8, 2, 9, 10** (visual left-to-right, top-to-bottom)
- **3, 5, 7, 8, 11, 2, 9, 10** (smallest-numbered available vertex first)
- **5, 7, 3, 8, 11, 10, 9, 2** (fewest edges first)
- **7, 5, 11, 3, 10, 8, 9, 2** (largest-numbered availabled vertex first)
- **7, 5, 11, 2, 3, 8, 9, 10** (attempting top-to-bottom, left-to-right)
- **3, 7, 8, 5, 11, 10, 2, 9** (arbitrary; no specific secondary criteria)

This example was taken from 
[Wikipedia's article on topological sorting](https://en.wikipedia.org/wiki/Topological_sorting).

## Limitations and Design Decisions

In the current version, all functions to generate the topological sort does not use any secondary criteria for 
ordering; all results can be considered arbitrary.

## Usage

```go
package main

import (
    "fmt"
    "gitlab.com/otaviokr/topological-sort"
)

func main() {
    unsorted := map[string][]string{"0": []string{"1", "4"}, "1": []string{"3", "5"}, "2": []string{"5"},
		"3": []string{"5", "7"}, "4": []string{}, "5": []string{"6"}, "6": []string{"7"}, "7": []string{}}
		
	sorted := Sort(unsorted)
	fmt.Printf("Result: %v\n", sorted)
}
```

```go
package main

import (
    "fmt"
    "gitlab.com/otaviokr/topological-sort"
)

func main() {
    unsorted := map[string][]string{"0": []string{"1", "4"}, "1": []string{"3", "5"}, "2": []string{"5"},
		"3": []string{"5", "7"}, "4": []string{}, "5": []string{"6"}, "6": []string{"7"}, "7": []string{}}
		
	sorted := ReversedSort(unsorted)
	fmt.Printf("Result: %v\n", sorted)
}
```
