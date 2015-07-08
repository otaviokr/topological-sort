package sort

// ReversedSort returns the sorted list of nodes in the reverse order compared to Sort.
func ReversedSort(tree map[string][]string) []string {
sorted := Sort(tree)
reversed := []string{}

for i := len(sorted); i > 0; i-- {
reversed = append(reversed, sorted[i-1])
}

return reversed
}

// Sort receives a description of a search tree and returns an array with the elements sorted.
// Example for tree: [map[rule1: [rule2 rule3], rule2: [], rule3: [rule2]]]
func Sort(tree map[string][]string) []string {
sorted := []string{}
indegree := map[string]int{}

// 01. Calculate this.indegree of all vertices by going through every edge of the graph;
// Each child gets indegree++ during breadth-first run.
for element, children := range tree {
if indegree[element] == 0 {
indegree[element] = 0
}
for _, child := range children {
//fmt.Printf("Element: %s - Children: %v - Child: %s - InDegree: %d\n", element, children, child, indegree[child])
indegree[child]++
}
}

//fmt.Printf("InDegree: %v\n", indegree)

// 02. Collect all vertices with indegree==0 onto a stack;
stack := []string{}
for rule, value := range indegree {
if value == 0 {
stack = append(stack, rule)
indegree[rule] = -1
}
}

// 03. While zero-degree-stack is not empty:
for ; len(stack) > 0; {
// 03.01. Pop element from zero-degree-stack and append it to topological order;
var rule string
rule, stack = stack[len(stack)-1], stack[:len(stack)-1]

// 03.02. Find all children of element and decrease indegree. If indegree becomes 0, add to zero-degree-stack;
for _, children := range tree[rule] {
for child := range children {
element := string(children[child])
indegree[element]--
if indegree[element] == 0 {
stack = append(stack, element)
indegree[element] = -1
}
}
}

// 03.03. Append to the sorted list.
sorted = append(sorted, rule)
}
return sorted
}
