package expression

import (
	"testing"

	"github.com/JekaMas/go-mutesting/test"
)

func TestMutatorRemoveTerm(t *testing.T) {
	test.Mutator(
		t,
		MutatorRemoveTerm,
		"../../testdata/expression/remove.go",
		6,
	)
}
