package main

func main() {

}

func FindElementInArray(arr []int, key, low, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if arr[mid] == key {
		return mid
	}
	// assume this half is sorted
	if arr[low] <= arr[mid] {
		// key is in here
		if key >= arr[low] && key <= arr[mid] {
			return FindElementInArray(arr, key, low, mid-1)
		}
		// Key must be in the other half
		return FindElementInArray(arr, key, mid+1, high)
	}
	// this must be the sorted half then?
	if key >= arr[mid] && key <= arr[high] {
		return FindElementInArray(arr, key, mid+1, high)
	}
	//Must be in the other half
	return FindElementInArray(arr, key, low, mid-1)
}
