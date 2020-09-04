package toposort_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/otaviokr/topological-sort/toposort"
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
		{
			T: map[string][]string{
				"Parent": {"Child"},
				"Child":  {}},
			R: [][]string{
				{"Parent", "Child"}}},

		// Empty tree.
		{
			T: map[string][]string{},
			R: [][]string{{}}},

		// Single node in the tree.
		{
			T: map[string][]string{
				"A": {}},
			R: [][]string{
				{"A"}}},

		// Two nodes in a single tree.
		{
			T: map[string][]string{
				"A": {},
				"B": {"A"}},
			R: [][]string{
				{"B", "A"}}},

		// Complex tree with multiple nodes.
		{
			T: map[string][]string{
				"0": {"1", "4"},
				"1": {"3", "5"},
				"2": {"5"},
				"3": {"5", "7"},
				"4": {},
				"5": {"6"},
				"6": {"7"},
				"7": {}},
			R: [][]string{
				{"2", "0", "4", "1", "3", "5", "6", "7"},
				{"0", "4", "1", "3", "2", "5", "6", "7"}}}}
}

func TestKahnSort(t *testing.T) {
	for i, testCase := range TestableTrees {
		sorted, err := toposort.KahnSort(testCase.T)
		if err != nil {
			t.Error(err.Error())
		}
		//fmt.Println("Sorted: ", sorted)

		result := compareResults(sorted, testCase.R)
		if len(result) > 0 {
			errMsgTmpl := "Error on iteration %d - %s:\nExpected: %+v\nFound: %+v"
			errMsg := fmt.Sprintf(errMsgTmpl, i, result, testCase.R, sorted)
			t.Fatal(errMsg)
		}
	}
}

func TestReverse(t *testing.T) {
	for i, testCase := range TestableTrees {
		reversed, err := toposort.ReverseKahn(testCase.T)
		if err != nil {
			t.Error(err.Error())
		}
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
