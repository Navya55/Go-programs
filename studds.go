package main
import "fmt"
type student struct {
	Rollno string
	marks int
}
func main(){
	fmt.Println(student{"CS101",85})
	fmt.Println(student{Rollno:"CS102",marks:92})
	fmt.Println(student{Rollno:"CS103"})
	fmt.Println(&student{Rollno:"CS104",marks:90})
	s:=student{Rollno:"CS105",marks:89}
	fmt.Println(s.Rollno)
	sp:=&s 
	fmt.Println(sp.marks)
	sp.marks=89
	fmt.Println(sp.marks)
}