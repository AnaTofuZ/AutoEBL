package AutoEBL

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
)

type Options struct {
	Help    bool `short:"h" long:"helo" description:"show this message"`
	Version bool `short:"vj" long:"version" description:"show the version"`
}

func (opts *Options) parse(argv []string) ([]string, error) {
	parser := flags.NewParser(opts, flags.PrintErrors)
	args, err := parser.ParseArgs(argv)

	if err != nil {
		parser.WriteHelp(os.Stderr)
		return nil, errors.Wrap(err, "invalid command line options")
	}
	return args, nil
}

func (auebl *AutoEBL) parseOptions() {
	return
}
