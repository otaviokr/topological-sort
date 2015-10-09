# topological-sort
A simple implementation of the topological sort algorithm.

[![Coverage Status](https://coveralls.io/repos/otaviokr/topological-sort/badge.svg?branch=master&service=github)](https://coveralls.io/github/otaviokr/topological-sort?branch=master) [![Build Status](https://travis-ci.org/otaviokr/topological-sort.svg)](https://travis-ci.org/otaviokr/topological-sort)

## Overview

Suppose you have a directed graph. To sort it topologically is to generate a linear ordering of its vertices such that 
for every edge *UV* from vertex *U* to vertex *V*, *U* comes before *V* in the ordering. It is important to notice that 
a topological sort can only be done if - and only if - the **graph has no directed cycles**; such graph is usually 
called a **directed acyclic graph**, or DAG.

A directed graph may have more than one topological ordering.

## Example

![Example of a directed acyclic graph](images/dag.png)

All the following orderings are valid topological sort outcomes of the example graph above. Each one has an arbitrary, 
secondary criteria to consistently visit the nodes. Because this module does not follow any particular secondary 
criteria, executing it on this graph could result in any of that outcome randomly.

- **7, 5, 3, 11, 8, 2, 9, 10** (visual left-to-right, top-to-bottom)
- **3, 5, 7, 8, 11, 2, 9, 10** (smallest-numbered available vertex first)
- **5, 7, 3, 8, 11, 10, 9, 2** (fewest edges first)
- **7, 5, 11, 3, 10, 8, 9, 2** (largest-numbered availabled vertex first)
- **7, 5, 11, 2, 3, 8, 9, 10** (attempting top-to-bottom, left-to-right)
- **3, 7, 8, 5, 11, 10, 2, 9** (arbitrary; no specific secondary criteria)

This example was taken from 
[Wikipedia's article on topological sorting](https://en.wikipedia.org/wiki/Topological_sorting).

## Limitations and Design Decisions

In the current version, all functions to generate the topological sort does **not** use any secondary criteria for 
ordering - for example, if there's two equivalent child nodes, it will choose which one to visit first randomly.

## Usage

The example below is a complete example (just run it!) with both functions being called for the following graph:

![The graph used in the example code below](images/complex_tree_test.png)

```go
package main

import (
    "fmt"
    "gitlab.com/otaviokr/topological-sort"
)

func main() {
    unsorted := map[string][]string{
        "0": []string{"1", "4"}, 
        "1": []string{"3", "5"}, 
        "2": []string{"5"},
		"3": []string{"5", "7"}, 
		"4": []string{}, 
		"5": []string{"6"}, 
		"6": []string{"7"}, 
		"7": []string{}}
		
	sorted := Sort(unsorted)
	fmt.Printf("Result: %v\n", sorted)
	
	reversed := ReversedSort(unsorted)
    fmt.Printf("Reversed result: %v\n", reversed)
}
```
