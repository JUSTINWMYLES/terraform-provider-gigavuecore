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

func testAccAppVisibilityResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_app_visibility\" \"example\" {\n  app_export_config = {\n    cache_config = {\n      advance_hash = true\n      alias = \"example\"\n      description = \"example\"\n      dpi_inject_limit = 0\n      event = \"txnEnd\"\n      exporters = [\"example\"]\n      flow_behavior = \"unidir\"\n      match = {\n        datalink = {\n          mac_dst = true\n          mac_src = true\n          vlan = true\n        }\n        interface = {\n          in_name_width = 1\n          in_physical_width = 2\n        }\n        ip = {\n          version = true\n        }\n        ipv4 = {\n          destination = {\n            prefix_min_mask = \"example\"\n          }\n          dscp = true\n          fragmentation = {\n            flags = true\n            offset = true\n          }\n          header_len = true\n          option_map = true\n          precedence = true\n          protocol = true\n          section = {\n            header_size = 1\n            payload_size = 1\n          }\n          source = {\n            prefix_min_mask = \"example\"\n          }\n          tos = true\n          total_length = true\n          ttl = true\n        }\n        ipv6 = {\n          destination = {\n            prefix_min_mask = \"example\"\n          }\n          dscp = true\n          extension_map = true\n          flow_label = true\n          fragmentation = {\n            flags = true\n            offset = true\n          }\n          hop_limit = true\n          length = {\n            header = true\n            payload = true\n            total = true\n          }\n          next_header = true\n          precedence = true\n          section = {\n            header_size = 1\n            payload_size = 1\n          }\n          source = {\n            prefix_min_mask = \"example\"\n          }\n          traffic_class = true\n        }\n        transport = {\n          dst_port = true\n          icmp = {\n            ipv4_code = true\n            ipv4_type = true\n            ipv6_code = true\n            ipv6_type = true\n          }\n          src_port = true\n          tcp = {\n            ack_number = true\n            dst_port = true\n            flags = true\n            header_len = true\n            seq_number = true\n            src_port = true\n            urgent_ptr = true\n            window_size = true\n          }\n          udp = {\n            dst_port = true\n            msg_len = true\n            src_port = true\n          }\n        }\n      }\n      multi_collect = true\n      network_profiles = [\"example\"]\n      observation_domain_id = 0\n      sampling = {\n        mode = \"multiRate\"\n        single_sampling_rate = 10\n      }\n      size = {\n        flows = 1\n      }\n      timeout = {\n        idle = 1\n      }\n    }\n    destination_configs = [{\n      application_names = [{\n        attributes = [{\n          name = \"example\"\n          value = \"example\"\n        }]\n        is_user_defined = true\n        name = \"example\"\n      }]\n      destination_name = \"example\"\n      export_meta_app_profile = {\n        alias = \"example\"\n        application_id = true\n        applications = [{\n          attributes = [{\n            name = \"example\"\n            value = \"example\"\n          }]\n          is_user_defined = true\n          name = \"example\"\n        }]\n        counter = {\n          bytes = true\n          bytes_long = true\n          inner_byte = true\n          inner_byte_long = true\n          packets = true\n          packets_long = true\n        }\n        datalink = {\n          mac_dst = true\n          mac_src = true\n          vlan = true\n        }\n        description = \"example\"\n        flow = {\n          end_reason = true\n        }\n        gtpu = {\n          qfi = true\n          teid = true\n        }\n        interface = {\n          in_name_width = 1\n          in_physical_width = 2\n          out_physical_width = 2\n        }\n        ip = {\n          version = true\n        }\n        ipv4 = {\n          destination = {\n            prefix_min_mask = \"example\"\n          }\n          dscp = true\n          fragmentation = {\n            flags = true\n            offset = true\n          }\n          header_len = true\n          option_map = true\n          precedence = true\n          protocol = true\n          section = {\n            header_size = 1\n            payload_size = 1\n          }\n          source = {\n            prefix_min_mask = \"example\"\n          }\n          tos = true\n          total_length = true\n          ttl = true\n        }\n        ipv6 = {\n          destination = {\n            prefix_min_mask = \"example\"\n          }\n          dscp = true\n          extension_map = true\n          flow_label = true\n          fragmentation = {\n            flags = true\n            offset = true\n          }\n          hop_limit = true\n          length = {\n            header = true\n            payload = true\n            total = true\n          }\n          next_header = true\n          precedence = true\n          section = {\n            header_size = 1\n            payload_size = 1\n          }\n          source = {\n            prefix_min_mask = \"example\"\n          }\n          traffic_class = true\n        }\n        outer_ipv4 = {\n          destination = true\n          source = true\n        }\n        outer_ipv6 = {\n          destination = true\n          source = true\n        }\n        timestamp = {\n          flow_end_msec = true\n          flow_endsec = true\n          flow_start_msec = true\n          flow_startsec = true\n          sys_up_time_first = true\n          sys_up_time_last = true\n        }\n        transport = {\n          dst_port = true\n          icmp = {\n            ipv4_code = true\n            ipv4_type = true\n            ipv6_code = true\n            ipv6_type = true\n          }\n          src_port = true\n          tcp = {\n            ack_number = true\n            dst_port = true\n            flags = true\n            header_len = true\n            seq_number = true\n            src_port = true\n            urgent_ptr = true\n            window_size = true\n          }\n          udp = {\n            dst_port = true\n            msg_len = true\n            src_port = true\n          }\n        }\n        type = \"export\"\n      }\n      export_meta_app_profile_alias = \"example\"\n      exporter_alias = \"example\"\n      exporter_config = {\n        alias = \"example\"\n        application_profiles = [\"example\"]\n        cef = {\n          active_timeout = 1\n          inactive_timeout = 1\n        }\n        description = \"example\"\n        destination = {\n          dscp = 0\n          ipv4_address = \"example\"\n          l4_port_dst = 1\n          l4_port_src = 1\n          l4_protocol = \"udp\"\n          ttl = 1\n        }\n        max_pkt_size = 0\n        mobility_sam = {\n          encoding = \"example\"\n          encoding_format = \"hierarchy\"\n          event_enable = {\n            modify = true\n            update = true\n          }\n          trigger = \"example\"\n        }\n        monitor = {\n          timeout = 60\n        }\n        netflow = {\n          active_timeout = 1\n          inactive_timeout = 1\n          template_refresh = 1\n          template_type = \"cohesive\"\n          version = \"v5\"\n        }\n        snmp = {\n          enabled = true\n        }\n        source = {\n          ip_interface = \"example\"\n        }\n        type = \"cef\"\n      }\n      iface = \"example\"\n    }]\n  }\n  app_filter_config = {\n    egress_traffic_configs = [{\n      drop_application_names = [{\n        attributes = [{\n          name = \"example\"\n          value = \"example\"\n        }]\n        is_user_defined = true\n        name = \"example\"\n      }]\n      pass_application_names = [{\n        attributes = [{\n          name = \"example\"\n          value = \"example\"\n        }]\n        is_user_defined = true\n        name = \"example\"\n      }]\n      tunnel_aliases = [\"example\"]\n    }]\n  }\n  conn_id = \"%s\"\n  distribute_traffic = true\n  dynamic_scale_unit = true\n  env_id = \"example\"\n  ingress_traffic_configs = [{\n    source_selector_aliases = [\"example\"]\n    src_raw_end_point_interfaces = [\"example\"]\n    tunnel_aliases = [\"example\"]\n    tunnel_interface_mappings = [{\n      iface = \"example\"\n      tunnel_alias = \"example\"\n    }]\n  }]\n  monitor_solution_config = {\n    destination_config = {\n      iface = \"example\"\n    }\n    mgmt_interface = \"internal\"\n  }\n  scale_unit = 0\n  solution_alias = \"example\"\n  solution_desc = \"example\"\n}\n", serverURL, name)
}

// newAppVisibilityResourceMockServer returns an httptest server that stubs the AppVisibilityResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newAppVisibilityResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/intent/appVisibility"), "/")
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
			if _, ok := body["solutionAlias"]; !ok {
				body["solutionAlias"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["solutionAlias"])
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
			if _, ok := body["solutionAlias"]; !ok {
				body["solutionAlias"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["solutionAlias"])
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
	mux.HandleFunc("/intent/appVisibility", handler0)
	mux.HandleFunc("/intent/appVisibility/", handler0)
	return httptest.NewServer(mux)
}

// TestAccAppVisibilityResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccAppVisibilityResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newAppVisibilityResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccAppVisibilityResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_app_visibility.example", "solution_alias"), resource.TestCheckResourceAttr("gigavuecore_app_visibility.example", "conn_id", "example"))}, resource.TestStep{Config: testAccAppVisibilityResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_app_visibility.example", "solution_alias"), resource.TestCheckResourceAttr("gigavuecore_app_visibility.example", "conn_id", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_app_visibility.example", ImportState: true, ImportStateId: "imported-solution_alias"}}})
}
