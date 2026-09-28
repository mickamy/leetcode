package main

import (
	"math"
)

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	var targetCount [128]int
	for i := 0; i < len(t); i++ {
		targetCount[t[i]]++
	}

	var windowCount [128]int

	required := 0
	for _, c := range targetCount {
		if c > 0 {
			required++
		}
	}

	var formed, l, start int
	minLen := math.MaxInt

	for r := 0; r < len(s); r++ {
		charR := s[r]
		windowCount[charR]++

		if targetCount[charR] > 0 && windowCount[charR] == targetCount[charR] {
			formed++
		}

		for formed == required {
			if r-l+1 < minLen {
				minLen = r - l + 1
				start = l
			}

			charL := s[l]
			windowCount[charL]--
			if targetCount[charL] > 0 && windowCount[charL] < targetCount[charL] {
				formed--
			}
			l++
		}
	}

	if minLen == math.MaxInt {
		return ""
	}
	return s[start : start+minLen]
}

//func main() {
//	fmt.Println(minWindow(
//		"rsfvquycmabtdxcmgwnoxiicpzxczxnspungqokcolwtlahrzvhwfqawraytxoloibuzpgdfsbbdeiwdddbivcenefidttrbuoclugtmjurncpqitssqwzdcelfxhwadkdrhlktueniicaqxulosktuohbnqantmsktdupaaeilgkdgfowzapuoyxdoxriklufprurtabehlipsylszampzltpjmxvxucolzfezglgutvmtgesjsikedzppzkotmdelkknrvvnqrzkzeekwmucpwrgvdvosaufkgsdeoquhqggwltuxpxplovguswmssrdzkidyzzfrgnnrghqwghjvfqpxodshuvgefjeeijvqzkjhsafyhpvohmmvpecmbkqgvxnggkwhvgppiwprlngubuqryufbfvjcsyibjhkgpbxxyoolbwgdqwwqsbtzydctdmwovgukkthiiytjwzonxkfsnzroevhhifydnsqxozhwzbhwcwlciquyrlmvhvfradwpxmnnoujnxtjcpueznyzellcuijnkokanypbmlmcjhllyiryxpxdclfxtqkeewguxvxlglnptbcrjtpikplnatmvjybqvhkucjfjutanufffawnnlacntoldmmtjufdwo",
//		"jhkeiyapffpbu",
//	))
//}
