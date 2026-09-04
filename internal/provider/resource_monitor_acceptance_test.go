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

func testAccMonitorResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_monitor\" \"example\" {\n  alias = \"example\"\n  cache = {\n    export_triggers = {\n      event = \"txnEnd\"\n      timeout_active = 1\n      timeout_inactive = 1\n    }\n    type = \"normal\"\n  }\n  cluster_id = \"%s\"\n  description = \"example\"\n  records = [\"example\"]\n  sampling = {\n    mode = \"multi-rate\"\n    single_sampling_rate = 10\n  }\n  sampling_space = 0\n  ssl_port_restrictions = {\n    ports = [0]\n    ssl_ports = \"all\"\n  }\n}\n", serverURL, name)
}

// newMonitorResourceMockServer returns an httptest server that stubs the MonitorResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newMonitorResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/apps/netflow/monitors"), "/")
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
	mux.HandleFunc("/apps/netflow/monitors", handler0)
	mux.HandleFunc("/apps/netflow/monitors/", handler0)
	return httptest.NewServer(mux)
}

// TestAccMonitorResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccMonitorResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newMonitorResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccMonitorResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_monitor.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_monitor.example", "cluster_id", "example"))}, resource.TestStep{Config: testAccMonitorResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_monitor.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_monitor.example", "cluster_id", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_monitor.example", ImportState: true, ImportStateId: "imported-alias/imported-cluster_id"}}})
}
