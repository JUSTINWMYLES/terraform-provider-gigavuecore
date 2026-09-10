package provider

import (
	json "encoding/json"
	"fmt"
	"io"
	http "net/http"
	httptest "net/http/httptest"
	"strings"
	"sync"
	"testing"
)
import (
	providerserver "github.com/hashicorp/terraform-plugin-framework/providerserver"
	tfprotov6 "github.com/hashicorp/terraform-plugin-go/tfprotov6"
	resource "github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccActivationResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_activation\" \"example\" {\n  eli_id = \"example\"\n  page = \"%s\"\n  quantity = 0\n  sort = \"example\"\n}\n", serverURL, name)
}

// newActivationResourceMockServer returns an httptest server that stubs the ActivationResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newActivationResourceMockServer() *httptest.Server {
	mux := http.NewServeMux()
	state0 := make(map[string]map[string]interface{})
	var mu0 sync.Mutex
	lastKey0 := ""
	handler0 := func(w http.ResponseWriter, r *http.Request) {
		if bu0, bp0, bok0 := r.BasicAuth(); !bok0 || bu0 != "example" || bp0 != "example" {
			http.Error(w, "missing basic auth credential", http.StatusUnauthorized)
			return
		}
		mu0.Lock()
		defer mu0.Unlock()
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/licensing/ems/activations"), "/")
		if id == "" {
			id = "example-id"
		}
		switch r.Method {
		case http.MethodPost:
			body := make(map[string]interface{})
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil && err != io.EOF {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if _, ok := body["eliId"]; !ok {
				body["eliId"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["eliId"])
			state0[id] = body
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			_ = json.NewEncoder(w).Encode(body)
			lastKey0 = id
			return
		case http.MethodGet:
			body, ok := state0[id]
			if !ok && lastKey0 != "" {
				body = state0[lastKey0]
			}
			if body == nil {
				http.NotFound(w, r)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			_ = json.NewEncoder(w).Encode(body)
			return
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
	mux.HandleFunc("/licensing/ems/activations", handler0)
	mux.HandleFunc("/licensing/ems/activations/", handler0)
	state1 := make(map[string]map[string]interface{})
	var mu1 sync.Mutex
	handler1 := func(w http.ResponseWriter, r *http.Request) {
		if bu0, bp0, bok0 := r.BasicAuth(); !bok0 || bu0 != "example" || bp0 != "example" {
			http.Error(w, "missing basic auth credential", http.StatusUnauthorized)
			return
		}
		mu1.Lock()
		defer mu1.Unlock()
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/licensing/ems/reclaim"), "/")
		if id == "" {
			id = "example-id"
		}
		switch r.Method {
		case http.MethodDelete:
			delete(state1, id)
			w.WriteHeader(200)
			return
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
	mux.HandleFunc("/licensing/ems/reclaim", handler1)
	mux.HandleFunc("/licensing/ems/reclaim/", handler1)
	return httptest.NewServer(mux)
}

// TestAccActivationResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccActivationResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newActivationResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccActivationResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_activation.example", "eli_id"), resource.TestCheckResourceAttr("gigavuecore_activation.example", "page", "example"))}, resource.TestStep{ResourceName: "gigavuecore_activation.example", ImportState: true, ImportStateId: "imported-eli_id"}}})
}
