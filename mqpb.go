package main
import "fmt"
func main(){
	A:= [5][5] string{{"a","b","c","d","e"},
		        {"f","g","h","i","j"},
		        {"k","l","m","n","o"},
		        {"p","q","r","s","t"},
	           {"u","v","w","x","y"}}
	var i,j,c,d,a,b int
	var moves int
	fmt.Println("Enter the element ")
	fmt.Scanf("%d",&i)
	fmt.Println("Enter the position of element")
	fmt.Scanf("%d%d",&c&d)
	fmt.Println("Enter the element which is blocked")
	fmt.Scanf("%d",&j)
	fmt.Println("Enter the position of element")
	fmt.Scanf("%d%d",&a&b)
	moves(c,d,a,b)
	}
		
		
func moves(c,d,a,b){
	if((c==0&&d==0)||(c==4&&d==0)||(c==0&&d==4)||(c==4&&d==4)){
		moves=3
		}else if((c==1&&d==0)||(c==2&&d==0)||(c==3&&d==0)||(c==1&&d==4)||(c==2&&d==4)||(c==3&&d==4)){
		moves=5
	} else{
		moves=8
	}
	var s1,s2 int=0,0
	s1=c-a
	s2=d-b
	if((s1==0&&s2==1)||(s1==1&&s2==0)
	{
		moves=moves - 1
	}
	fmt.Println("moves = %d",moves)
	}
	
	
	

