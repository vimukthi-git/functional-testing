package utils

import (
	"os"
	"strings"
	"github.com/gavv/httpexpect"
	"testing"
)

const (
	NODE1 = "node1"
	NODE2 = "node2"
)

var Nodes map[string]string
var Network string

func SetupEnvironment() {
	nodesEnv := os.Getenv("NODES")
	nodesSlice := SplitString(nodesEnv, ",")

	if len(nodesSlice) == 0 {
		nodesSlice = append(nodesSlice, "https://localhost:8082", "https://localhost:8083")
	}

	Nodes = map[string]string{
		NODE1: nodesSlice[0],
		NODE2: nodesSlice[1],
	}

	Network = os.Getenv("NETWORK")
	if Network == "" {
		Network = "testing"
	}

}

func GetInsecureClient(t *testing.T, nodeId string) *httpexpect.Expect {
	SetupEnvironment()
	return CreateInsecureClient(t, Nodes[nodeId])
}

func SplitString(data string, del string) []string {
	result := strings.Split(data, ",")
	if result[0] == "" {
		return []string{}
	}

	return result
}
