package tests

import (
	"testing"
	"github.com/centrifuge/functional-testing/go/utils"
	"net/http"
)

func TestSendInvoiceToOwnNodeOnly(t *testing.T) {
	e := utils.GetInsecureClient(t, utils.NODE1)

	payload := map[string]interface{}{
		"document": map[string]interface{}{
			"data": map[string]interface{}{
				"currency": "USD",
				"net_amount": "1501",
			},
		},
	}

	obj := e.POST("/legacy/invoice/send").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(payload).
		Expect().Status(http.StatusOK).JSON().Object()

	docIdentifier := obj.Value("core_document").Path("$.document_identifier").String().Raw()

	getPayload := map[string]interface{}{
		"document_identifier": docIdentifier,
	}

	e.POST("/legacy/invoice/get").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(getPayload).
		Expect().Status(http.StatusOK).JSON().NotNull()
}

func TestSendPurchaseOrderToOwnNodeOnly(t *testing.T) {
	e := utils.GetInsecureClient(t, utils.NODE1)

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
	e := utils.GetInsecureClient(t, utils.NODE1)

	payload := map[string]interface{}{
		"document": map[string]interface{}{
			"data": map[string]interface{}{
				"currency": "USD",
				"net_amount": "1501",
			},
		},
		"recipients": []string{"JP5lVb65"},
	}

	obj := e.POST("/legacy/invoice/send").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(payload).
		Expect().Status(http.StatusOK).JSON().Object()

	docIdentifier := obj.Value("core_document").Path("$.document_identifier").String().Raw()

	// Sender has document
	getPayload := map[string]interface{}{
		"document_identifier": docIdentifier,
	}

	e.POST("/legacy/invoice/get").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(getPayload).
		Expect().Status(http.StatusOK).JSON().NotNull()

	// Receiver has document
	e1 := utils.GetInsecureClient(t, utils.NODE2)

	e1.POST("/legacy/invoice/get").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(getPayload).
		Expect().Status(http.StatusOK).JSON().NotNull()
}

func TestSendPurchaseOrderToCollaborator(t *testing.T) {
	e := utils.GetInsecureClient(t, utils.NODE1)

	payload := map[string]interface{}{
		"document": map[string]interface{}{
			"data": map[string]interface{}{
				"currency": "USD",
				"net_amount": "1501",
			},
		},
		"recipients": []string{"JP5lVb65"},
	}

	obj := e.POST("/legacy/purchaseorder/send").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(payload).
		Expect().Status(http.StatusOK).JSON().Object()

	docIdentifier := obj.Value("core_document").Path("$.document_identifier").String().Raw()

	// Sender has document
	getPayload := map[string]interface{}{
		"document_identifier": docIdentifier,
	}

	e.POST("/legacy/purchaseorder/get").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(getPayload).
		Expect().Status(http.StatusOK).JSON().NotNull()

	// Receiver has document
	e1 := utils.GetInsecureClient(t, utils.NODE2)

	e1.POST("/legacy/purchaseorder/get").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(getPayload).
		Expect().Status(http.StatusOK).JSON().NotNull()
}
