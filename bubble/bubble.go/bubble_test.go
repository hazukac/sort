package bubble

import (
	"testing"

	"github.com/hazukac/sort/utils/utils.go"
	. "github.com/otiai10/mint"
)

func TestSort(t *testing.T) {
	src := utils.GetRandomArray(10)
	Expect(t, Sort(src)).ToBe([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}
