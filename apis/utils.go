package apis

func swap(x, y *uint) {
	if *x < *y {
		*x, *y = *y, *x
	}
}

func gcd(x, y uint) uint {
	swap(&x, &y)
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	}
	return y
}

// GetGCD : get greatest common divisor
func GetGCD(allWeight *[]uint, count uint) uint {
	if count == 1 {
		return (*allWeight)[0]
	}
	return gcd((*allWeight)[count-1], GetGCD(allWeight, count-1))
}
