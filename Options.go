package AutoEBL

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
)

type Options struct {
	Help     bool   `short:"h" long:"help" description:"show this message"`
	Version  bool   `short:"v" long:"version" description:"show the version"`
	Username string `short:"u" long:"username" description:"set your user name ex.e1557xx"`
	Password string `short:"p" long:"password" description:"set your password for ie"`
}

func (opts *Options) parse(parser *flags.Parser, argv []string) ([]string, error) {
	args, err := parser.ParseArgs(argv)

	if err != nil {
		parser.WriteHelp(os.Stderr)
		return nil, errors.Wrap(err, "invalid command line options")
	}
	return args, nil
}

func (autoebl *AutoEBL) parseOptions(osargs []string) error {
	var opts Options
	parser := flags.NewParser(&opts, flags.PrintErrors)
	args, err := opts.parse(parser, osargs)

	if err != nil {
		return err
	}

	if opts.Version {
		fmt.Println("Auto EBL", version)
		os.Exit(0)
	}

	if opts.Help || (!(autoebl.Define()) && len(args) == 0) {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	autoebl.check_set(&autoebl.username, opts.Username)
	autoebl.check_set(&autoebl.password, opts.Password)

	return nil
}
