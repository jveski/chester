package main

import (
	"flag"
	"github.com/jolshevski/chester/api"
)

func main() {
	app := api.New()

	// Define configuration flags
	modulepath := flag.String("modulepath", "", "Directory containing module release tarballs to serve. Required.")
	binding := flag.String("binding", "", "Golang ListenAndServe binding. Defaults to :8080.")
	flag.Parse()

	// Configure the API
	app.Config["modulepath"] = *modulepath
	app.Config["binding"] = *binding

	// Start the API server
	app.Listen()
}
