package ispalindrome

func Ispalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x <= 10 {
		return true
	}
	x_c := x
	y := 0
	temp := 0
	for x >= 10 {
		temp = x % 10
		y = 10*y + temp
		x = x / 10
	}
	y = 10*y + x
	if x_c == y {
		return true
	} else {
		return false
	}
}
