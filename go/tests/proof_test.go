package tests

import (
	"net/http"
	"testing"

	"github.com/centrifuge/functional-testing/go/utils"
	"github.com/gavv/httpexpect"
)

func TestProofGenerationWithMultipleFields(t *testing.T) {
	e := utils.GetInsecureClient(t, utils.NODE1)

	currency := "USD"
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"number":       "12324",
			"date_due":     "2018-09-26T23:12:37.902198664Z",
			"gross_amount": "40",
			"currency":     currency,
			"net_amount":   "40",
		},
	}

	obj := CreateDocument(t, utils.INVOICE, e, utils.Nodes[utils.NODE1].ID, payload, http.StatusAccepted)

	docIdentifier := obj.Value("header").Path("$.document_id").String().NotEmpty().Raw()

	proofPayload := map[string]interface{}{
		"fields": []string{"invoice.net_amount", "invoice.currency"},
	}

	objProof := GetProof(t, e, utils.Nodes[utils.NODE1].ID, docIdentifier, proofPayload)
	objProof.Path("$.header.document_id").String().Equal(docIdentifier)
	objProof.Path("$.field_proofs[0].property").String().Equal("0x000100000000002d") // invoice.net_amount
	objProof.Path("$.field_proofs[0].sorted_hashes").NotNull()
	objProof.Path("$.field_proofs[1].property").String().Equal("0x000100000000000d") // invoice.currency
	objProof.Path("$.field_proofs[1].sorted_hashes").NotNull()
}

func GetProof(t *testing.T, e *httpexpect.Expect, auth string, documentID string, payload map[string]interface{}) *httpexpect.Object {
	obj := utils.AddCommonHeaders(e.POST("/v1/documents/"+documentID+"/proofs"), auth).
		WithJSON(payload).
		Expect().Status(http.StatusOK)
	assertOkResponse(t, obj, http.StatusOK)
	return obj.JSON().Object()
}
