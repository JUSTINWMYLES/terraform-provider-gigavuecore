package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestTrafficPolicyGraphResource_Create_Happy exercises TrafficPolicyGraphResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestTrafficPolicyGraphResource_Create_Happy(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTrafficPolicyGraphResource_Create_NilClient exercises TrafficPolicyGraphResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTrafficPolicyGraphResource_Create_NilClient(t *testing.T) {
	r := &TrafficPolicyGraphResource{}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTrafficPolicyGraphResource_Create_BuildError exercises TrafficPolicyGraphResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTrafficPolicyGraphResource_Create_BuildError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMalformedBaseURLClient(t)}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTrafficPolicyGraphResource_Create_SendError exercises TrafficPolicyGraphResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestTrafficPolicyGraphResource_Create_SendError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newTransportErrorClient(t)}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTrafficPolicyGraphResource_Create_APIError exercises TrafficPolicyGraphResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTrafficPolicyGraphResource_Create_APIError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_traffic_policy_graph")
}

// TestTrafficPolicyGraphResource_Create_APIErrorReadBody exercises TrafficPolicyGraphResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTrafficPolicyGraphResource_Create_APIErrorReadBody(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientReadErrorBody(t, 501)}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTrafficPolicyGraphResource_Create_InvalidJSON exercises TrafficPolicyGraphResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTrafficPolicyGraphResource_Create_InvalidJSON(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 201, "{{")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTrafficPolicyGraphResource_Create_MapError exercises TrafficPolicyGraphResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTrafficPolicyGraphResource_Create_MapError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTrafficPolicyGraphResource_Create_MissingID exercises TrafficPolicyGraphResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestTrafficPolicyGraphResource_Create_MissingID(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 201, "{}")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestTrafficPolicyGraphResource_Create_LocationFallback exercises TrafficPolicyGraphResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestTrafficPolicyGraphResource_Create_LocationFallback(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestTrafficPolicyGraphResource_Read_Happy exercises TrafficPolicyGraphResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestTrafficPolicyGraphResource_Read_Happy(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 200, "{}")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestTrafficPolicyGraphResource_Read_NilClient exercises TrafficPolicyGraphResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTrafficPolicyGraphResource_Read_NilClient(t *testing.T) {
	r := &TrafficPolicyGraphResource{}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTrafficPolicyGraphResource_Read_BuildError exercises TrafficPolicyGraphResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTrafficPolicyGraphResource_Read_BuildError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMalformedBaseURLClient(t)}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTrafficPolicyGraphResource_Read_SendError exercises TrafficPolicyGraphResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestTrafficPolicyGraphResource_Read_SendError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newTransportErrorClient(t)}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTrafficPolicyGraphResource_Read_NotFound exercises TrafficPolicyGraphResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestTrafficPolicyGraphResource_Read_NotFound(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 404, "")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestTrafficPolicyGraphResource_Read_APIError exercises TrafficPolicyGraphResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTrafficPolicyGraphResource_Read_APIError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_traffic_policy_graph")
}

// TestTrafficPolicyGraphResource_Read_APIErrorReadBody exercises TrafficPolicyGraphResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTrafficPolicyGraphResource_Read_APIErrorReadBody(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientReadErrorBody(t, 501)}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTrafficPolicyGraphResource_Read_InvalidJSON exercises TrafficPolicyGraphResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTrafficPolicyGraphResource_Read_InvalidJSON(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 200, "{{")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTrafficPolicyGraphResource_Read_MapError exercises TrafficPolicyGraphResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTrafficPolicyGraphResource_Read_MapError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTrafficPolicyGraphResource_Update_Happy exercises TrafficPolicyGraphResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestTrafficPolicyGraphResource_Update_Happy(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 200, "{}")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTrafficPolicyGraphResource_Update_NilClient exercises TrafficPolicyGraphResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTrafficPolicyGraphResource_Update_NilClient(t *testing.T) {
	r := &TrafficPolicyGraphResource{}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTrafficPolicyGraphResource_Update_BuildError exercises TrafficPolicyGraphResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTrafficPolicyGraphResource_Update_BuildError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMalformedBaseURLClient(t)}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTrafficPolicyGraphResource_Update_SendError exercises TrafficPolicyGraphResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestTrafficPolicyGraphResource_Update_SendError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newTransportErrorClient(t)}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTrafficPolicyGraphResource_Update_APIError exercises TrafficPolicyGraphResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTrafficPolicyGraphResource_Update_APIError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_traffic_policy_graph")
}

// TestTrafficPolicyGraphResource_Update_APIErrorReadBody exercises TrafficPolicyGraphResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTrafficPolicyGraphResource_Update_APIErrorReadBody(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientReadErrorBody(t, 501)}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestTrafficPolicyGraphResource_Update_InvalidJSON exercises TrafficPolicyGraphResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestTrafficPolicyGraphResource_Update_InvalidJSON(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 200, "{{")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestTrafficPolicyGraphResource_Update_MapError exercises TrafficPolicyGraphResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestTrafficPolicyGraphResource_Update_MapError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestTrafficPolicyGraphResource_Delete_Happy exercises TrafficPolicyGraphResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestTrafficPolicyGraphResource_Delete_Happy(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 204, "")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTrafficPolicyGraphResource_Delete_NilClient exercises TrafficPolicyGraphResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTrafficPolicyGraphResource_Delete_NilClient(t *testing.T) {
	r := &TrafficPolicyGraphResource{}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTrafficPolicyGraphResource_Delete_BuildError exercises TrafficPolicyGraphResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTrafficPolicyGraphResource_Delete_BuildError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMalformedBaseURLClient(t)}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTrafficPolicyGraphResource_Delete_SendError exercises TrafficPolicyGraphResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestTrafficPolicyGraphResource_Delete_SendError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newTransportErrorClient(t)}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTrafficPolicyGraphResource_Delete_NotFoundSuccess exercises TrafficPolicyGraphResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestTrafficPolicyGraphResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 404, "")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTrafficPolicyGraphResource_Delete_APIError exercises TrafficPolicyGraphResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTrafficPolicyGraphResource_Delete_APIError(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_traffic_policy_graph")
}

// TestTrafficPolicyGraphResource_Delete_APIErrorReadBody exercises TrafficPolicyGraphResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTrafficPolicyGraphResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &TrafficPolicyGraphResource{client: newMockClientReadErrorBody(t, 501)}
	m := TrafficPolicyGraphResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
