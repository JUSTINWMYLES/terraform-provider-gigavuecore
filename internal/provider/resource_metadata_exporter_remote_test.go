package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestMetadataExporterResource_Create_Happy exercises MetadataExporterResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestMetadataExporterResource_Create_Happy(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := MetadataExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMetadataExporterResource_Create_NilClient exercises MetadataExporterResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMetadataExporterResource_Create_NilClient(t *testing.T) {
	r := &MetadataExporterResource{}
	m := MetadataExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMetadataExporterResource_Create_BuildError exercises MetadataExporterResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMetadataExporterResource_Create_BuildError(t *testing.T) {
	r := &MetadataExporterResource{client: newMalformedBaseURLClient(t)}
	m := MetadataExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMetadataExporterResource_Create_SendError exercises MetadataExporterResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestMetadataExporterResource_Create_SendError(t *testing.T) {
	r := &MetadataExporterResource{client: newTransportErrorClient(t)}
	m := MetadataExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMetadataExporterResource_Create_APIError exercises MetadataExporterResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMetadataExporterResource_Create_APIError(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MetadataExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_metadata_exporter")
}

// TestMetadataExporterResource_Create_APIErrorReadBody exercises MetadataExporterResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMetadataExporterResource_Create_APIErrorReadBody(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientReadErrorBody(t, 501)}
	m := MetadataExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMetadataExporterResource_Create_InvalidJSON exercises MetadataExporterResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMetadataExporterResource_Create_InvalidJSON(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 201, "{{")}
	m := MetadataExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMetadataExporterResource_Create_MapError exercises MetadataExporterResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMetadataExporterResource_Create_MapError(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := MetadataExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMetadataExporterResource_Create_MissingID exercises MetadataExporterResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestMetadataExporterResource_Create_MissingID(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 201, "{}")}
	m := MetadataExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestMetadataExporterResource_Create_LocationFallback exercises MetadataExporterResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestMetadataExporterResource_Create_LocationFallback(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := MetadataExporterResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestMetadataExporterResource_Read_Happy exercises MetadataExporterResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestMetadataExporterResource_Read_Happy(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 200, "{}")}
	m := MetadataExporterResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestMetadataExporterResource_Read_NilClient exercises MetadataExporterResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMetadataExporterResource_Read_NilClient(t *testing.T) {
	r := &MetadataExporterResource{}
	m := MetadataExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMetadataExporterResource_Read_BuildError exercises MetadataExporterResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMetadataExporterResource_Read_BuildError(t *testing.T) {
	r := &MetadataExporterResource{client: newMalformedBaseURLClient(t)}
	m := MetadataExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMetadataExporterResource_Read_SendError exercises MetadataExporterResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestMetadataExporterResource_Read_SendError(t *testing.T) {
	r := &MetadataExporterResource{client: newTransportErrorClient(t)}
	m := MetadataExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMetadataExporterResource_Read_NotFound exercises MetadataExporterResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestMetadataExporterResource_Read_NotFound(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 404, "")}
	m := MetadataExporterResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestMetadataExporterResource_Read_APIError exercises MetadataExporterResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMetadataExporterResource_Read_APIError(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MetadataExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_metadata_exporter")
}

// TestMetadataExporterResource_Read_APIErrorReadBody exercises MetadataExporterResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMetadataExporterResource_Read_APIErrorReadBody(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientReadErrorBody(t, 501)}
	m := MetadataExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMetadataExporterResource_Read_InvalidJSON exercises MetadataExporterResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMetadataExporterResource_Read_InvalidJSON(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 200, "{{")}
	m := MetadataExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMetadataExporterResource_Read_MapError exercises MetadataExporterResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMetadataExporterResource_Read_MapError(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := MetadataExporterResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMetadataExporterResource_Update_Happy exercises MetadataExporterResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestMetadataExporterResource_Update_Happy(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 200, "{}")}
	m := MetadataExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMetadataExporterResource_Update_NilClient exercises MetadataExporterResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMetadataExporterResource_Update_NilClient(t *testing.T) {
	r := &MetadataExporterResource{}
	m := MetadataExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMetadataExporterResource_Update_BuildError exercises MetadataExporterResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMetadataExporterResource_Update_BuildError(t *testing.T) {
	r := &MetadataExporterResource{client: newMalformedBaseURLClient(t)}
	m := MetadataExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMetadataExporterResource_Update_SendError exercises MetadataExporterResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestMetadataExporterResource_Update_SendError(t *testing.T) {
	r := &MetadataExporterResource{client: newTransportErrorClient(t)}
	m := MetadataExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMetadataExporterResource_Update_APIError exercises MetadataExporterResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMetadataExporterResource_Update_APIError(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MetadataExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_metadata_exporter")
}

// TestMetadataExporterResource_Update_APIErrorReadBody exercises MetadataExporterResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMetadataExporterResource_Update_APIErrorReadBody(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientReadErrorBody(t, 501)}
	m := MetadataExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMetadataExporterResource_Update_InvalidJSON exercises MetadataExporterResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMetadataExporterResource_Update_InvalidJSON(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 200, "{{")}
	m := MetadataExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMetadataExporterResource_Update_MapError exercises MetadataExporterResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMetadataExporterResource_Update_MapError(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := MetadataExporterResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMetadataExporterResource_Delete_Happy exercises MetadataExporterResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestMetadataExporterResource_Delete_Happy(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 204, "")}
	m := MetadataExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMetadataExporterResource_Delete_NilClient exercises MetadataExporterResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMetadataExporterResource_Delete_NilClient(t *testing.T) {
	r := &MetadataExporterResource{}
	m := MetadataExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMetadataExporterResource_Delete_BuildError exercises MetadataExporterResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMetadataExporterResource_Delete_BuildError(t *testing.T) {
	r := &MetadataExporterResource{client: newMalformedBaseURLClient(t)}
	m := MetadataExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMetadataExporterResource_Delete_SendError exercises MetadataExporterResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestMetadataExporterResource_Delete_SendError(t *testing.T) {
	r := &MetadataExporterResource{client: newTransportErrorClient(t)}
	m := MetadataExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMetadataExporterResource_Delete_NotFoundSuccess exercises MetadataExporterResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestMetadataExporterResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 404, "")}
	m := MetadataExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMetadataExporterResource_Delete_APIError exercises MetadataExporterResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMetadataExporterResource_Delete_APIError(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MetadataExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_metadata_exporter")
}

// TestMetadataExporterResource_Delete_APIErrorReadBody exercises MetadataExporterResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMetadataExporterResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &MetadataExporterResource{client: newMockClientReadErrorBody(t, 501)}
	m := MetadataExporterResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
