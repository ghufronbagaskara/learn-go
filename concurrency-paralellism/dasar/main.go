package main

import (
	"fmt"
	"sync"
)


func main(){
	var wg sync.WaitGroup	

	wg.Add(1)
	go doPrint(1, &wg)
	
	
	wg.Add(1)
	go doPrint(2, &wg)

	wg.Wait()
	fmt.Println("Done")

}

func doPrint(index int, wg *sync.WaitGroup){

	for i := 0; i <= 5; i++ {
		fmt.Println(index, ":", i)
	}

	wg.Done()



}