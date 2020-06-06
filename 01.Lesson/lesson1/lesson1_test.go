package lesson1

// package lesson1
import "testing"

func TestTwoSum(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9

	got := twoSum(numbers, target)
	want := []int{0, 1}

	if got[0] != want[0] && got[1] != want[1] {
		t.Error("got", got, "want", want)
	}
}

func TestTwoSum2(t *testing.T) {
	numbers := []int{15, 2, 8, 11, 12, 20}
	target := 14

	got := twoSum(numbers, target)
	want := []int{1, 4}

	if got[0] != want[0] && got[1] != want[1] {
		t.Error("got", got, "want", want)
	}
}
