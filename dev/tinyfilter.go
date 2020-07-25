package main

import (
	"errors"
	"os"
	"tinyfilter/dev/app"
	"tinyfilter/dev/app/args"
	"tinyfilter/dev/command/reload"
	"tinyfilter/dev/command/youtube"
	"tinyfilter/dev/log"
)

const (
	version = "1.0.0"
)

func main() {

	banner()

	cmd, err := args.Parse(os.Args[1:])
	if err != nil {
		log.Err("Error parsing arguments>", err)
		if cmd.Cmd == args.Help {
			log.Eol()
			usage()
		}
		os.Exit(1)
	}

	err = runCommand(cmd)
	if err != nil {
		log.ErrIfErr(err, "Error executing command:", cmd.Cmd)
		os.Exit(2)
	}
	os.Exit(0)
}

func runCommand(cmd args.CmdArgs) error {

	switch cmd.Cmd {

	case args.Reload:
		return reload.Exec()

	case args.Youtube:
		return youtube.Exec(cmd.Opt == args.OptsYtOn)

	case args.Web:
		return app.RunWeb()

	case args.Help:
		usage()
		return nil
	}

	return errors.New("Command " + string(cmd.Cmd) + " not implemented yet")
}

func banner() {
	text := "tinyfilter v." + version + ", app for manage tinyproxy configuration"
	log.Out(text)
}

func usage() {
	usage := `Usage: tinyfilter <command> [<options>]

Commands:
  help,h       Show this screen and exit
  web          Start web-server for listen remote commands
  reload,r     Reload tinyproxy
  youtube,yt allow|deny
               Allow or deny youtube links by proxy`
	log.Out(usage)
}
