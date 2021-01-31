package server

import (
	"log"
	"os/exec"
)

// Exector specify exector function
type Exector interface {
	Exec(string, string) error
}

// tc only impl qdisc.delay
type tc struct {
	cmd          string
	netInterface string
}

// NewExector return exector
func NewExector(ni string) Exector {
	return &tc{cmd: "tc", netInterface: ni}
}

// define actions
const (
	Add    = "add"
	Change = "change"
	Del    = "del"
)

// Exec action: add, change, del
func (tc *tc) Exec(action, t string) error {
	cmd := exec.Command(tc.cmd, "qdisc", action, "dev",
		tc.netInterface, "root", "netem", "delay", t)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	log.Printf("exec result: %s\n", out)
	return nil
}
