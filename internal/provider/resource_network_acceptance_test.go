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

func testAccNetworkResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_network\" \"example\" {\n  alias = \"example\"\n  cluster_id = \"%s\"\n  comment = \"example\"\n  forwarding_state = \"physicalBypass\"\n  health_state = \"green\"\n  health_state_reasons = [{\n    message = \"example\"\n    severity = \"green\"\n    traffic_health_state_computation_type = \"PORT_LOW_UTIL\"\n  }]\n  heartbeat = {\n    enabled = true\n  }\n  lfp = true\n  physical_bypass = true\n  port_a = \"example\"\n  port_b = \"example\"\n  redundancy_control_state = \"neutral\"\n  redundancy_profile = \"example\"\n  traffic_path = \"drop\"\n  type = \"protected\"\n}\n", serverURL, name)
}

// newNetworkResourceMockServer returns an httptest server that stubs the NetworkResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newNetworkResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/inline/networks"), "/")
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
	mux.HandleFunc("/inline/networks", handler0)
	mux.HandleFunc("/inline/networks/", handler0)
	return httptest.NewServer(mux)
}

// TestAccNetworkResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccNetworkResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newNetworkResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccNetworkResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_network.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_network.example", "cluster_id", "example"))}, resource.TestStep{Config: testAccNetworkResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_network.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_network.example", "cluster_id", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_network.example", ImportState: true, ImportStateId: "imported-alias/imported-cluster_id"}}})
}
