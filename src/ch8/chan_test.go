package ch8

import (
	"fmt"
	"testing"
	"time"
)

func TaskJob() {
	time.Sleep(3 * time.Second)
	fmt.Println("TaskJob Complete")

}

func TaskService() {
	fmt.Println("TaskService Start")
	time.Sleep(2 * time.Second)
	fmt.Println("TaskService End")
}

func TestNormalTask(t *testing.T) {
	TaskJob()
	TaskService()
}

func TestAsyncTask(t *testing.T) {
	result := FutureTask()

	TaskService()

	fmt.Printf("(<-result): %v\n", (<-result))
}

func FutureTask() chan bool {
	taskChan := make(chan bool)
	go func() {
		TaskJob()
		taskChan <- true
	}()
	return taskChan
}

func TestChan(t *testing.T) {
	c := make(chan bool, 1)

	fmt.Println(c)
	// <-c
	c <- true
	c <- true
	<-c
	t.Log("完成")
}
