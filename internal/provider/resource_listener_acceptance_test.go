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

func testAccListenerResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_listener\" \"example\" {\n  alias = \"example\"\n  cluster_id = \"%s\"\n  description = \"example\"\n  gs_group_associated = [\"example\"]\n  ip_interface = [\"example\"]\n  l3 = {\n    dscp = 0\n    protocol = \"ipv4\"\n    ttl = 1\n  }\n  l4 = {\n    port_list = [1]\n    protocol = \"tcp\"\n  }\n  mode = \"promiscuous\"\n  ssl_profile = \"example\"\n  status = \"active\"\n  tags = [{\n    tag_key = \"example\"\n    tag_values = [\"example\"]\n  }]\n  tcp_profile = \"example\"\n  type = \"gtp-cups\"\n}\n", serverURL, name)
}

// newListenerResourceMockServer returns an httptest server that stubs the ListenerResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newListenerResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/apps/listener"), "/")
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
	mux.HandleFunc("/apps/listener", handler0)
	mux.HandleFunc("/apps/listener/", handler0)
	return httptest.NewServer(mux)
}

// TestAccListenerResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccListenerResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newListenerResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccListenerResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_listener.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_listener.example", "cluster_id", "example"))}, resource.TestStep{Config: testAccListenerResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_listener.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_listener.example", "cluster_id", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_listener.example", ImportState: true, ImportStateId: "imported-alias/imported-cluster_id"}}})
}
