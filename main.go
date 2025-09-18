package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var counter int

type ContextData struct {
	Ctx    context.Context
	Cancel context.CancelFunc
}

var contextHolder []ContextData

const (
	SHOW_COUNTER         = 1
	ADD_NEW_GOROUTINE    = 2
	REMOVE_GOROUTINE     = 3
	CLEAR_ALL_GOROUTINES = 4
)

func main() {
	var allContext ContextData
	allContext.Ctx, allContext.Cancel = context.WithCancel(context.Background())

	contextHolder = append(contextHolder, allContext)

	startCounting(allContext)
	handleMenu(allContext)
}

func startCounting(ctx ContextData) {
	go func() {
		for {
			select {
			case <-ctx.Ctx.Done():
				fmt.Print("\nCancelled...")
				return
			default:
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func handleMenu(ctx ContextData) {
	choice := getUserChoice()

	switch choice {
	case SHOW_COUNTER:
		showCounter(ctx)
	case ADD_NEW_GOROUTINE:
		addNewGoroutine()
	case REMOVE_GOROUTINE:
		removeGoroutine(ctx)
	case CLEAR_ALL_GOROUTINES:
		clearAllGoroutines(ctx)
	default:
		handleMenu(ctx)
	}
}

func getUserChoice() int {
	var number int
	fmt.Print("\nEnter number: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		number, _ = strconv.Atoi(text)
	}

	return number
}

func showCounter(ctx ContextData) {
	fmt.Println("Counter:", counter)
	handleMenu(ctx)
}

func removeGoroutine(ctx ContextData) {
	var number int

	for i, d := range contextHolder {
		fmt.Printf("%d) %v\n", i, d)
	}

	fmt.Print("\nEnter number to remove: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		number, _ = strconv.Atoi(text)
	}

	contextHolder[number].Cancel()
	contextHolder = append(contextHolder[:number], contextHolder[number+1:]...)

	handleMenu(ctx)
}

func addNewGoroutine() {
	var allContext ContextData
	allContext.Ctx, allContext.Cancel = context.WithCancel(context.Background())

	contextHolder = append(contextHolder, allContext)

	fmt.Println("Len of contextHolder =>", len(contextHolder))

	startCounting(allContext)
	handleMenu(allContext)
}

func clearAllGoroutines(ctx ContextData) {
	log.Println("Count of contexts =>", len(contextHolder))

	for _, c := range contextHolder {
		c.Cancel()
	}

	contextHolder = nil
	handleMenu(ctx)
}