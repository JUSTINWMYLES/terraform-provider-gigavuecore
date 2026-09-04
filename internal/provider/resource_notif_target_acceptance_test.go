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

func testAccNotifTargetResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_notif_target\" \"example\" {\n  cluster_id = \"%s\"\n  enabled = true\n  host = \"example\"\n  notify_config = {\n    auth_key = \"example-value\"\n    auth_protocol = \"md5\"\n    community = \"example\"\n    engine_id = \"example\"\n    port = 0\n    priv_key = \"example-value\"\n    priv_protocol = \"des\"\n    v3_user = \"example\"\n    version = \"v1\"\n  }\n  notify_type = \"trap\"\n}\n", serverURL, name)
}

// newNotifTargetResourceMockServer returns an httptest server that stubs the NotifTargetResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newNotifTargetResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/system/snmp/notifTargets"), "/")
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
			if _, ok := body["notif_target_address"]; !ok {
				body["notif_target_address"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["notif_target_address"])
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
			if _, ok := body["notif_target_address"]; !ok {
				body["notif_target_address"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["notif_target_address"])
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
	mux.HandleFunc("/system/snmp/notifTargets", handler0)
	mux.HandleFunc("/system/snmp/notifTargets/", handler0)
	return httptest.NewServer(mux)
}

// TestAccNotifTargetResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccNotifTargetResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newNotifTargetResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccNotifTargetResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_notif_target.example", "notif_target_address"), resource.TestCheckResourceAttr("gigavuecore_notif_target.example", "cluster_id", "example"))}, resource.TestStep{Config: testAccNotifTargetResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_notif_target.example", "notif_target_address"), resource.TestCheckResourceAttr("gigavuecore_notif_target.example", "cluster_id", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_notif_target.example", ImportState: true, ImportStateId: "imported-notif_target_address/imported-cluster_id"}}})
}
