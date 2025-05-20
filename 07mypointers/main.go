package main
import "fmt"


func main(){
	var ptr *int ; 
	fmt.Println("here is pointer val", ptr)
	
	myNumber :=  23;

	var ptr2 = &myNumber

	fmt.Println("Pointer is :", ptr2)
	fmt.Println("Value is :", *ptr2)

	*ptr2 = *ptr2 + 2 
	fmt.Println(" Trick :", myNumber)


}
