package main

//last float64 is the return
func add(x float64, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	//Static typing, there is no dynamic type
	// var num1 float64 = 10.5
	// var num2 float64 = 11.5
	// fmt.Println(add(num1, num2))

	//String
	// w1, w2 := "Hey", "there"
	// fmt.Println(multiple(w1, w2))

	var a int = 62
	var b float64 = float64(a)
}
