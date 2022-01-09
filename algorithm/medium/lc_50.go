package medium

// 50. Pow(x, n)

func myPow(x float64, n int) float64 {
	var result float64
	if n < 0 {
		result = 1 / helper(x, -n)
	} else {
		result = helper(x, n)
	}

	// result, _ = strconv.ParseFloat(fmt.Sprintf("%.5f", result), 64)
	return result
}

func helper(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	result := helper(x, n/2)
	if n%2 == 1 {
		return result * result * x
	}
	return result * result
}

func TestLC50() {
	println(myPow(2, 10) == 1024)
	println(myPow(2.1, 3))        //  == 9.261
	println(myPow(2, -2) == 0.25) //  == 0.25
	println(myPow(-2, 2) == 4)
	println(myPow(-2, 3) == -8)
	println(myPow(2, -3) == 0.125) // == 0.125
	println(myPow(3, -1))
}

/*
x^7 = x^3 * x^3 * x^1
     = x^1
*/
