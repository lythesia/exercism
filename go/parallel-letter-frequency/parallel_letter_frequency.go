package letter

func ConcurrentFrequency(strs []string) FreqMap {
	ans := FreqMap{}

	ch := make(chan FreqMap)
	for _, s := range strs {
		go func(_s string) {
			ch <- Frequency(_s)
		}(s)
	}
	for i := 0; i < len(strs); i++ {
		h := <-ch
		for k, v := range h {
			ans[k] += v // 0 as default
		}
	}
	return ans
}
