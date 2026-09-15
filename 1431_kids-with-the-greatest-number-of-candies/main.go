package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	var maximum int
	for _, candy := range candies {
		maximum = max(maximum, candy)
	}

	for i := range candies {
		res[i] = candies[i]+extraCandies >= maximum
	}

	return res
}
