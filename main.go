package main

import (
	"embed"
	"log"
	"os"

	"github.com/profsmallpine/books/app"
)

var files embed.FS

func main() {
	logger := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	ranger, err := app.New(logger, files)
	if err != nil {
		logger.Fatal(err)
	}

	// start the web server until receiving a signal to stop.
	if err := ranger.Guide(); err != nil {
		logger.Fatal("could not guide: ", err)
	}
}
