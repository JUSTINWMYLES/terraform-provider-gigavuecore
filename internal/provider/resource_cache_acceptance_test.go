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

func testAccCacheResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_cache\" \"example\" {\n  advance_hash = true\n  alias = \"example\"\n  description = \"%s\"\n  dpi_inject_limit = 0\n  event = \"txnEnd\"\n  exporters = [\"example\"]\n  flow_behavior = \"unidir\"\n  match = {\n    datalink = {\n      mac_dst = true\n      mac_src = true\n      vlan = true\n    }\n    interface = {\n      in_name_width = 1\n      in_physical_width = 2\n    }\n    ip = {\n      version = true\n    }\n    ipv4 = {\n      destination = {\n        prefix_min_mask = \"example\"\n      }\n      dscp = true\n      fragmentation = {\n        flags = true\n        offset = true\n      }\n      header_len = true\n      option_map = true\n      precedence = true\n      protocol = true\n      section = {\n        header_size = 1\n        payload_size = 1\n      }\n      source = {\n        prefix_min_mask = \"example\"\n      }\n      tos = true\n      total_length = true\n      ttl = true\n    }\n    ipv6 = {\n      destination = {\n        prefix_min_mask = \"example\"\n      }\n      dscp = true\n      extension_map = true\n      flow_label = true\n      fragmentation = {\n        flags = true\n        offset = true\n      }\n      hop_limit = true\n      length = {\n        header = true\n        payload = true\n        total = true\n      }\n      next_header = true\n      precedence = true\n      section = {\n        header_size = 1\n        payload_size = 1\n      }\n      source = {\n        prefix_min_mask = \"example\"\n      }\n      traffic_class = true\n    }\n    transport = {\n      dst_port = true\n      icmp = {\n        ipv4_code = true\n        ipv4_type = true\n        ipv6_code = true\n        ipv6_type = true\n      }\n      src_port = true\n      tcp = {\n        ack_number = true\n        dst_port = true\n        flags = true\n        header_len = true\n        seq_number = true\n        src_port = true\n        urgent_ptr = true\n        window_size = true\n      }\n      udp = {\n        dst_port = true\n        msg_len = true\n        src_port = true\n      }\n    }\n  }\n  multi_collect = true\n  network_profiles = [\"example\"]\n  observation_domain_id = 0\n  sampling = {\n    mode = \"multiRate\"\n    single_sampling_rate = 10\n  }\n  size = {\n    flows = 1\n  }\n  timeout = {\n    idle = 1\n  }\n}\n", serverURL, name)
}

// newCacheResourceMockServer returns an httptest server that stubs the CacheResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newCacheResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/apps/metadata/cache"), "/")
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
	mux.HandleFunc("/apps/metadata/cache", handler0)
	mux.HandleFunc("/apps/metadata/cache/", handler0)
	return httptest.NewServer(mux)
}

// TestAccCacheResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccCacheResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newCacheResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccCacheResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_cache.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_cache.example", "description", "example"))}, resource.TestStep{Config: testAccCacheResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_cache.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_cache.example", "description", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_cache.example", ImportState: true, ImportStateId: "imported-alias"}}})
}
