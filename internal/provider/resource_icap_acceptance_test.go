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

func testAccIcapResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_icap\" \"example\" {\n  alias = \"example\"\n  cluster_id = \"%s\"\n  config_status = \"SUCCESS\"\n  config_status_reasons = [\"example\"]\n  gs_engines = [\"example\"]\n  gs_grp_alias = \"example\"\n  gsop_alias = \"example\"\n  health_state = \"green\"\n  health_state_reasons = [{\n    message = \"example\"\n    severity = \"green\"\n    traffic_health_state_computation_type = \"PORT_LOW_UTIL\"\n  }]\n  icap_map_alias = \"example\"\n  icap_profile_config = {\n    alias = \"example\"\n    cluster_id = \"example\"\n    exceed_action = \"drop\"\n    http_req_buf = 10\n    inactivity_timeout = 2\n    preview = 0\n    resp_mod = \"enable\"\n    resp_timeout = 5\n    resp_timeout_action = \"drop\"\n    server_group = \"example\"\n    src_max_l4_port = 10000\n    src_min_l4_port = 10000\n  }\n  icap_server_grp_alias = \"example\"\n  icap_servers = [{\n    alias = \"example\"\n    cluster_id = \"example\"\n    comment = \"example\"\n    l3_address = \"example\"\n    l4_port = 0\n    options_service_url = \"example\"\n    reqmod_service_url = \"example\"\n    respmod_service_url = \"example\"\n  }]\n  ing_alias = \"example\"\n  inline_networks = [\"example\"]\n  ip_interface_alias = \"example\"\n}\n", serverURL, name)
}

// newIcapResourceMockServer returns an httptest server that stubs the IcapResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newIcapResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/icap"), "/")
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
	mux.HandleFunc("/icap", handler0)
	mux.HandleFunc("/icap/", handler0)
	return httptest.NewServer(mux)
}

// TestAccIcapResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccIcapResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newIcapResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccIcapResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_icap.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_icap.example", "cluster_id", "example"))}, resource.TestStep{Config: testAccIcapResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_icap.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_icap.example", "cluster_id", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_icap.example", ImportState: true, ImportStateId: "imported-alias"}}})
}
