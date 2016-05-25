package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	x, y int
}

type Matrix [][]int

func New(s string) (Matrix, error) {
	lines := strings.Split(s, "\n")
	m, n := len(lines), len(strings.Split(lines[0], " "))
	mat := make([][]int, m)
	for i, line := range lines {
		line = strings.TrimSpace(line)
		for _, e := range strings.Split(line, " ") {
			if num, err := strconv.Atoi(e); err != nil {
				return nil, fmt.Errorf("invalid integer: %s", e)
			} else {
				mat[i] = append(mat[i], num)
			}
		}
		if len(mat[i]) != n {
			return nil, fmt.Errorf("non matrix")
		}
	}
	return mat, nil
}

func (m Matrix) Rows() [][]int {
	mat := make([][]int, len(m))
	for i, row := range m {
		mat[i] = make([]int, len(m[0]))
		copy(mat[i], row)
	}
	return mat
}

func (m Matrix) Cols() [][]int {
	rows, cols := len(m[0]), len(m)
	mat := make([][]int, rows)
	for i := 0; i < rows; i++ {
		mat[i] = make([]int, cols)
	}
	for i, row := range m {
		for j, e := range row {
			mat[j][i] = e
		}
	}
	return mat
}

func minmax(a []int, cmp func(int, int) bool) int {
	v := a[0]
	for _, e := range a[1:] {
		if cmp(e, v) {
			v = e
		}
	}
	return v
}

func index(a []int, t int) []int {
	var ans []int
	for i, e := range a {
		if e == t {
			ans = append(ans, i)
		}
	}
	return ans
}

func (m Matrix) Saddle() []Pair {
	var ans []Pair
	cols := m.Cols()
	for i, row := range m {
		rowMax := index(row, minmax(row, func(e, v int) bool { return e > v }))
		for _, j := range rowMax {
			col := cols[j]
			v := m[i][j]
			for _, e := range col {
				if v > e {
					goto next
				}
			}
			ans = append(ans, Pair{i, j})
		next:
		}
	}
	return ans
}
