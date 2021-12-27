package selection_sort

import (
	"fmt"
	"github.com/panutat-p/data-structures-algorithms-beginner/sorts/bogo_sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	sl := []int{5, 3, 7, 0, 14, 8, 1, 13, 21, 82, 44}
	sortedSl := SelectionSort(sl)
	fmt.Println(sortedSl)

	if !bogo_sort.IsSorted(sortedSl) {
		fmt.Println(sortedSl)
		t.Errorf("got %v, but expect %v", bogo_sort.IsSorted(sortedSl), true)
	}
}
