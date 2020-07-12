package args

type CmdArgs struct {
	// Cmd is command type
	Cmd Command
	// Opt is a command option value
	Opt Option
}

var CmdHelp = CmdArgs{
	Cmd: Help, Opt: OptsNone,
}
