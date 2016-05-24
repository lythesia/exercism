package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (mat Matrix) Set(row, col, val int) bool {
	m, n := len(mat), len(mat[0])
	if row < 0 || row >= m || col < 0 || col >= n {
		return false
	}
	mat[row][col] = val
	return true
}
