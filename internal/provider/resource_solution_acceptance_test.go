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

func testAccSolutionResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_solution\" \"example\" {\n  app_export_config = {\n    cache_config = {\n      advance_hash = true\n      alias = \"example\"\n      description = \"example\"\n      dpi_inject_limit = 0\n      event = \"txnEnd\"\n      exporters = [\"example\"]\n      flow_behavior = \"unidir\"\n      match = {\n        datalink = {\n          mac_dst = true\n          mac_src = true\n          vlan = true\n        }\n        interface = {\n          in_name_width = 1\n          in_physical_width = 2\n        }\n        ip = {\n          version = true\n        }\n        ipv4 = {\n          destination = {\n            prefix_min_mask = \"example\"\n          }\n          dscp = true\n          fragmentation = {\n            flags = true\n            offset = true\n          }\n          header_len = true\n          option_map = true\n          precedence = true\n          protocol = true\n          section = {\n            header_size = 1\n            payload_size = 1\n          }\n          source = {\n            prefix_min_mask = \"example\"\n          }\n          tos = true\n          total_length = true\n          ttl = true\n        }\n        ipv6 = {\n          destination = {\n            prefix_min_mask = \"example\"\n          }\n          dscp = true\n          extension_map = true\n          flow_label = true\n          fragmentation = {\n            flags = true\n            offset = true\n          }\n          hop_limit = true\n          length = {\n            header = true\n            payload = true\n            total = true\n          }\n          next_header = true\n          precedence = true\n          section = {\n            header_size = 1\n            payload_size = 1\n          }\n          source = {\n            prefix_min_mask = \"example\"\n          }\n          traffic_class = true\n        }\n        transport = {\n          dst_port = true\n          icmp = {\n            ipv4_code = true\n            ipv4_type = true\n            ipv6_code = true\n            ipv6_type = true\n          }\n          src_port = true\n          tcp = {\n            ack_number = true\n            dst_port = true\n            flags = true\n            header_len = true\n            seq_number = true\n            src_port = true\n            urgent_ptr = true\n            window_size = true\n          }\n          udp = {\n            dst_port = true\n            msg_len = true\n            src_port = true\n          }\n        }\n      }\n      multi_collect = true\n      network_profiles = [\"example\"]\n      observation_domain_id = 0\n      sampling = {\n        mode = \"multiRate\"\n        single_sampling_rate = 10\n      }\n      size = {\n        flows = 1\n      }\n      timeout = {\n        idle = 1\n      }\n    }\n    cache_config_alias = \"example\"\n    destination_configs = [{\n      application_names = [{\n        attributes = [{\n          name = \"example\"\n          value = \"example\"\n        }]\n        is_user_defined = true\n        name = \"example\"\n      }]\n      destination_name = \"example\"\n      export_ip_interface = \"example\"\n      export_meta_app_profile = {\n        alias = \"example\"\n        application_id = true\n        applications = [{\n          attributes = [{\n            name = \"example\"\n            value = \"example\"\n          }]\n          is_user_defined = true\n          name = \"example\"\n        }]\n        counter = {\n          bytes = true\n          bytes_long = true\n          inner_byte = true\n          inner_byte_long = true\n          packets = true\n          packets_long = true\n        }\n        datalink = {\n          mac_dst = true\n          mac_src = true\n          vlan = true\n        }\n        description = \"example\"\n        flow = {\n          end_reason = true\n        }\n        gtpu = {\n          qfi = true\n          teid = true\n        }\n        interface = {\n          in_name_width = 1\n          in_physical_width = 2\n          out_physical_width = 2\n        }\n        ip = {\n          version = true\n        }\n        ipv4 = {\n          destination = {\n            prefix_min_mask = \"example\"\n          }\n          dscp = true\n          fragmentation = {\n            flags = true\n            offset = true\n          }\n          header_len = true\n          option_map = true\n          precedence = true\n          protocol = true\n          section = {\n            header_size = 1\n            payload_size = 1\n          }\n          source = {\n            prefix_min_mask = \"example\"\n          }\n          tos = true\n          total_length = true\n          ttl = true\n        }\n        ipv6 = {\n          destination = {\n            prefix_min_mask = \"example\"\n          }\n          dscp = true\n          extension_map = true\n          flow_label = true\n          fragmentation = {\n            flags = true\n            offset = true\n          }\n          hop_limit = true\n          length = {\n            header = true\n            payload = true\n            total = true\n          }\n          next_header = true\n          precedence = true\n          section = {\n            header_size = 1\n            payload_size = 1\n          }\n          source = {\n            prefix_min_mask = \"example\"\n          }\n          traffic_class = true\n        }\n        outer_ipv4 = {\n          destination = true\n          source = true\n        }\n        outer_ipv6 = {\n          destination = true\n          source = true\n        }\n        timestamp = {\n          flow_end_msec = true\n          flow_endsec = true\n          flow_start_msec = true\n          flow_startsec = true\n          sys_up_time_first = true\n          sys_up_time_last = true\n        }\n        transport = {\n          dst_port = true\n          icmp = {\n            ipv4_code = true\n            ipv4_type = true\n            ipv6_code = true\n            ipv6_type = true\n          }\n          src_port = true\n          tcp = {\n            ack_number = true\n            dst_port = true\n            flags = true\n            header_len = true\n            seq_number = true\n            src_port = true\n            urgent_ptr = true\n            window_size = true\n          }\n          udp = {\n            dst_port = true\n            msg_len = true\n            src_port = true\n          }\n        }\n        type = \"export\"\n      }\n      export_meta_app_profile_alias = \"example\"\n      exporter_alias = \"example\"\n      exporter_config = {\n        alias = \"example\"\n        application_profiles = [\"example\"]\n        cef = {\n          active_timeout = 1\n          inactive_timeout = 1\n        }\n        description = \"example\"\n        destination = {\n          dscp = 0\n          ipv4_address = \"example\"\n          l4_port_dst = 1\n          l4_port_src = 1\n          l4_protocol = \"udp\"\n          ttl = 1\n        }\n        max_pkt_size = 0\n        mobility_sam = {\n          encoding = \"example\"\n          encoding_format = \"hierarchy\"\n          event_enable = {\n            modify = true\n            update = true\n          }\n          trigger = \"example\"\n        }\n        monitor = {\n          timeout = 60\n        }\n        netflow = {\n          active_timeout = 1\n          inactive_timeout = 1\n          template_refresh = 1\n          template_type = \"cohesive\"\n          version = \"v5\"\n        }\n        snmp = {\n          enabled = true\n        }\n        source = {\n          ip_interface = \"example\"\n        }\n        type = \"cef\"\n      }\n    }]\n    gsop_alias = \"example\"\n    gsop_config = {\n      alias = \"example\"\n      cluster_id = \"example\"\n      gs_apps = {\n        apf = {\n          enabled = \"enabled\"\n        }\n        dedup = {\n          enabled = \"enabled\"\n        }\n        diameter_whitelist = {\n          enabled = \"enabled\"\n        }\n        flow_filter = {\n          type = \"gtp\"\n        }\n        flow_sampling = {\n          type = \"ip\"\n        }\n        gseries_header_add = {\n          types = [\"srcid\"]\n        }\n        gseries_header_remove = {\n          enabled = \"enabled\"\n        }\n        gseries_load_balance = {\n          fixed_offset = {\n            hash = \"checksum\"\n            length = 1\n            offset = 0\n          }\n          variable_offset = {\n            end_delim = \"example\"\n            hash = \"checksum\"\n            start_delim = \"example\"\n            start_field = \"example\"\n          }\n        }\n        gseries_pattern_match = {\n          fixed_offset = {\n            length = 1\n            offset = 0\n          }\n          variable_offset = {\n            end_delim = \"example\"\n            start_delim = \"example\"\n          }\n        }\n        gtp_whitelist = {\n          enabled = \"enabled\"\n        }\n        header_add = {\n          vlan = 0\n        }\n        header_remove = {\n          ah1 = \"none\"\n          ah2 = \"none\"\n          custom_len = 1\n          erspan_flow_id = 0\n          fp_dst_switch_id = 0\n          fp_src_switch_id = 0\n          header_count = 1\n          offset = \"start\"\n          offset_range_value = 0\n          protocol = \"gtp\"\n          timestamp_format = \"gigasmart\"\n          vlan_header = \"all\"\n          vxlan_id = 0\n        }\n        icap = {\n          icap_profile = \"example\"\n        }\n        inline_ssl = {\n          inline_ssl_profile = \"example\"\n        }\n        load_balance = {\n          enhanced = {\n            elb_alias = \"example\"\n          }\n          stateful = {\n            app_type = \"gtp\"\n            diameter_key_hash_type = \"sessionId\"\n            diameter_key_multi_hash_type = [{\n              avp_codevalue = 0\n              key = \"sessionId\"\n            }]\n            gtp_key_hash_type = \"imsi\"\n            lb_type = \"leastBw\"\n            sip_key_hash_type = \"callerId\"\n          }\n          stateless = {\n            field_location = \"inner\"\n            hash_fields = \"ipOnly\"\n          }\n        }\n        masking = {\n          content_type = \"message_cpim\"\n          length = 1\n          offset = 0\n          pattern = \"a1\"\n          protocol = \"none\"\n        }\n        metadata_export = {\n          cache = \"example\"\n        }\n        netflow = {\n          enabled = \"enabled\"\n        }\n        sa_apf = {\n          enabled = \"enabled\"\n        }\n        sip_whitelist = {\n          enabled = \"enabled\"\n        }\n        slicing = {\n          enhanced = \"example\"\n          offset = 4\n          protocol = \"none\"\n        }\n        ssl_decrypt = {\n          in_port = 0\n          out_port = 0\n        }\n        trailer_add = {\n          types = [\"crc\"]\n        }\n        trailer_remove = {\n          enabled = \"enabled\"\n        }\n        tunnel_decap = {\n          custom = {\n            port_dst = 0\n            port_src = 0\n          }\n          erspan_flow_id = 0\n          gmip_port = 0\n          l2_gre_key = 0\n          tls_pcapng = {\n            decap_key = \"example\"\n            listener = \"example\"\n          }\n          type = \"gmip\"\n          vxlan = {\n            port_dst = 1\n            port_src = 0\n            vni = 0\n          }\n        }\n        tunnel_encap = {\n          gmip_config = {\n            dscp = 0\n            dst_ip = \"example\"\n            dst_port = 0\n            flow_label = 0\n            prec = 0\n            src_port = 0\n            ttl = 1\n          }\n          l2_gre_config = {\n            dscp = 0\n            dst_ip = \"example\"\n            flow_label = 0\n            key = 0\n            pg_dst = \"example\"\n            prec = 0\n            session_field = \"fiveTupleIpv4\"\n            session_pos = \"inner\"\n            ttl = 1\n          }\n          tls_pcapng = {\n            exporter = \"example\"\n            exporter_group = \"example\"\n          }\n          type = \"gmip\"\n          vxlan_config = {\n            dscp = 0\n            dst_ip = \"example\"\n            dst_port = 4789\n            src_port = 0\n            ttl = 1\n            vni = 1\n          }\n        }\n      }\n      gs_group = \"example\"\n      health_state = \"green\"\n      health_state_reasons = [{\n        message = \"example\"\n        severity = \"green\"\n        traffic_health_state_computation_type = \"PORT_LOW_UTIL\"\n      }]\n    }\n  }\n  app_filter_config = {\n    egress_traffic_config = null\n    sapf_profile = {\n      alias = \"example\"\n      bidi = true\n      buffering = {\n        buffer_count_before_match = 3\n        enabled = true\n        protocol = \"tcp\"\n      }\n      cluster_id = \"example\"\n      packet_count = 0\n      session_fields = [{\n        pos = 1\n        type = \"ipv4Addr\"\n      }]\n      timeout = 10\n    }\n    sapf_profile_alias = \"example\"\n  }\n  associated_monitor_solution_alias = \"%s\"\n  cluster_id = \"example\"\n  config_status = \"example\"\n  delete_monitor_sol = true\n  egress_map_aliases_to_delete = [\"example\"]\n  exporter_aliases_to_delete = [\"example\"]\n  health_state = \"green\"\n  health_state_reasons = [{\n    message = \"example\"\n    severity = \"green\"\n    traffic_health_state_computation_type = \"PORT_LOW_UTIL\"\n  }]\n  ingress_map_aliases_to_delete = [\"example\"]\n  ingress_traffic_configs = null\n  monitor_solution_config = {\n    action = \"example\"\n    cluster_id = \"example\"\n    config_status = \"example\"\n    error_message = \"example\"\n    exporter_alias = \"example\"\n    exporter_config = {\n      alias = \"example\"\n      application_profiles = [\"example\"]\n      cef = {\n        active_timeout = 1\n        inactive_timeout = 1\n      }\n      description = \"example\"\n      destination = {\n        dscp = 0\n        ipv4_address = \"example\"\n        l4_port_dst = 1\n        l4_port_src = 1\n        l4_protocol = \"udp\"\n        ttl = 1\n      }\n      max_pkt_size = 0\n      mobility_sam = {\n        encoding = \"example\"\n        encoding_format = \"hierarchy\"\n        event_enable = {\n          modify = true\n          update = true\n        }\n        trigger = \"example\"\n      }\n      monitor = {\n        timeout = 60\n      }\n      netflow = {\n        active_timeout = 1\n        inactive_timeout = 1\n        template_refresh = 1\n        template_type = \"cohesive\"\n        version = \"v5\"\n      }\n      snmp = {\n        enabled = true\n      }\n      source = {\n        ip_interface = \"example\"\n      }\n      type = \"cef\"\n    }\n    gs_group = \"example\"\n    monitor_ip_interface = \"example\"\n    ref_sols = [{\n      alias = \"example\"\n      associated_map = \"example\"\n      type = \"example\"\n    }]\n    solution_alias = \"example\"\n    vport_alias = \"example\"\n    vport_config = {\n      alias = \"example\"\n      deferred_binding = true\n      fail_over_action = \"vport-bypass\"\n      gs_group = \"example\"\n      inline_status = \"up\"\n      inner_traffic_path = \"to-inline-tool\"\n      metadata_monitoring = {\n        action = \"enable\"\n        exporters = [\"example\"]\n      }\n      mode = \"none\"\n      outer_traffic_path = \"to-inline-tool\"\n      sa_apf_profile = \"example\"\n    }\n  }\n  solution_alias = \"example\"\n  solution_desc = \"example\"\n  solution_status = \"example\"\n  solution_type = \"monitor\"\n}\n", serverURL, name)
}

// newSolutionResourceMockServer returns an httptest server that stubs the SolutionResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newSolutionResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/appsVisibility/solutions"), "/")
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
	mux.HandleFunc("/appsVisibility/solutions", handler0)
	mux.HandleFunc("/appsVisibility/solutions/", handler0)
	return httptest.NewServer(mux)
}

// TestAccSolutionResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccSolutionResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newSolutionResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccSolutionResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_solution.example", "solution_alias"), resource.TestCheckResourceAttr("gigavuecore_solution.example", "associated_monitor_solution_alias", "example"))}, resource.TestStep{Config: testAccSolutionResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_solution.example", "solution_alias"), resource.TestCheckResourceAttr("gigavuecore_solution.example", "associated_monitor_solution_alias", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_solution.example", ImportState: true, ImportStateId: "imported-solution_alias"}}})
}
