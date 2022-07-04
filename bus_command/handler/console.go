package handler

import (
	"errors"
	"github.com/io-da/command"
	"github.com/io-da/event"
	"log"
	"main/bus_command"
	"os/exec"
	"regexp"
	"sync"
)

type Console struct {
	wg     *sync.WaitGroup
	evtBus *event.Bus
}

func NewConsole(
	wg *sync.WaitGroup,
	evtBus *event.Bus,
) *Console {
	return &Console{
		wg:     wg,
		evtBus: evtBus,
	}
}

func (hdl *Console) Handle(cmd command.Command) error {
	if cmd, is := cmd.(*bus_command.Console); is {
		defer hdl.wg.Done()
		stdout, err := exec.Command("/bin/sh", "-c", cmd.Cmd).Output()
		if err != nil {
			if ee, ok := err.(*exec.ExitError); ok {
				stderr := string(ee.Stderr)
				re := regexp.MustCompile("^/bin/sh:\\s*\\d+:\\s*")
				stderr = re.ReplaceAllString(stderr, "")
				log.Printf("%s error occurred\n'%s'\n%s\n\n", cmd.Id, cmd.Cmd, stderr)
				return errors.New(stderr)
			}
			log.Printf("%s error occurred\n'%s'\n%s\n\n", cmd.Id, cmd.Cmd, err)
			return err
		}
		log.Printf("%s execution successfully finished\n'%s'\n%s\n\n", cmd.Id, cmd.Cmd, stdout)
	}
	return nil
}
