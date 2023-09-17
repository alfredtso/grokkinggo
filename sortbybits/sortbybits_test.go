package sortbybits

import (
	"reflect"
	"testing"
)

func TestSortByBits(t *testing.T) {
	res := Hello("alfred")
	if res != "Hello alfred" {
		t.Errorf("Hello() = %s; want Hello alfred", res)
	}

}

func TestSortBits1(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	res := sortByBits(arr)
	expected := []int{0, 1, 2, 4, 8, 3, 5, 6, 7}
	if reflect.DeepEqual(res, expected) != true {
		t.Errorf("sortByBits() = %d; want [0, 1, 2, 4, 8, 3, 5, 6, 7]", res)
	}
}
func TestSortBits2(t *testing.T) {
	arr := []int{10, 100, 1000, 10000}
	res := sortByBits(arr)
	expected := []int{10, 100, 1000, 10000}
	if reflect.DeepEqual(arr, expected) != true {
		t.Errorf("sortByBits() = %d; want [10 100 10000 1000]", res)
	}
}
