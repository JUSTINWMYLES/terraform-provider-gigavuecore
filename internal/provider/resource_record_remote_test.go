package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestRecordResource_Create_Happy exercises RecordResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestRecordResource_Create_Happy(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := RecordResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRecordResource_Create_NilClient exercises RecordResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRecordResource_Create_NilClient(t *testing.T) {
	r := &RecordResource{}
	m := RecordResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRecordResource_Create_BuildError exercises RecordResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRecordResource_Create_BuildError(t *testing.T) {
	r := &RecordResource{client: newMalformedBaseURLClient(t)}
	m := RecordResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRecordResource_Create_SendError exercises RecordResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestRecordResource_Create_SendError(t *testing.T) {
	r := &RecordResource{client: newTransportErrorClient(t)}
	m := RecordResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRecordResource_Create_APIError exercises RecordResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRecordResource_Create_APIError(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RecordResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_record")
}

// TestRecordResource_Create_APIErrorReadBody exercises RecordResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRecordResource_Create_APIErrorReadBody(t *testing.T) {
	r := &RecordResource{client: newMockClientReadErrorBody(t, 501)}
	m := RecordResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestRecordResource_Create_InvalidJSON exercises RecordResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestRecordResource_Create_InvalidJSON(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 201, "{{")}
	m := RecordResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestRecordResource_Create_MapError exercises RecordResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestRecordResource_Create_MapError(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := RecordResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestRecordResource_Create_MissingID exercises RecordResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestRecordResource_Create_MissingID(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 201, "{}")}
	m := RecordResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestRecordResource_Create_LocationFallback exercises RecordResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestRecordResource_Create_LocationFallback(t *testing.T) {
	r := &RecordResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := RecordResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestRecordResource_Read_Happy exercises RecordResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestRecordResource_Read_Happy(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 200, "{}")}
	m := RecordResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestRecordResource_Read_NilClient exercises RecordResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRecordResource_Read_NilClient(t *testing.T) {
	r := &RecordResource{}
	m := RecordResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRecordResource_Read_BuildError exercises RecordResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRecordResource_Read_BuildError(t *testing.T) {
	r := &RecordResource{client: newMalformedBaseURLClient(t)}
	m := RecordResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRecordResource_Read_SendError exercises RecordResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestRecordResource_Read_SendError(t *testing.T) {
	r := &RecordResource{client: newTransportErrorClient(t)}
	m := RecordResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRecordResource_Read_NotFound exercises RecordResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestRecordResource_Read_NotFound(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 404, "")}
	m := RecordResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestRecordResource_Read_APIError exercises RecordResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRecordResource_Read_APIError(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RecordResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_record")
}

// TestRecordResource_Read_APIErrorReadBody exercises RecordResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRecordResource_Read_APIErrorReadBody(t *testing.T) {
	r := &RecordResource{client: newMockClientReadErrorBody(t, 501)}
	m := RecordResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestRecordResource_Read_InvalidJSON exercises RecordResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestRecordResource_Read_InvalidJSON(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 200, "{{")}
	m := RecordResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestRecordResource_Read_MapError exercises RecordResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestRecordResource_Read_MapError(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := RecordResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestRecordResource_Update_Happy exercises RecordResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestRecordResource_Update_Happy(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 200, "{}")}
	m := RecordResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRecordResource_Update_NilClient exercises RecordResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRecordResource_Update_NilClient(t *testing.T) {
	r := &RecordResource{}
	m := RecordResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRecordResource_Update_BuildError exercises RecordResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRecordResource_Update_BuildError(t *testing.T) {
	r := &RecordResource{client: newMalformedBaseURLClient(t)}
	m := RecordResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRecordResource_Update_SendError exercises RecordResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestRecordResource_Update_SendError(t *testing.T) {
	r := &RecordResource{client: newTransportErrorClient(t)}
	m := RecordResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRecordResource_Update_APIError exercises RecordResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRecordResource_Update_APIError(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RecordResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_record")
}

// TestRecordResource_Update_APIErrorReadBody exercises RecordResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRecordResource_Update_APIErrorReadBody(t *testing.T) {
	r := &RecordResource{client: newMockClientReadErrorBody(t, 501)}
	m := RecordResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestRecordResource_Update_InvalidJSON exercises RecordResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestRecordResource_Update_InvalidJSON(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 200, "{{")}
	m := RecordResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestRecordResource_Update_MapError exercises RecordResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestRecordResource_Update_MapError(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := RecordResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestRecordResource_Delete_Happy exercises RecordResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestRecordResource_Delete_Happy(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 204, "")}
	m := RecordResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRecordResource_Delete_NilClient exercises RecordResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRecordResource_Delete_NilClient(t *testing.T) {
	r := &RecordResource{}
	m := RecordResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRecordResource_Delete_BuildError exercises RecordResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRecordResource_Delete_BuildError(t *testing.T) {
	r := &RecordResource{client: newMalformedBaseURLClient(t)}
	m := RecordResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRecordResource_Delete_SendError exercises RecordResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestRecordResource_Delete_SendError(t *testing.T) {
	r := &RecordResource{client: newTransportErrorClient(t)}
	m := RecordResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRecordResource_Delete_NotFoundSuccess exercises RecordResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestRecordResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 404, "")}
	m := RecordResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRecordResource_Delete_APIError exercises RecordResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRecordResource_Delete_APIError(t *testing.T) {
	r := &RecordResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RecordResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_record")
}

// TestRecordResource_Delete_APIErrorReadBody exercises RecordResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRecordResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &RecordResource{client: newMockClientReadErrorBody(t, 501)}
	m := RecordResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
