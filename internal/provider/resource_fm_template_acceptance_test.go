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

func testAccFmTemplateResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_fm_template\" \"example\" {\n  config = {\n    aaa_auth_config = {\n      auth_sequence = [\"local\"]\n      external_login_mapping = {\n        default_local_user = \"example\"\n        user_map_order = \"localOnly\"\n      }\n    }\n    acme_certificate = [{\n      acme_server_url = \"example\"\n      algorithm = \"rsa-2048\"\n      operation_type = \"issue\"\n      renew_days = 0\n    }]\n    device_ssl_certificate_configs = [{\n      issuer = \"example\"\n      not_after = \"example\"\n      not_before = \"example\"\n      operation_type = \"add\"\n      signature_algorithm = \"example\"\n      subject = \"example\"\n      trusted_ca = {\n        name = \"example\"\n      }\n      upload_spec = {\n        info = {\n          comment = \"example\"\n          name = \"example\"\n          passphrase = \"example\"\n          type = \"privateKey\"\n        }\n        pem = \"example\"\n      }\n    }]\n    export_metadata_app_profile = {\n      alias = \"example\"\n      application_id = true\n      applications = [{\n        attributes = [{\n          name = \"example\"\n          value = \"example\"\n        }]\n        is_user_defined = true\n        name = \"example\"\n      }]\n      counter = {\n        bytes = true\n        bytes_long = true\n        inner_byte = true\n        inner_byte_long = true\n        packets = true\n        packets_long = true\n      }\n      datalink = {\n        mac_dst = true\n        mac_src = true\n        vlan = true\n      }\n      description = \"example\"\n      flow = {\n        end_reason = true\n      }\n      gtpu = {\n        qfi = true\n        teid = true\n      }\n      interface = {\n        in_name_width = 1\n        in_physical_width = 2\n        out_physical_width = 2\n      }\n      ip = {\n        version = true\n      }\n      ipv4 = {\n        destination = {\n          prefix_min_mask = \"example\"\n        }\n        dscp = true\n        fragmentation = {\n          flags = true\n          offset = true\n        }\n        header_len = true\n        option_map = true\n        precedence = true\n        protocol = true\n        section = {\n          header_size = 1\n          payload_size = 1\n        }\n        source = {\n          prefix_min_mask = \"example\"\n        }\n        tos = true\n        total_length = true\n        ttl = true\n      }\n      ipv6 = {\n        destination = {\n          prefix_min_mask = \"example\"\n        }\n        dscp = true\n        extension_map = true\n        flow_label = true\n        fragmentation = {\n          flags = true\n          offset = true\n        }\n        hop_limit = true\n        length = {\n          header = true\n          payload = true\n          total = true\n        }\n        next_header = true\n        precedence = true\n        section = {\n          header_size = 1\n          payload_size = 1\n        }\n        source = {\n          prefix_min_mask = \"example\"\n        }\n        traffic_class = true\n      }\n      outer_ipv4 = {\n        destination = true\n        source = true\n      }\n      outer_ipv6 = {\n        destination = true\n        source = true\n      }\n      timestamp = {\n        flow_end_msec = true\n        flow_endsec = true\n        flow_start_msec = true\n        flow_startsec = true\n        sys_up_time_first = true\n        sys_up_time_last = true\n      }\n      transport = {\n        dst_port = true\n        icmp = {\n          ipv4_code = true\n          ipv4_type = true\n          ipv6_code = true\n          ipv6_type = true\n        }\n        src_port = true\n        tcp = {\n          ack_number = true\n          dst_port = true\n          flags = true\n          header_len = true\n          seq_number = true\n          src_port = true\n          urgent_ptr = true\n          window_size = true\n        }\n        udp = {\n          dst_port = true\n          msg_len = true\n          src_port = true\n        }\n      }\n      type = \"export\"\n    }\n    giga_port_neighbors_discovery_config = {\n      resource_configs = null\n    }\n    giga_stream_threshold_config = {\n      giga_stream_type_thresholds = [{\n        type = \"all\"\n        variance_threshold = 1.0\n      }]\n    }\n    giga_user_defined_application_config = {\n      alias = \"example\"\n      app_id = 0\n      priority = 0\n      rules = {\n        rule = [{\n          address = \"example\"\n          code = \"example\"\n          common_name = \"example\"\n          content = \"example\"\n          cts_cookie = \"example\"\n          cts_page_url = \"example\"\n          cts_referer = \"example\"\n          cts_server = \"example\"\n          cts_uri = \"example\"\n          cts_user_agent = \"example\"\n          dscp = \"example\"\n          mime_type = \"example\"\n          mindata = 0\n          port = \"example\"\n          resolv_name = \"example\"\n          stc_location = \"example\"\n          stc_server_agent = \"example\"\n          stc_subject_alt_name = \"example\"\n          stream = \"example\"\n          typeval = \"example\"\n          user_agent = \"example\"\n        }]\n      }\n    }\n    ldap_server_system_config = null\n    ldap_servers = [{\n      order = \"example\"\n      server_address = \"example\"\n    }]\n    metadata_exporter = {\n      alias = \"example\"\n      application_profiles = [\"example\"]\n      cef = {\n        active_timeout = 1\n        inactive_timeout = 1\n      }\n      description = \"example\"\n      destination = {\n        dscp = 0\n        ipv4_address = \"example\"\n        l4_port_dst = 1\n        l4_port_src = 1\n        l4_protocol = \"udp\"\n        ttl = 1\n      }\n      max_pkt_size = 0\n      mobility_sam = {\n        encoding = \"example\"\n        encoding_format = \"hierarchy\"\n        event_enable = {\n          modify = true\n          update = true\n        }\n        trigger = \"example\"\n      }\n      monitor = {\n        timeout = 60\n      }\n      netflow = {\n        active_timeout = 1\n        inactive_timeout = 1\n        template_refresh = 1\n        template_type = \"cohesive\"\n        version = \"v5\"\n      }\n      snmp = {\n        enabled = true\n      }\n      source = {\n        ip_interface = \"example\"\n      }\n      type = \"cef\"\n    }\n    port_packet_threshold_config = {\n      drop_threshold = {\n        rx = [{\n          count_ = 0\n          percent = 1.0\n          port_type = \"all\"\n        }]\n        tx = [{\n          count_ = 0\n          percent = 1.0\n          port_type = \"all\"\n        }]\n      }\n      error_threshold = {\n        rx = [{\n          count_ = 0\n          percent = 1.0\n          port_type = \"all\"\n        }]\n        tx = [{\n          count_ = 0\n          percent = 1.0\n          port_type = \"all\"\n        }]\n      }\n    }\n    proxy_server_profile = {\n      alias = \"example\"\n      auth_type = \"none\"\n      comment = \"example\"\n      password = \"example\"\n      periodic_ping = \"enable\"\n      periodic_ping_failure_retry = 1\n      periodic_ping_interval = 1\n      periodic_ping_type = \"http-connect\"\n      port = 1\n      protocol = \"http\"\n      proxy_address = \"example\"\n      ssl_apps = {\n        cluster_name = [\"example\"]\n      }\n      username = \"example\"\n    }\n    snmp_trap_event_configs = [{\n      enabled = true\n      notify_event = \"example\"\n    }]\n    snmp_v3_users_config = {\n      snmp_v3_user = [{\n        auth_key = \"example\"\n        auth_protocol = \"md5\"\n        min_sw_version = \"example\"\n        previous_username = \"example\"\n        priv_key = \"example\"\n        priv_protocol = \"des\"\n        username = \"example\"\n      }]\n    }\n    ssh_ciphers_config = {\n      classic = {\n        client_ciphers = [\"default\"]\n        client_hostkey = [\"default\"]\n        client_kex = [\"default\"]\n        client_macs = [\"default\"]\n        server_ciphers = [\"default\"]\n        server_hostkey = [\"default\"]\n        server_kex = [\"default\"]\n        server_macs = [\"default\"]\n      }\n      crypto = {\n        client_ciphers = [\"default\"]\n        client_hostkey = [\"default\"]\n        client_kex = [\"default\"]\n        client_macs = [\"default\"]\n        server_ciphers = [\"default\"]\n        server_hostkey = [\"default\"]\n        server_kex = [\"default\"]\n        server_macs = [\"default\"]\n      }\n      fips = {\n        client_ciphers = [\"default\"]\n        client_hostkey = [\"default\"]\n        client_kex = [\"default\"]\n        client_macs = [\"default\"]\n        server_ciphers = [\"default\"]\n        server_hostkey = [\"default\"]\n        server_kex = [\"default\"]\n        server_macs = [\"default\"]\n      }\n    }\n  }\n  config_level = \"%s\"\n  config_level_value = [\"example\"]\n  config_resource = null\n  config_type = \"SNMPTRAPS\"\n  modifiable = true\n  ref_count = 0\n  ref_object = {\n    ref_object_type = \"PHYSICAL\"\n  }\n  template_name = \"example\"\n  update_time = \"example\"\n}\n", serverURL, name)
}

// newFmTemplateResourceMockServer returns an httptest server that stubs the FmTemplateResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newFmTemplateResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/fm/templates"), "/")
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
			if _, ok := body["templateName"]; !ok {
				body["templateName"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["templateName"])
			state0[id] = body
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
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
			if _, ok := body["templateName"]; !ok {
				body["templateName"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["templateName"])
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
	mux.HandleFunc("/fm/templates", handler0)
	mux.HandleFunc("/fm/templates/", handler0)
	return httptest.NewServer(mux)
}

// TestAccFmTemplateResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccFmTemplateResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newFmTemplateResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccFmTemplateResourceConfig(server.URL, "GLOBAL"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_fm_template.example", "template_name"), resource.TestCheckResourceAttr("gigavuecore_fm_template.example", "config_level", "GLOBAL"))}, resource.TestStep{Config: testAccFmTemplateResourceConfig(server.URL, "SITE"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_fm_template.example", "template_name"), resource.TestCheckResourceAttr("gigavuecore_fm_template.example", "config_level", "SITE"))}, resource.TestStep{ResourceName: "gigavuecore_fm_template.example", ImportState: true, ImportStateId: "imported-template_name"}}})
}
