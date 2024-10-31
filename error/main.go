package main

import "fmt"

func divide(a , b float64) (float64, error){
	if b == 0 {
		return 0, fmt.Errorf("denomination not zero")
	}
	return a/b, nil
}

func main() {

	fmt.Println("Dwaipayan");

	// ans,err := divide(10,8)
	
	// if err != nil{
	// 	fmt.Println("Error Handling")
	// }
	// // '_' 
	// fmt.Println("Division of two numbers is: ",ans)


	
	ans,_ := divide(10,8)

	fmt.Println("Division of two numbers is: ",ans)

}