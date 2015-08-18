package bubble

import (
	"testing"

	"github.com/hazukac/sort/utils/utils.go"
	. "github.com/otiai10/mint"
)

func TestSort(t *testing.T) {
	Expect(t, utils.GetRandomArray(10)).ToBe([]int{1, 2, 3, 4, 5})
}
