package tests

import (
	"testing"
	"net/http"
	"github.com/centrifuge/functional-testing/go/utils"
)

func TestPing(t *testing.T) {
	e := utils.GetInsecureClient(t, utils.NODE1)
	obj := e.GET("/ping").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.Value("network").Equal(utils.Network)
	obj.ContainsKey("version")
}
