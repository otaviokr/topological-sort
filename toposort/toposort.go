// Package toposort is a Lib to perform topological sorting
//
// Topological sorting is a linear ordering of a directed graph, such that for every directed edge 'uv' from vertex 'u' to vertex 'v'
// 'u' comes before 'v' in the ordering. For more info, see https://en.wikipedia.org/wiki/Topological_sorting
//
// This lib offers 4 different algorithms of topological sorting: Kahn, Tarjan (aka depth-first()), Reverse Kahn and Reverse Tarjan.
package toposort

import (
	"fmt"
	"sort"
	"strings"
)

// KahnSort receives a description of a search tree and returns an array with the elements sorted.
//
// The Kahn's Algorithm creates an "orphan-list" of all nodes that has no parents. Then, it puts
// the first element of that list in the sorted list and removes all edges from that node to
// other nodes; if any of those nodes has no other parents connected, it is appended to the
// orphan-list. The analysis starts again for the first element in the orphan-list.
func KahnSort(tree map[string][]string) ([]string, error) {
	sorted := []string{}
	inDegree := map[string]int{}

	// Make sure all nodes are referred in the map.
	normalizedTree := normalizeTree(tree)

	// 01. Calculate this.indegree of all vertices by going through every edge of the graph;
	// Each child gets indegree++ during breadth-first run.
	for element, children := range normalizedTree {
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
		for _, child := range normalizedTree[node] {
			inDegree[child]--
			if inDegree[child] == 0 {
				stack = append(stack, child)
				inDegree[child] = -1
			}
		}

		// 03.03. Append to the sorted list.
		sorted = append(sorted, node)
	}

	if len(normalizedTree) != len(sorted) {
		//cycle := getCycle(sorted)
		// It seems that there's a directed cycle. Toposort won't work.
		cycle := []string{}
		for element, value := range inDegree {
			if value > 0 {
				cycle = append(cycle, element)
			}
		}

		sort.Slice(cycle, func(i, j int) bool {
			return cycle[i] < cycle[j]
		})

		return []string{}, fmt.Errorf("Cycle involving elements: %s", strings.Join(cycle, ", "))
	}

	return sorted, nil
}

// TarjanSort receives a description of a search tree and returns an array with the elements sorted.
//
// The Tarjan's Algorithm creates an "orphan-list" of all nodes that has no parents. Then, it puts
// the first element of that list in the sorted list and removes all edges from that node to
// other nodes; if any of those nodes has no other parents connected, it is appended to the
// orphan-list. The analysis starts again for the first element in the orphan-list.
func TarjanSort(tree map[string][]string) ([]string, error) {

	// Make sure all nodes are referred in the map.
	normalizedTree := normalizeTree(tree)

	/*
			L ← Empty list that will contain the sorted nodes
		while there are unmarked nodes do
		    select an unmarked node n
		    visit(n)
		function visit(node n)
		    if n has a temporary mark then stop (not a DAG)
		    if n is not marked (i.e. has not been visited yet) then
		        mark n temporarily
		        for each node m with an edge from n to m do
		            visit(m)
		        mark n permanently
		        unmark n temporarily
		        add n to head of L
	*/
	var visitFunc func(string) error
	auxSorted := make([]string, len(normalizedTree))
	index := len(normalizedTree)
	temporary := map[string]bool{}
	visited := map[string]bool{}

	visitFunc = func(node string) error {
		switch {
		case temporary[node]:
			// Cycle found!
			return fmt.Errorf("Found cycle at node: %s", node)
		case visited[node]:
			// Already visited. Moving on...
			return nil
		}

		temporary[node] = true // Mark as temporary to check for cycles...
		for _, child := range normalizedTree[node] {
			err := visitFunc(child) // Visit all children of a node
			if err != nil {
				return err
			}
		}

		delete(temporary, node)
		visited[node] = true
		index--
		auxSorted[index] = node
		return nil
	}

	for element := range normalizedTree {
		if visited[element] {
			continue
		}

		err := visitFunc(element)
		if err != nil {
			return []string{}, err
		}
	}

	sorted := []string{}
	for _, node := range auxSorted {
		if len(node) > 0 {
			sorted = append(sorted, node)
		}
	}

	return sorted, nil
}

// ReverseKahn inverts the order of the elements in the resulting sorted list.
func ReverseKahn(tree map[string][]string) ([]string, error) {
	return reverse(tree, "Kahn")
}

// ReverseTarjan inverts the order of the elements in the resulting sorted list.
func ReverseTarjan(tree map[string][]string) ([]string, error) {
	return reverse(tree, "Tarjan")
}

// Reverse is just a wrapper to invert the order of the elements in a sorted list.
func reverse(tree map[string][]string, algorithm string) ([]string, error) {
	reversed := []string{}

	var sorted []string
	var err error
	switch algorithm {
	case "Kahn":
		sorted, err = KahnSort(tree)
		break
	case "Tarjan":
		sorted, err = TarjanSort(tree)
		break
	}

	if err != nil {
		return []string{}, err
	}

	for i := len(sorted); i > 0; i-- {
		reversed = append(reversed, sorted[i-1])
	}

	return reversed, nil
}

// NormalizeTree will check if all nodes referred in the slices are present in the map as key too.
// If not, it will create the entry to make sure all nodes are accounted for.
func normalizeTree(source map[string][]string) map[string][]string {
	normalized := map[string][]string {}

	for key, values := range source {
		// Copy the valid entry from the source map to the normalized map.
		normalized[key] = values
		for _, node := range values {
			if _, found := source[node]; !found {
				// Current node is in the slice, but not as a key in map.
				// This means we need to treat it as a leaf-node.
				normalized[node] = []string{}
			}
		}
	}

	return normalized
}

// GetCycle will check if there is a cycle in the
// func GetCycle(sorted []string) []string {
// 	for i, e := range(sorted) {
// 		var cycle := []string{e}
// 		for _, f := range(sorted[i:]) {
// 			cycle = append(cycle, f)
// 			if e == f {
// 				// We found a closed cycle.
// 				return cycle
// 			}
// 		}
// 	}

// 	// If we arrive here, there is no cycles!
// 	return []string{}
// }
