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

func testAccProxyServerProfileResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_proxy_server_profile\" \"example\" {\n  alias = \"example\"\n  auth_type = \"%s\"\n  comment = \"example\"\n  password = \"example\"\n  periodic_ping = \"enable\"\n  periodic_ping_failure_retry = 1\n  periodic_ping_interval = 1\n  periodic_ping_type = \"http-connect\"\n  port = 1\n  protocol = \"http\"\n  proxy_address = \"example\"\n  ssl_apps = {\n    cluster_name = [\"example\"]\n  }\n  username = \"example\"\n}\n", serverURL, name)
}

// newProxyServerProfileResourceMockServer returns an httptest server that stubs the ProxyServerProfileResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newProxyServerProfileResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/apps/proxyServer/profiles"), "/")
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
	mux.HandleFunc("/apps/proxyServer/profiles", handler0)
	mux.HandleFunc("/apps/proxyServer/profiles/", handler0)
	return httptest.NewServer(mux)
}

// TestAccProxyServerProfileResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccProxyServerProfileResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newProxyServerProfileResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccProxyServerProfileResourceConfig(server.URL, "none"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_proxy_server_profile.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_proxy_server_profile.example", "auth_type", "none"))}, resource.TestStep{Config: testAccProxyServerProfileResourceConfig(server.URL, "basic"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_proxy_server_profile.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_proxy_server_profile.example", "auth_type", "basic"))}, resource.TestStep{ResourceName: "gigavuecore_proxy_server_profile.example", ImportState: true, ImportStateId: "imported-alias"}}})
}
