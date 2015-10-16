package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("concurrency exercise 1")
	fmt.Println("Let's go for a walk")
	msgChannel := make(chan string)

	bobOut := startRoutine(msgChannel, "Bob")
	aliceOut := startRoutine(msgChannel, "Alice")
	for i := 0; i < 4; i++ {
		fmt.Println(<-msgChannel)
	}
	fmt.Println("Arming alarm.")
	bobOut <- "carry on"
	go startAlarm(msgChannel)
	aliceOut <- "carry on"
	for i := 0; i < 5; i++ {
		fmt.Println(<-msgChannel)
	}
	fmt.Println("Exiting and looking the door")
	fmt.Println(<-msgChannel)
}

func startRoutine(msgChannel chan string, name string) chan string {
	out := make(chan string)
	readySpentTime := random(60, 90)
	shoesSpentTime := random(35, 45)
	go func() {
		msgChannel <- fmt.Sprintf("%s started to get ready", name)
		time.Sleep(time.Second * time.Duration(readySpentTime))
		msgChannel <- fmt.Sprintf("%s Spent %d seconds getting ready", name, readySpentTime)
		<-out
		msgChannel <- fmt.Sprintf("%s started putting on shoes", name)
		time.Sleep(time.Second * time.Duration(shoesSpentTime))
		msgChannel <- fmt.Sprintf("%s Spent %d seconds puttin on shoes", name, shoesSpentTime)
	}()
	return out
}

func startAlarm(msgChannel chan string) {
	msgChannel <- "Alarm is counting down"
	time.Sleep(time.Second * time.Duration(random(60, 90)))
	msgChannel <- "Alarm is armed"
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
