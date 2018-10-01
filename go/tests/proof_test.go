package tests

import (
	"testing"
	"github.com/centrifuge/functional-testing/go/utils"
	"net/http"
)

func TestProofGenerationWithMultipleFields(t *testing.T) {
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


	proofPayload := map[string]interface{}{
		"documentIdentifier": docIdentifier,
		"fields": []string{"net_amount", "currency"},
	}
	obj = e.POST("/legacy/invoice/proof").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(proofPayload).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Value("document_identifier").String().Equal(docIdentifier)

}
