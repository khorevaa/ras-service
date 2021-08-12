package main

import (
	"golang.org/x/sys/windows/svc"

	"os"
	"sync"
	"time"

	"github.com/khorevaa/logos"
	"github.com/urfave/cli/v2"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

var log = logos.New("github.com/v8platform/onec-util")

func main() {

	app := &cli.App{
		Name:     "onec-util",
		Version:  version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name: "Aleksey Khorev",
			},
		},
		Usage:     "Command line utilities for server 1C.Enterprise",
		UsageText: "onec-util command [command options] [arguments...]",
		Copyright: "(c) 2021 Khorevaa",
		//Description: "Command line utilities for server 1S.Enterprise",
	}

	for _, command := range cmd.Commands {
		app.Commands = append(app.Commands, command.Cmd())
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// program implements svc.Service
type program struct {
	wg   sync.WaitGroup
	quit chan struct{}
}

// func main() {
// 	prg := &program{}
//
// 	// Call svc.Run to start your program/service.
// 	if err := svc.Run(prg); err != nil {
// 		log.Fatal(err)
// 	}
// }

func (p *program) Init(env svc.Environment) error {
	log.Printf("is win service? %v\n", env.IsWindowsService())
	return nil
}

func (p *program) Start() error {
	// The Start method must not block, or Windows may assume your service failed
	// to start. Launch a Goroutine here to do something interesting/blocking.

	p.quit = make(chan struct{})

	p.wg.Add(1)
	go func() {
		log.Println("Starting...")
		<-p.quit
		log.Println("Quit signal received...")
		p.wg.Done()
	}()

	return nil
}

func (p *program) Stop() error {
	// The Stop method is invoked by stopping the Windows service, or by pressing Ctrl+C on the console.
	// This method may block, but it's a good idea to finish quickly or your process may be killed by
	// Windows during a shutdown/reboot. As a general rule you shouldn't rely on graceful shutdown.

	log.Println("Stopping...")
	close(p.quit)
	p.wg.Wait()
	log.Println("Stopped.")
	return nil
}
