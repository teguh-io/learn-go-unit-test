package main

import (
	"fmt"
	"sesi11/helpers"
)

func main() {
	petarung := []int{2, 6, 3, 9}
	kekuatan := []int{2, 2, 3, 5}

	res := helpers.FinalPower(petarung, kekuatan, 2)
	fmt.Println(res)
}
