package mostwater

// MaxArea is a function that to judge most water container
/* func MaxArea(height []int) int {
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			// area = vh
			v := height[j]
			if height[i] < height[j] {
				v = height[i]
			}
			temp := (j - i) * v
			if temp > max {
				max = temp
			}
		}
	}
	return max
}
*/
// gready algorithm
func MaxArea(height []int) int {
	max := 0
	i := 0
	j := len(height) - 1
	for i < j {
		v := height[j]
		if height[i] < height[j] {
			v = height[i]
		}
		temp := (j - i) * v
		if temp > max {
			max = temp
		}
		for height[i] <= v && i < j {
			i++
		}
		for height[j] <= v && i < j {
			j--
		}
	}
	return max
}
