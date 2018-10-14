package AutoEBL

import (
	"fmt"
	"os"
)

const (
	version = "0.01"
)

type AutoEBL struct {
	URL  string
	args []string
}

func New() *AutoEBL {
	return &AutoEBL{
		URL:  "http://www.lib.u-ryukyu.ac.jp",
		args: os.Args[1:],
	}
}

func (auebl *AutoEBL) Run() error {
	_, err := fmt.Println("Hello")
	return err
}
