package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSolutionResource_Create_Happy exercises SolutionResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestSolutionResource_Create_Happy(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 201, "{\"solution_alias\":\"example-id\"}")}
	m := SolutionResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSolutionResource_Create_NilClient exercises SolutionResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSolutionResource_Create_NilClient(t *testing.T) {
	r := &SolutionResource{}
	m := SolutionResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSolutionResource_Create_BuildError exercises SolutionResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSolutionResource_Create_BuildError(t *testing.T) {
	r := &SolutionResource{client: newMalformedBaseURLClient(t)}
	m := SolutionResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSolutionResource_Create_SendError exercises SolutionResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestSolutionResource_Create_SendError(t *testing.T) {
	r := &SolutionResource{client: newTransportErrorClient(t)}
	m := SolutionResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSolutionResource_Create_APIError exercises SolutionResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSolutionResource_Create_APIError(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SolutionResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_solution")
}

// TestSolutionResource_Create_APIErrorReadBody exercises SolutionResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSolutionResource_Create_APIErrorReadBody(t *testing.T) {
	r := &SolutionResource{client: newMockClientReadErrorBody(t, 501)}
	m := SolutionResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSolutionResource_Create_InvalidJSON exercises SolutionResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSolutionResource_Create_InvalidJSON(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 201, "{{")}
	m := SolutionResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSolutionResource_Create_MapError exercises SolutionResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSolutionResource_Create_MapError(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 201, "{\"solution_alias\":12345}")}
	m := SolutionResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSolutionResource_Create_MissingID exercises SolutionResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestSolutionResource_Create_MissingID(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 201, "{}")}
	m := SolutionResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestSolutionResource_Create_LocationFallback exercises SolutionResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestSolutionResource_Create_LocationFallback(t *testing.T) {
	r := &SolutionResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := SolutionResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.SolutionAlias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.SolutionAlias.ValueString(), "example-id")
	}
}

// TestSolutionResource_Read_Happy exercises SolutionResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestSolutionResource_Read_Happy(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 200, "{}")}
	m := SolutionResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSolutionResource_Read_NilClient exercises SolutionResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSolutionResource_Read_NilClient(t *testing.T) {
	r := &SolutionResource{}
	m := SolutionResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSolutionResource_Read_BuildError exercises SolutionResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSolutionResource_Read_BuildError(t *testing.T) {
	r := &SolutionResource{client: newMalformedBaseURLClient(t)}
	m := SolutionResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSolutionResource_Read_SendError exercises SolutionResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestSolutionResource_Read_SendError(t *testing.T) {
	r := &SolutionResource{client: newTransportErrorClient(t)}
	m := SolutionResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSolutionResource_Read_NotFound exercises SolutionResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestSolutionResource_Read_NotFound(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 404, "")}
	m := SolutionResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSolutionResource_Read_APIError exercises SolutionResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSolutionResource_Read_APIError(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SolutionResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_solution")
}

// TestSolutionResource_Read_APIErrorReadBody exercises SolutionResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSolutionResource_Read_APIErrorReadBody(t *testing.T) {
	r := &SolutionResource{client: newMockClientReadErrorBody(t, 501)}
	m := SolutionResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSolutionResource_Read_InvalidJSON exercises SolutionResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSolutionResource_Read_InvalidJSON(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 200, "{{")}
	m := SolutionResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSolutionResource_Read_MapError exercises SolutionResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSolutionResource_Read_MapError(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 200, "{\"solution_alias\":12345}")}
	m := SolutionResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSolutionResource_Update_Happy exercises SolutionResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestSolutionResource_Update_Happy(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 200, "{}")}
	m := SolutionResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSolutionResource_Update_NilClient exercises SolutionResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSolutionResource_Update_NilClient(t *testing.T) {
	r := &SolutionResource{}
	m := SolutionResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSolutionResource_Update_BuildError exercises SolutionResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSolutionResource_Update_BuildError(t *testing.T) {
	r := &SolutionResource{client: newMalformedBaseURLClient(t)}
	m := SolutionResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSolutionResource_Update_SendError exercises SolutionResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestSolutionResource_Update_SendError(t *testing.T) {
	r := &SolutionResource{client: newTransportErrorClient(t)}
	m := SolutionResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSolutionResource_Update_APIError exercises SolutionResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSolutionResource_Update_APIError(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SolutionResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_solution")
}

// TestSolutionResource_Update_APIErrorReadBody exercises SolutionResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSolutionResource_Update_APIErrorReadBody(t *testing.T) {
	r := &SolutionResource{client: newMockClientReadErrorBody(t, 501)}
	m := SolutionResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSolutionResource_Update_InvalidJSON exercises SolutionResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSolutionResource_Update_InvalidJSON(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 200, "{{")}
	m := SolutionResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSolutionResource_Update_MapError exercises SolutionResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSolutionResource_Update_MapError(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 200, "{\"solution_alias\":12345}")}
	m := SolutionResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSolutionResource_Delete_Happy exercises SolutionResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestSolutionResource_Delete_Happy(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 204, "")}
	m := SolutionResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSolutionResource_Delete_NilClient exercises SolutionResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSolutionResource_Delete_NilClient(t *testing.T) {
	r := &SolutionResource{}
	m := SolutionResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSolutionResource_Delete_BuildError exercises SolutionResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSolutionResource_Delete_BuildError(t *testing.T) {
	r := &SolutionResource{client: newMalformedBaseURLClient(t)}
	m := SolutionResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSolutionResource_Delete_SendError exercises SolutionResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestSolutionResource_Delete_SendError(t *testing.T) {
	r := &SolutionResource{client: newTransportErrorClient(t)}
	m := SolutionResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSolutionResource_Delete_NotFoundSuccess exercises SolutionResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestSolutionResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 404, "")}
	m := SolutionResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSolutionResource_Delete_APIError exercises SolutionResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSolutionResource_Delete_APIError(t *testing.T) {
	r := &SolutionResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SolutionResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_solution")
}

// TestSolutionResource_Delete_APIErrorReadBody exercises SolutionResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSolutionResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &SolutionResource{client: newMockClientReadErrorBody(t, 501)}
	m := SolutionResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
