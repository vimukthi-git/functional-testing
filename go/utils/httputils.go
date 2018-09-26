package utils

import (
	"github.com/gavv/httpexpect"
	"net/http"
	"crypto/tls"
	"testing"
	"time"
)

func CreateInsecureClient(t *testing.T, baseURL string) *httpexpect.Expect {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	config := httpexpect.Config{
		BaseURL:  baseURL,
		Client:   &http.Client{
			Transport: tr,
			Timeout: time.Second * 600,
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewCompactPrinter(t),
		},
	}
	return httpexpect.WithConfig(config)
}
