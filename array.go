//! Author: Mochammad Ihza Rizky Karim

package main
import ("fmt")

func main(){

	// var name = [length]data type{value}
	var arr = [2]int{1,2}

	fmt.Println(arr)

	var arrs = [3]string{"satu","dua","tiga"}

	fmt.Println(arrs)

	// access specific data

	var abc = [4]int{1,23,4,4}

	fmt.Println(abc[2])

	// not initialized

	bca := [4]int{}

	// half initialized

	bac := [5]int{1,2,3}

	// initialized specific data

	bbc := [5]int{1:11,3:2}

	fmt.Println(bca)
	fmt.Println(bac)
	fmt.Println(bbc)

	// find length

	fmt.Println(len(bbc))
}