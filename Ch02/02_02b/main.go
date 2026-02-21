package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	num := 42
	fmt.Println(str1, str2, str3)
	stringLength, err := fmt.Println("This is the number", num)
	if err == nil {
		fmt.Println("String Length is ", stringLength)
	}
	//formatted Printing
	fmt.Printf("the number formated as variable is : %v\n", num)
	//Formating with digit place holder
	fmt.Printf("the number formated as digit is : %d\n", num)
	//printing data type of variables
	fmt.Printf("This is the Data Type for the number: %T\n", num)
	fmt.Printf("This is the Data Type for the text: %T\n", str1)
}
