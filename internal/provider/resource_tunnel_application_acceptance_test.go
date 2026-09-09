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

func testAccTunnelApplicationResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_tunnel_application\" \"example\" {\n  alias = \"example\"\n  cluster_id = \"%s\"\n  config_status = \"example\"\n  config_status_reasons = [\"example\"]\n  decap = null\n  description = \"example\"\n  dscp = 0\n  encap = {\n    additonalgsops = {\n      apf = {\n        enabled = \"enabled\"\n      }\n      dedup = {\n        enabled = \"enabled\"\n      }\n      diameter_whitelist = {\n        enabled = \"enabled\"\n      }\n      flow_filter = {\n        type = \"gtp\"\n      }\n      flow_sampling = {\n        type = \"ip\"\n      }\n      gseries_header_add = {\n        types = [\"srcid\"]\n      }\n      gseries_header_remove = {\n        enabled = \"enabled\"\n      }\n      gseries_load_balance = {\n        fixed_offset = {\n          hash = \"checksum\"\n          length = 1\n          offset = 0\n        }\n        variable_offset = {\n          end_delim = \"example\"\n          hash = \"checksum\"\n          start_delim = \"example\"\n          start_field = \"example\"\n        }\n      }\n      gseries_pattern_match = {\n        fixed_offset = {\n          length = 1\n          offset = 0\n        }\n        variable_offset = {\n          end_delim = \"example\"\n          start_delim = \"example\"\n        }\n      }\n      gtp_whitelist = {\n        enabled = \"enabled\"\n      }\n      header_add = {\n        vlan = 0\n      }\n      header_remove = {\n        ah1 = \"none\"\n        ah2 = \"none\"\n        custom_len = 1\n        erspan_flow_id = 0\n        fp_dst_switch_id = 0\n        fp_src_switch_id = 0\n        header_count = 1\n        offset = \"start\"\n        offset_range_value = 0\n        protocol = \"gtp\"\n        timestamp_format = \"gigasmart\"\n        vlan_header = \"all\"\n        vxlan_id = 0\n      }\n      icap = {\n        icap_profile = \"example\"\n      }\n      inline_ssl = {\n        inline_ssl_profile = \"example\"\n      }\n      load_balance = {\n        enhanced = {\n          elb_alias = \"example\"\n        }\n        stateful = {\n          app_type = \"gtp\"\n          diameter_key_hash_type = \"sessionId\"\n          diameter_key_multi_hash_type = [{\n            avp_codevalue = 0\n            key = \"sessionId\"\n          }]\n          gtp_key_hash_type = \"imsi\"\n          lb_type = \"leastBw\"\n          sip_key_hash_type = \"callerId\"\n        }\n        stateless = {\n          field_location = \"inner\"\n          hash_fields = \"ipOnly\"\n        }\n      }\n      masking = {\n        content_type = \"message_cpim\"\n        length = 1\n        offset = 0\n        pattern = \"a1\"\n        protocol = \"none\"\n      }\n      metadata_export = {\n        cache = \"example\"\n      }\n      netflow = {\n        enabled = \"enabled\"\n      }\n      sa_apf = {\n        enabled = \"enabled\"\n      }\n      sip_whitelist = {\n        enabled = \"enabled\"\n      }\n      slicing = {\n        enhanced = \"example\"\n        offset = 4\n        protocol = \"none\"\n      }\n      ssl_decrypt = {\n        in_port = 0\n        out_port = 0\n      }\n      trailer_add = {\n        types = [\"crc\"]\n      }\n      trailer_remove = {\n        enabled = \"enabled\"\n      }\n      tunnel_decap = {\n        custom = {\n          port_dst = 0\n          port_src = 0\n        }\n        erspan_flow_id = 0\n        gmip_port = 0\n        l2_gre_key = 0\n        tls_pcapng = {\n          decap_key = \"example\"\n          listener = \"example\"\n        }\n        type = \"gmip\"\n        vxlan = {\n          port_dst = 1\n          port_src = 0\n          vni = 0\n        }\n      }\n      tunnel_encap = {\n        gmip_config = {\n          dscp = 0\n          dst_ip = \"example\"\n          dst_port = 0\n          flow_label = 0\n          prec = 0\n          src_port = 0\n          ttl = 1\n        }\n        l2_gre_config = {\n          dscp = 0\n          dst_ip = \"example\"\n          flow_label = 0\n          key = 0\n          pg_dst = \"example\"\n          prec = 0\n          session_field = \"fiveTupleIpv4\"\n          session_pos = \"inner\"\n          ttl = 1\n        }\n        tls_pcapng = {\n          exporter = \"example\"\n          exporter_group = \"example\"\n        }\n        type = \"gmip\"\n        vxlan_config = {\n          dscp = 0\n          dst_ip = \"example\"\n          dst_port = 4789\n          src_port = 0\n          ttl = 1\n          vni = 1\n        }\n      }\n    }\n    export_configs = [{\n      exporter_alias = \"example\"\n      remote_application_port = 0\n      remote_ip = \"example\"\n      source_application_port = 0\n    }]\n    exporter_group_alias = \"example\"\n    gsop_alias = \"example\"\n    map_alias = \"example\"\n    rules = {\n      drop_rules = null\n      pass_rules = null\n    }\n    source_port = [\"example\"]\n  }\n  gsgroup = \"example\"\n  ip_interface_in = \"example\"\n  ip_interface_out = \"example\"\n  local_key_alias = \"example\"\n  remote_key_alias = \"example\"\n  ssl_profile = {\n    cipher = \"example\"\n    mtls = \"enable\"\n    version = \"example\"\n  }\n  ssl_profile_alias = \"example\"\n  tcp_profile = {\n    keep_alive_timer = 30\n    selective_ack = \"enable\"\n    syn_retries = 1\n  }\n  tcp_profile_alias = \"example\"\n  traffic_dir = \"IN\"\n  ttl = 1\n  tunnel_type = \"example\"\n}\n", serverURL, name)
}

// newTunnelApplicationResourceMockServer returns an httptest server that stubs the TunnelApplicationResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newTunnelApplicationResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/apps/tunnelApplication"), "/")
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
	mux.HandleFunc("/apps/tunnelApplication", handler0)
	mux.HandleFunc("/apps/tunnelApplication/", handler0)
	return httptest.NewServer(mux)
}

// TestAccTunnelApplicationResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccTunnelApplicationResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newTunnelApplicationResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccTunnelApplicationResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_tunnel_application.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_tunnel_application.example", "cluster_id", "example"))}, resource.TestStep{Config: testAccTunnelApplicationResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_tunnel_application.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_tunnel_application.example", "cluster_id", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_tunnel_application.example", ImportState: true, ImportStateId: "imported-alias/imported-cluster_id"}}})
}
