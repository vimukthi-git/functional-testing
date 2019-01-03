package tests

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/centrifuge/functional-testing/go/utils"
	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndUpdateInvoiceFromOrigin(t *testing.T) {
	// nodes
	e := utils.GetInsecureClient(t, utils.NODE1)
	e1 := utils.GetInsecureClient(t, utils.NODE2)

	// create invoice
	currency := "USD"
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"invoice_number": "12324",
			"due_date":       "2018-09-26T23:12:37.902198664Z",
			"gross_amount":   "40",
			"currency":       currency,
			"net_amount":     "40",
		},
		"collaborators": []string{utils.Nodes[utils.NODE2].ID},
	}

	obj := CreateDocument(t, utils.INVOICE, e, payload)

	docIdentifier := obj.Value("header").Path("$.document_id").String().NotEmpty().Raw()

	doc := GetDocument(t, utils.INVOICE, e, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// Receiver has document
	doc = GetDocument(t, utils.INVOICE, e1, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// update invoice
	payload = map[string]interface{}{
		"data": map[string]interface{}{
			"invoice_number": "12324",
			"due_date":       "2018-09-26T23:12:37.902198664Z",
			"gross_amount":   "41",
			"currency":       currency,
			"net_amount":     "41",
		},
		"collaborators": []string{utils.Nodes[utils.NODE2].ID},
	}

	obj = UpdateDocument(t, utils.INVOICE, e, docIdentifier, payload)

	// check updated gross amount
	obj.Value("data").Path("$.gross_amount").String().Equal("41")
	GetDocument(t, utils.INVOICE, e, docIdentifier)

	// Receiver has document
	doc = GetDocument(t, utils.INVOICE, e1, docIdentifier)
	doc.Path("$.data.gross_amount").String().Equal("41")
}

func TestCreateAndUpdateInvoiceFromCollaborator(t *testing.T) {
	// nodes
	e := utils.GetInsecureClient(t, utils.NODE1)
	e1 := utils.GetInsecureClient(t, utils.NODE2)

	// create invoice
	currency := "USD"
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"invoice_number": "12324",
			"due_date":       "2018-09-26T23:12:37.902198664Z",
			"gross_amount":   "40",
			"currency":       currency,
			"net_amount":     "40",
		},
		"collaborators": []string{utils.Nodes[utils.NODE2].ID},
	}

	obj := CreateDocument(t, utils.INVOICE, e, payload)

	docIdentifier := obj.Value("header").Path("$.document_id").String().NotEmpty().Raw()

	doc := GetDocument(t, utils.INVOICE, e, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// Receiver has document
	doc = GetDocument(t, utils.INVOICE, e1, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// update invoice
	payload = map[string]interface{}{
		"data": map[string]interface{}{
			"invoice_number": "12324",
			"due_date":       "2018-09-26T23:12:37.902198664Z",
			"gross_amount":   "41",
			"currency":       currency,
			"net_amount":     "41",
		},
		"collaborators": []string{utils.Nodes[utils.NODE1].ID},
	}

	obj = UpdateDocument(t, utils.INVOICE, e1, docIdentifier, payload)

	// check updated gross amount
	obj.Value("data").Path("$.gross_amount").String().Equal("41")
	GetDocument(t, utils.INVOICE, e1, docIdentifier)

	// Receiver has document
	doc = GetDocument(t, utils.INVOICE, e, docIdentifier)
	doc.Path("$.data.gross_amount").String().Equal("41")
}

func TestCreateAndUpdatePurchaseOrderFromOrigin(t *testing.T) {
	// nodes
	e := utils.GetInsecureClient(t, utils.NODE1)
	e1 := utils.GetInsecureClient(t, utils.NODE2)

	// create invoice
	currency := "USD"
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"po_number":     "12324",
			"delivery_date": "2018-09-26T23:12:37.902198664Z",
			"tax_amount":    "40",
			"currency":      currency,
			"net_amount":    "40",
		},
		"collaborators": []string{utils.Nodes[utils.NODE2].ID},
	}

	obj := CreateDocument(t, utils.PURCHASEORDER, e, payload)

	docIdentifier := obj.Value("header").Path("$.document_id").String().NotEmpty().Raw()

	doc := GetDocument(t, utils.PURCHASEORDER, e, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// Receiver has document
	doc = GetDocument(t, utils.PURCHASEORDER, e1, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// update invoice
	payload = map[string]interface{}{
		"data": map[string]interface{}{
			"po_number":     "12324",
			"delivery_date": "2018-09-26T23:12:37.902198664Z",
			"tax_amount":    "41",
			"currency":      currency,
			"net_amount":    "41",
		},
		"collaborators": []string{utils.Nodes[utils.NODE2].ID},
	}

	obj = UpdateDocument(t, utils.PURCHASEORDER, e, docIdentifier, payload)

	// check updated gross amount
	obj.Value("data").Path("$.net_amount").String().Equal("41")
	GetDocument(t, utils.PURCHASEORDER, e, docIdentifier)

	// Receiver has document
	doc = GetDocument(t, utils.PURCHASEORDER, e1, docIdentifier)
	doc.Path("$.data.net_amount").String().Equal("41")
}

func TestCreateAndUpdatePurchaseOrderFromCollaborator(t *testing.T) {
	// nodes
	e := utils.GetInsecureClient(t, utils.NODE1)
	e1 := utils.GetInsecureClient(t, utils.NODE2)

	// create invoice
	currency := "USD"
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"po_number":    "12324",
			"due_date":     "2018-09-26T23:12:37.902198664Z",
			"gross_amount": "40",
			"currency":     currency,
			"net_amount":   "40",
		},
		"collaborators": []string{utils.Nodes[utils.NODE2].ID},
	}

	obj := CreateDocument(t, utils.PURCHASEORDER, e, payload)

	docIdentifier := obj.Value("header").Path("$.document_id").String().NotEmpty().Raw()

	doc := GetDocument(t, utils.PURCHASEORDER, e, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// Receiver has document
	doc = GetDocument(t, utils.PURCHASEORDER, e1, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// update invoice
	payload = map[string]interface{}{
		"data": map[string]interface{}{
			"po_number":     "12324",
			"delivery_date": "2018-09-26T23:12:37.902198664Z",
			"tax_amount":    "41",
			"currency":      currency,
			"net_amount":    "41",
		},
		"collaborators": []string{utils.Nodes[utils.NODE1].ID},
	}

	obj = UpdateDocument(t, utils.PURCHASEORDER, e1, docIdentifier, payload)

	// check updated gross amount
	obj.Value("data").Path("$.net_amount").String().Equal("41")
	GetDocument(t, utils.PURCHASEORDER, e1, docIdentifier)

	// Receiver has document
	doc = GetDocument(t, utils.PURCHASEORDER, e, docIdentifier)
	doc.Path("$.data.net_amount").String().Equal("41")
}

func GetDocument(t *testing.T, docType string, e *httpexpect.Expect, docIdentifier string) *httpexpect.Value {
	objGet := e.GET(fmt.Sprintf("/%s/%s", docType, docIdentifier)).
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		Expect().Status(http.StatusOK)
	assertOkResponse(t, objGet)
	objGet.JSON().Path("$.header.document_id").String().Equal(docIdentifier)
	return objGet.JSON()
}

func CreateDocument(t *testing.T, docType string, e *httpexpect.Expect, payload map[string]interface{}) *httpexpect.Object {
	path := fmt.Sprintf("/%s", docType)
	method := "POST"
	resp := getResponse(method, path, e, payload).Status(http.StatusOK)
	assertOkResponse(t, resp)
	obj := resp.JSON().Object()
	txID := getTransactionID(t, obj)
	waitTillSuccess(t, e, txID)
	return obj
}

func UpdateDocument(t *testing.T, docType string, e *httpexpect.Expect, documentID string, payload map[string]interface{}) *httpexpect.Object {
	path := fmt.Sprintf("/%s/%s", docType, documentID)
	method := "PUT"
	resp := getResponse(method, path, e, payload).Status(http.StatusOK)
	assertOkResponse(t, resp)
	obj := resp.JSON().Object()
	txID := getTransactionID(t, obj)
	waitTillSuccess(t, e, txID)
	return obj
}

func getResponse(method, path string, e *httpexpect.Expect, payload map[string]interface{}) *httpexpect.Response {
	return e.Request(method, path).
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(payload).
		Expect()
}

func assertOkResponse(t *testing.T, response *httpexpect.Response) {
	if response.Raw().StatusCode != http.StatusOK {
		assert.Fail(t, "Response Payload: ", response.Body().Raw())
	}
}

func getTransactionID(t *testing.T, resp *httpexpect.Object) string {
	txID := resp.Value("header").Path("$.transaction_id").String().Raw()
	if txID == "" {
		t.Error("transaction ID empty")
	}

	return txID
}

func waitTillSuccess(t *testing.T, e *httpexpect.Expect, txID string) {
	for {
		resp := e.GET("/transactions/" + txID).Expect().Status(200).JSON().Object()
		status := resp.Path("$.status").String().Raw()
		if status == "pending" {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		if status == "failed" {
			t.Error(resp.Path("$.message").String().Raw())
		}

		break
	}
}
