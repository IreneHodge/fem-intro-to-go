package main

//can name return values
func printAge(age1, age2 int) (ageOfSally, ageOfBob int) {
	ageOfSally = age1
	ageOfBob = age2
	return
}

func main() {
		x, y := printAge(10, 21)
	// 	fmt.Println(x)
	// 	fmt.Println(y)
}
