package main
import "fmt"

func main(){

		var (
			sum = 1.00
			prod = 1.00
			i = 1.00
			dig = 1.00
		)
	 
        
		fmt.Println("digite el numero: ")
		fmt.Scan(&dig)
 
        for i=1 ; i <= dig ; i++{
             prod = prod * i;
             sum += + (1 / prod);
         }
        fmt.Println("El numero e es: ",sum)
  
}