package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNetflowExporterResource_Create_Happy exercises NetflowExporterResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestNetflowExporterResource_Create_Happy(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := NetflowExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetflowExporterResource_Create_NilClient exercises NetflowExporterResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetflowExporterResource_Create_NilClient(t *testing.T) {
	r := &NetflowExporterResource{}
	m := NetflowExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetflowExporterResource_Create_BuildError exercises NetflowExporterResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetflowExporterResource_Create_BuildError(t *testing.T) {
	r := &NetflowExporterResource{client: newMalformedBaseURLClient(t)}
	m := NetflowExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetflowExporterResource_Create_SendError exercises NetflowExporterResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetflowExporterResource_Create_SendError(t *testing.T) {
	r := &NetflowExporterResource{client: newTransportErrorClient(t)}
	m := NetflowExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetflowExporterResource_Create_APIError exercises NetflowExporterResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetflowExporterResource_Create_APIError(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetflowExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_netflow_exporter")
}

// TestNetflowExporterResource_Create_APIErrorReadBody exercises NetflowExporterResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetflowExporterResource_Create_APIErrorReadBody(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetflowExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNetflowExporterResource_Create_InvalidJSON exercises NetflowExporterResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNetflowExporterResource_Create_InvalidJSON(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 201, "{{")}
	m := NetflowExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNetflowExporterResource_Create_MapError exercises NetflowExporterResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNetflowExporterResource_Create_MapError(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := NetflowExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNetflowExporterResource_Create_MissingID exercises NetflowExporterResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestNetflowExporterResource_Create_MissingID(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 201, "{}")}
	m := NetflowExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestNetflowExporterResource_Create_LocationFallback exercises NetflowExporterResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestNetflowExporterResource_Create_LocationFallback(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := NetflowExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestNetflowExporterResource_Read_Happy exercises NetflowExporterResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestNetflowExporterResource_Read_Happy(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 200, "{}")}
	m := NetflowExporterResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetflowExporterResource_Read_NilClient exercises NetflowExporterResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetflowExporterResource_Read_NilClient(t *testing.T) {
	r := &NetflowExporterResource{}
	m := NetflowExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetflowExporterResource_Read_BuildError exercises NetflowExporterResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetflowExporterResource_Read_BuildError(t *testing.T) {
	r := &NetflowExporterResource{client: newMalformedBaseURLClient(t)}
	m := NetflowExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetflowExporterResource_Read_SendError exercises NetflowExporterResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetflowExporterResource_Read_SendError(t *testing.T) {
	r := &NetflowExporterResource{client: newTransportErrorClient(t)}
	m := NetflowExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetflowExporterResource_Read_NotFound exercises NetflowExporterResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestNetflowExporterResource_Read_NotFound(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 404, "")}
	m := NetflowExporterResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetflowExporterResource_Read_APIError exercises NetflowExporterResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetflowExporterResource_Read_APIError(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetflowExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_netflow_exporter")
}

// TestNetflowExporterResource_Read_APIErrorReadBody exercises NetflowExporterResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetflowExporterResource_Read_APIErrorReadBody(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetflowExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNetflowExporterResource_Read_InvalidJSON exercises NetflowExporterResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNetflowExporterResource_Read_InvalidJSON(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 200, "{{")}
	m := NetflowExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNetflowExporterResource_Read_MapError exercises NetflowExporterResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNetflowExporterResource_Read_MapError(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := NetflowExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNetflowExporterResource_Update_Happy exercises NetflowExporterResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestNetflowExporterResource_Update_Happy(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 200, "{}")}
	m := NetflowExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetflowExporterResource_Update_NilClient exercises NetflowExporterResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetflowExporterResource_Update_NilClient(t *testing.T) {
	r := &NetflowExporterResource{}
	m := NetflowExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetflowExporterResource_Update_BuildError exercises NetflowExporterResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetflowExporterResource_Update_BuildError(t *testing.T) {
	r := &NetflowExporterResource{client: newMalformedBaseURLClient(t)}
	m := NetflowExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetflowExporterResource_Update_SendError exercises NetflowExporterResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetflowExporterResource_Update_SendError(t *testing.T) {
	r := &NetflowExporterResource{client: newTransportErrorClient(t)}
	m := NetflowExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetflowExporterResource_Update_APIError exercises NetflowExporterResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetflowExporterResource_Update_APIError(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetflowExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_netflow_exporter")
}

// TestNetflowExporterResource_Update_APIErrorReadBody exercises NetflowExporterResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetflowExporterResource_Update_APIErrorReadBody(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetflowExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestNetflowExporterResource_Update_InvalidJSON exercises NetflowExporterResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestNetflowExporterResource_Update_InvalidJSON(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 200, "{{")}
	m := NetflowExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestNetflowExporterResource_Update_MapError exercises NetflowExporterResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestNetflowExporterResource_Update_MapError(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := NetflowExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestNetflowExporterResource_Delete_Happy exercises NetflowExporterResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestNetflowExporterResource_Delete_Happy(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 204, "")}
	m := NetflowExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetflowExporterResource_Delete_NilClient exercises NetflowExporterResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNetflowExporterResource_Delete_NilClient(t *testing.T) {
	r := &NetflowExporterResource{}
	m := NetflowExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNetflowExporterResource_Delete_BuildError exercises NetflowExporterResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNetflowExporterResource_Delete_BuildError(t *testing.T) {
	r := &NetflowExporterResource{client: newMalformedBaseURLClient(t)}
	m := NetflowExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNetflowExporterResource_Delete_SendError exercises NetflowExporterResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestNetflowExporterResource_Delete_SendError(t *testing.T) {
	r := &NetflowExporterResource{client: newTransportErrorClient(t)}
	m := NetflowExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNetflowExporterResource_Delete_NotFoundSuccess exercises NetflowExporterResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestNetflowExporterResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 404, "")}
	m := NetflowExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNetflowExporterResource_Delete_APIError exercises NetflowExporterResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNetflowExporterResource_Delete_APIError(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NetflowExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_netflow_exporter")
}

// TestNetflowExporterResource_Delete_APIErrorReadBody exercises NetflowExporterResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNetflowExporterResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &NetflowExporterResource{client: newMockClientReadErrorBody(t, 501)}
	m := NetflowExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
