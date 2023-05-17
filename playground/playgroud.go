package playgroud
import (
	"fmt"
	"example/hello"
	"example.com/greetings"

)
	


func Playgroud() {

	arr := []int{7, 2, 8, -9, 4, 0}
	bubbleSort(arr)
	fmt.Printf("arr:: %v\n", arr)
	//
  	message , err := greetings.Hello("fdsa")		
	if (err == nil) {
		fmt.Printf("%v",message)
	}
    hello.Demo()			

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	fmt.Printf("start go1\n")
	go sum(s[:len(s)/2], c)
	fmt.Printf("end go1\n")
	fmt.Printf("start go2\n")
	go sum(s[len(s)/2:], c)
	fmt.Printf("end go2\n")
	fmt.Printf("START X\n")
	x := <-c // receive from c
	fmt.Printf("end x\n")
	close(c)

	fmt.Printf("START Y\n")
	y := <-c
	fmt.Printf("END Y\n")
	fmt.Println(x, y, x+y)
}
func assignDemo() {	

	num := 1
	var num2 int = 1
	var num3 int
	num3 = 1
	fmt.Printf("num:%v\n%v\n%v\n", num,num2,num3)
	
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Printf("before write:: %v\n", s)
	c <- sum // send sum to c

	// fmt.Printf("after write:: %v\n", s)
}

// 冒泡算法
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
	