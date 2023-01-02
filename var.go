//! Author: Mochammad Ihza Rizky Karim

package main
import ("fmt")

func main(){

	// define var
	// var name data type
	var a int = 12
	var x string = "hello string"
	var ok = 123 // inferred type (Not Define Data Type Variable)

	fmt.Println(a)
	fmt.Println(x)
	fmt.Println(ok)

	// multiple var

	var b,d,e,f int = 1,2,3,4 

	fmt.Println(d)
	fmt.Println(b)
	fmt.Print(e,("\n")) // new line
	fmt.Print(f,("\n")) // new line


}