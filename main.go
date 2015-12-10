package main

import (
	"flag"
	"log"
	"net/http"
	"runtime"

	"github.com/dakiva/widget-sample/widget"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("main")

// Represents the version number of the server.
// This value should not be written to programmatically and is set at compile time through -ldflags
var version string = "<local dev build>"

func main() {
	// export the version in a package outside main
	widget.ServiceVersion = version
	http.Handle("/favicon.ico", http.NotFoundHandler())
	runtime.GOMAXPROCS(runtime.NumCPU())

	configFile := flag.String("conf", "conf/widget_service.conf", "Name and location of the widget configuration file.")
	flag.Parse()
	serviceConfig, err := widget.LoadServiceConfig(*configFile)
	if err != nil {
		log.Fatalln("Error loading the widget configuration file.", err)
	}
	err = serviceConfig.Validate()
	if err != nil {
		log.Fatalln(err)
	}
	err = serviceConfig.Initialize()
	if err != nil {
		log.Fatalln(err)
	}
	logger.Notice("%v (v%v) started on %s", widget.ServiceName, version, serviceConfig.GetHostAddress())
	logger.Fatal(http.ListenAndServe(serviceConfig.GetHostAddress(), widget.Container))
}
