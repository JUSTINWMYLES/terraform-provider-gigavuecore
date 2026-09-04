package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestExportTargetResource_Create_Happy exercises ExportTargetResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestExportTargetResource_Create_Happy(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ExportTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExportTargetResource_Create_NilClient exercises ExportTargetResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExportTargetResource_Create_NilClient(t *testing.T) {
	r := &ExportTargetResource{}
	m := ExportTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExportTargetResource_Create_BuildError exercises ExportTargetResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExportTargetResource_Create_BuildError(t *testing.T) {
	r := &ExportTargetResource{client: newMalformedBaseURLClient(t)}
	m := ExportTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExportTargetResource_Create_SendError exercises ExportTargetResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestExportTargetResource_Create_SendError(t *testing.T) {
	r := &ExportTargetResource{client: newTransportErrorClient(t)}
	m := ExportTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExportTargetResource_Create_APIError exercises ExportTargetResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExportTargetResource_Create_APIError(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ExportTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_export_target")
}

// TestExportTargetResource_Create_APIErrorReadBody exercises ExportTargetResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExportTargetResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientReadErrorBody(t, 501)}
	m := ExportTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestExportTargetResource_Create_InvalidJSON exercises ExportTargetResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestExportTargetResource_Create_InvalidJSON(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 201, "{{")}
	m := ExportTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestExportTargetResource_Create_MapError exercises ExportTargetResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestExportTargetResource_Create_MapError(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ExportTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestExportTargetResource_Create_MissingID exercises ExportTargetResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestExportTargetResource_Create_MissingID(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 201, "{}")}
	m := ExportTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestExportTargetResource_Create_LocationFallback exercises ExportTargetResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestExportTargetResource_Create_LocationFallback(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ExportTargetResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestExportTargetResource_Read_Happy exercises ExportTargetResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestExportTargetResource_Read_Happy(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 200, "{}")}
	m := ExportTargetResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestExportTargetResource_Read_NilClient exercises ExportTargetResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExportTargetResource_Read_NilClient(t *testing.T) {
	r := &ExportTargetResource{}
	m := ExportTargetResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExportTargetResource_Read_BuildError exercises ExportTargetResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExportTargetResource_Read_BuildError(t *testing.T) {
	r := &ExportTargetResource{client: newMalformedBaseURLClient(t)}
	m := ExportTargetResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExportTargetResource_Read_SendError exercises ExportTargetResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestExportTargetResource_Read_SendError(t *testing.T) {
	r := &ExportTargetResource{client: newTransportErrorClient(t)}
	m := ExportTargetResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExportTargetResource_Read_NotFound exercises ExportTargetResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestExportTargetResource_Read_NotFound(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 404, "")}
	m := ExportTargetResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestExportTargetResource_Read_APIError exercises ExportTargetResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExportTargetResource_Read_APIError(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ExportTargetResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_export_target")
}

// TestExportTargetResource_Read_APIErrorReadBody exercises ExportTargetResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExportTargetResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientReadErrorBody(t, 501)}
	m := ExportTargetResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestExportTargetResource_Read_InvalidJSON exercises ExportTargetResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestExportTargetResource_Read_InvalidJSON(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 200, "{{")}
	m := ExportTargetResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestExportTargetResource_Read_MapError exercises ExportTargetResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestExportTargetResource_Read_MapError(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ExportTargetResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestExportTargetResource_Update_Happy exercises ExportTargetResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestExportTargetResource_Update_Happy(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 200, "{}")}
	m := ExportTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExportTargetResource_Update_NilClient exercises ExportTargetResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExportTargetResource_Update_NilClient(t *testing.T) {
	r := &ExportTargetResource{}
	m := ExportTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExportTargetResource_Update_BuildError exercises ExportTargetResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExportTargetResource_Update_BuildError(t *testing.T) {
	r := &ExportTargetResource{client: newMalformedBaseURLClient(t)}
	m := ExportTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExportTargetResource_Update_SendError exercises ExportTargetResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestExportTargetResource_Update_SendError(t *testing.T) {
	r := &ExportTargetResource{client: newTransportErrorClient(t)}
	m := ExportTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExportTargetResource_Update_APIError exercises ExportTargetResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExportTargetResource_Update_APIError(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ExportTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_export_target")
}

// TestExportTargetResource_Update_APIErrorReadBody exercises ExportTargetResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExportTargetResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientReadErrorBody(t, 501)}
	m := ExportTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestExportTargetResource_Update_InvalidJSON exercises ExportTargetResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestExportTargetResource_Update_InvalidJSON(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 200, "{{")}
	m := ExportTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestExportTargetResource_Update_MapError exercises ExportTargetResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestExportTargetResource_Update_MapError(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ExportTargetResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestExportTargetResource_Delete_Happy exercises ExportTargetResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestExportTargetResource_Delete_Happy(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 204, "")}
	m := ExportTargetResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExportTargetResource_Delete_NilClient exercises ExportTargetResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExportTargetResource_Delete_NilClient(t *testing.T) {
	r := &ExportTargetResource{}
	m := ExportTargetResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExportTargetResource_Delete_BuildError exercises ExportTargetResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExportTargetResource_Delete_BuildError(t *testing.T) {
	r := &ExportTargetResource{client: newMalformedBaseURLClient(t)}
	m := ExportTargetResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExportTargetResource_Delete_SendError exercises ExportTargetResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestExportTargetResource_Delete_SendError(t *testing.T) {
	r := &ExportTargetResource{client: newTransportErrorClient(t)}
	m := ExportTargetResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExportTargetResource_Delete_NotFoundSuccess exercises ExportTargetResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestExportTargetResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 404, "")}
	m := ExportTargetResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExportTargetResource_Delete_APIError exercises ExportTargetResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExportTargetResource_Delete_APIError(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ExportTargetResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_export_target")
}

// TestExportTargetResource_Delete_APIErrorReadBody exercises ExportTargetResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExportTargetResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ExportTargetResource{client: newMockClientReadErrorBody(t, 501)}
	m := ExportTargetResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
