package utils

import (
	"os"
	"strings"
	"testing"

	"github.com/gavv/httpexpect"
)

const (
	NODE1         = "node1"
	NODE2         = "node2"
	INVOICE       = "invoice"
	PURCHASEORDER = "purchaseorder"
)

var Nodes map[string]node
var Network string

type node struct {
	ID   string
	HOST string
}

func SetupEnvironment() {
	nodesEnv := os.Getenv("NODES")
	idsEnv := os.Getenv("IDS")
	nodesSlice := SplitString(nodesEnv)
	idsSlice := SplitString(idsEnv)

	if len(nodesSlice) == 0 {
		nodesSlice = append(nodesSlice, "https://localhost:8082", "https://localhost:8083")
	}

	if len(idsSlice) == 0 {
		idsSlice = append(idsSlice, "0xf77402CA591ba52d51949FF12dC3c1f5c3e0f790", "0xD3527C484020947040e471c16442eA5a485fd46C")
	}

	Nodes = map[string]node{
		NODE1: {
			idsSlice[0],
			nodesSlice[0],
		},
		NODE2: {
			idsSlice[1],
			nodesSlice[1],
		},
	}

	Network = os.Getenv("NETWORK")
	if Network == "" {
		Network = "testing"
	}

}

func GetInsecureClient(t *testing.T, nodeId string) *httpexpect.Expect {
	SetupEnvironment()
	return CreateInsecureClient(t, Nodes[nodeId].HOST)
}

func SplitString(data string) []string {
	result := strings.Split(data, ",")
	if result[0] == "" {
		return []string{}
	}

	return result
}
