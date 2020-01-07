// package soter implements a remote API interface for a running ipfs daemon
package soter

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ipfs/go-ipfs-api"
	"io"
	"io/ioutil"
	gohttp "net/http"
	"strings"
	"time"

	files "github.com/ipfs/go-ipfs-files"
	tar "github.com/whyrusleeping/tar-utils"
)

const (
	DefaultPathName = ".ipfs"
	DefaultPathRoot = "~/" + DefaultPathName
	DefaultApiFile  = "api"
	EnvDir          = "IPFS_PATH"
)

type Soter struct {
	url        string
	httpcli    gohttp.Client
	userAddr   string
	privateKey string
}

func NewSoter(url string) *Soter {
	c := &gohttp.Client{
		Transport: &gohttp.Transport{
			Proxy:             gohttp.ProxyFromEnvironment,
			DisableKeepAlives: true,
		},
	}

	return NewSoterWithClient(url, c)
}

func NewSoterWithClient(url string, useraddr string, private string, c *gohttp.Client) *Soter {
	var st Soter
	st.url = url
	st.httpcli = *c
	st.userAddr = useraddr
	st.privateKey = private
	// We don't support redirects.
	st.httpcli.CheckRedirect = func(_ *gohttp.Request, _ []*gohttp.Request) error {
		return fmt.Errorf("unexpected redirect")
	}
	return &st
}

func (s *Soter) SetTimeout(d time.Duration) {
	s.httpcli.Timeout = d
}

func (s *Soter) Request(command string, args ...string) *RequestBuilder {
	return &RequestBuilder{
		command: command,
		args:    args,
		soter:   s,
	}
}

func (s *Soter) QueryUserBalance() (result string, err error) {
	err = s.Request("balance", 具体值).
		ExecGet(context.Background(), &result)
	return
}

// Cat the content at the given path. Callers need to drain and close the returned reader after usage.
func (s *Soter) Cat(path string) (io.ReadCloser, error) {
	resp, err := s.Request("cat", path).Send(context.Background())
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, resp.Error
	}

	return resp.Output, nil
}

const (
	TRaw = iota
	TDirectory
	TFile
	TMetadata
	TSymlink
)

// List entries at the given path
func (s *Soter) List(path string) ([]*LsLink, error) {
	var out struct{ Objects []LsObject }
	err := s.Request("ls", path).ExecPost(context.Background(), &out)
	if err != nil {
		return nil, err
	}
	if len(out.Objects) != 1 {
		return nil, errors.New("bad response from server")
	}
	return out.Objects[0].Links, nil
}
