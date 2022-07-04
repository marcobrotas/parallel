package bus_command

type Console struct {
	Id  []byte
	Cmd string
}

func (cmd *Console) ID() []byte {
	return cmd.Id
}
