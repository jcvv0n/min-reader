package main

import (
	"github.com/jcvv0n/min-reader/router"
)

func main() {
	r := router.Router()
	r.Run(":7676")
}
