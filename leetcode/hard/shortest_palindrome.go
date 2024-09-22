package hard

import (
	"fmt"
)

func shortestPalindrome(s string) string {
	if isPalindrome(s) {
		return s
	}
	from := indexLeastPalindrome(s)
	res := s[0:from]
	for i := from; i < len(s); i++ {
		res = string(s[i]) + res + string(s[i])
	}
	return res
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func indexLeastPalindrome(s string) int {
	index_max := 0
	for i := 1; i < len(s); i++ {
		if isPalindrome(s[0:i]) {
			index_max = i
		}
	}
	return index_max
}

func TestSTPDR() {
	// task consists only lower case letters
	case_one := "aacecaaa"
	res_one := "aaacecaaa"
	case_two := "abcd"
	res_two := "dcbabcd"
	case_three := "abcdcba"
	res_three := "abcdcba"
	case_four := "abc"
	res_four := "cbabc"
	case_five := "abbacd"
	res_five := "dcabbacd"
	passed := 0
	total := 5
	if shortestPalindrome(case_one) == res_one {
		passed += 1
	}
	if shortestPalindrome(case_two) == res_two {
		passed += 1
	}
	if shortestPalindrome(case_three) == res_three {
		passed += 1
	}
	if shortestPalindrome(case_four) == res_four {
		passed += 1
	}
	if shortestPalindrome(case_five) == res_five {
		passed += 1
	}
	if passed == total {
		fmt.Println("All tests has been passed")
	} else {
		output := fmt.Sprintf("Something went wrong, completed %d of %d", passed, total)
		fmt.Println(output)
	}
	fmt.Println(shortestPalindrome(case_one))
	fmt.Println(shortestPalindrome(case_two))
	fmt.Println(shortestPalindrome(case_three))
	fmt.Println(shortestPalindrome(case_four))
	fmt.Println((shortestPalindrome(case_five)))
}
