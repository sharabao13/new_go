package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)
}
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
	1, s[0],s[5] =s[5],s[0]  //i=1,j=4
    2. s[1],s[4] =s[4],s[1]  //i=2,j=3
	3. s[2],s[3] =s[3],s[2]  //i=3,j=2
*/
