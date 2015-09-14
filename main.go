package main

import (
	"flag"
	"github.com/jolshevski/chester/api"
)

func main() {
	app := api.New()

	// Define configuration flags
	modulepath := flag.String("modulepath", "", "Directory containing module release tarballs to serve. Required.")
	binding := flag.String("binding", ":8080", "Golang ListenAndServe binding")
	flag.Parse()

	// Configure the API
	app.Config["modulepath"] = *modulepath
	app.Config["binding"] = *binding

	// Start the API server
	app.Listen()
}
