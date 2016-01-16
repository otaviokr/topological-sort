package sort

import (
	"fmt"
	"os"
)

// Sort is DEPRECATED. This function will be removed from the library. It used to be the
// KahnSort function, but the name is too generic, so it will be removed. Feel free
// to use any sort function in this library, but if you wish the same behaviour,
// use KahnSort.
func Sort(tree map[string][]string) []string {
	msg := "github.com/otaviokr/sort.Sort function is deprecated. Please update it to KahnSort"
	fmt.Fprintln(os.Stderr, msg)
	return KahnSort(tree)
}

// ReversedSort is DEPRECATED. It used a deprecated function (Sort), and with new algorithms
// this function makes no sense. I suggest to use the new Reverse() function with the sorting
// function of your choice.
func ReversedSort(tree map[string][]string) []string {
	msg := "github.com/otaviokr/sort.ReversedSort function is deprecated. Please update it to Reverse"
	fmt.Fprintln(os.Stderr, msg)
	sorted := Sort(tree)
	reversed := []string{}

	for i := len(sorted); i > 0; i-- {
		reversed = append(reversed, sorted[i-1])
	}

	return reversed
}

// KahnSort receives a description of a search tree and returns an array with the elements sorted.
// The Kahn's Algorithm creates an "orphan-list" of all nodes that has no parents. Then, it puts
// the first element of that list in the sorted list and removes all edges from that node to
// other nodes; if any of those nodes has no other parents connected, it is appended to the
// orphan-list. The analysis starts again for the first element in the orphan-list.
// Example for tree: map["A": ["B", "C"], "B": [], "C": ["B"]]]. Meaning A to B, A to C and C to B.
func KahnSort(tree map[string][]string) []string {
	sorted := []string{}
	inDegree := map[string]int{}

	// 01. Calculate this.indegree of all vertices by going through every edge of the graph;
	// Each child gets indegree++ during breadth-first run.
	for element, children := range tree {
		if _, exists := inDegree[element]; !exists {
			inDegree[element] = 0 // So far, element does not have any parent.
		}

		for _, child := range children {
			if _, exists := inDegree[child]; !exists {
				inDegree[child] = 1 // Being a child of an element, it is already a inDegree 1.
			} else {
				inDegree[child]++
			}
		}
	}

	// 02. Collect all vertices with indegree==0 onto a stack;
	stack := []string{}
	for element, value := range inDegree {
		if value == 0 {
			stack = append(stack, element)
			inDegree[element] = -1
		}
	}

	// 03. While zero-degree-stack is not empty:
	for len(stack) > 0 {
		// 03.01. Pop element from zero-degree-stack and append it to topological order;
		var node string
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 03.02. Find all children of element and decrease indegree. If indegree becomes 0, add to zero-degree-stack;
		for _, child := range tree[node] {
			inDegree[child]--
			if inDegree[child] == 0 {
				stack = append(stack, child)
				inDegree[child] = -1
			}
		}

		// 03.03. Append to the sorted list.
		sorted = append(sorted, node)
	}

	if len(tree) != len(sorted) {
		// It seems that there's a directed cycle. Toposort won't work.
		return []string{}
	}

	return sorted
}

// Reverse is just a wrapper to invert the order of the elements in a sorted list.
func Reverse(sorted []string) []string {
	reversed := []string{}

	for i := len(sorted); i > 0; i-- {
		reversed = append(reversed, sorted[i-1])
	}

	return reversed
}
