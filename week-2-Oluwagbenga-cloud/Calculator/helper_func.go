package Calculator

func Multiplication(arr []float64) float64 {
	var proResult = float64(1)

	for i := 0; i < len(arr); i++ {
		proResult *= arr[i]
	}
	return proResult
}

func Addition(arr []float64) float64 {
	var sum float64 = 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func Subtraction(arr []float64) float64 {
	var diffResult float64 = arr[0]
	// {1,2,4,5}
	for i := 1; i < len(arr); i++ {
		diffResult -= arr[i]
	}
	return diffResult
}

func Division(arr []float64) float64 {
	var divResult float64 = arr[0]

	for i := 1; i < len(arr); i++ {
		//divResult /= arr[i]
		divResult = divResult / arr[i]
	}
	return divResult
}
