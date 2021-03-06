package main

// import "fmt"

// Repeat func repeats a string for 5 times.
func Repeat(s string) string {

	t := ""
	for i := 0; i < 5; i++ {
		t += s
	}
	return t
}

// func main(){
// 	fmt.Println(Repeat("a"))
// }
