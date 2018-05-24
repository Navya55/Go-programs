package main
import "fmt"
func main(){
	var A [5]int 
	var lsum,rsum,t,n int
	fmt.Println("Enter the number of test cases")
	fmt.Scanf("%d",&t)
	fmt.Scanf("%d",&t)
	for i:=0;i<t;i++ {
		fmt.Println("Enter the number of elements")
		fmt.Scanf("%d",&n)
		fmt.Scanf("%d",&n)
		for j:=0;j<n;j++ {
			fmt.Println("Enter the element")
			fmt.Scanf("%d",&A[j])
			fmt.Scanf("%d",&A[j])
		}//fmt.Scanf("%d",&n)

		
		check:=false
		for i:=0;i<len(A);i++{
			lsum=0
			rsum=0
		
			for j:=i+1;j<len(A);j++{
				rsum+=A[j]
			}
			for k:=0;k<i;k++{
				lsum+=A[k]
			}
		
			if(lsum==rsum && rsum!=0){
				check=true
			}
			fmt.Println("rsum",rsum)
			fmt.Println("lsum",lsum)	
			
		
		}
	

		if(check==true){
			fmt.Println("Yes")
		} else{
			fmt.Println("No")
			}
	}
}