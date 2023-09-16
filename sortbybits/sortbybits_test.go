package sortbybits

import "testing"

func TestSortByBits(t *testing.T) {
	res := Hello("alfred")
	if res != "Hello alfred" {
		t.Errorf("Hello() = %s; want Hello alfred", res)
	}

}
