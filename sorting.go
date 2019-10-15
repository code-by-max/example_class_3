package main

import (
	"fmt"
	"math/rand"
	"time"
)

func split(a []int) ([]int, []int) {
	return a[0 : len(a)/2], a[len(a)/2:]
}

//merge sort method
func mergeSort(arr []int, l int, r int) {
	if l < r {
		m := l + (r-l)/2

		mergeSort(arr[:], l, m)
		mergeSort(arr[:], m+1, r)

		merge(arr[:], l, m, r)
	}
}

func merge(arr []int, l int, m int, r int) {
	var i, j, k int
	n1 := m - l + 1
	n2 := r - m

	L := make([]int, n1)
	R := make([]int, n2)

	for i = 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j = 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	i = 0
	j = 0
	k = l

	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

func alt_mergeSort(arr []int, l int, r int, s int) {
	if (r - l) > s {
		mid := (l + r) / 2
		alt_mergeSort(arr, l, mid, s)
		alt_mergeSort(arr, mid+1, r, s)

		merge(arr, l, mid, r)
	} else {
		insertionSort(arr, l, r)
	}
}

//insertion sort method
func insertionSort(a []int, first int, last int) []int {
	for i := first; i <= last && i < len(a); i++ {
		n := a[i]
		j := i
		for j > first && a[j-1] > n {
			a[j] = a[j-1]
			j--
		}
		a[j] = n
	}
	return a
}

func genArr(length int, maxNum int) (result []int) {
	result = make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = rand.Intn(maxNum)
	}

	return
}

func main() {
	arr := genArr(1000000, 10000)
	start := time.Now()
	mergeSort(arr[:], 0, len(arr)-1)
	elapsed := time.Since(start)
	fmt.Println("MergeSort: ", elapsed)

	arr = genArr(1000000, 10000)
	start = time.Now()
	alt_mergeSort(arr[:], 0, len(arr)-1, 100)
	elapsed = time.Since(start)
	fmt.Println("Alt MergeSort: ", elapsed)
}
