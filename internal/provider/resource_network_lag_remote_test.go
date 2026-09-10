package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNetworkLagResource_Create_Happy exercises NetworkLagResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestNetworkLagResource_Create_Happy(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := NetworkLagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkLagResource_Create_NilClient exercises NetworkLagResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetworkLagResource_Create_NilClient(t *testing.T) {
	r := &NetworkLagResource{}
	m := NetworkLagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetworkLagResource_Create_BuildError exercises NetworkLagResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetworkLagResource_Create_BuildError(t *testing.T) {
	r := &NetworkLagResource{client: newMalformedBaseURLClient(t)}
	m := NetworkLagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetworkLagResource_Create_SendError exercises NetworkLagResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetworkLagResource_Create_SendError(t *testing.T) {
	r := &NetworkLagResource{client: newTransportErrorClient(t)}
	m := NetworkLagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetworkLagResource_Create_APIError exercises NetworkLagResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetworkLagResource_Create_APIError(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetworkLagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_network_lag")
}

// TestNetworkLagResource_Create_APIErrorReadBody exercises NetworkLagResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetworkLagResource_Create_APIErrorReadBody(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetworkLagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNetworkLagResource_Create_InvalidJSON exercises NetworkLagResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNetworkLagResource_Create_InvalidJSON(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 201, "{{")}
	m := NetworkLagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNetworkLagResource_Create_MapError exercises NetworkLagResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNetworkLagResource_Create_MapError(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := NetworkLagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNetworkLagResource_Create_MissingID exercises NetworkLagResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestNetworkLagResource_Create_MissingID(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 201, "{}")}
	m := NetworkLagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestNetworkLagResource_Create_LocationFallback exercises NetworkLagResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestNetworkLagResource_Create_LocationFallback(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := NetworkLagResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestNetworkLagResource_Read_Happy exercises NetworkLagResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestNetworkLagResource_Read_Happy(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 200, "{}")}
	m := NetworkLagResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkLagResource_Read_NilClient exercises NetworkLagResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetworkLagResource_Read_NilClient(t *testing.T) {
	r := &NetworkLagResource{}
	m := NetworkLagResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetworkLagResource_Read_BuildError exercises NetworkLagResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetworkLagResource_Read_BuildError(t *testing.T) {
	r := &NetworkLagResource{client: newMalformedBaseURLClient(t)}
	m := NetworkLagResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetworkLagResource_Read_SendError exercises NetworkLagResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetworkLagResource_Read_SendError(t *testing.T) {
	r := &NetworkLagResource{client: newTransportErrorClient(t)}
	m := NetworkLagResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetworkLagResource_Read_NotFound exercises NetworkLagResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestNetworkLagResource_Read_NotFound(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 404, "")}
	m := NetworkLagResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkLagResource_Read_APIError exercises NetworkLagResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetworkLagResource_Read_APIError(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetworkLagResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_network_lag")
}

// TestNetworkLagResource_Read_APIErrorReadBody exercises NetworkLagResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetworkLagResource_Read_APIErrorReadBody(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetworkLagResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNetworkLagResource_Read_InvalidJSON exercises NetworkLagResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNetworkLagResource_Read_InvalidJSON(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 200, "{{")}
	m := NetworkLagResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNetworkLagResource_Read_MapError exercises NetworkLagResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNetworkLagResource_Read_MapError(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := NetworkLagResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNetworkLagResource_Update_Happy exercises NetworkLagResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestNetworkLagResource_Update_Happy(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 200, "{}")}
	m := NetworkLagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkLagResource_Update_NilClient exercises NetworkLagResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetworkLagResource_Update_NilClient(t *testing.T) {
	r := &NetworkLagResource{}
	m := NetworkLagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetworkLagResource_Update_BuildError exercises NetworkLagResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetworkLagResource_Update_BuildError(t *testing.T) {
	r := &NetworkLagResource{client: newMalformedBaseURLClient(t)}
	m := NetworkLagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetworkLagResource_Update_SendError exercises NetworkLagResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetworkLagResource_Update_SendError(t *testing.T) {
	r := &NetworkLagResource{client: newTransportErrorClient(t)}
	m := NetworkLagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetworkLagResource_Update_APIError exercises NetworkLagResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetworkLagResource_Update_APIError(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetworkLagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_network_lag")
}

// TestNetworkLagResource_Update_APIErrorReadBody exercises NetworkLagResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetworkLagResource_Update_APIErrorReadBody(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetworkLagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNetworkLagResource_Update_InvalidJSON exercises NetworkLagResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNetworkLagResource_Update_InvalidJSON(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 200, "{{")}
	m := NetworkLagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNetworkLagResource_Update_MapError exercises NetworkLagResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNetworkLagResource_Update_MapError(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := NetworkLagResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNetworkLagResource_Delete_Happy exercises NetworkLagResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestNetworkLagResource_Delete_Happy(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 204, "")}
	m := NetworkLagResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkLagResource_Delete_NilClient exercises NetworkLagResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetworkLagResource_Delete_NilClient(t *testing.T) {
	r := &NetworkLagResource{}
	m := NetworkLagResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetworkLagResource_Delete_BuildError exercises NetworkLagResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetworkLagResource_Delete_BuildError(t *testing.T) {
	r := &NetworkLagResource{client: newMalformedBaseURLClient(t)}
	m := NetworkLagResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetworkLagResource_Delete_SendError exercises NetworkLagResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetworkLagResource_Delete_SendError(t *testing.T) {
	r := &NetworkLagResource{client: newTransportErrorClient(t)}
	m := NetworkLagResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetworkLagResource_Delete_NotFoundSuccess exercises NetworkLagResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestNetworkLagResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 404, "")}
	m := NetworkLagResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetworkLagResource_Delete_APIError exercises NetworkLagResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetworkLagResource_Delete_APIError(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetworkLagResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_network_lag")
}

// TestNetworkLagResource_Delete_APIErrorReadBody exercises NetworkLagResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetworkLagResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &NetworkLagResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetworkLagResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
