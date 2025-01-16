package main

import "fmt"

func main(){
	s := make([]int, 5,6)
	fmt.Println(len(s), cap(s))
	for i:= 0; i < len(s); i++ {
		s[i] = -1
	}
	fmt.Println(len(s), cap(s))
	s = append(s, -2,-3)
	fmt.Println(len(s), cap(s))
}
