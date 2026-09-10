package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestTrafficFlowResource_Create_Happy exercises TrafficFlowResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestTrafficFlowResource_Create_Happy(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := TrafficFlowResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTrafficFlowResource_Create_NilClient exercises TrafficFlowResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTrafficFlowResource_Create_NilClient(t *testing.T) {
	r := &TrafficFlowResource{}
	m := TrafficFlowResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTrafficFlowResource_Create_BuildError exercises TrafficFlowResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTrafficFlowResource_Create_BuildError(t *testing.T) {
	r := &TrafficFlowResource{client: newMalformedBaseURLClient(t)}
	m := TrafficFlowResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTrafficFlowResource_Create_SendError exercises TrafficFlowResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestTrafficFlowResource_Create_SendError(t *testing.T) {
	r := &TrafficFlowResource{client: newTransportErrorClient(t)}
	m := TrafficFlowResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTrafficFlowResource_Create_APIError exercises TrafficFlowResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTrafficFlowResource_Create_APIError(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TrafficFlowResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_traffic_flow")
}

// TestTrafficFlowResource_Create_APIErrorReadBody exercises TrafficFlowResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTrafficFlowResource_Create_APIErrorReadBody(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientReadErrorBody(t, 501)}
	m := TrafficFlowResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTrafficFlowResource_Create_InvalidJSON exercises TrafficFlowResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTrafficFlowResource_Create_InvalidJSON(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 201, "{{")}
	m := TrafficFlowResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTrafficFlowResource_Create_MapError exercises TrafficFlowResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTrafficFlowResource_Create_MapError(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := TrafficFlowResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTrafficFlowResource_Create_MissingID exercises TrafficFlowResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestTrafficFlowResource_Create_MissingID(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 201, "{}")}
	m := TrafficFlowResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestTrafficFlowResource_Create_LocationFallback exercises TrafficFlowResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestTrafficFlowResource_Create_LocationFallback(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := TrafficFlowResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestTrafficFlowResource_Read_Happy exercises TrafficFlowResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestTrafficFlowResource_Read_Happy(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 200, "{}")}
	m := TrafficFlowResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestTrafficFlowResource_Read_NilClient exercises TrafficFlowResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTrafficFlowResource_Read_NilClient(t *testing.T) {
	r := &TrafficFlowResource{}
	m := TrafficFlowResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTrafficFlowResource_Read_BuildError exercises TrafficFlowResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTrafficFlowResource_Read_BuildError(t *testing.T) {
	r := &TrafficFlowResource{client: newMalformedBaseURLClient(t)}
	m := TrafficFlowResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTrafficFlowResource_Read_SendError exercises TrafficFlowResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestTrafficFlowResource_Read_SendError(t *testing.T) {
	r := &TrafficFlowResource{client: newTransportErrorClient(t)}
	m := TrafficFlowResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTrafficFlowResource_Read_NotFound exercises TrafficFlowResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestTrafficFlowResource_Read_NotFound(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 404, "")}
	m := TrafficFlowResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestTrafficFlowResource_Read_APIError exercises TrafficFlowResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTrafficFlowResource_Read_APIError(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TrafficFlowResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_traffic_flow")
}

// TestTrafficFlowResource_Read_APIErrorReadBody exercises TrafficFlowResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTrafficFlowResource_Read_APIErrorReadBody(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientReadErrorBody(t, 501)}
	m := TrafficFlowResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTrafficFlowResource_Read_InvalidJSON exercises TrafficFlowResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTrafficFlowResource_Read_InvalidJSON(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 200, "{{")}
	m := TrafficFlowResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTrafficFlowResource_Read_MapError exercises TrafficFlowResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTrafficFlowResource_Read_MapError(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := TrafficFlowResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTrafficFlowResource_Update_Happy exercises TrafficFlowResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestTrafficFlowResource_Update_Happy(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 200, "{}")}
	m := TrafficFlowResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTrafficFlowResource_Update_NilClient exercises TrafficFlowResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTrafficFlowResource_Update_NilClient(t *testing.T) {
	r := &TrafficFlowResource{}
	m := TrafficFlowResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTrafficFlowResource_Update_BuildError exercises TrafficFlowResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTrafficFlowResource_Update_BuildError(t *testing.T) {
	r := &TrafficFlowResource{client: newMalformedBaseURLClient(t)}
	m := TrafficFlowResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTrafficFlowResource_Update_SendError exercises TrafficFlowResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestTrafficFlowResource_Update_SendError(t *testing.T) {
	r := &TrafficFlowResource{client: newTransportErrorClient(t)}
	m := TrafficFlowResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTrafficFlowResource_Update_APIError exercises TrafficFlowResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTrafficFlowResource_Update_APIError(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TrafficFlowResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_traffic_flow")
}

// TestTrafficFlowResource_Update_APIErrorReadBody exercises TrafficFlowResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTrafficFlowResource_Update_APIErrorReadBody(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientReadErrorBody(t, 501)}
	m := TrafficFlowResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTrafficFlowResource_Update_InvalidJSON exercises TrafficFlowResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTrafficFlowResource_Update_InvalidJSON(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 200, "{{")}
	m := TrafficFlowResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTrafficFlowResource_Update_MapError exercises TrafficFlowResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTrafficFlowResource_Update_MapError(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := TrafficFlowResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTrafficFlowResource_Delete_Happy exercises TrafficFlowResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestTrafficFlowResource_Delete_Happy(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 204, "")}
	m := TrafficFlowResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTrafficFlowResource_Delete_NilClient exercises TrafficFlowResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTrafficFlowResource_Delete_NilClient(t *testing.T) {
	r := &TrafficFlowResource{}
	m := TrafficFlowResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTrafficFlowResource_Delete_BuildError exercises TrafficFlowResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTrafficFlowResource_Delete_BuildError(t *testing.T) {
	r := &TrafficFlowResource{client: newMalformedBaseURLClient(t)}
	m := TrafficFlowResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTrafficFlowResource_Delete_SendError exercises TrafficFlowResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestTrafficFlowResource_Delete_SendError(t *testing.T) {
	r := &TrafficFlowResource{client: newTransportErrorClient(t)}
	m := TrafficFlowResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTrafficFlowResource_Delete_NotFoundSuccess exercises TrafficFlowResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestTrafficFlowResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 404, "")}
	m := TrafficFlowResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTrafficFlowResource_Delete_APIError exercises TrafficFlowResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTrafficFlowResource_Delete_APIError(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TrafficFlowResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_traffic_flow")
}

// TestTrafficFlowResource_Delete_APIErrorReadBody exercises TrafficFlowResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTrafficFlowResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &TrafficFlowResource{client: newMockClientReadErrorBody(t, 501)}
	m := TrafficFlowResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
