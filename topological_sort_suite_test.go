package sort_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"testing"
)

func TestTopologicalSort(t *testing.T) {
	gomega.RegisterFailHandler(Fail)
	RunSpecs(t, "TopologicalSort Suite")
}
