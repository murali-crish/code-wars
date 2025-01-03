package pyramidarray

import (
	"reflect"
	"testing"
)

func TestGetPyramid(t *testing.T) {
	t.Run("My Solution", func(t *testing.T) {
		assertEqual(getPyramid(0), [][]int{}, t)
		assertEqual(getPyramid(1), [][]int{{1}}, t)
		assertEqual(getPyramid(2), [][]int{{1}, {1, 1}}, t)
		assertEqual(getPyramid(3), [][]int{{1}, {1, 1}, {1, 1, 1}}, t)
	})
	t.Run("Best Solution", func(t *testing.T) {
		assertEqual(BestSolutionPyramid(0), [][]int{}, t)
		assertEqual(BestSolutionPyramid(1), [][]int{{1}}, t)
		assertEqual(BestSolutionPyramid(2), [][]int{{1}, {1, 1}}, t)
		assertEqual(BestSolutionPyramid(3), [][]int{{1}, {1, 1}, {1, 1, 1}}, t)
	})
}

func assertEqual(got, expected [][]int, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
