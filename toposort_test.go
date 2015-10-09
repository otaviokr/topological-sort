package sort

import (
	"testing"
	"fmt"
	"reflect"
)

type Testable struct {
	T map[string][]string
	R [][]string
}

var (
	TestableTrees []Testable
)

func init() {
	TestableTrees = []Testable{
		Testable{
			T: map[string][]string{
				"Parent": []string{"Child"},
				"Child": []string{}},
			R: [][]string{
				[]string{"Parent", "Child"}}},

		// Empty tree.
		Testable{
			T: map[string][]string{},
			R: [][]string{[]string{}}},

		// Single node in the tree.
		Testable{
			T: map[string][]string{
				"A": []string{}},
			R: [][]string{
				[]string{"A"}}},

		// Two nodes in a single tree.
		Testable{
			T: map[string][]string{
				"A": []string{},
				"B": []string{"A"}},
			R: [][]string{
				[]string{"B", "A"}}},

		// Complex tree with multiple nodes.
		Testable{
			T: map[string][]string{
				"0": []string{"1", "4"},
				"1": []string{"3", "5"},
				"2": []string{"5"},
				"3": []string{"5", "7"},
				"4": []string{},
				"5": []string{"6"},
				"6": []string{"7"},
				"7": []string{}},
			R: [][]string{
				[]string{"2", "0", "4", "1", "3", "5", "6", "7"},
				[]string{"0", "4", "1", "3", "2", "5", "6", "7"}}}}
}

func TestSort(t *testing.T) {
	for i, testCase := range TestableTrees {
		sorted := Sort(testCase.T)
		//fmt.Println("Sorted: ", sorted)

		result := compareResults(sorted, testCase.R)
		if len(result) > 0 {
			t.Fatal(i, result)
		}
	}
}

func TestReversedSort(t *testing.T) {
	for i, testCase := range TestableTrees {
		reversed := ReversedSort(testCase.T)
		//fmt.Println("Sorted: ", reversed)

		expected := [][]string{}
		for _, a := range testCase.R {
			sub := []string{}
			for j := len(a); j > 0; j-- {
				sub = append(sub, a[j-1])
			}
			expected = append(expected, sub)
		}

		result := compareResults(reversed, expected)
		if len(result) > 0 {
			t.Fatal(testCase, i, result)
		}
	}
}

func compareResults(testable []string, expected [][]string) string {
	for _, r := range expected {
		if len(testable) != len(r) {
			errorMsg := "Sorted array has a different number of elements - Expected %d, found %d."
			return fmt.Sprintf(errorMsg, len(r), len(testable))
		}
	}

	if len(testable) == 0 {
		return ""
	}

	for _, r := range expected {
		if reflect.DeepEqual(r, testable) {
			return ""
		}
	}

	return "Not a match."
}
