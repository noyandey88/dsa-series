package slidingwindow

func MaxWatchTime(arr []int, size int) int {
	current := 0

	for i := range size {
		current += arr[i]
	}

	maxTime := current

	for i := 1; i <= len(arr)-size; i++ {
		current = current - arr[i-1] + arr[i+size-1]

		if current > maxTime {
			maxTime = current
		}
	}
	return maxTime
}
