package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNotifTargetResource_Create_Happy exercises NotifTargetResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestNotifTargetResource_Create_Happy(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 201, "{\"notif_target_address\":\"example-id\"}")}
	m := NotifTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNotifTargetResource_Create_NilClient exercises NotifTargetResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNotifTargetResource_Create_NilClient(t *testing.T) {
	r := &NotifTargetResource{}
	m := NotifTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNotifTargetResource_Create_BuildError exercises NotifTargetResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNotifTargetResource_Create_BuildError(t *testing.T) {
	r := &NotifTargetResource{client: newMalformedBaseURLClient(t)}
	m := NotifTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNotifTargetResource_Create_SendError exercises NotifTargetResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestNotifTargetResource_Create_SendError(t *testing.T) {
	r := &NotifTargetResource{client: newTransportErrorClient(t)}
	m := NotifTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNotifTargetResource_Create_APIError exercises NotifTargetResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNotifTargetResource_Create_APIError(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NotifTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_notif_target")
}

// TestNotifTargetResource_Create_APIErrorReadBody exercises NotifTargetResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNotifTargetResource_Create_APIErrorReadBody(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientReadErrorBody(t, 501)}
	m := NotifTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNotifTargetResource_Create_InvalidJSON exercises NotifTargetResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNotifTargetResource_Create_InvalidJSON(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 201, "{{")}
	m := NotifTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNotifTargetResource_Create_MapError exercises NotifTargetResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNotifTargetResource_Create_MapError(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 201, "{\"notif_target_address\":12345}")}
	m := NotifTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNotifTargetResource_Create_MissingID exercises NotifTargetResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestNotifTargetResource_Create_MissingID(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 201, "{}")}
	m := NotifTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestNotifTargetResource_Create_LocationFallback exercises NotifTargetResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestNotifTargetResource_Create_LocationFallback(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := NotifTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.NotifTargetAddress.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.NotifTargetAddress.ValueString(), "example-id")
	}
}

// TestNotifTargetResource_Read_Happy exercises NotifTargetResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestNotifTargetResource_Read_Happy(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 200, "{}")}
	m := NotifTargetResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNotifTargetResource_Read_NilClient exercises NotifTargetResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNotifTargetResource_Read_NilClient(t *testing.T) {
	r := &NotifTargetResource{}
	m := NotifTargetResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNotifTargetResource_Read_BuildError exercises NotifTargetResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNotifTargetResource_Read_BuildError(t *testing.T) {
	r := &NotifTargetResource{client: newMalformedBaseURLClient(t)}
	m := NotifTargetResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNotifTargetResource_Read_SendError exercises NotifTargetResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestNotifTargetResource_Read_SendError(t *testing.T) {
	r := &NotifTargetResource{client: newTransportErrorClient(t)}
	m := NotifTargetResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNotifTargetResource_Read_NotFound exercises NotifTargetResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestNotifTargetResource_Read_NotFound(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 404, "")}
	m := NotifTargetResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNotifTargetResource_Read_APIError exercises NotifTargetResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNotifTargetResource_Read_APIError(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NotifTargetResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_notif_target")
}

// TestNotifTargetResource_Read_APIErrorReadBody exercises NotifTargetResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNotifTargetResource_Read_APIErrorReadBody(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientReadErrorBody(t, 501)}
	m := NotifTargetResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNotifTargetResource_Read_InvalidJSON exercises NotifTargetResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNotifTargetResource_Read_InvalidJSON(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 200, "{{")}
	m := NotifTargetResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNotifTargetResource_Read_MapError exercises NotifTargetResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNotifTargetResource_Read_MapError(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 200, "{\"notif_target_address\":12345}")}
	m := NotifTargetResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNotifTargetResource_Update_Happy exercises NotifTargetResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestNotifTargetResource_Update_Happy(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 200, "{}")}
	m := NotifTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNotifTargetResource_Update_NilClient exercises NotifTargetResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNotifTargetResource_Update_NilClient(t *testing.T) {
	r := &NotifTargetResource{}
	m := NotifTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNotifTargetResource_Update_BuildError exercises NotifTargetResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNotifTargetResource_Update_BuildError(t *testing.T) {
	r := &NotifTargetResource{client: newMalformedBaseURLClient(t)}
	m := NotifTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNotifTargetResource_Update_SendError exercises NotifTargetResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestNotifTargetResource_Update_SendError(t *testing.T) {
	r := &NotifTargetResource{client: newTransportErrorClient(t)}
	m := NotifTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNotifTargetResource_Update_APIError exercises NotifTargetResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNotifTargetResource_Update_APIError(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NotifTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_notif_target")
}

// TestNotifTargetResource_Update_APIErrorReadBody exercises NotifTargetResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNotifTargetResource_Update_APIErrorReadBody(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientReadErrorBody(t, 501)}
	m := NotifTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNotifTargetResource_Update_InvalidJSON exercises NotifTargetResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNotifTargetResource_Update_InvalidJSON(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 200, "{{")}
	m := NotifTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNotifTargetResource_Update_MapError exercises NotifTargetResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNotifTargetResource_Update_MapError(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 200, "{\"notif_target_address\":12345}")}
	m := NotifTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNotifTargetResource_Delete_Happy exercises NotifTargetResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestNotifTargetResource_Delete_Happy(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 204, "")}
	m := NotifTargetResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNotifTargetResource_Delete_NilClient exercises NotifTargetResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNotifTargetResource_Delete_NilClient(t *testing.T) {
	r := &NotifTargetResource{}
	m := NotifTargetResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNotifTargetResource_Delete_BuildError exercises NotifTargetResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNotifTargetResource_Delete_BuildError(t *testing.T) {
	r := &NotifTargetResource{client: newMalformedBaseURLClient(t)}
	m := NotifTargetResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNotifTargetResource_Delete_SendError exercises NotifTargetResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestNotifTargetResource_Delete_SendError(t *testing.T) {
	r := &NotifTargetResource{client: newTransportErrorClient(t)}
	m := NotifTargetResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNotifTargetResource_Delete_NotFoundSuccess exercises NotifTargetResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestNotifTargetResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 404, "")}
	m := NotifTargetResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNotifTargetResource_Delete_APIError exercises NotifTargetResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNotifTargetResource_Delete_APIError(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NotifTargetResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_notif_target")
}

// TestNotifTargetResource_Delete_APIErrorReadBody exercises NotifTargetResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNotifTargetResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &NotifTargetResource{client: newMockClientReadErrorBody(t, 501)}
	m := NotifTargetResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
