package opts

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
)

// Opts contains command line options
type Opts struct {
	// Webserver options
	WebServer string `short:"w" long:"webserver" description:"Listen address for webserver" default:":8888" value-name:"<[host]:port>"`

	// UDP listener
	UDPListenAddress string `short:"u" long:"udp-listener" description:"Listen address for UDP listener" default:":9191" value-name:"<[host]:port>"`
	UDPBufferSize    int    `short:"b" long:"udp-buffer-size" description:"Size of UDP read buffer" default:"1024" value-name:"<num bytes>"`

	// Database options
	DBFilename string `short:"d" long:"db" description:"Data storage file" default:"ds.db" value-name:"<file>"`
}

var parsedOpts Opts

// Parse parses the command line options
func Parse() *Opts {
	parser := flags.NewParser(&parsedOpts, flags.Default)

	_, err := parser.Parse()
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		log.Fatalf("Error parsing flags: %v", err)
	}
	return &parsedOpts
}

// ParsedOpts returns a reference to the parsedOpts structure
func ParsedOpts() *Opts {
	return &parsedOpts
}
