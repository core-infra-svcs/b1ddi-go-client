package acl

import (
	"context"
	"crypto/tls"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/infobloxopen/b1ddi-go-client/runtimetest"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestClient(t *testing.T) {
	testCases := []struct {
		testMethodName   string
		testMethodParams interface{}
		expectedRequest  http.Request
	}{
		{
			"ACLCreate",
			&ACLCreateParams{
				Body: &models.ConfigACL{
					Name: swag.String("acl_test_name"),
				},
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/acl"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"name\":\"acl_test_name\"}\n")),
			},
		},
		{
			"ACLDelete",
			&ACLDeleteParams{
				ID:      "acl-delete-id",
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/acl/acl-delete-id"},
				Method: http.MethodDelete,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"ACLList",
			&ACLListParams{
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/acl"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"ACLRead",
			&ACLReadParams{
				ID:      "acl-read-id",
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/acl/acl-read-id"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
		},
		{
			"ACLUpdate",
			&ACLUpdateParams{
				ID: "acl-update-id",
				Body: &models.ConfigACL{
					Name:    swag.String("acl_test_name"),
					Comment: "Updated comment",
				},
				Context: context.TODO()},
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/dns/acl/acl-update-id"},
				Method: http.MethodPatch,
				Body:   io.NopCloser(strings.NewReader("{\"comment\":\"Updated comment\",\"name\":\"acl_test_name\"}\n")),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testMethodName, func(t *testing.T) {
			// Get the mock server response body
			resp := reflect.ValueOf(&Client{}).MethodByName(tc.testMethodName).Type().Out(0)
			// Initialize mock server
			s := runtimetest.InitTestServer(t, tc.expectedRequest, reflect.New(resp))

			// Initialize the client
			c := initACLTestClient(s.URL)

			methodParams := []reflect.Value{
				reflect.ValueOf(tc.testMethodParams),
				reflect.New(reflect.TypeOf((*runtime.ClientAuthInfoWriter)(nil)).Elem()).Elem(),
			}

			// Call the method by the name specified in the test case
			results := reflect.ValueOf(c).MethodByName(tc.testMethodName).Call(methodParams)
			if err := results[1].Interface(); err != nil {
				t.Fatal(err)
			}

			// Close the mock server
			s.Close()
		})
	}
}

func initACLTestClient(url string) ClientService {
	transport := httptransport.New(
		strings.TrimPrefix(url, "https://"), "api/ddi/v1", nil,
	)
	// Disable TLS verify for the mock server
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	transport.Transport = tr
	// create the Address API client
	return New(transport, strfmt.Default)
}
