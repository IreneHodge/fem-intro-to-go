package main

import "fmt"

func main() {
   var name string = "Beyonce"
   var name1, name2 = "Irene", "Ann"
   //can't happen outside functions
   name3 := "John"
   fmt.Println(name)
   fmt.Println(name1, name2)
   fmt.Println(name3)
}