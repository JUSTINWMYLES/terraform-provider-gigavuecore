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

func testAccInlineSslAppResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_inline_ssl_app\" \"example\" {\n  alias = \"example\"\n  app_intent_configs = {\n    black_list_config = {\n      operation = \"ADD\"\n      profile_list = {\n        file_source = {\n          hostname = \"example\"\n          password = \"example\"\n          path = \"example\"\n          protocol = \"scp\"\n          username = \"example\"\n        }\n        list = \"example\"\n      }\n    }\n    global_default_configs = {\n      caching = {\n        persistence = {\n          enable = true\n        }\n      }\n      cluster_id = \"example\"\n      dhe_ciphersuit = \"disable\"\n      monitor = {\n        enable = true\n      }\n      resumption = {\n        client = {\n          enable = true\n        }\n      }\n      ssl_versions = {\n        connection_reset_action_for_max_version = \"no-decrypt\"\n        connection_reset_action_for_min_version = \"no-decrypt\"\n        max_version = \"sslv3\"\n        min_version = \"sslv3\"\n      }\n      start_tls = {\n        enable = true\n      }\n    }\n    gs_engines = [\"example\"]\n    gs_group_alias = \"example\"\n    gs_group_param_configs = {\n      hsm_group = \"example\"\n      session_logging = {\n        interface = \"example\"\n        log_level = \"err\"\n        remote_syslog_ip = \"example\"\n        remote_syslog_port = 0\n      }\n    }\n    gsop_alias = \"example\"\n    inline_ssl = {\n      standalone = true\n    }\n    key_store_configs = {\n      deployment_type = \"Inbound\"\n      inbound_keys = [{\n        key_alias = \"example\"\n        server_domain_alias = \"example\"\n      }]\n      outboundkeys = [{\n        key_alias = \"example\"\n        signing_for = \"Primary\"\n      }]\n    }\n    m_tls_configs = {\n      primary_signing = \"example\"\n      secondary_signing = \"example\"\n      trust_store = \"example\"\n    }\n    network_access_configs = [{\n      network_access = {\n        cluster_id = \"example\"\n        dhcp = true\n        dns = \"example\"\n        eport = \"example\"\n        gateway = \"example\"\n        hw_address = \"example\"\n        interface = \"eth2\"\n        ip_address = \"example\"\n        ip_mask = \"example\"\n        mtu = 68\n        proxy_server_profile = \"example\"\n        status = \"up\"\n        vlan = 20\n      }\n      operation = \"Add\"\n    }]\n    ssl_path = [{\n      alias = \"example\"\n      flex_inline_map = {\n        a_to_b = {\n          ib_pathway = \"example\"\n          tools = [\"example\"]\n          type = \"bypass\"\n        }\n        b_to_a = {\n          ib_pathway = \"example\"\n          tools = [\"example\"]\n          type = \"bypass\"\n        }\n        oob_copy = [{\n          direction = \"aToB\"\n          dst_ports = [\"example\"]\n          src_ports = [\"example\"]\n          tag = {\n            type = \"none\"\n          }\n        }]\n        svt_mode = true\n        svt_tag = 0\n        tag = {\n          tag_protocol_id = \"0x8100\"\n          type = \"auto\"\n          vlan_id = 0\n        }\n      }\n    }]\n    ssl_profile_alias = \"example\"\n    ssl_profile_config = {\n      alias = \"example\"\n      certificate = {\n        expired = \"decrypt\"\n        invalid = \"decrypt\"\n        revocation = {\n          crl = {\n            defer = 20\n            enabled = true\n            fail = \"soft\"\n          }\n          ocsp = {\n            defer = 20\n            enabled = true\n            fail = \"soft\"\n          }\n        }\n        self_signed = \"decrypt\"\n        unknown_ca = \"decrypt\"\n      }\n      client_auth = {\n        expired = \"decrypt\"\n        invalid = \"decrypt\"\n        revocation = {\n          crl = {\n            defer = 20\n            enabled = true\n            fail = \"soft\"\n          }\n          ocsp = {\n            defer = 20\n            enabled = true\n            fail = \"soft\"\n          }\n        }\n        self_signed = \"decrypt\"\n        unknown_ca = \"decrypt\"\n      }\n      cluster_id = \"example\"\n      decrypt = {\n        tcp = {\n          inactive_timeout = 2\n          port_map = {\n            default_out_port = 0\n            ports = [{\n              in_port = 1\n              out_port = 1\n              rule_id = 0\n            }]\n          }\n        }\n        tool_bypass = {\n          enable = true\n        }\n      }\n      default_action = \"decrypt\"\n      high_avail = {\n        active_standby = {\n          enable = true\n        }\n      }\n      key_map = [{\n        hostname = \"example\"\n        key = \"example\"\n        rule_id = 0\n      }]\n      monitor = \"enable\"\n      network_group = {\n        multiple_entry = {\n          enable = true\n        }\n      }\n      no_decrypt = {\n        tool_bypass = {\n          enable = true\n        }\n      }\n      non_ssl_tcp = {\n        tool_bypass = {\n          enable = true\n        }\n      }\n      rules = null\n      split_proxy = {\n        mode = {\n          enable = true\n        }\n        server_non_pfs_ciphers = {\n          enable = true\n        }\n      }\n      start_tls = {\n        l4_port = [0]\n      }\n      tcp = {\n        delayed_ack = true\n        syn_retries = 0\n        timewait_timeout = 0\n      }\n      tool = {\n        early_engage = true\n        fail_action = \"fail-open\"\n      }\n      url_cache = {\n        miss_action = \"decrypt\"\n        timeout = 1\n      }\n    }\n    tag_protocol_id = \"0x8100\"\n    trust_store_configs = {\n      trust_store_append_configs = {\n        file = \"example\"\n        file_source = {\n          hostname = \"example\"\n          password = \"example\"\n          path = \"example\"\n          protocol = \"scp\"\n          username = \"example\"\n        }\n      }\n      trust_store_replace_configs = {\n        file = \"example\"\n        file_source = {\n          hostname = \"example\"\n          password = \"example\"\n          path = \"example\"\n          protocol = \"scp\"\n          username = \"example\"\n        }\n      }\n    }\n    vlan_id = 0\n    vport_alias = \"example\"\n    vport_config = {\n      alias = \"example\"\n      deferred_binding = true\n      fail_over_action = \"vport-bypass\"\n      gs_group = \"example\"\n      health_state = \"green\"\n      health_state_reasons = [{\n        message = \"example\"\n        severity = \"green\"\n        traffic_health_state_computation_type = \"PORT_LOW_UTIL\"\n      }]\n      inline_status = \"up\"\n      inner_traffic_path = \"to-inline-tool\"\n      metadata_monitoring = {\n        action = \"enable\"\n        exporters = [\"example\"]\n      }\n      mode = \"none\"\n      outer_traffic_path = \"to-inline-tool\"\n      sa_apf_profile = \"example\"\n    }\n    white_list_config = {\n      operation = \"ADD\"\n      profile_list = {\n        file_source = {\n          hostname = \"example\"\n          password = \"example\"\n          path = \"example\"\n          protocol = \"scp\"\n          username = \"example\"\n        }\n        list = \"example\"\n      }\n    }\n  }\n  cluster_id = \"%s\"\n  cluster_name = \"example\"\n  config_status = \"SUCCESS\"\n  config_status_reasons = [\"example\"]\n  health_state = \"green\"\n  health_state_reasons = [{\n    message = \"example\"\n    severity = \"green\"\n    traffic_health_state_computation_type = \"PORT_LOW_UTIL\"\n  }]\n  m_tls = \"false\"\n  ria_configs = [{\n    cluster_name = \"example\"\n    gs_engines = [\"example\"]\n    gs_group_alias = \"example\"\n    gsop_alias = \"example\"\n    network_access_configs = [{\n      network_access = {\n        cluster_id = \"example\"\n        dhcp = true\n        dns = \"example\"\n        eport = \"example\"\n        gateway = \"example\"\n        hw_address = \"example\"\n        interface = \"eth2\"\n        ip_address = \"example\"\n        ip_mask = \"example\"\n        mtu = 68\n        proxy_server_profile = \"example\"\n        status = \"up\"\n        vlan = 20\n      }\n      operation = \"Add\"\n    }]\n    ssl_profile_alias = \"example\"\n    vport_alias = \"example\"\n  }]\n  ria_enabled = \"false\"\n}\n", serverURL, name)
}

// newInlineSslAppResourceMockServer returns an httptest server that stubs the InlineSslAppResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newInlineSslAppResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/flexInline/inlineSslApp"), "/")
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
	mux.HandleFunc("/flexInline/inlineSslApp", handler0)
	mux.HandleFunc("/flexInline/inlineSslApp/", handler0)
	return httptest.NewServer(mux)
}

// TestAccInlineSslAppResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccInlineSslAppResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newInlineSslAppResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccInlineSslAppResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_inline_ssl_app.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_inline_ssl_app.example", "cluster_id", "example"))}, resource.TestStep{Config: testAccInlineSslAppResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_inline_ssl_app.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_inline_ssl_app.example", "cluster_id", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_inline_ssl_app.example", ImportState: true, ImportStateId: "imported-alias"}}})
}
