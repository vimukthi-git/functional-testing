package tests

import (
	"testing"
	"os"
	"github.com/centrifuge/functional-testing/go/utils"
	"net/http"
)

func TestSendInvoiceToOwnNodeOnly(t *testing.T) {
	nodeURL := os.Getenv("NODE1URL")
	if nodeURL == "" {
		nodeURL = "https://35.184.66.29:8082"
	}
	e := utils.CreateInsecureClient(t, nodeURL)

	payload := map[string]interface{}{
		"document": map[string]interface{}{
			"data": map[string]interface{}{
				"currency": "USD",
				"net_amount": "1501",
			},
		},
	}

	obj := e.POST("/invoice/send").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(payload).
		Expect().Status(http.StatusOK).JSON().Object()

	docIdentifier := obj.Value("core_document").Path("$.document_identifier").String().Raw()

	getPayload := map[string]interface{}{
		"document_identifier": docIdentifier,
	}

	e.POST("/invoice/get").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(getPayload).
		Expect().Status(http.StatusOK).JSON().NotNull()
}

func TestSendPurchaseOrderToOwnNodeOnly(t *testing.T) {
	nodeURL := os.Getenv("NODE1URL")
	if nodeURL == "" {
		nodeURL = "https://35.184.66.29:8082"
	}
	e := utils.CreateInsecureClient(t, nodeURL)

	payload := map[string]interface{}{
		"document": map[string]interface{}{
			"data": map[string]interface{}{
				"currency": "USD",
				"net_amount": "1501",
			},
		},
	}

	obj := e.POST("/purchaseorder/send").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(payload).
		Expect().Status(http.StatusOK).JSON().Object()

	docIdentifier := obj.Value("core_document").Path("$.document_identifier").String().Raw()

	getPayload := map[string]interface{}{
		"document_identifier": docIdentifier,
	}

	e.POST("/purchaseorder/get").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(getPayload).
		Expect().Status(http.StatusOK).JSON().NotNull()
}

func TestSendInvoiceToCollaborator(t *testing.T) {
	nodeURL1 := os.Getenv("NODE1URL")
	nodeURL2 := os.Getenv("NODE2URL")
	if nodeURL1 == "" {
		nodeURL1 = "https://35.184.66.29:8082"
	}
	if nodeURL2 == "" {
		nodeURL2 = "https://35.184.39.100:8082"
	}
	e := utils.CreateInsecureClient(t, nodeURL1)

	payload := map[string]interface{}{
		"document": map[string]interface{}{
			"data": map[string]interface{}{
				"currency": "USD",
				"net_amount": "1501",
			},
		},
		"recipients": []string{"JP5lVb65"},
	}

	obj := e.POST("/invoice/send").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(payload).
		Expect().Status(http.StatusOK).JSON().Object()

	docIdentifier := obj.Value("core_document").Path("$.document_identifier").String().Raw()

	// Sender has document
	getPayload := map[string]interface{}{
		"document_identifier": docIdentifier,
	}

	e.POST("/invoice/get").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(getPayload).
		Expect().Status(http.StatusOK).JSON().NotNull()

	// Receiver has document
	e1 := utils.CreateInsecureClient(t, nodeURL2)

	e1.POST("/invoice/get").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(getPayload).
		Expect().Status(http.StatusOK).JSON().NotNull()
}

func TestSendPurchaseOrderToCollaborator(t *testing.T) {
	nodeURL1 := os.Getenv("NODE1URL")
	nodeURL2 := os.Getenv("NODE2URL")
	if nodeURL1 == "" {
		nodeURL1 = "https://35.184.66.29:8082"
	}
	if nodeURL2 == "" {
		nodeURL2 = "https://35.184.39.100:8082"
	}
	e := utils.CreateInsecureClient(t, nodeURL1)

	payload := map[string]interface{}{
		"document": map[string]interface{}{
			"data": map[string]interface{}{
				"currency": "USD",
				"net_amount": "1501",
			},
		},
		"recipients": []string{"JP5lVb65"},
	}

	obj := e.POST("/purchaseorder/send").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(payload).
		Expect().Status(http.StatusOK).JSON().Object()

	docIdentifier := obj.Value("core_document").Path("$.document_identifier").String().Raw()

	// Sender has document
	getPayload := map[string]interface{}{
		"document_identifier": docIdentifier,
	}

	e.POST("/purchaseorder/get").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(getPayload).
		Expect().Status(http.StatusOK).JSON().NotNull()

	// Receiver has document
	e1 := utils.CreateInsecureClient(t, nodeURL2)

	e1.POST("/purchaseorder/get").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(getPayload).
		Expect().Status(http.StatusOK).JSON().NotNull()
}
