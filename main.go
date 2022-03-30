package main

import (
	"flag"

	"github.com/qobbysam/socketnotify/internal"
)

//socket main will be a rest server that will recieve notifications from socketlabs and notify local business interests

func main() {

	start := flag.String("st", "", "start operation, server or emailtest")

	path_to_config := flag.String("p", "", "path to config file")

	resourcename := flag.String("rn", "", "nameof resource")

	alive := flag.Bool("al", false, "should this be alive")
	flag.Parse()

	app := internal.InternalStruct{}

	app.StartApplication(*start, *path_to_config, *resourcename, *alive)
}
