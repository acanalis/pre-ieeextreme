package main

import "fmt"

func main() {
	//var seq = []int32{2,7,4,3,8} // result: 2,7,8 = 3
	var seq = []int32{2,4,3,7,4,5} // result 2,3,4,5 = 4
	fmt.Print(longestIncreasingSubsequence(seq))
}

func longestIncreasingSubsequence(arr []int32) int32 {
	var longest int32
	for i,u := range arr {
		
			var currentLength int32 = 1
			for j := i + 1; j < len(arr); j++ {
				v := arr[j]
				if v > u {
					u = v
					currentLength++
				}
				if longest > int32(len(arr[i:]))+currentLength {
					break
				}
			}
			if currentLength > longest {
				longest = currentLength
			}
		}

	return longest
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
