package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// cached on Redis
var (
	TopDriverIDs []uuid.UUID
	ReliableDriverIDs []uuid.UUID
	NormalDriverIDs []uuid.UUID
)
func main() {
	fmt.Println("program start...")
	FindNearbyDriverIDs()
	
	fmt.Println("program finished...")
	
}

func FindNearbyDriverIDs() {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	doneCH := make(chan struct{}, 3)
	
	// 	findTopDriverIDs - optional
	go func () {
		findTopDriverIDs(ctx)
		doneCH <- struct{}{}
		}()
		
	// findRelibleDriverIDs - optional
	go func () {
		findRelibleDriverIDs(ctx)
		doneCH <- struct{}{}
	}()

	// findNormalDriverIDs - mandatory
	go func () {
		findNormalDriverIDs(ctx)
		doneCH <- struct{}{}
	}() 

	for i := 0; i < 3; i++ {
		<- doneCH
	}

	fmt.Println("all goroutine finished")

}

func findTopDriverIDs(ctx context.Context) (driverIDs []uuid.UUID) {
	// simulate to have a long running query
	select {
	case <- time.After(1500 * time.Millisecond) :
		fmt.Printf("top driver found")
		return []uuid.UUID{uuid.New(), uuid.New()}
		
	case <- ctx.Done():
		fmt.Println("find top driver cancelled")
		return []uuid.UUID{} 
	}

}


func findRelibleDriverIDs(ctx context.Context) (driverIDs []uuid.UUID) {
	// simulate to have a long running query
	select {
	case <- time.After(300 * time.Millisecond) :
		fmt.Printf("reliable driver found")
		return []uuid.UUID{uuid.New(), uuid.New()}
		
	case <- ctx.Done():
		fmt.Println("find reliable driver cancelled")
		return []uuid.UUID{} 
	}
}


func findNormalDriverIDs(ctx context.Context) (driverIDs []uuid.UUID) {
	fmt.Println("normal drivers found") 
	return []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}
}