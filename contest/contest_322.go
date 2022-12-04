package contest

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/contest/weekly-contest-322/
func Test322() {
	// fmt.Println(isCircularSentence("leetcode exercises sound delightful"))
	// fmt.Println(isCircularSentence("I like Leetcode"))

	// fmt.Println(dividePlayers([]int{3, 2, 5, 1, 3, 4}))

	fmt.Println(minScore(4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}))
	fmt.Println(minScore(4, [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}))
}

//https://leetcode.cn/problems/minimum-score-of-a-path-between-two-cities/
// 并查集， BFS
func minScore(n int, roads [][]int) int {

	sort.Slice(roads, func(i, j int) bool {
		return roads[i][0] < roads[j][0]
	})

	maps := make(map[int]bool)
	maps[1] = true
	minLen := 1000001
	for i := 0; i < len(roads); i++ {
		_, ok := maps[roads[i][0]]
		if ok {
			maps[roads[i][1]] = true
			if roads[i][2] < minLen {
				minLen = roads[i][2]
			}
		}
	}
	return minLen
}

func dividePlayers(skill []int) int64 {
	sum := 0
	for _, num := range skill {
		sum += num
	}
	if sum%(len(skill)/2) != 0 {
		return -1
	}
	twoSum := sum / (len(skill) / 2)

	maps := make(map[int]int)

	for _, num := range skill {
		value, ok := maps[num]
		if ok {
			maps[num] = value + 1
		} else {
			maps[num] = 1
		}
	}

	var result int64 = 0
	for _, num := range skill {
		value, ok := maps[twoSum-num]
		if ok {
			if value <= 0 {
				return -1
			} else {
				maps[twoSum-num] = value - 1
				result = int64(num*(twoSum-num)) + result
			}
		} else {
			return -1
		}
	}

	return result / 2
}

func isCircularSentence(sentence string) bool {
	if len(sentence) == 0 {
		return false
	}

	start := sentence[0]
	end := sentence[len(sentence)-1]

	if start != end {
		return false
	}
	for i := 0; i < len(sentence); i++ {
		ch := sentence[i]
		if ch == ' ' && sentence[i-1] != sentence[i+1] {
			return false
		}
	}

	return true
}
