package AutoEBL

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"log"
	"os"
	"time"
)

const (
	version = "0.01"
	url     = "https://opac.lib.u-ryukyu.ac.jp/portal/portal/selectLogin"
)

type AutoEBL struct {
	username string
	password string
}

func New() *AutoEBL {
	return &AutoEBL{}
}

func (auebl *AutoEBL) Run() error {
	auebl.setEnvParams()
	err := auebl.parseOptions(os.Args[1:])

	//if err != nil { return err }

	if !auebl.Define() {
		return errors.New("failed to not input username and password")
	}

	auebl.doEBL()

	return err
}

// setEnvParams is set username and password for  enviroment values
// this function didn't happen exception by that os.Getenv return  not return erorr
func (auebl *AutoEBL) setEnvParams() {
	auebl.check_set(&auebl.username, os.Getenv("AUEBL_USERNAME"))
	auebl.check_set(&auebl.password, os.Getenv("AUEBL_PASSWORD"))
}

func (auebl *AutoEBL) check_set(auebl_conf *string, set_val string) {
	if set_val == "" {
		return
	}
	*auebl_conf = set_val
}

func (auebl *AutoEBL) Define() bool {
	if auebl.username == "" || auebl.password == "" {
		return false
	}
	return true
}

func (auebl *AutoEBL) doEBL() error {
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	chrome, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))

	if err != nil {
		return err
	}

	err = chrome.Run(ctxt, auebl.sendkeys())

	if err != nil {
		return err
	}

	chrome.Shutdown(ctxt)

	return err
}

func (auebl *AutoEBL) sendkeys() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		// LDAP login
		chromedp.Click("javascript:$('input[name=lang]').val('ja');$('#frmSso').submit();"),
		chromedp.SetValue(`#username`, auebl.username),
		chromedp.SetValue(`#password`, auebl.password),
		chromedp.Click("_eventId_proceed"),
		// library menue
		chromedp.WaitVisible("#menu-inner"),
		chromedp.Click(`//*[@id="menu-inner"]/ul/li[2]/a`),
		chromedp.Click(`//*[@id="menu-inner"]/ul/li[2]/ul/li[1]/a`),
		// extention
		chromedp.WaitVisible(`//*[@id="lblTh1"]`),
		chromedp.Click(`#itemCheckA000`),
		chromedp.Click(`//*[@id="lblExtend"]`),
		chromedp.Sleep(30 * time.Second),
	}
}
