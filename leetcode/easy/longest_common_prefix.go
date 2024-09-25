package easy

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for _, stri := range strs {
		for j := 0; j < len(res) && j < len(stri); j++ {
			if res[j] != stri[j] {
				res = res[:j]
				break
			}
		}
		if len(res) > len(stri) {
			res = stri
		}
	}
	return res
}

func TestLCP() {
	var input = []string{"flower", "flaska", "flight"}
	assert := "fl"
	input2 := []string{"", "b"}
	assert2 := ""
	input3 := []string{"aaa", "aa", "aaa"}
	assert3 := "aa"
	if longestCommonPrefix(input) != assert {
		fmt.Println("TestLCP failed")
	}
	if longestCommonPrefix(input2) != assert2 {
		fmt.Println("TestLCP failed")
	}
	if longestCommonPrefix(input3) != assert3 {
		fmt.Println("TestLCP failed")
	}
	fmt.Println("All tests r done")
}
