package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestMobilityResource_Create_Happy exercises MobilityResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestMobilityResource_Create_Happy(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 207, "{\"solution_alias\":\"example-id\"}")}
	m := MobilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMobilityResource_Create_NilClient exercises MobilityResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMobilityResource_Create_NilClient(t *testing.T) {
	r := &MobilityResource{}
	m := MobilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMobilityResource_Create_BuildError exercises MobilityResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMobilityResource_Create_BuildError(t *testing.T) {
	r := &MobilityResource{client: newMalformedBaseURLClient(t)}
	m := MobilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMobilityResource_Create_SendError exercises MobilityResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestMobilityResource_Create_SendError(t *testing.T) {
	r := &MobilityResource{client: newTransportErrorClient(t)}
	m := MobilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMobilityResource_Create_APIError exercises MobilityResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMobilityResource_Create_APIError(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MobilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_mobility")
}

// TestMobilityResource_Create_APIErrorReadBody exercises MobilityResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMobilityResource_Create_APIErrorReadBody(t *testing.T) {
	r := &MobilityResource{client: newMockClientReadErrorBody(t, 501)}
	m := MobilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMobilityResource_Create_InvalidJSON exercises MobilityResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMobilityResource_Create_InvalidJSON(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 207, "{{")}
	m := MobilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMobilityResource_Create_MapError exercises MobilityResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMobilityResource_Create_MapError(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 207, "{\"solution_alias\":12345}")}
	m := MobilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMobilityResource_Create_MissingID exercises MobilityResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestMobilityResource_Create_MissingID(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 207, "{}")}
	m := MobilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestMobilityResource_Create_LocationFallback exercises MobilityResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestMobilityResource_Create_LocationFallback(t *testing.T) {
	r := &MobilityResource{client: newMockClientWithLocation(t, 207, "http://example.test/folders/example-id", "{}")}
	m := MobilityResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.SolutionAlias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.SolutionAlias.ValueString(), "example-id")
	}
}

// TestMobilityResource_Read_Happy exercises MobilityResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestMobilityResource_Read_Happy(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 200, "{}")}
	m := MobilityResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestMobilityResource_Read_NilClient exercises MobilityResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMobilityResource_Read_NilClient(t *testing.T) {
	r := &MobilityResource{}
	m := MobilityResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMobilityResource_Read_BuildError exercises MobilityResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMobilityResource_Read_BuildError(t *testing.T) {
	r := &MobilityResource{client: newMalformedBaseURLClient(t)}
	m := MobilityResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMobilityResource_Read_SendError exercises MobilityResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestMobilityResource_Read_SendError(t *testing.T) {
	r := &MobilityResource{client: newTransportErrorClient(t)}
	m := MobilityResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMobilityResource_Read_NotFound exercises MobilityResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestMobilityResource_Read_NotFound(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 404, "")}
	m := MobilityResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestMobilityResource_Read_APIError exercises MobilityResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMobilityResource_Read_APIError(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MobilityResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_mobility")
}

// TestMobilityResource_Read_APIErrorReadBody exercises MobilityResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMobilityResource_Read_APIErrorReadBody(t *testing.T) {
	r := &MobilityResource{client: newMockClientReadErrorBody(t, 501)}
	m := MobilityResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMobilityResource_Read_InvalidJSON exercises MobilityResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMobilityResource_Read_InvalidJSON(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 200, "{{")}
	m := MobilityResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMobilityResource_Read_MapError exercises MobilityResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMobilityResource_Read_MapError(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 200, "{\"solution_alias\":12345}")}
	m := MobilityResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMobilityResource_Update_Happy exercises MobilityResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestMobilityResource_Update_Happy(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 207, "{}")}
	m := MobilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMobilityResource_Update_NilClient exercises MobilityResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMobilityResource_Update_NilClient(t *testing.T) {
	r := &MobilityResource{}
	m := MobilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMobilityResource_Update_BuildError exercises MobilityResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMobilityResource_Update_BuildError(t *testing.T) {
	r := &MobilityResource{client: newMalformedBaseURLClient(t)}
	m := MobilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMobilityResource_Update_SendError exercises MobilityResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestMobilityResource_Update_SendError(t *testing.T) {
	r := &MobilityResource{client: newTransportErrorClient(t)}
	m := MobilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMobilityResource_Update_APIError exercises MobilityResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMobilityResource_Update_APIError(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MobilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_mobility")
}

// TestMobilityResource_Update_APIErrorReadBody exercises MobilityResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMobilityResource_Update_APIErrorReadBody(t *testing.T) {
	r := &MobilityResource{client: newMockClientReadErrorBody(t, 501)}
	m := MobilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestMobilityResource_Update_InvalidJSON exercises MobilityResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestMobilityResource_Update_InvalidJSON(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 207, "{{")}
	m := MobilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestMobilityResource_Update_MapError exercises MobilityResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestMobilityResource_Update_MapError(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 207, "{\"solution_alias\":12345}")}
	m := MobilityResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestMobilityResource_Delete_Happy exercises MobilityResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestMobilityResource_Delete_Happy(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 207, "")}
	m := MobilityResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMobilityResource_Delete_NilClient exercises MobilityResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMobilityResource_Delete_NilClient(t *testing.T) {
	r := &MobilityResource{}
	m := MobilityResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMobilityResource_Delete_BuildError exercises MobilityResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMobilityResource_Delete_BuildError(t *testing.T) {
	r := &MobilityResource{client: newMalformedBaseURLClient(t)}
	m := MobilityResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMobilityResource_Delete_SendError exercises MobilityResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestMobilityResource_Delete_SendError(t *testing.T) {
	r := &MobilityResource{client: newTransportErrorClient(t)}
	m := MobilityResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMobilityResource_Delete_NotFoundSuccess exercises MobilityResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestMobilityResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 404, "")}
	m := MobilityResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMobilityResource_Delete_APIError exercises MobilityResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMobilityResource_Delete_APIError(t *testing.T) {
	r := &MobilityResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MobilityResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_mobility")
}

// TestMobilityResource_Delete_APIErrorReadBody exercises MobilityResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMobilityResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &MobilityResource{client: newMockClientReadErrorBody(t, 501)}
	m := MobilityResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
