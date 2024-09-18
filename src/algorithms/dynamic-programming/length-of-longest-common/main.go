package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// calcula o comprimento da maior subsequência comum entre duas strings
func longestCommonSubsequence(arr1, arr2 string) int {
	m := len(arr1)
	n := len(arr2)

	lcs := make([][]int, m+1)
	for i := range lcs {
		lcs[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if arr1[i-1] == arr2[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}

	return lcs[m][n]
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter the 1st string: ")
	arr1, _ := reader.ReadString('\n')
	arr1 = strings.TrimSpace(arr1)

	fmt.Println("enter the 2nd string: ")
	arr2, _ := reader.ReadString('\n')
	arr2 = strings.TrimSpace(arr2)

	fmt.Printf("length of the Longest Common Subsequence is: %d\n", longestCommonSubsequence(arr1, arr2))
}
