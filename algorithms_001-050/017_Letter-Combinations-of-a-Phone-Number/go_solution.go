package main

import "fmt"

func main() {
	cases := []struct {
		digits string
	}{
		{digits: "23"},
		{digits: ""},
		{digits: "2"},
	}

	for _, c := range cases {
		out := letterCombinations(c.digits)
		fmt.Printf("Input: digits = \"%s\"\nOutput: %v \n\n", c.digits, out)
	}
}

// leetcode start

func letterCombinations(digits string) []string {
	outLen := 1
	letterMap := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	for _, d := range digits {
		if letters, ok := letterMap[d]; ok {
			outLen *= len(letters)
		}
	}
	if outLen <= 1 {
		return []string{}
	}

	out := make([]string, outLen)
	pushed := 0
	for _, d := range digits {
		baseLen := pushed

		if letters, ok := letterMap[d]; ok {
			// init first digit's letter
			if baseLen == 0 {
				for i, l := range letters {
					out[i] = string(l)
					pushed++
				}
				continue
			}

			// copy & append letter of next digit
			for loop := 1; loop < len(letters); loop++ {
				appendLetter := string(letters[loop])

				for j := 0; j < baseLen; j++ {
					out[baseLen*(loop)+j] = out[j] + appendLetter
					pushed++
				}
			}
			for j := 0; j < baseLen; j++ {
				out[j] = out[j] + string(letters[0])
			}
		}
	}

	return out
}
