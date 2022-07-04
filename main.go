package main

import (
	"fmt"
	"github.com/io-da/command"
	"github.com/io-da/event"
	"log"
	"main/bus_command"
	busError "main/bus_command/error"
	"main/bus_command/handler"
	"os"
	"runtime"
	"sync"
)

var (
	// core
	cmdBus = command.NewBus()
	evtBus = event.NewBus()
	errs   = busError.NewErrors()
	wg     = &sync.WaitGroup{}
)

func initializeCommandBus() {
	cmdBus.WorkerPoolSize(16)
	cmdBus.ErrorHandlers(errs)
	cmdBus.Initialize(
		handler.NewConsole(wg, evtBus),
	)
}

func initialize() {
	initializeCommandBus()
}

func main() {
	//goland:noinspection GoBoolExpressions
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
		return
	}
	cmds := os.Args[1:]
	if len(cmds) <= 0 {
		fmt.Println("Please provide one or more commands to execute in parallel.")
		fmt.Println("If the commands have parameters or spaces, enclose them in ''.")
		return
	}
	initialize()
	msg := "Executing the following commands in parallel:"
	for i, cmd := range cmds {
		id := fmt.Sprintf("Command %d", i+1)
		c := &bus_command.Console{Id: []byte(id), Cmd: cmd}
		wg.Add(1)
		if err := cmdBus.HandleAsync(c); err != nil {
			log.Fatal(err)
		}
		msg += fmt.Sprintf("\n - %s '%s'", c.Id, c.Cmd)
	}
	log.Println(msg)
	wg.Wait()

	if !errs.IsEmpty() {
		log.Fatalf("Execution finished but the following commands returned errors: %s\n", errs.GetErrorIds())
	}

	log.Println("Done!")
}
