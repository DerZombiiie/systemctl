package main

import (
	"fmt"
	"os"
	"os/exec"
)

func usage() {
	fmt.Printf("openrc-systemctl\nUSAGE: systemctl COMMAND UNIT...\n\nCOMMAND := { start stop restart }\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "--help" {
		usage()
	}

	switch os.Args[1] {
	case "start":
		run(os.Args[2:], "rc-service", "%s", "start")
	case "stop":
		run(os.Args[2:], "rc-service", "%s", "stop")
	case "restart":
		run(os.Args[2:], "rc-service", "%s", "restart")
	default:
		usage()
	}
}

func run(services []string, cmd string, args ...string) {
	for _, service := range services {
		for k := range args {
			if args[k] == "%s" {
				args[k] = service
			}
		}

		cmd := exec.Command(cmd, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Failed on service %s: %s\n", service, err)
			os.Exit(1)
		}
	}
}
