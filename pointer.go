package main

import "fmt"

func main() {
   var aa int = 20  
   var ip *int      
   ip = &aa  

   fmt.Printf("Address of a variable: %x\n", &aa  )

  
   fmt.Printf("Address stored in ip variable: %x\n", ip )

   fmt.Printf("Value of *ip variable: %d\n", *ip )
}