package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/bububa/ratelimitd/app"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())
	server, deferFunc := app.NewApp()
	defer deferFunc()
	err := server.Run(os.Args)
	if err != nil {
		log.Println(err)
	}
}
