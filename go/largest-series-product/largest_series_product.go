package lsproduct

import "fmt"

const testVersion = 3

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func each_cons(arr []int64, n, span int) chan []int64 {
	c := make(chan []int64)

	go func() {
		for i := 0; i <= n-span; i++ {
			c <- arr[i : i+span]
		}
		close(c)
	}()

	return c
}

func prod(cons []int64) int64 {
	p := int64(1)
	for _, v := range cons {
		if v != 0 {
			p *= v
		} else {
			p = 0
			break
		}
	}
	return p
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	n := len(digits)
	if n < span {
		return -1, fmt.Errorf("digits too short: %d < %d.", n, span)
	}

	if span < 0 {
		return -1, fmt.Errorf("negative span: %d.", span)
	} else if span == 0 {
		return 1, nil
	}

	arr := make([]int64, n)
	for i, v := range digits {
		if v >= '0' && v <= '9' {
			arr[i] = int64(v - '0')
		} else {
			return 0, fmt.Errorf("invalid digits: %s.", digits)
		}
	}

	ans := int64(0)
	for cons := range each_cons(arr, n, span) {
		ans = max(ans, prod(cons))
	}

	return ans, nil
}
