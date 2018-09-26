package tests

import (
	"testing"
	"net/http"
	"github.com/centrifuge/functional-testing/go/utils"
	"os"
)

func TestPing(t *testing.T) {
	network := os.Getenv("NETWORK")
	if network == "" {
		network = "centrifugeRussianhillEthRinkeby"
	}
	nodeURL := os.Getenv("NODE1URL")
	if nodeURL == "" {
		nodeURL = "https://35.184.66.29:8082"
	}
	e := utils.CreateInsecureClient(t, nodeURL)
	obj := e.GET("/ping").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.Value("network").Equal(network)
	obj.ContainsKey("version")
}
