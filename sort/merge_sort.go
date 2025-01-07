package sort

func merge(arr []int, low int, mid int, high int) {
	temp := make([]int, high+1)
	left := low
	right := mid + 1
	ctr := 0
	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp[ctr] = arr[left]
			left++
		} else {
			temp[ctr] = arr[right]
			right++
		}
		ctr = ctr + 1

	}
	for left <= mid {
		temp[ctr] = arr[left]
		left++
		ctr++
	}
	for right <= high {
		temp[ctr] = arr[right]
		right++
		ctr++
	}
	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}
}
func sort(arr []int, low int, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	sort(arr, low, mid)
	sort(arr, mid+1, high)
	merge(arr, low, mid, high)
}
func MergeSort(arr []int) []int {
	sort(arr, 0, len(arr)-1)
	return arr
}
