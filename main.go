package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func main() {

	start := time.Now()
	totalOrders := 5
	maxPool := 2

	var wg sync.WaitGroup
	syncOrderCoffee := func() {
		orderCoffee()
		wg.Done()
	}

	// Set Options
	p, _ := ants.NewPool(maxPool)

	for i := 0; i < totalOrders; i++ {
		wg.Add(1)
		p.Submit(syncOrderCoffee)
	}

	defer p.Release()
	wg.Wait()

	fmt.Printf("Running goroutines: %d\n", ants.Running())
	fmt.Printf("Finish all tasks. \n")

	elapsed := time.Since(start)
	fmt.Printf("Took a %s seconds", elapsed)
}

func orderCoffee(){
	makePayment()
	processCoffee()
	processMilk()
	processCup()
	serveOrder()
}

func makePayment(){
	time.Sleep(time.Second*2)
	fmt.Println("===== makePayment 💰 =====") 
}

func processCoffee(){
	time.Sleep(time.Second*2)
	fmt.Println("===== processCoffee ➰ =====") 
}

func processMilk(){
	time.Sleep(time.Second*2)
	fmt.Println("===== processMilk 🥛 =====")
}

func processCup(){
	time.Sleep(time.Second*2)
	fmt.Println("===== processCup 🥃 =====") 
}

func serveOrder(){
	time.Sleep(time.Second*2)
	fmt.Println("===== serveOrder 🖐 =====") 
}