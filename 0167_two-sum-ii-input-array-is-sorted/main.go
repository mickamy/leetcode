package main

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		l, r := numbers[left], numbers[right]
		if l+r == target {
			return []int{left + 1, right + 1}
		}
		if l+r < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

//func main() {
//	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
//	fmt.Println(twoSum([]int{2, 3, 4}, 6))
//}
