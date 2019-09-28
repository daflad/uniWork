package main

import (
	"fmt"

	"github.com/daflad/uniWork/models"
)

func main() {
	test := 0.0
	result := models.CtoF(test)
	check := models.FtoC(result)
	fmt.Println(result)
	fmt.Println(check)
}
