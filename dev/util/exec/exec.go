package exec

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var (
	shell       string
	shellArgs   []string
	shellQuotes bool
	shellQuote  string
)

func init() {

	switch runtime.GOOS {
	default:
		panic("Cannot find shell interpreter for " + runtime.GOOS)

	case "freebsd", "netbsd", "dragonfly", "openbsd", "linux":
		shell = "bash"
		shellArgs = append(shellArgs, "-c")
		shellQuotes = true
		shellQuote = "\\\""

	case "windows":
		shell = "cmd.exe"
		shellArgs = append(shellArgs, "/C")
		shellQuotes = false
	}

}

func PrepareExecArgs(cmd string, args ...string) (string, []string) {

	cmdArgs := make([]string, 0, len(shellArgs)+1)
	cmdArgs = append(cmdArgs, shellArgs...)
	cmdText := cmd + " " + strings.Join(args, " ")
	if shellQuotes {
		if strings.Index(cmdText, "\"") >= 0 {
			cmdText = strings.ReplaceAll(cmdText, "\"", shellQuote)
		}
		// cmdText = "\"" + cmdText + "\""
	}
	cmdArgs = append(cmdArgs, cmdText)
	return shell, cmdArgs
}

func Exec(isShell bool, cmd string, args ...string) (err error, out string) {

	cmdName := cmd
	cmdArgs := args
	if isShell {
		cmdName, cmdArgs = PrepareExecArgs(cmd, args...)
	}

	// Command
	command := exec.Command(cmdName, cmdArgs...)
	command.Env = os.Environ()

	var bOut bytes.Buffer
	var bErr bytes.Buffer
	command.Stdout = &bOut
	command.Stderr = &bErr

	err = command.Run()
	out = bOut.String()
	if err != nil {
		return err, out
	}
	if command.ProcessState != nil && !command.ProcessState.Success() {
		return errors.New("command " + cmd + " process state is not a success"), out
	}
	return nil, out
}
