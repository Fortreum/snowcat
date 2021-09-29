package netscan

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

var (
	InsecureClient = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
)

type httpScanner struct {
	tls bool
}

func (s *httpScanner) Scan(addr string, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	u := &url.URL{
		Scheme: "http",
		Host:   addr,
	}
	if s.tls {
		u.Scheme = "https"
	}

	req, err := http.NewRequest("HEAD", u.String(), nil)
	if err != nil {
		log.Printf("invalid url: %s", err)
	}

	_, err = http.DefaultClient.Do(req.WithContext(ctx))
	return err == nil
}
