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

func testAccMetadataExporterResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_metadata_exporter\" \"example\" {\n  alias = \"example\"\n  application_profiles = [\"example\"]\n  cef = {\n    active_timeout = 1\n    inactive_timeout = 1\n  }\n  description = \"%s\"\n  destination = {\n    dscp = 0\n    ipv4_address = \"example\"\n    l4_port_dst = 1\n    l4_port_src = 1\n    l4_protocol = \"udp\"\n    ttl = 1\n  }\n  max_pkt_size = 0\n  mobility_sam = {\n    encoding = \"example\"\n    encoding_format = \"hierarchy\"\n    event_enable = {\n      modify = true\n      update = true\n    }\n    trigger = \"example\"\n  }\n  monitor = {\n    timeout = 60\n  }\n  netflow = {\n    active_timeout = 1\n    inactive_timeout = 1\n    template_refresh = 1\n    template_type = \"cohesive\"\n    version = \"v5\"\n  }\n  snmp = {\n    enabled = true\n  }\n  source = {\n    ip_interface = \"example\"\n  }\n  type = \"cef\"\n}\n", serverURL, name)
}

// newMetadataExporterResourceMockServer returns an httptest server that stubs the MetadataExporterResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newMetadataExporterResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/apps/metadata/exporters"), "/")
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
	mux.HandleFunc("/apps/metadata/exporters", handler0)
	mux.HandleFunc("/apps/metadata/exporters/", handler0)
	return httptest.NewServer(mux)
}

// TestAccMetadataExporterResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccMetadataExporterResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newMetadataExporterResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccMetadataExporterResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_metadata_exporter.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_metadata_exporter.example", "description", "example"))}, resource.TestStep{Config: testAccMetadataExporterResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_metadata_exporter.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_metadata_exporter.example", "description", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_metadata_exporter.example", ImportState: true, ImportStateId: "imported-alias"}}})
}
