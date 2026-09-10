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

func testAccInlineSslProfileResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_inline_ssl_profile\" \"example\" {\n  alias = \"example\"\n  certificate = {\n    expired = \"decrypt\"\n    invalid = \"decrypt\"\n    revocation = {\n      crl = {\n        defer = 20\n        enabled = true\n        fail = \"soft\"\n      }\n      ocsp = {\n        defer = 20\n        enabled = true\n        fail = \"soft\"\n      }\n    }\n    self_signed = \"decrypt\"\n    unknown_ca = \"decrypt\"\n  }\n  cluster_id = \"%s\"\n  decrypt = {\n    tcp = {\n      inactive_timeout = 2\n      port_map = {\n        default_out_port = 0\n        ports = [{\n          in_port = 1\n          out_port = 1\n          rule_id = 0\n        }]\n      }\n    }\n    tool_bypass = {\n      enable = true\n    }\n  }\n  default_action = \"decrypt\"\n  high_avail = {\n    active_standby = {\n      enable = true\n    }\n  }\n  inbound_tool_early_inspect = {\n    connection_timeout = 1\n    mode = {\n      enable = true\n    }\n  }\n  key_map = [{\n    hostname = \"example\"\n    key = \"example\"\n    rule_id = 0\n  }]\n  monitor = \"enable\"\n  network_group = {\n    multiple_entry = {\n      enable = true\n    }\n  }\n  no_decrypt = {\n    tool_bypass = {\n      enable = true\n    }\n  }\n  non_ssl_tcp = {\n    tool_bypass = {\n      enable = true\n    }\n  }\n  one_arm = \"enable\"\n  resilient_inline = {\n    mode = {\n      enable = true\n    }\n  }\n  rules = null\n  split_proxy = {\n    mode = {\n      enable = true\n    }\n    server_non_pfs_ciphers = {\n      enable = true\n    }\n  }\n  start_tls = {\n    l4_port = [0]\n  }\n  tcp = {\n    delayed_ack = true\n    syn_retries = 0\n    timewait_timeout = 0\n  }\n  tool = {\n    early_engage = true\n    fail_action = \"fail-open\"\n  }\n  tool_l3 = {\n    cache_server_cert_timeout = 1\n    http2_downgrade = {\n      enable = true\n    }\n    nat_pat = {\n      enable = true\n    }\n  }\n  url_cache = {\n    miss_action = \"decrypt\"\n    timeout = 1\n  }\n}\n", serverURL, name)
}

// newInlineSslProfileResourceMockServer returns an httptest server that stubs the InlineSslProfileResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newInlineSslProfileResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/apps/inlineSsl/profiles"), "/")
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
	mux.HandleFunc("/apps/inlineSsl/profiles", handler0)
	mux.HandleFunc("/apps/inlineSsl/profiles/", handler0)
	return httptest.NewServer(mux)
}

// TestAccInlineSslProfileResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccInlineSslProfileResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newInlineSslProfileResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccInlineSslProfileResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_inline_ssl_profile.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_inline_ssl_profile.example", "cluster_id", "example"))}, resource.TestStep{Config: testAccInlineSslProfileResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_inline_ssl_profile.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_inline_ssl_profile.example", "cluster_id", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_inline_ssl_profile.example", ImportState: true, ImportStateId: "imported-alias:imported-cluster_id"}}})
}
