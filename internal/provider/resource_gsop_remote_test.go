package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestGsopResource_Create_Happy exercises GsopResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestGsopResource_Create_Happy(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := GsopResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGsopResource_Create_NilClient exercises GsopResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGsopResource_Create_NilClient(t *testing.T) {
	r := &GsopResource{}
	m := GsopResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGsopResource_Create_BuildError exercises GsopResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGsopResource_Create_BuildError(t *testing.T) {
	r := &GsopResource{client: newMalformedBaseURLClient(t)}
	m := GsopResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGsopResource_Create_SendError exercises GsopResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestGsopResource_Create_SendError(t *testing.T) {
	r := &GsopResource{client: newTransportErrorClient(t)}
	m := GsopResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGsopResource_Create_APIError exercises GsopResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGsopResource_Create_APIError(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GsopResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_gsop")
}

// TestGsopResource_Create_APIErrorReadBody exercises GsopResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGsopResource_Create_APIErrorReadBody(t *testing.T) {
	r := &GsopResource{client: newMockClientReadErrorBody(t, 501)}
	m := GsopResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGsopResource_Create_InvalidJSON exercises GsopResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGsopResource_Create_InvalidJSON(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 201, "{{")}
	m := GsopResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGsopResource_Create_MapError exercises GsopResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGsopResource_Create_MapError(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := GsopResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGsopResource_Create_MissingID exercises GsopResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestGsopResource_Create_MissingID(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 201, "{}")}
	m := GsopResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestGsopResource_Create_LocationFallback exercises GsopResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestGsopResource_Create_LocationFallback(t *testing.T) {
	r := &GsopResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := GsopResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestGsopResource_Read_Happy exercises GsopResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestGsopResource_Read_Happy(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 200, "{}")}
	m := GsopResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestGsopResource_Read_NilClient exercises GsopResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGsopResource_Read_NilClient(t *testing.T) {
	r := &GsopResource{}
	m := GsopResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGsopResource_Read_BuildError exercises GsopResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGsopResource_Read_BuildError(t *testing.T) {
	r := &GsopResource{client: newMalformedBaseURLClient(t)}
	m := GsopResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGsopResource_Read_SendError exercises GsopResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGsopResource_Read_SendError(t *testing.T) {
	r := &GsopResource{client: newTransportErrorClient(t)}
	m := GsopResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGsopResource_Read_NotFound exercises GsopResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestGsopResource_Read_NotFound(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 404, "")}
	m := GsopResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestGsopResource_Read_APIError exercises GsopResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGsopResource_Read_APIError(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GsopResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_gsop")
}

// TestGsopResource_Read_APIErrorReadBody exercises GsopResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGsopResource_Read_APIErrorReadBody(t *testing.T) {
	r := &GsopResource{client: newMockClientReadErrorBody(t, 501)}
	m := GsopResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGsopResource_Read_InvalidJSON exercises GsopResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGsopResource_Read_InvalidJSON(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 200, "{{")}
	m := GsopResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGsopResource_Read_MapError exercises GsopResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGsopResource_Read_MapError(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := GsopResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGsopResource_Update_Happy exercises GsopResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestGsopResource_Update_Happy(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 200, "{}")}
	m := GsopResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGsopResource_Update_NilClient exercises GsopResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGsopResource_Update_NilClient(t *testing.T) {
	r := &GsopResource{}
	m := GsopResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGsopResource_Update_BuildError exercises GsopResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGsopResource_Update_BuildError(t *testing.T) {
	r := &GsopResource{client: newMalformedBaseURLClient(t)}
	m := GsopResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGsopResource_Update_SendError exercises GsopResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestGsopResource_Update_SendError(t *testing.T) {
	r := &GsopResource{client: newTransportErrorClient(t)}
	m := GsopResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGsopResource_Update_APIError exercises GsopResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGsopResource_Update_APIError(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GsopResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_gsop")
}

// TestGsopResource_Update_APIErrorReadBody exercises GsopResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGsopResource_Update_APIErrorReadBody(t *testing.T) {
	r := &GsopResource{client: newMockClientReadErrorBody(t, 501)}
	m := GsopResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGsopResource_Update_InvalidJSON exercises GsopResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGsopResource_Update_InvalidJSON(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 200, "{{")}
	m := GsopResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestGsopResource_Update_MapError exercises GsopResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestGsopResource_Update_MapError(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := GsopResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestGsopResource_Delete_Happy exercises GsopResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestGsopResource_Delete_Happy(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 204, "")}
	m := GsopResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGsopResource_Delete_NilClient exercises GsopResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGsopResource_Delete_NilClient(t *testing.T) {
	r := &GsopResource{}
	m := GsopResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGsopResource_Delete_BuildError exercises GsopResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGsopResource_Delete_BuildError(t *testing.T) {
	r := &GsopResource{client: newMalformedBaseURLClient(t)}
	m := GsopResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGsopResource_Delete_SendError exercises GsopResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestGsopResource_Delete_SendError(t *testing.T) {
	r := &GsopResource{client: newTransportErrorClient(t)}
	m := GsopResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGsopResource_Delete_NotFoundSuccess exercises GsopResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestGsopResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 404, "")}
	m := GsopResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGsopResource_Delete_APIError exercises GsopResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGsopResource_Delete_APIError(t *testing.T) {
	r := &GsopResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GsopResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_gsop")
}

// TestGsopResource_Delete_APIErrorReadBody exercises GsopResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGsopResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &GsopResource{client: newMockClientReadErrorBody(t, 501)}
	m := GsopResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
