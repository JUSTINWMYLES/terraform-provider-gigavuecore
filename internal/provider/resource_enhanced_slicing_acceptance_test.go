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

func testAccEnhancedSlicingResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_enhanced_slicing\" \"example\" {\n  alias = \"example\"\n  hash_field = \"%s\"\n  max_sessions = 4\n  protocol_fields = [{\n    gtp = {\n      flow_session = {\n        action = \"slice\"\n        skip_pkt_count = 1\n        timeout = 10\n        value = \"inner\"\n      }\n      offset = 0\n      protocol = \"gtp\"\n    }\n    gtp_u = {\n      flow_session = {\n        action = \"slice\"\n        skip_pkt_count = 1\n        timeout = 10\n        value = \"inner\"\n      }\n      l4_port = 1\n      offset = 0\n      protocol = \"gtpu-tcp\"\n    }\n    ip = {\n      flow_session = {\n        action = \"slice\"\n        skip_pkt_count = 1\n        timeout = 10\n        value = \"inner\"\n      }\n      offset = 0\n      protocol = \"ip\"\n      value = \"inner\"\n    }\n    none = {\n      flow_session = {\n        action = \"slice\"\n        skip_pkt_count = 1\n        timeout = 10\n        value = \"inner\"\n      }\n      offset = 0\n    }\n    transport = {\n      flow_session = {\n        action = \"slice\"\n        skip_pkt_count = 1\n        timeout = 10\n        value = \"inner\"\n      }\n      l4_port = 1\n      offset = 0\n      protocol = \"tcp\"\n      value = \"inner\"\n    }\n  }]\n}\n", serverURL, name)
}

// newEnhancedSlicingResourceMockServer returns an httptest server that stubs the EnhancedSlicingResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newEnhancedSlicingResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/apps/enhancedSlicing"), "/")
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
			if _, ok := body["alias"]; !ok {
				body["alias"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["alias"])
			state0[id] = body
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(201)
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
		case http.MethodPut, http.MethodPatch:
			body := make(map[string]interface{})
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil && err != io.EOF {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if _, ok := body["alias"]; !ok {
				body["alias"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["alias"])
			state0[id] = body
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			_ = json.NewEncoder(w).Encode(body)
			lastKey0 = id
			return
		case http.MethodDelete:
			delete(state0, id)
			w.WriteHeader(204)
			return
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
	mux.HandleFunc("/apps/enhancedSlicing", handler0)
	mux.HandleFunc("/apps/enhancedSlicing/", handler0)
	return httptest.NewServer(mux)
}

// TestAccEnhancedSlicingResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccEnhancedSlicingResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newEnhancedSlicingResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccEnhancedSlicingResourceConfig(server.URL, "5tuple"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_enhanced_slicing.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_enhanced_slicing.example", "hash_field", "5tuple"))}, resource.TestStep{Config: testAccEnhancedSlicingResourceConfig(server.URL, "7tuple"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_enhanced_slicing.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_enhanced_slicing.example", "hash_field", "7tuple"))}, resource.TestStep{ResourceName: "gigavuecore_enhanced_slicing.example", ImportState: true, ImportStateId: "imported-alias"}}})
}
