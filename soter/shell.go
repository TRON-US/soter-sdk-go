package soter

// package shell implements a remote API interface for a running ipfs daemon

import (
	"crypto/tls"
	"fmt"
	gohttp "net/http"
	"time"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multiaddr-net"
)

type Shell struct {
	privateKey  string
	userAddress string
	url         string
	httpcli     gohttp.Client
}

func NewShell(privateKey, userAddress, url string) *Shell {
	c := &gohttp.Client{
		Transport: &gohttp.Transport{
			Proxy:             gohttp.ProxyFromEnvironment,
			DisableKeepAlives: true,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	return NewShellWithClient(privateKey, userAddress, url, c)
}

func NewShellWithClient(privateKey, userAddress, url string, c *gohttp.Client) *Shell {
	var sh Shell
	sh.privateKey = privateKey
	sh.userAddress = userAddress
	if a, err := ma.NewMultiaddr(url); err == nil {
		_, host, err := manet.DialArgs(a)
		if err == nil {
			url = host
		}
	}
	sh.url = url
	sh.httpcli = *c
	// We don't support redirects.
	sh.httpcli.CheckRedirect = func(_ *gohttp.Request, _ []*gohttp.Request) error {
		return fmt.Errorf("unexpected redirect")
	}
	return &sh
}

func (s *Shell) SetTimeout(d time.Duration) {
	s.httpcli.Timeout = d
}

func (s *Shell) Request(command string, args ...string) *RequestBuilder {
	return &RequestBuilder{
		command: command,
		args:    args,
		shell:   s,
	}
}
