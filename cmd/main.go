package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/oc-robot/pingcap-interview/server"
)

type options struct {
	port int
}

func (o *options) Validate() error {
	return nil
}

func gatherOptions(fs *flag.FlagSet, args ...string) options {
	var o options
	fs.IntVar(&o.port, "port", 2332, "Port to serve")
	fs.Parse(args)
	return o
}

func main() {
	o := gatherOptions(flag.NewFlagSet(os.Args[0], flag.ExitOnError), os.Args[1:]...)
	if err := o.Validate(); err != nil {
		log.Fatalf("options invalid, Err: %+v", err)
	}

	exector := server.NewExector("eth0")
	if err := exector.Exec(server.Add, "0ms"); err != nil {
		log.Fatalf("create tc.qdisc failed, Err: %+v", err)
	}
	defer func() {
		if err := exector.Exec(server.Del, "0ms"); err != nil {
			log.Fatalf("del tc.qdisc failed, Err: %+v", err)
		}
	}()

	http.Handle("/latency/", server.NewServer(exector))
	if err := http.ListenAndServe(":"+strconv.Itoa(o.port), nil); err != nil {
		log.Fatal(err)
	}
}
