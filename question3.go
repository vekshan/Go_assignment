package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

var semPop = make(chan int, 1)
var semPush = make(chan int, 1)

type Stack struct {
	triangles []Triangle
}

func newStack() *Stack{
	var triangles []Triangle

	return &Stack{triangles:triangles}
}

func pop(stack *Stack) Triangle{

	if isEmpty(stack){
		return Triangle{}
	}

	triangle := peek(stack)

	semPop<-1
	stack.triangles = stack.triangles[:len(stack.triangles)-1]
	<-semPop

	return triangle

}
func peek(stack *Stack) Triangle{

	if isEmpty(stack){
		return Triangle{}
	}

	top := len(stack.triangles) - 1


	return stack.triangles[top]

}

func push(stack *Stack, triangle Triangle){
	semPush <- 1
	stack.triangles = append(stack.triangles, triangle)
	<-semPush
}

func size(stack *Stack) int{
	return len(stack.triangles)
}

func isEmpty(stack *Stack) bool{
	if len(stack.triangles) == 0 {
		return true
	}else{
		return false
	}
}

type Point struct {
	x float64
	y float64
}
type Triangle struct {
	A Point
	B Point
	C Point
}


func triangles10000() (result [10000]Triangle) {
	rand.Seed(2120)
	for i := 0; i < 10000; i++ {
		result[i].A= Point{rand.Float64()*100.,rand.Float64()*100.}
		result[i].B= Point{rand.Float64()*100.,rand.Float64()*100.}
		result[i].C= Point{rand.Float64()*100.,rand.Float64()*100.}
	}
	return
}

func (t Triangle) Perimeter() float64 {
	return math.Sqrt(math.Pow(t.A.x - t.B.x ,2) + math.Pow(t.A.y - t.B.y,2)) + math.Sqrt(math.Pow(t.A.x - t.C.x ,2) + math.Pow(t.A.y - t.C.y,2)) +  math.Sqrt(math.Pow(t.B.x - t.C.x ,2) + math.Pow(t.B.y - t.C.y,2) )

}

func (t Triangle) Area() float64 {
	return 0.5 * math.Abs((t.B.x - t.A.x ) * (t.C.y - t.A.y ) - (t.C.x - t.A.x ) * (t.B.y - t.A.y ))
}

func classifyTriangles(highRatio *Stack, lowRatio *Stack, ratioThreshold float64, triangles []Triangle){

	for _,t := range triangles{
		if t.Perimeter()/t.Area() > ratioThreshold {
			push(highRatio, t)
		} else{
			push(lowRatio, t)
		}
	}
}


func main() {

	t := Triangle{Point{2.0,3.0},Point{3.0,4.0},Point{5.0,4.0}}
	fmt.Println(t.Area())
	fmt.Println(t.Perimeter())

	triangles:= triangles10000()


	start := 0
	end := 1000

	highRatio := newStack()
	lowRatio := newStack()

	var wg sync.WaitGroup

	wg.Add(10)
	for i:=0; i < 10; i++{

		sl := triangles[start:end]
		go func(){
			classifyTriangles(highRatio,lowRatio,1,sl)
			wg.Done()
		}()


		start += 1000
		end += 1000
	}

	wg.Wait()


	fmt.Println(peek(lowRatio))
	fmt.Println(peek(highRatio))
	fmt.Println(size(lowRatio)+size(highRatio))
	fmt.Println(pop(lowRatio))
	fmt.Println(size(lowRatio)+size(highRatio))



}
