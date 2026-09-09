package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestExporterResource_Create_Happy exercises ExporterResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestExporterResource_Create_Happy(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExporterResource_Create_NilClient exercises ExporterResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExporterResource_Create_NilClient(t *testing.T) {
	r := &ExporterResource{}
	m := ExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExporterResource_Create_BuildError exercises ExporterResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExporterResource_Create_BuildError(t *testing.T) {
	r := &ExporterResource{client: newMalformedBaseURLClient(t)}
	m := ExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExporterResource_Create_SendError exercises ExporterResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestExporterResource_Create_SendError(t *testing.T) {
	r := &ExporterResource{client: newTransportErrorClient(t)}
	m := ExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExporterResource_Create_APIError exercises ExporterResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExporterResource_Create_APIError(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_exporter")
}

// TestExporterResource_Create_APIErrorReadBody exercises ExporterResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExporterResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ExporterResource{client: newMockClientReadErrorBody(t, 501)}
	m := ExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestExporterResource_Create_InvalidJSON exercises ExporterResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestExporterResource_Create_InvalidJSON(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 201, "{{")}
	m := ExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestExporterResource_Create_MapError exercises ExporterResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestExporterResource_Create_MapError(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestExporterResource_Create_MissingID exercises ExporterResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestExporterResource_Create_MissingID(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 201, "{}")}
	m := ExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestExporterResource_Create_LocationFallback exercises ExporterResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestExporterResource_Create_LocationFallback(t *testing.T) {
	r := &ExporterResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestExporterResource_Read_Happy exercises ExporterResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestExporterResource_Read_Happy(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 200, "{}")}
	m := ExporterResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestExporterResource_Read_NilClient exercises ExporterResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExporterResource_Read_NilClient(t *testing.T) {
	r := &ExporterResource{}
	m := ExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExporterResource_Read_BuildError exercises ExporterResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExporterResource_Read_BuildError(t *testing.T) {
	r := &ExporterResource{client: newMalformedBaseURLClient(t)}
	m := ExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExporterResource_Read_SendError exercises ExporterResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestExporterResource_Read_SendError(t *testing.T) {
	r := &ExporterResource{client: newTransportErrorClient(t)}
	m := ExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExporterResource_Read_NotFound exercises ExporterResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestExporterResource_Read_NotFound(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 404, "")}
	m := ExporterResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestExporterResource_Read_APIError exercises ExporterResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExporterResource_Read_APIError(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_exporter")
}

// TestExporterResource_Read_APIErrorReadBody exercises ExporterResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExporterResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ExporterResource{client: newMockClientReadErrorBody(t, 501)}
	m := ExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestExporterResource_Read_InvalidJSON exercises ExporterResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestExporterResource_Read_InvalidJSON(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 200, "{{")}
	m := ExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestExporterResource_Read_MapError exercises ExporterResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestExporterResource_Read_MapError(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestExporterResource_Update_Happy exercises ExporterResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestExporterResource_Update_Happy(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 200, "{}")}
	m := ExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExporterResource_Update_NilClient exercises ExporterResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExporterResource_Update_NilClient(t *testing.T) {
	r := &ExporterResource{}
	m := ExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExporterResource_Update_BuildError exercises ExporterResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExporterResource_Update_BuildError(t *testing.T) {
	r := &ExporterResource{client: newMalformedBaseURLClient(t)}
	m := ExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExporterResource_Update_SendError exercises ExporterResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestExporterResource_Update_SendError(t *testing.T) {
	r := &ExporterResource{client: newTransportErrorClient(t)}
	m := ExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExporterResource_Update_APIError exercises ExporterResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExporterResource_Update_APIError(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_exporter")
}

// TestExporterResource_Update_APIErrorReadBody exercises ExporterResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExporterResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ExporterResource{client: newMockClientReadErrorBody(t, 501)}
	m := ExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestExporterResource_Update_InvalidJSON exercises ExporterResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestExporterResource_Update_InvalidJSON(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 200, "{{")}
	m := ExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestExporterResource_Update_MapError exercises ExporterResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestExporterResource_Update_MapError(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestExporterResource_Delete_Happy exercises ExporterResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestExporterResource_Delete_Happy(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 204, "")}
	m := ExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExporterResource_Delete_NilClient exercises ExporterResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExporterResource_Delete_NilClient(t *testing.T) {
	r := &ExporterResource{}
	m := ExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExporterResource_Delete_BuildError exercises ExporterResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExporterResource_Delete_BuildError(t *testing.T) {
	r := &ExporterResource{client: newMalformedBaseURLClient(t)}
	m := ExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExporterResource_Delete_SendError exercises ExporterResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestExporterResource_Delete_SendError(t *testing.T) {
	r := &ExporterResource{client: newTransportErrorClient(t)}
	m := ExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExporterResource_Delete_NotFoundSuccess exercises ExporterResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestExporterResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 404, "")}
	m := ExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExporterResource_Delete_APIError exercises ExporterResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExporterResource_Delete_APIError(t *testing.T) {
	r := &ExporterResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_exporter")
}

// TestExporterResource_Delete_APIErrorReadBody exercises ExporterResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExporterResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ExporterResource{client: newMockClientReadErrorBody(t, 501)}
	m := ExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
