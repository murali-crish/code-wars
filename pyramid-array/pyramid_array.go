package pyramidarray

func getPyramid(size int) [][]int {
	pyramid := [][]int{}
	if size > 0 {
		for i := 1; i <= size; i++ {
			var dyn []int
			for j := 0; j < i; j++ {
				dyn = append(dyn, 1)
			}
			pyramid = append(pyramid, dyn)
		}
	}
	return pyramid
}

func BestSolutionPyramid(n int) [][]int {
    row := [][]int{}
    cell := []int{}
    
    for i := 0; i < n; i++ {
      cell = append(cell, 1)
      row = append(row, cell)
    }
    
    return row
}