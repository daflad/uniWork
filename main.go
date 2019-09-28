package main

import (
	"fmt"
	"uniWork/models"
)

func main() {
	test := 0.0
	result := models.CtoF(test)
	check := models.FtoC(result)
	fmt.Println(result)
	fmt.Println(check)
}
