package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestExporterGroupResource_Create_Happy exercises ExporterGroupResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestExporterGroupResource_Create_Happy(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ExporterGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExporterGroupResource_Create_NilClient exercises ExporterGroupResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExporterGroupResource_Create_NilClient(t *testing.T) {
	r := &ExporterGroupResource{}
	m := ExporterGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExporterGroupResource_Create_BuildError exercises ExporterGroupResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExporterGroupResource_Create_BuildError(t *testing.T) {
	r := &ExporterGroupResource{client: newMalformedBaseURLClient(t)}
	m := ExporterGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExporterGroupResource_Create_SendError exercises ExporterGroupResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestExporterGroupResource_Create_SendError(t *testing.T) {
	r := &ExporterGroupResource{client: newTransportErrorClient(t)}
	m := ExporterGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExporterGroupResource_Create_APIError exercises ExporterGroupResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExporterGroupResource_Create_APIError(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ExporterGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_exporter_group")
}

// TestExporterGroupResource_Create_APIErrorReadBody exercises ExporterGroupResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExporterGroupResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := ExporterGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestExporterGroupResource_Create_InvalidJSON exercises ExporterGroupResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestExporterGroupResource_Create_InvalidJSON(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := ExporterGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestExporterGroupResource_Create_MapError exercises ExporterGroupResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestExporterGroupResource_Create_MapError(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ExporterGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestExporterGroupResource_Create_MissingID exercises ExporterGroupResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestExporterGroupResource_Create_MissingID(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := ExporterGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestExporterGroupResource_Create_LocationFallback exercises ExporterGroupResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestExporterGroupResource_Create_LocationFallback(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ExporterGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestExporterGroupResource_Read_Happy exercises ExporterGroupResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestExporterGroupResource_Read_Happy(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := ExporterGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestExporterGroupResource_Read_NilClient exercises ExporterGroupResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExporterGroupResource_Read_NilClient(t *testing.T) {
	r := &ExporterGroupResource{}
	m := ExporterGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExporterGroupResource_Read_BuildError exercises ExporterGroupResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExporterGroupResource_Read_BuildError(t *testing.T) {
	r := &ExporterGroupResource{client: newMalformedBaseURLClient(t)}
	m := ExporterGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExporterGroupResource_Read_SendError exercises ExporterGroupResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestExporterGroupResource_Read_SendError(t *testing.T) {
	r := &ExporterGroupResource{client: newTransportErrorClient(t)}
	m := ExporterGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExporterGroupResource_Read_NotFound exercises ExporterGroupResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestExporterGroupResource_Read_NotFound(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 404, "")}
	m := ExporterGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestExporterGroupResource_Read_APIError exercises ExporterGroupResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExporterGroupResource_Read_APIError(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ExporterGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_exporter_group")
}

// TestExporterGroupResource_Read_APIErrorReadBody exercises ExporterGroupResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExporterGroupResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := ExporterGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestExporterGroupResource_Read_InvalidJSON exercises ExporterGroupResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestExporterGroupResource_Read_InvalidJSON(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := ExporterGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestExporterGroupResource_Read_MapError exercises ExporterGroupResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestExporterGroupResource_Read_MapError(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ExporterGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestExporterGroupResource_Update_Happy exercises ExporterGroupResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestExporterGroupResource_Update_Happy(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := ExporterGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExporterGroupResource_Update_NilClient exercises ExporterGroupResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExporterGroupResource_Update_NilClient(t *testing.T) {
	r := &ExporterGroupResource{}
	m := ExporterGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExporterGroupResource_Update_BuildError exercises ExporterGroupResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExporterGroupResource_Update_BuildError(t *testing.T) {
	r := &ExporterGroupResource{client: newMalformedBaseURLClient(t)}
	m := ExporterGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExporterGroupResource_Update_SendError exercises ExporterGroupResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestExporterGroupResource_Update_SendError(t *testing.T) {
	r := &ExporterGroupResource{client: newTransportErrorClient(t)}
	m := ExporterGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExporterGroupResource_Update_APIError exercises ExporterGroupResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExporterGroupResource_Update_APIError(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ExporterGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_exporter_group")
}

// TestExporterGroupResource_Update_APIErrorReadBody exercises ExporterGroupResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExporterGroupResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := ExporterGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestExporterGroupResource_Update_InvalidJSON exercises ExporterGroupResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestExporterGroupResource_Update_InvalidJSON(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := ExporterGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestExporterGroupResource_Update_MapError exercises ExporterGroupResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestExporterGroupResource_Update_MapError(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ExporterGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestExporterGroupResource_Delete_Happy exercises ExporterGroupResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestExporterGroupResource_Delete_Happy(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 204, "")}
	m := ExporterGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExporterGroupResource_Delete_NilClient exercises ExporterGroupResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExporterGroupResource_Delete_NilClient(t *testing.T) {
	r := &ExporterGroupResource{}
	m := ExporterGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExporterGroupResource_Delete_BuildError exercises ExporterGroupResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExporterGroupResource_Delete_BuildError(t *testing.T) {
	r := &ExporterGroupResource{client: newMalformedBaseURLClient(t)}
	m := ExporterGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExporterGroupResource_Delete_SendError exercises ExporterGroupResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestExporterGroupResource_Delete_SendError(t *testing.T) {
	r := &ExporterGroupResource{client: newTransportErrorClient(t)}
	m := ExporterGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExporterGroupResource_Delete_NotFoundSuccess exercises ExporterGroupResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestExporterGroupResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 404, "")}
	m := ExporterGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExporterGroupResource_Delete_APIError exercises ExporterGroupResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExporterGroupResource_Delete_APIError(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ExporterGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_exporter_group")
}

// TestExporterGroupResource_Delete_APIErrorReadBody exercises ExporterGroupResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExporterGroupResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ExporterGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := ExporterGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
