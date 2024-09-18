package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
  A DynamicProgramming based solution for Edit Distance problem
  Description of Edit Distance with an Example:

  Edit distance is a way of quantifying how dissimilar two strings (e.g., words) are to one another,
  by counting the minimum number of operations required to transform one string into the other. The
  distance operations are the removal, insertion, or substitution of a character in the string.


  The Distance between "kitten" and "sitting" is 3. A minimal edit script that transforms the former into the latter is:

  kitten → sitten (substitution of "s" for "k")
  sitten → sittin (substitution of "i" for "e")
  sittin → sitting (insertion of "g" at the end).


*/

// calcula a distância mínima de edit entre duas strings
func minDistance(word1, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)

	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	for i := 0; i <= len1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				replace := dp[i][j] + 1
				insert := dp[i][j+1] + 1
				delete := dp[i+1][j] + 1
				dp[i+1][j+1] = min(replace, insert, delete)
			}
		}
	}


	return dp[len1][len2]
}


func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter with the first string:")
	s1, _ := reader.ReadString('\n')
	s1 = strings.TrimSpace(s1)

	fmt.Println("enter the second sring:")
	s2, _ := reader.ReadString('\n')
	s2 = strings.TrimSpace(s2)

	ans := minDistance(s1, s2)
	fmt.Printf("the minimum edit distance between \"%s\" and \"%s\" is %d\n", s1, s2, ans)
}
