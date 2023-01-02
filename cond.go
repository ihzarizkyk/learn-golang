//! Author: Mochammad Ihza Rizky Karim

package main
import ("fmt")

func main(){

	// if

	if 10>5{
		fmt.Println("true")
	}

	// if else if

	var a int = 10

	if a > 100{
		fmt.Println("false")
	} else if a == 10 {
		fmt.Println("same")
	} else{
		fmt.Println("true")
	}
}