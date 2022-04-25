package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/Farengier/myofficestorage/src/storage"
	"github.com/Farengier/myofficestorage/src/www"
)

type opts struct {
	Listen string `short:"l" long:"listen" description:"listen address" default:"0.0.0.0:8080"`
}

func main() {
	fmt.Println("Reading options")
	o := initOpts()

	st := storage.MemoryStorage()
	srv := www.Server(log.New(os.Stderr, "[Server] ", log.LstdFlags|log.Lmsgprefix), o.Listen, st)
	err := srv.Start()
	fmt.Println(err)
}

func initOpts() opts {
	o := opts{}
	_, err := flags.NewParser(&o, flags.HelpFlag|flags.PassDoubleDash).Parse()
	if err != nil {
		panic(err)
	}
	return o
}
