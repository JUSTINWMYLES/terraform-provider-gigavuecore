package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestMonitorResource_Create_Happy exercises MonitorResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestMonitorResource_Create_Happy(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := MonitorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMonitorResource_Create_NilClient exercises MonitorResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMonitorResource_Create_NilClient(t *testing.T) {
	r := &MonitorResource{}
	m := MonitorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMonitorResource_Create_BuildError exercises MonitorResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMonitorResource_Create_BuildError(t *testing.T) {
	r := &MonitorResource{client: newMalformedBaseURLClient(t)}
	m := MonitorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMonitorResource_Create_SendError exercises MonitorResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestMonitorResource_Create_SendError(t *testing.T) {
	r := &MonitorResource{client: newTransportErrorClient(t)}
	m := MonitorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMonitorResource_Create_APIError exercises MonitorResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMonitorResource_Create_APIError(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MonitorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_monitor")
}

// TestMonitorResource_Create_APIErrorReadBody exercises MonitorResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMonitorResource_Create_APIErrorReadBody(t *testing.T) {
	r := &MonitorResource{client: newMockClientReadErrorBody(t, 501)}
	m := MonitorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMonitorResource_Create_InvalidJSON exercises MonitorResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMonitorResource_Create_InvalidJSON(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 201, "{{")}
	m := MonitorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMonitorResource_Create_MapError exercises MonitorResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMonitorResource_Create_MapError(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := MonitorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMonitorResource_Create_MissingID exercises MonitorResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestMonitorResource_Create_MissingID(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 201, "{}")}
	m := MonitorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestMonitorResource_Create_LocationFallback exercises MonitorResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestMonitorResource_Create_LocationFallback(t *testing.T) {
	r := &MonitorResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := MonitorResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestMonitorResource_Read_Happy exercises MonitorResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestMonitorResource_Read_Happy(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 200, "{}")}
	m := MonitorResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestMonitorResource_Read_NilClient exercises MonitorResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMonitorResource_Read_NilClient(t *testing.T) {
	r := &MonitorResource{}
	m := MonitorResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMonitorResource_Read_BuildError exercises MonitorResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMonitorResource_Read_BuildError(t *testing.T) {
	r := &MonitorResource{client: newMalformedBaseURLClient(t)}
	m := MonitorResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMonitorResource_Read_SendError exercises MonitorResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestMonitorResource_Read_SendError(t *testing.T) {
	r := &MonitorResource{client: newTransportErrorClient(t)}
	m := MonitorResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMonitorResource_Read_NotFound exercises MonitorResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestMonitorResource_Read_NotFound(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 404, "")}
	m := MonitorResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestMonitorResource_Read_APIError exercises MonitorResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMonitorResource_Read_APIError(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MonitorResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_monitor")
}

// TestMonitorResource_Read_APIErrorReadBody exercises MonitorResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMonitorResource_Read_APIErrorReadBody(t *testing.T) {
	r := &MonitorResource{client: newMockClientReadErrorBody(t, 501)}
	m := MonitorResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMonitorResource_Read_InvalidJSON exercises MonitorResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMonitorResource_Read_InvalidJSON(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 200, "{{")}
	m := MonitorResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMonitorResource_Read_MapError exercises MonitorResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMonitorResource_Read_MapError(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := MonitorResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMonitorResource_Update_Happy exercises MonitorResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestMonitorResource_Update_Happy(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 200, "{}")}
	m := MonitorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMonitorResource_Update_NilClient exercises MonitorResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMonitorResource_Update_NilClient(t *testing.T) {
	r := &MonitorResource{}
	m := MonitorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMonitorResource_Update_BuildError exercises MonitorResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMonitorResource_Update_BuildError(t *testing.T) {
	r := &MonitorResource{client: newMalformedBaseURLClient(t)}
	m := MonitorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMonitorResource_Update_SendError exercises MonitorResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestMonitorResource_Update_SendError(t *testing.T) {
	r := &MonitorResource{client: newTransportErrorClient(t)}
	m := MonitorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMonitorResource_Update_APIError exercises MonitorResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMonitorResource_Update_APIError(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MonitorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_monitor")
}

// TestMonitorResource_Update_APIErrorReadBody exercises MonitorResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMonitorResource_Update_APIErrorReadBody(t *testing.T) {
	r := &MonitorResource{client: newMockClientReadErrorBody(t, 501)}
	m := MonitorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMonitorResource_Update_InvalidJSON exercises MonitorResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMonitorResource_Update_InvalidJSON(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 200, "{{")}
	m := MonitorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMonitorResource_Update_MapError exercises MonitorResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMonitorResource_Update_MapError(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := MonitorResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMonitorResource_Delete_Happy exercises MonitorResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestMonitorResource_Delete_Happy(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 204, "")}
	m := MonitorResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMonitorResource_Delete_NilClient exercises MonitorResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMonitorResource_Delete_NilClient(t *testing.T) {
	r := &MonitorResource{}
	m := MonitorResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMonitorResource_Delete_BuildError exercises MonitorResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMonitorResource_Delete_BuildError(t *testing.T) {
	r := &MonitorResource{client: newMalformedBaseURLClient(t)}
	m := MonitorResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMonitorResource_Delete_SendError exercises MonitorResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestMonitorResource_Delete_SendError(t *testing.T) {
	r := &MonitorResource{client: newTransportErrorClient(t)}
	m := MonitorResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMonitorResource_Delete_NotFoundSuccess exercises MonitorResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestMonitorResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 404, "")}
	m := MonitorResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMonitorResource_Delete_APIError exercises MonitorResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMonitorResource_Delete_APIError(t *testing.T) {
	r := &MonitorResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MonitorResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_monitor")
}

// TestMonitorResource_Delete_APIErrorReadBody exercises MonitorResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMonitorResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &MonitorResource{client: newMockClientReadErrorBody(t, 501)}
	m := MonitorResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
