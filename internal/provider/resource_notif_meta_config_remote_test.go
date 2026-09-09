package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNotifMetaConfigResource_Create_Happy exercises NotifMetaConfigResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestNotifMetaConfigResource_Create_Happy(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 200, "{\"task_id\":\"example-id\"}")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNotifMetaConfigResource_Create_NilClient exercises NotifMetaConfigResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNotifMetaConfigResource_Create_NilClient(t *testing.T) {
	r := &NotifMetaConfigResource{}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNotifMetaConfigResource_Create_BuildError exercises NotifMetaConfigResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNotifMetaConfigResource_Create_BuildError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMalformedBaseURLClient(t)}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNotifMetaConfigResource_Create_SendError exercises NotifMetaConfigResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestNotifMetaConfigResource_Create_SendError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newTransportErrorClient(t)}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNotifMetaConfigResource_Create_APIError exercises NotifMetaConfigResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNotifMetaConfigResource_Create_APIError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_notif_meta_config")
}

// TestNotifMetaConfigResource_Create_APIErrorReadBody exercises NotifMetaConfigResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNotifMetaConfigResource_Create_APIErrorReadBody(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientReadErrorBody(t, 501)}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNotifMetaConfigResource_Create_InvalidJSON exercises NotifMetaConfigResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNotifMetaConfigResource_Create_InvalidJSON(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 200, "{{")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNotifMetaConfigResource_Create_MapError exercises NotifMetaConfigResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNotifMetaConfigResource_Create_MapError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 200, "{\"task_id\":12345}")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNotifMetaConfigResource_Create_MissingID exercises NotifMetaConfigResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestNotifMetaConfigResource_Create_MissingID(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 200, "{}")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestNotifMetaConfigResource_Create_LocationFallback exercises NotifMetaConfigResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestNotifMetaConfigResource_Create_LocationFallback(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientWithLocation(t, 200, "http://example.test/folders/example-id", "{}")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.TaskId.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.TaskId.ValueString(), "example-id")
	}
}

// TestNotifMetaConfigResource_Read_Happy exercises NotifMetaConfigResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestNotifMetaConfigResource_Read_Happy(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 200, "{}")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNotifMetaConfigResource_Read_NilClient exercises NotifMetaConfigResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNotifMetaConfigResource_Read_NilClient(t *testing.T) {
	r := &NotifMetaConfigResource{}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNotifMetaConfigResource_Read_BuildError exercises NotifMetaConfigResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNotifMetaConfigResource_Read_BuildError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMalformedBaseURLClient(t)}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNotifMetaConfigResource_Read_SendError exercises NotifMetaConfigResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestNotifMetaConfigResource_Read_SendError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newTransportErrorClient(t)}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNotifMetaConfigResource_Read_NotFound exercises NotifMetaConfigResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestNotifMetaConfigResource_Read_NotFound(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 404, "")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNotifMetaConfigResource_Read_APIError exercises NotifMetaConfigResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNotifMetaConfigResource_Read_APIError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_notif_meta_config")
}

// TestNotifMetaConfigResource_Read_APIErrorReadBody exercises NotifMetaConfigResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNotifMetaConfigResource_Read_APIErrorReadBody(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientReadErrorBody(t, 501)}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNotifMetaConfigResource_Read_InvalidJSON exercises NotifMetaConfigResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNotifMetaConfigResource_Read_InvalidJSON(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 200, "{{")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNotifMetaConfigResource_Read_MapError exercises NotifMetaConfigResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNotifMetaConfigResource_Read_MapError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 200, "{\"task_id\":12345}")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNotifMetaConfigResource_Update_Happy exercises NotifMetaConfigResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestNotifMetaConfigResource_Update_Happy(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 200, "{}")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNotifMetaConfigResource_Update_NilClient exercises NotifMetaConfigResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNotifMetaConfigResource_Update_NilClient(t *testing.T) {
	r := &NotifMetaConfigResource{}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNotifMetaConfigResource_Update_BuildError exercises NotifMetaConfigResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNotifMetaConfigResource_Update_BuildError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMalformedBaseURLClient(t)}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNotifMetaConfigResource_Update_SendError exercises NotifMetaConfigResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestNotifMetaConfigResource_Update_SendError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newTransportErrorClient(t)}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNotifMetaConfigResource_Update_APIError exercises NotifMetaConfigResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNotifMetaConfigResource_Update_APIError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_notif_meta_config")
}

// TestNotifMetaConfigResource_Update_APIErrorReadBody exercises NotifMetaConfigResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNotifMetaConfigResource_Update_APIErrorReadBody(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientReadErrorBody(t, 501)}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNotifMetaConfigResource_Update_InvalidJSON exercises NotifMetaConfigResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNotifMetaConfigResource_Update_InvalidJSON(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 200, "{{")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNotifMetaConfigResource_Update_MapError exercises NotifMetaConfigResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNotifMetaConfigResource_Update_MapError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 200, "{\"task_id\":12345}")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNotifMetaConfigResource_Delete_Happy exercises NotifMetaConfigResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestNotifMetaConfigResource_Delete_Happy(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 200, "")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNotifMetaConfigResource_Delete_NilClient exercises NotifMetaConfigResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNotifMetaConfigResource_Delete_NilClient(t *testing.T) {
	r := &NotifMetaConfigResource{}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNotifMetaConfigResource_Delete_BuildError exercises NotifMetaConfigResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNotifMetaConfigResource_Delete_BuildError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMalformedBaseURLClient(t)}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNotifMetaConfigResource_Delete_SendError exercises NotifMetaConfigResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestNotifMetaConfigResource_Delete_SendError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newTransportErrorClient(t)}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNotifMetaConfigResource_Delete_NotFoundSuccess exercises NotifMetaConfigResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestNotifMetaConfigResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 404, "")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNotifMetaConfigResource_Delete_APIError exercises NotifMetaConfigResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNotifMetaConfigResource_Delete_APIError(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_notif_meta_config")
}

// TestNotifMetaConfigResource_Delete_APIErrorReadBody exercises NotifMetaConfigResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNotifMetaConfigResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &NotifMetaConfigResource{client: newMockClientReadErrorBody(t, 501)}
	m := NotifMetaConfigResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
