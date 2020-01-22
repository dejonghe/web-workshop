package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	logger.Printf("Starting up! Press CTRL-C to stop!")

    Waiter(logger)

	// 0 means no errors
	os.Exit(0)
}

func Waiter(logger *log.Logger) {
	// make a quit channel for operating system signals, buffered to size 1
	quit := make(chan os.Signal, 1)

	// listen for all interrupt signals, send them to quit channel
	signal.Notify(quit,
		os.Interrupt,    // interrupt = SIGINT = Ctrl+C
		syscall.SIGQUIT, // Ctrl-\
		syscall.SIGTERM, // "the normal way to politely ask a program to terminate"
	)
	logger.Printf("Just going to wait here until you press control-C")
	// block, waiting for receive on quit channel
	sig := <-quit
	logger.Printf("Shutting down after receiving %v signal!", sig)
}
