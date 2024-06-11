package main

func relativeSortArray(arr1 []int, arr2 []int) []int {
    frequency:=make([]int,1001)

	for _, number := range arr1 {
		frequency[number]++
	}

	result:=make([]int,len(arr1))
	index:=0

	for _, number := range arr2 {
		for frequency[number]>0{
			result[index]=number
			index++
			frequency[number]--
		}
	}

	for i := 0; i < len(frequency); i++ {
		for frequency[i]>0{
			result[index]=i
			index++
			frequency[i]--
		}
	}

	return result
}