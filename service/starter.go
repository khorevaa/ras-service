package service

import (
	"github.com/ericcornelissen/stringsx"
	"github.com/khorevaa/logos"
	"golang.org/x/text/encoding/charmap"
	"os"
	"os/exec"
)

type RasStarter struct {
	cmd *exec.Cmd
}

func NewRasStarter(cmd string, args []string) *RasStarter {
	return &RasStarter{
		cmd: exec.Command(cmd, args...),
	}
}

func (s *RasStarter) Start() error {

	s.cmd.Stdout = LogWriter{}
	s.cmd.Stderr = ErrorLogWriter{}

	log.Debug("Exec command", logos.String("cmd", s.cmd.Path))

	err := s.cmd.Start()
	if err != nil {
		return err
	}

	return nil
}

func (s *RasStarter) Stop() error {

	err := s.cmd.Process.Signal(os.Kill)
	if err != nil {
		return err
	}

	return nil
}

type ErrorLogWriter struct{}

func (lw ErrorLogWriter) Write(p []byte) (n int, err error) {

	p, _ = decodeOutBytes(p)

	log.Error(string(p), logos.String("prosess", "ras"))
	return len(p), nil
}

type LogWriter struct{}

func (lw LogWriter) Write(p []byte) (n int, err error) {

	p, _ = decodeOutBytes(p)

	log.Info(string(p), logos.String("prosess", "ras"))
	return len(p), nil
}

func decodeOutBytes(in []byte) ([]byte, error) {

	if stringsx.IsValidUTF8(string(in)) {
		return in, nil
	}

	dec := charmap.CodePage866.NewDecoder()
	out, err := dec.Bytes(in)

	return out, err
}
