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

func testAccInterfaceResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_interface\" \"example\" {\n  alias = \"example\"\n  attach = [\"example\"]\n  cluster_id = \"%s\"\n  comment = \"example\"\n  gateway = \"example\"\n  gs_groups = [\"example\"]\n  hw_address = \"example\"\n  ip_address = \"example\"\n  ip_mask = \"example\"\n  mtu = 1500\n  netflow_exporters = [\"example\"]\n  tags = [{\n    tag_key = \"example\"\n    tag_values = [\"example\"]\n  }]\n}\n", serverURL, name)
}

// newInterfaceResourceMockServer returns an httptest server that stubs the InterfaceResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newInterfaceResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/ip/interfaces"), "/")
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
	mux.HandleFunc("/ip/interfaces", handler0)
	mux.HandleFunc("/ip/interfaces/", handler0)
	return httptest.NewServer(mux)
}

// TestAccInterfaceResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccInterfaceResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newInterfaceResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccInterfaceResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_interface.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_interface.example", "cluster_id", "example"))}, resource.TestStep{Config: testAccInterfaceResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_interface.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_interface.example", "cluster_id", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_interface.example", ImportState: true, ImportStateId: "imported-alias/imported-cluster_id"}}})
}
