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

func GetMaxSales(sales []int, days int) int {
	current := 0

	for i := range days {
		current += sales[i]
	}

	maxSales := current

	for i := 1; i < len(sales)-days; i++ {
		current = current - sales[i-1] + sales[i+days-1]

		if current > maxSales {
			maxSales = current
		}
	}
	return maxSales
}
