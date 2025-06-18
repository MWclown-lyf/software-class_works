package main
import "fmt"
import ("math/rand"
		"time"
)

func bubble_sort(x []int){
	n :=len(x)
	for i :=0; i<n-1; i++{
		for j :=0; j<n-i-1; j++{
			x[j], x[j+1] = x[j+1], x[j]
		}
	}
}

func main(){
	rand.Seed(time.Now().UnixNano())
	arraylength :=10
	randomArray :=make([]int,arraylength)
	for i :=0;i<arraylength;i++{
		randomArray[i]=rand.Intn(100)
	}
	fmt.Println("before",randomArray)
	bubble_sort(randomArray)
	fmt.Println("after",randomArray)

}