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

func testAccNotifMetaConfigResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_notif_meta_config\" \"example\" {\n  allow_attachment = true\n  attachment_limit = 0\n  comment = \"%s\"\n  email_subject_prefix = \"example\"\n  enabled = true\n  event_details = [{\n    description = \"example\"\n    display_name = \"example\"\n    event_type = \"example\"\n    name = \"example\"\n    scope = \"example\"\n    severity = [\"Clear\"]\n    severity_type = \"critical\"\n    sub_type = \"example\"\n  }]\n  external_trap_receivers = [\"example\"]\n  instant_rate_limit = 0\n  recipients = [\"example\"]\n  recurring_schedule = \"example\"\n  send_mail_if_empty = true\n  severity = [\"Clear\"]\n  tags = [{\n    tag_key = \"example\"\n    tag_values = [\"example\"]\n  }]\n  task_id = \"example\"\n  task_name = \"example\"\n  template_details = [\"example\"]\n  time_interval = 1\n  time_left = \"example\"\n  type = \"instant\"\n}\n", serverURL, name)
}

// newNotifMetaConfigResourceMockServer returns an httptest server that stubs the NotifMetaConfigResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newNotifMetaConfigResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/notification/event/notifMetaConfig/instant"), "/")
		if id == "" {
			id = "example-id"
		}
		switch r.Method {
		case http.MethodPut:
			body := make(map[string]interface{})
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil && err != io.EOF {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if _, ok := body["taskId"]; !ok {
				body["taskId"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["taskId"])
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
		case http.MethodPatch:
			body := make(map[string]interface{})
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil && err != io.EOF {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if _, ok := body["taskId"]; !ok {
				body["taskId"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["taskId"])
			state0[id] = body
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			_ = json.NewEncoder(w).Encode(body)
			lastKey0 = id
			return
		case http.MethodDelete:
			delete(state0, id)
			w.WriteHeader(200)
			return
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
	mux.HandleFunc("/notification/event/notifMetaConfig/instant", handler0)
	mux.HandleFunc("/notification/event/notifMetaConfig/instant/", handler0)
	return httptest.NewServer(mux)
}

// TestAccNotifMetaConfigResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccNotifMetaConfigResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newNotifMetaConfigResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccNotifMetaConfigResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_notif_meta_config.example", "task_id"), resource.TestCheckResourceAttr("gigavuecore_notif_meta_config.example", "comment", "example"))}, resource.TestStep{Config: testAccNotifMetaConfigResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_notif_meta_config.example", "task_id"), resource.TestCheckResourceAttr("gigavuecore_notif_meta_config.example", "comment", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_notif_meta_config.example", ImportState: true, ImportStateId: "instant/imported-task_id"}}})
}
