package service

import (
	"context"
	"github.com/judwhite/go-svc"
	"github.com/khorevaa/logos"
	"os"
	"sync"
	"syscall"
)

var log = logos.New("github.com/khorevaa/ras-service/service").Sugar()

// Control implements service.Control
type Control struct {
	Starter *RasStarter

	wg   sync.WaitGroup
	quit chan struct{}
}

func (p *Control) Run(ctx context.Context) error {

	err := svc.Run(p, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)
	if err != nil {
		log.Fatalf("Control error", err)
	}

	return err
}

// func main() {
// 	prg := &Control{}
//
// 	// Call service.Run to start your Control/service.
// 	if err := service.Run(prg); err != nil {
// 		log.Fatal(err)
// 	}
// }

func (p *Control) Init(env svc.Environment) error {
	log.Infof("is win service? %v\n", env.IsWindowsService())
	return nil
}

func (p *Control) Start() error {
	// The Start method must not block, or Windows may assume your service failed
	// to start. Launch a Goroutine here to do something interesting/blocking.

	p.quit = make(chan struct{})

	p.wg.Add(1)
	go func() {
		log.Info("Starting...")
		err := p.Starter.Start()
		if err != nil {
			log.Error(err.Error())
		}

		<-p.quit
		p.wg.Done()
	}()

	return nil
}

func (p *Control) Stop() error {
	// The Stop method is invoked by stopping the Windows service, or by pressing Ctrl+C on the console.
	// This method may block, but it's a good idea to finish quickly or your process may be killed by
	// Windows during a shutdown/reboot. As a general rule you shouldn't rely on graceful shutdown.

	log.Info("Stopping...")
	close(p.quit)
	err := p.Starter.Stop()
	if err != nil {
		log.Error(err.Error())
	}
	p.wg.Wait()
	log.Info("Stopped.")
	return nil
}
