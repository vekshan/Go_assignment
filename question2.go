package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	NumRoutines = 3
	NumRequests = 1000
)
// global semaphore monitoring the number of routines
var semRout = make(chan int, NumRoutines)
// global semaphore monitoring console
var semDisp = make(chan int, 1)
// Waitgroups to ensure that main does not exit until all done
var wgRout sync.WaitGroup
var wgDisp sync.WaitGroup

type Task struct {
	a, b float32
	disp chan float32
}

func solve(t *Task){
	timer := (rand.Float64() * 14 ) + 1
	time.Sleep( time.Duration(timer) * time.Second)
	res  := t.a + t.b
	t.disp <- res

}

func handleReq(t *Task){
	solve(t)
}

func ComputeServer()(chan *Task){

	reqChan := make(chan *Task, NumRequests)

		go func() {

			for elem:=range reqChan {

				wgRout.Add(1)

				semRout <- 1

				handleReq(elem)

				<-semRout


			}
		}()





	return reqChan
}

func DisplayServer() (chan float32){

	dispChan := make(chan float32, NumRequests)

		go func() {

			for elem:= range dispChan{


			wgDisp.Add(1)

			fmt.Println("\nResult: ", elem)
			fmt.Println("-------")

			wgRout.Done()
			wgDisp.Done()

			}

		}()

	return dispChan


}

func main() {
	dispChan := DisplayServer()
	reqChan := ComputeServer()

	for {
		var a, b float32
		// make sure to use semDisp
		// …


		semDisp <- 1
		fmt.Print("Enter two numbers: ")
		fmt.Scanf("%f %f \n", &a, &b)
		fmt.Printf("%f %f \n", a, b)
		<- semDisp




		if a == 0 && b == 0 {
			break
		}
		// Create task and send to ComputeServer
		// …
		task := &Task{a, b, dispChan}
		reqChan <- task

		time.Sleep( 1e9 )
	}
	// Don’t exit until all is done
	wgDisp.Wait()
	close(reqChan)
	wgRout.Wait()



}

