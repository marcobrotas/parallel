package error

import (
	"github.com/io-da/command"
	"sync"
)

type Errors struct {
	sync.Mutex
	errs map[string]error
}

func NewErrors() *Errors {
	return &Errors{
		errs: make(map[string]error),
	}
}

func (hdl *Errors) Handle(cmd command.Command, err error) {
	hdl.Lock()
	hdl.errs[hdl.key(cmd)] = err
	hdl.Unlock()
}

func (hdl *Errors) Get(cmd command.Command) error {
	hdl.Lock()
	defer hdl.Unlock()
	if err, hasError := hdl.errs[hdl.key(cmd)]; hasError {
		return err
	}
	return nil
}

func (hdl *Errors) IsEmpty() bool {
	return len(hdl.errs) <= 0
}

func (hdl *Errors) GetErrorIds() []string {
	ids := make([]string, 0, len(hdl.errs))
	for id := range hdl.errs {
		ids = append(ids, id)
	}
	return ids
}

func (hdl *Errors) key(cmd command.Command) string {
	if cmd == nil {
		return "nil"
	}
	return string(cmd.ID())
}
