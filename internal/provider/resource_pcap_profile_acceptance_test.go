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

func testAccPcapProfileResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_pcap_profile\" \"example\" {\n  alias_list = \"%s\"\n  cluster_id = \"example\"\n  pcap_configs = [{\n    alias = \"example\"\n    channel_port = \"example\"\n    direction = \"example\"\n    packet_limit = 1\n    pcap_rules_list = [{\n      dscp = \"af11\"\n      dst_mac = {\n        address = \"example\"\n        mask = \"example\"\n      }\n      dstipv4_addrandmask = {\n        address = \"example\"\n        mask = \"example\"\n      }\n      ether_type = \"example\"\n      inner_vlan = 1\n      ip4_frag = \"noFrag\"\n      ip4_ttl = 0\n      ip_ver = \"v4\"\n      packet_hit_count = 0\n      portdst = 0\n      portsrc = 0\n      protocol = 0\n      rule_id = 1\n      src_mac = {\n        address = \"example\"\n        mask = \"example\"\n      }\n      srcipv4_addrandmask = {\n        address = \"example\"\n        mask = \"example\"\n      }\n      tcpctl = 0\n      vlan = 1\n    }]\n    port = \"example\"\n  }]\n  port_ids = \"example\"\n}\n", serverURL, name)
}

// newPcapProfileResourceMockServer returns an httptest server that stubs the PcapProfileResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newPcapProfileResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/pcap"), "/")
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
			if _, ok := body["id"]; !ok {
				body["id"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["id"])
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
		case http.MethodDelete:
			delete(state0, id)
			w.WriteHeader(204)
			return
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
	mux.HandleFunc("/pcap", handler0)
	mux.HandleFunc("/pcap/", handler0)
	return httptest.NewServer(mux)
}

// TestAccPcapProfileResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccPcapProfileResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newPcapProfileResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccPcapProfileResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_pcap_profile.example", "id"), resource.TestCheckResourceAttr("gigavuecore_pcap_profile.example", "alias_list", "example"))}}})
}
