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

func testAccApplicationProfileResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_application_profile\" \"example\" {\n  alias = \"example\"\n  application_id = true\n  applications = [{\n    attributes = [{\n      name = \"example\"\n      value = \"example\"\n    }]\n    is_user_defined = true\n    name = \"example\"\n  }]\n  counter = {\n    bytes = true\n    bytes_long = true\n    inner_byte = true\n    inner_byte_long = true\n    packets = true\n    packets_long = true\n  }\n  datalink = {\n    mac_dst = true\n    mac_src = true\n    vlan = true\n  }\n  description = \"%s\"\n  flow = {\n    end_reason = true\n  }\n  gtpu = {\n    qfi = true\n    teid = true\n  }\n  interface = {\n    in_name_width = 1\n    in_physical_width = 2\n    out_physical_width = 2\n  }\n  ip = {\n    version = true\n  }\n  ipv4 = {\n    destination = {\n      prefix_min_mask = \"example\"\n    }\n    dscp = true\n    fragmentation = {\n      flags = true\n      offset = true\n    }\n    header_len = true\n    option_map = true\n    precedence = true\n    protocol = true\n    section = {\n      header_size = 1\n      payload_size = 1\n    }\n    source = {\n      prefix_min_mask = \"example\"\n    }\n    tos = true\n    total_length = true\n    ttl = true\n  }\n  ipv6 = {\n    destination = {\n      prefix_min_mask = \"example\"\n    }\n    dscp = true\n    extension_map = true\n    flow_label = true\n    fragmentation = {\n      flags = true\n      offset = true\n    }\n    hop_limit = true\n    length = {\n      header = true\n      payload = true\n      total = true\n    }\n    next_header = true\n    precedence = true\n    section = {\n      header_size = 1\n      payload_size = 1\n    }\n    source = {\n      prefix_min_mask = \"example\"\n    }\n    traffic_class = true\n  }\n  outer_ipv4 = {\n    destination = true\n    source = true\n  }\n  outer_ipv6 = {\n    destination = true\n    source = true\n  }\n  timestamp = {\n    flow_end_msec = true\n    flow_endsec = true\n    flow_start_msec = true\n    flow_startsec = true\n    sys_up_time_first = true\n    sys_up_time_last = true\n  }\n  transport = {\n    dst_port = true\n    icmp = {\n      ipv4_code = true\n      ipv4_type = true\n      ipv6_code = true\n      ipv6_type = true\n    }\n    src_port = true\n    tcp = {\n      ack_number = true\n      dst_port = true\n      flags = true\n      header_len = true\n      seq_number = true\n      src_port = true\n      urgent_ptr = true\n      window_size = true\n    }\n    udp = {\n      dst_port = true\n      msg_len = true\n      src_port = true\n    }\n  }\n  type = \"export\"\n}\n", serverURL, name)
}

// newApplicationProfileResourceMockServer returns an httptest server that stubs the ApplicationProfileResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newApplicationProfileResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/apps/metadata/applicationProfiles"), "/")
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
	mux.HandleFunc("/apps/metadata/applicationProfiles", handler0)
	mux.HandleFunc("/apps/metadata/applicationProfiles/", handler0)
	return httptest.NewServer(mux)
}

// TestAccApplicationProfileResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccApplicationProfileResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newApplicationProfileResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccApplicationProfileResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_application_profile.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_application_profile.example", "description", "example"))}, resource.TestStep{Config: testAccApplicationProfileResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_application_profile.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_application_profile.example", "description", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_application_profile.example", ImportState: true, ImportStateId: "imported-alias"}}})
}
