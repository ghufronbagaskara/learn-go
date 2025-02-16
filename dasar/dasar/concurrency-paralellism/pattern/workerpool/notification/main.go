package main

import (
	"fmt"
	"time"
)

func main()  {

	taskQueueCH := make(chan Notification, 10)
	dispatcher := Dispatcher{
		MaxWorker: 3,
		TaskQueue: taskQueueCH,
	}

	dispatcher.StartWorker()

	go func() {
		for i := 1; ; i++ {
			notificationContent := Notification{
				ID: i,
				Message: fmt.Sprintf("Notification with ID: %d", i),
			}

			dispatcher.AssignTask(notificationContent)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// worker wait forever
	select {}
	
}


// tasks channel
type Notification struct {
	ID int 
	Message string
}


func (n Notification) Send() {
	fmt.Printf("Sending message with content: %s\n", n.Message)
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Message sent")
}

// worker 
type Worker struct {
	ID int 
	TaskQueue chan Notification
}

// live worker
func (w Worker) Start() {
	go func() {
		for notification := range w.TaskQueue {
			fmt.Printf("Worker with ID: %d, is processing notification with ID: %d\n", w.ID, notification.ID)
			notification.Send()
			fmt.Printf("Worker with ID: %d, succesfully sending notification with ID: %d\n", w.ID, notification.ID)
		}
	}()
}

// dispatcher
type Dispatcher struct {
	TaskQueue chan Notification
	MaxWorker int
}


func (d *Dispatcher) StartWorker() {
	// start worker
	for i := 0; i < d.MaxWorker; i++ {
		worker := Worker{
			ID: i,
			TaskQueue: d.TaskQueue,
		}
		worker.Start()
	}
}


func (d *Dispatcher) AssignTask(notification Notification) {
	d.TaskQueue <- notification
}