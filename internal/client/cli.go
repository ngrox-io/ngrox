package client

import (
	"flag"
	"fmt"
	"os"

	"github.com/ngrox-io/ngrox/internal/version"
)

const usage1 string = `Usage: %s [OPTIONS] <local port or address>
Options:
`

const usage2 string = `
Examples:
	ngrox 80
	ngrox -subdomain=example 8080
	ngrox -proto=tcp 22
	ngrox -hostname="example.com" -httpauth="user:password" 10.0.0.1


Advanced usage: ngrox [OPTIONS] <command> [command args] [...]
Commands:
	ngrox start [tunnel] [...]    Start tunnels by name from config file
	ngork start-all               Start all tunnels defined in config file
	ngrox list                    List tunnel names from config file
	ngrox help                    Print help
	ngrox version                 Print ngrox version

Examples:
	ngrox start www api blog pubsub
	ngrox -log=stdout -config=ngrox.yml start ssh
	ngrox start-all
	ngrox version

`

type Options struct {
	config    string
	logto     string
	loglevel  string
	authtoken string
	httpauth  string
	hostname  string
	protocol  string
	subdomain string
	command   string
	args      []string
}

func ParseArgs() (opts *Options, err error) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage1, os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, usage2)
	}

	config := flag.String(
		"config",
		"",
		"Path to ngrox configuration file. (default: $HOME/.ngrox)")

	logto := flag.String(
		"log",
		"none",
		"Write log messages to this file. 'stdout' and 'none' have special meanings")

	loglevel := flag.String(
		"log-level",
		"DEBUG",
		"The level of messages to log. One of: DEBUG, INFO, WARNING, ERROR")

	authtoken := flag.String(
		"authtoken",
		"",
		"Authentication token for identifying an ngrox.com account")

	httpauth := flag.String(
		"httpauth",
		"",
		"username:password HTTP basic auth creds protecting the public tunnel endpoint")

	subdomain := flag.String(
		"subdomain",
		"",
		"Request a custom subdomain from the ngrox server. (HTTP only)")

	hostname := flag.String(
		"hostname",
		"",
		"Request a custom hostname from the ngrox server. (HTTP only) (requires CNAME of your DNS)")

	protocol := flag.String(
		"proto",
		"http+https",
		"The protocol of the traffic over the tunnel {'http', 'https', 'tcp'} (default: 'http+https')")

	flag.Parse()

	opts = &Options{
		config:    *config,
		logto:     *logto,
		loglevel:  *loglevel,
		httpauth:  *httpauth,
		subdomain: *subdomain,
		protocol:  *protocol,
		authtoken: *authtoken,
		hostname:  *hostname,
		command:   flag.Arg(0),
	}

	switch opts.command {
	case "list":
		opts.args = flag.Args()[1:]
	case "start":
		opts.args = flag.Args()[1:]
	case "start-all":
		opts.args = flag.Args()[1:]
	case "version":
		fmt.Println(version.MajorMinor())
		os.Exit(0)
	case "help":
		flag.Usage()
		os.Exit(0)
	case "":
		err = fmt.Errorf("Error: Specify a local port to tunnel to, or " +
			"an ngrox command.\n\nExample: To expose port 80, run " +
			"'ngrox 80'")
		return

	default:
		if len(flag.Args()) > 1 {
			err = fmt.Errorf("You may only specify one port to tunnel to on the command line, got %d: %v",
				len(flag.Args()),
				flag.Args())
			return
		}

		opts.command = "default"
		opts.args = flag.Args()
	}

	return
}
