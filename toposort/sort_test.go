package toposort_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/otaviokr/topological-sort/toposort"
)

// The tree passed as argument means
//
// - *A* is the head, and it is connected to *B* and *C*;
// - *B* is not connected to anything;
// - *C* is connected to bin.
func ExampleKahnSort() {
	result, err := toposort.KahnSort(map[string][]string{
		"A": {"B", "C"},
		"B": {},
		"C": {"B"},
	})

	if err != nil {
		fmt.Println("Example for KahnSort failed!")
		return
	}
	fmt.Println(strings.Join(result, ","))
	// Output: A,C,B
}

// The result is the opposite of the usual Kahn sorting algorithm. The tree passed as argument means
//
// - *A* is the head, and it is connected to *B* and *C*;
// - *B* is not connected to anything;
// - *C* is connected to bin.
func ExampleReverseKahn() {
	result, err := toposort.ReverseKahn(map[string][]string{
		"A": {"B", "C"},
		"B": {},
		"C": {"B"},
	})

	if err != nil {
		fmt.Println("Example for ReverseKahnt failed!")
		return
	}
	fmt.Println(strings.Join(result, ","))
	// Output: B,C,A
}

func ExampleTarjanSort() {
	result, err := toposort.TarjanSort(map[string][]string{
		"A": {"B", "C"},
		"B": {},
		"C": {"B"},
	})

	if err != nil {
		fmt.Println("Example for TarjanSort failed!")
		return
	}
	fmt.Println(strings.Join(result, ","))
	// Output: A,C,B
}

// The result is the opposite of the usual Kahn sorting algorithm. The tree passed as argument means
//
// - *A* is the head, and it is connected to *B* and *C*;
// - *B* is not connected to anything;
// - *C* is connected to bin.
func ExampleReverseTarjan() {
	result, err := toposort.ReverseTarjan(map[string][]string{
		"A": {"B", "C"},
		"B": {},
		"C": {"B"},
	})

	if err != nil {
		fmt.Println("Example for ReverseKahnt failed!")
		return
	}
	fmt.Println(strings.Join(result, ","))
	// Output: B,C,A
}

