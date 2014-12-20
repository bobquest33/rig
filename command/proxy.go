package command

import (
	"fmt"
	"net"
	"net/http"

	"github.com/mattbaird/gosaml"
	"github.com/mitchellh/cli"
)

// ProxyCommand is a Command implementation that starts proxy listener
type ProxyCommand struct {
	ShutdownCh      <-chan struct{}
	AppSettings     *saml.AppSettings
	AccountSettings *saml.AccountSettings
	UI              cli.Ui
}

// Help implements Commands Help documentation method
func (c *ProxyCommand) Help() string {
	return ""
}

// Run boots the proxy up, listens on the port, starts the handler, etc.
func (c *ProxyCommand) Run(_ []string) int {
	go c.Serve()

	select {
	case <-c.ShutdownCh:
		return 1
	}
}

// Serve wraps around http.Serve with the internal HandleFunc.
func (c *ProxyCommand) Serve() {
	l, err := net.Listen("unix", "/tmp/foo.bar.sock")
	if err != nil {
		c.UI.Output(fmt.Sprint(err))
	}
	defer l.Close()

	if err := http.Serve(l, c); err != nil {
		c.UI.Output(fmt.Sprint(err))
	}
}

// HandleFunc provides a http Handler implementation
func (c *ProxyCommand) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	authRequest := saml.NewAuthorizationRequest(
		*c.AppSettings,
		*c.AccountSettings,
	)

	// Return a SAML AuthnRequest as a string
	saml, err := authRequest.GetSignedRequest(
		false,
		"/path/to/publickey.cer",
		"/path/to/privatekey.pem",
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(saml)
}

// Synopsis implements a short descriptive string for the CLI library
func (c *ProxyCommand) Synopsis() string {
	return "Launch the Rig proxy"
}
