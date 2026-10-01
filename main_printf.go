package main

import "fmt"

func main() {
	//phi * jari * jari
	var phi float64 = 3.14
	jari := 12
	luas := phi * float64(jari) * float64(jari)

	fmt.Println("Luas Println: ", luas)
	fmt.Println("==================")
	fmt.Printf("Luas: %.2f\n", luas)
}