var _ = Describe("Sort", func() {
	Describe("Sorting using Kahn's", func() {
		Describe("Topological Sorting", func() {
			Context("Identifying direct cycles", func() {
				It("No direct cycles in the graph", func() {
					tree := map[string][]string{
						"0": {"1", "4"},
						"1": {"3", "5"},
						"2": {"5"},
						"3": {"5", "7"},
						"4": {},
						"5": {"6"},
						"6": {"7"},
						"7": {}}

					sorted, err := toposort.KahnSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Not(BeNil()))
					Expect(sorted).To(Not(BeEmpty()))
					Expect(sorted).To(HaveLen(8))
				})

				It("Graph is only a directed cycle", func() {
					tree := map[string][]string{
						"0": {"1"},
						"1": {"2"},
						"2": {"3"},
						"3": {"4"},
						"4": {"0"}}

					sorted, err := toposort.KahnSort(tree)
					Expect(sorted).To(BeEmpty())
					Expect(err).To(Not(BeNil()))
				})

				It("Contains a directed cycle in the graph", func() {
					tree := map[string][]string{
						"0": {"1", "4"},
						"1": {"3", "5"},
						"2": {"5"},
						"3": {},
						"4": {},
						"5": {"6"},
						"6": {"7"},
						"7": {"5"}}

					sorted, err := toposort.KahnSort(tree)
					Expect(sorted).To(BeEmpty())
					Expect(err).To(Not(BeNil()))
					Expect(err).To(Equal(fmt.Errorf("Cycle involving elements: 5, 6, 7")))
				})
			})

			Context("In-order sorting", func() {
				It("Using empty lists", func() {
					tree := map[string][]string{}

					sorted, err := toposort.KahnSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Equal([]string{}))
				})

				It("Using just one element", func() {
					tree := map[string][]string{"Single": {}}

					sorted, err := toposort.KahnSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Equal([]string{"Single"}))
				})

				It("Using two elements", func() {
					tree := map[string][]string{"Parent": {}, "Child": {"Parent"}}

					sorted, err := toposort.KahnSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Equal([]string{"Child", "Parent"}))
				})

				It("Using multiple elements", func() {
					tree := map[string][]string{
						"0": {"1", "4"},
						"1": {"3", "5"},
						"2": {"5"},
						"3": {"5", "7"},
						"4": {},
						"5": {"6"},
						"6": {"7"},
						"7": {}}

					sorted, err := toposort.KahnSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Or(
						Equal([]string{"2", "0", "4", "1", "3", "5", "6", "7"}),
						Equal([]string{"0", "4", "1", "3", "2", "5", "6", "7"})))
				})

				It("Element '2' is completely disconnected", func() {
					tree := map[string][]string{
						"0": {"1", "4"},
						"1": {"3", "5"},
						"2": {},
						"3": {"5", "7"},
						"4": {},
						"5": {"6"},
						"6": {"7"},
						"7": {}}

					sorted, err := toposort.KahnSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Or(
						Equal([]string{"0", "4", "1", "3", "5", "6", "7", "2"}),
						Equal([]string{"2", "0", "4", "1", "3", "5", "6", "7"})))
				})
			})

			Context("Reverse sorting", func() {
				It("Successful example", func() {
					tree := map[string][]string{
						"0": {"1", "4"},
						"1": {"3", "5"},
						"2": {"5"},
						"3": {"5", "7"},
						"4": {},
						"5": {"6"},
						"6": {"7"},
						"7": {}}

					sorted, err := toposort.ReverseKahn(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Or(
						Equal([]string{"7", "6", "5", "3", "1", "4", "0", "2"}),
						Equal([]string{"7", "6", "5", "2", "3", "1", "4", "0"})))
				})
			})
		})
	})

	Describe("Sorting using Tarjan's", func() {
		Describe("Topological Sorting", func() {
			Context("Identifying direct cycles", func() {
				It("No direct cycles in the graph", func() {
					tree := map[string][]string{
						"0": {"1", "4"},
						"1": {"3", "5"},
						"2": {"5"},
						"3": {"5", "7"},
						"4": {},
						"5": {"6"},
						"6": {"7"},
						"7": {}}

					sorted, err := toposort.TarjanSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Not(BeNil()))
					Expect(sorted).To(Not(BeEmpty()))
					Expect(sorted).To(HaveLen(8))
				})

				It("Graph is only a directed cycle", func() {
					tree := map[string][]string{
						"0": {"1"},
						"1": {"2"},
						"2": {"3"},
						"3": {"4"},
						"4": {"0"}}

					sorted, err := toposort.TarjanSort(tree)
					Expect(sorted).To(BeEmpty())
					Expect(err).To(Not(BeNil()))
				})

				It("Contains a directed cycle in the graph", func() {
					tree := map[string][]string{
						"0": {"1", "4"},
						"1": {"3", "5"},
						"2": {"5"},
						"3": {},
						"4": {},
						"5": {"6"},
						"6": {"7"},
						"7": {"5"}}

					sorted, err := toposort.TarjanSort(tree)
					Expect(sorted).To(BeEmpty())
					Expect(err).To(Not(BeNil()))
					Expect(err.Error()).To(MatchRegexp("Found cycle at node: [5-7]"))
				})
			})

			Context("In-order sorting", func() {
				It("Using empty lists", func() {
					tree := map[string][]string{}

					sorted, err := toposort.TarjanSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Equal([]string{}))
				})

				It("Using just one element", func() {
					tree := map[string][]string{"Single": {}}

					sorted, err := toposort.TarjanSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Equal([]string{"Single"}))
				})

				It("Using two elements", func() {
					tree := map[string][]string{"Parent": {}, "Child": {"Parent"}}

					sorted, err := toposort.TarjanSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Equal([]string{"Child", "Parent"}))
				})

				It("Using multiple elements", func() {
					tree := map[string][]string{
						"0": {"1", "4"},
						"1": {"3", "5"},
						"2": {"5"},
						"3": {"5", "7"},
						"4": {},
						"5": {"6"},
						"6": {"7"},
						"7": {}}

					sorted, err := toposort.TarjanSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Or(
						Equal([]string{"0", "1", "4", "3", "2", "5", "6", "7"}),
						Equal([]string{"2", "0", "4", "1", "3", "5", "6", "7"}),
						Equal([]string{"2", "0", "1", "4", "3", "5", "6", "7"}),
						Equal([]string{"2", "0", "1", "3", "5", "6", "7", "4"}),
						Equal([]string{"0", "4", "2", "1", "3", "5", "6", "7"})))
				})

				It("Element '2' is completely disconnected", func() {
					tree := map[string][]string{
						"0": {"1", "4"},
						"1": {"3", "5"},
						"2": {},
						"3": {"5", "7"},
						"4": {},
						"5": {"6"},
						"6": {"7"},
						"7": {}}

					sorted, err := toposort.TarjanSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Or(
						Equal([]string{"0", "1", "4", "3", "5", "6", "7", "2"}),
						Equal([]string{"0", "4", "2", "1", "3", "5", "6", "7"}),
						Equal([]string{"2", "0", "4", "1", "3", "5", "6", "7"}),
						Equal([]string{"2", "0", "1", "4", "3", "5", "6", "7"}),
						Equal([]string{"2", "0", "1", "3", "5", "6", "7", "4"})))
				})
			})

			Context("Reverse sorting", func() {
				It("Successful example", func() {
					tree := map[string][]string{
						"0": {"1", "4"},
						"1": {"3", "5"},
						"2": {"5"},
						"3": {"5", "7"},
						"4": {},
						"5": {"6"},
						"6": {"7"},
						"7": {}}

					sorted, err := toposort.ReverseTarjan(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Or(
						Equal([]string{"4", "7", "6", "5", "3", "1", "0", "2"}),
						Equal([]string{"7", "6", "5", "3", "1", "2", "4", "0"}),
						Equal([]string{"7", "6", "5", "3", "1", "4", "0", "2"}),
						Equal([]string{"7", "6", "5", "3", "4", "1", "0", "2"}),
						Equal([]string{"7", "6", "5", "2", "3", "4", "1", "0"})))
				})

				It("Fail reverse because Kahn sort failed", func() {
					tree := map[string][]string{
						"0": {"1", "4"},
						"1": {"3", "5"},
						"2": {"5"},
						"3": {},
						"4": {},
						"5": {"6"},
						"6": {"7"},
						"7": {"5"}}

					_, err := toposort.ReverseKahn(tree)
					Expect(err).To(Equal(fmt.Errorf("Cycle involving elements: 5, 6, 7")))
				})

				It("Fail reverse because Tarjan sort failed", func() {
					tree := map[string][]string{
						"0": {"1", "4"},
						"1": {"3", "5"},
						"2": {"5"},
						"3": {},
						"4": {},
						"5": {"6"},
						"6": {"7"},
						"7": {"5"}}

					_, err := toposort.ReverseTarjan(tree)
					Expect(err.Error()).To(MatchRegexp("Found cycle at node: [5-7]"))
				})
			})
		})
	})
})
