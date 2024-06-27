package main

func findCenter(edges [][]int) int {
    var n int=len(edges)+1
	degreeCount := make([]int, n+1)

	for _, edge := range edges {
		degreeCount[edge[0]]++;
        degreeCount[edge[1]]++;
	}

	for i := 1; i <= n; i++ {
		if degreeCount[i]==n-1 {
			return i
		}
	}

	return -1
}