package sort

import (
	"testing"
)

type expected struct {
	T map[string][]string
	R [][]string
}

var data expected

func init() {
	data = expected{T: map[string][]string{"0": []string{"1", "4"}, "1": []string{"3", "5"}, "2": []string{"5"},
		"3": []string{"5", "7"}, "4": []string{}, "5": []string{"6"}, "6": []string{"7"}, "7": []string{}},
		R: [][]string{[]string{"2", "0", "4", "1", "3", "5", "6", "7"},
			[]string{"0", "4", "1", "3", "2", "5", "6", "7"}}}
}

func TestSort(t *testing.T) {
	sorted := Sort(data.T)

	for _, r := range data.R {
		if len(sorted) != len(r) {
			t.Fatalf("Sorted array has a different number of elements - Expected %i, found %i.", len(r), len(sorted))
		}
	}

	match := []bool{true, true}
	for i, r := range data.R {
		for j, item := range sorted {
			if item != r[j] {
				match[i] = false
			}
		}
	}

	if !(match[0] || match[1]) {
		t.Fatal("Not a match.")
	}
}

func TestReversedSort(t *testing.T) {

	reversed := ReversedSort(data.T)

	for _, r := range data.R {
		if len(reversed) != len(r) {
			t.Fatalf("Sorted array has a different number of elements - Expected %i, found %i.", len(r), len(reversed))
		}
	}

	match := []bool{true, true}
	for i, r := range data.R {
		for j, item := range reversed {
			if item != r[len(r)-(j+1)] {
				match[i] = false
			}
		}
	}

	if !(match[0] || match[1]) {
		t.Fatal("Not a match.")
	}

}
