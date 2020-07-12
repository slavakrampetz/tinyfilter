package args

import (
	"errors"
	"fmt"
	"strings"
)

// Parse app arguments
func Parse(args []string) (cmd CmdArgs, err error) {

	nof := len(args)
	if nof == 0 {
		return CmdHelp, errors.New("command not specified")
	}

	s := strings.ToLower(args[0])
	switch s {

	case string(Help), cHelpShort:
		return CmdHelp, nil

	case string(Web):
		return CmdArgs{
			Cmd: Web, Opt: OptsNone,
		}, nil

	case string(Reload), cReloadShort:
		return CmdArgs{
			Cmd: Reload, Opt: OptsNone,
		}, nil

	case string(Youtube), cYoutube:
		return parseYoutube(args)
	}

	return CmdHelp, fmt.Errorf("unsupported command %s", s)
}

// Parse youtube command args
func parseYoutube(args []string) (cmd CmdArgs, err error) {

	cmd = CmdArgs{
		Cmd: Youtube, Opt: OptsNone,
	}

	nof := len(args)
	if nof < 2 {
		return cmd, errors.New("please specify Youtube action: on or off")
	}

	act := strings.ToLower(args[1])
	switch act {
	case ytOff:
		cmd.Opt = OptsYtOff
	case ytOn:
		cmd.Opt = OptsYtOn
	default:
		return cmd, errors.New("unknown Youtube action: " + act)
	}

	return cmd, nil
}
