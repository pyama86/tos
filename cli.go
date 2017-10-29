package main

import (
	"flag"
	"fmt"
	"io"
	"log/syslog"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

func init() {
	logger, err := syslog.New(syslog.LOG_NOTICE|syslog.LOG_USER, "tos")
	if err == nil {
		logrus.SetOutput(logger)
	}
}

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

var (
	version   string
	revision  string
	goversion string
	builddate string
	builduser string
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		ok       string
		warning  string
		critical string
		unknown  string
		no       string

		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&ok, "ok", "", "OK")
	flags.StringVar(&ok, "o", "", "OK(Short)")

	flags.StringVar(&warning, "warning", "", "WARNING")
	flags.StringVar(&warning, "w", "", "WARNING(Short)")

	flags.StringVar(&critical, "critical", "", "CRITICAL")
	flags.StringVar(&critical, "c", "", "CRITICAL(Short)")

	flags.StringVar(&unknown, "unknown", "", "Unknown")
	flags.StringVar(&unknown, "u", "", "Unknown(Short)")

	flags.StringVar(&no, "not_ok", "", "Not OK")
	flags.StringVar(&no, "no", "", "NotOK(Short)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		logrus.Error(err)
		return ExitCodeError
	}

	// Show version
	if version {
		printVersion()
		return ExitCodeOK
	}

	cmd := ""
	if no != "" && os.Getenv("MACKEREL_STATUS") != "" && os.Getenv("MACKEREL_STATUS") != "OK" {
		cmd = no
	} else {
		switch os.Getenv("MACKEREL_STATUS") {
		case "OK":
			cmd = ok
		case "WARNING":
			cmd = warning
		case "CRITICAL":
			cmd = critical
		case "UNKNOWN":
			cmd = unknown
		}
	}

	if cmd == "" {
		return ExitCodeOK
	}

	var c *exec.Cmd
	cmds := strings.Split(cmd, " ")
	if len(cmds) > 1 {
		c = exec.Command(cmds[0], cmds[1:]...)
	} else {
		c = exec.Command(cmds[0])
	}
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		logrus.Error(err)
		return ExitCodeError
	}

	return ExitCodeOK
}

func printVersion() {
	fmt.Printf("tos version: %s (%s)\n", version, revision)
	fmt.Printf("build at %s (with %s) by %s\n", builddate, goversion, builduser)
}
