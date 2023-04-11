package main

import "fmt"

func main() {
	cases := []struct {
		n int
		k int
	}{
		{n: 3, k: 3},
		{n: 4, k: 9},
		{n: 3, k: 1},
		{n: 9, k: 40},
	}

	for _, c := range cases {
		out := getPermutation(c.n, c.k)
		fmt.Printf("Input: n = %d, k = %d\nOutput: \"%s\"\n\n", c.n, c.k, out)
	}
}

// leetcode start

func getPermutation(n int, k int) string {
	sf := initFact(n)
	usedMap := make(map[int]bool, n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	// 1 2 3 : 1-1/ 2 = 0 : 1%2 = 1
	// 1 3 2 : 2-1/ 2 = 0 : 2%2 = 0 => 2
	// 2 1 3 : 3-1/ 2 = 1 : 3%2 = 1
	// 2 3 1 : 4-1/ 2 = 1 : 4%2 = 0 => 2
	// 3 1 2 : 5-1/ 2 = 2 : 5%2 = 1
	// 3 2 1 : 6-1/ 2 = 2 : 6%2 = 0 => 2
	out := ""
	for i := 1; i <= n; i++ {
		modFact := sf.get(n - i)
		pick := (k - 1) / modFact

		for j := 0; j < n; j++ {
			if usedMap[j] {
				pick++
				continue
			}

			if j == pick {
				out += string(rune(arr[j]) + '0')
				usedMap[j] = true
				break
			}
		}

		k = k % modFact
		if k == 0 {
			k = modFact
		}
	}

	return out
}

type smartFact struct {
	cache map[int]int
}

func initFact(n int) *smartFact {
	s := &smartFact{
		cache: map[int]int{},
	}

	s.cache[0] = 1
	tmp := 1
	for i := 1; i <= n; i++ {
		tmp *= i
		s.cache[i] = tmp
	}

	return s
}

func (s *smartFact) get(n int) int {
	return s.cache[n]
}
