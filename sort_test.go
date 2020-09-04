package sort_test

import (
	"fmt"

	. "github.com/otaviokr/topological-sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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

					sorted, err := KahnSort(tree)
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

					sorted, err := KahnSort(tree)
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

					sorted, err := KahnSort(tree)
					Expect(sorted).To(BeEmpty())
					Expect(err).To(Not(BeNil()))
					Expect(err).To(Equal(fmt.Errorf("Cycle involving elements: 5, 6, 7")))
				})
			})

			Context("In-order sorting", func() {
				It("Using empty lists", func() {
					tree := map[string][]string{}

					sorted, err := KahnSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Equal([]string{}))
				})

				It("Using just one element", func() {
					tree := map[string][]string{"Single": {}}

					sorted, err := KahnSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Equal([]string{"Single"}))
				})

				It("Using two elements", func() {
					tree := map[string][]string{"Parent": {}, "Child": {"Parent"}}

					sorted, err := KahnSort(tree)
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

					sorted, err := KahnSort(tree)
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

					sorted, err := KahnSort(tree)
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

					sorted, err := ReverseKahn(tree)
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

					sorted, err := TarjanSort(tree)
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

					sorted, err := TarjanSort(tree)
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

					sorted, err := TarjanSort(tree)
					Expect(sorted).To(BeEmpty())
					Expect(err).To(Not(BeNil()))
					Expect(err.Error()).To(MatchRegexp("Found cycle at node: [5-7]"))
				})
			})

			Context("In-order sorting", func() {
				It("Using empty lists", func() {
					tree := map[string][]string{}

					sorted, err := TarjanSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Equal([]string{}))
				})

				It("Using just one element", func() {
					tree := map[string][]string{"Single": {}}

					sorted, err := TarjanSort(tree)
					Expect(err).To(BeNil())
					Expect(sorted).To(Equal([]string{"Single"}))
				})

				It("Using two elements", func() {
					tree := map[string][]string{"Parent": {}, "Child": {"Parent"}}

					sorted, err := TarjanSort(tree)
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

					sorted, err := TarjanSort(tree)
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

					sorted, err := TarjanSort(tree)
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

					sorted, err := ReverseTarjan(tree)
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

					_, err := ReverseKahn(tree)
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

					_, err := ReverseTarjan(tree)
					Expect(err.Error()).To(MatchRegexp("Found cycle at node: [5-7]"))
				})
			})
		})
	})
})
