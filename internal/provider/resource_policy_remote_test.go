package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestPolicyResource_Create_Happy exercises PolicyResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestPolicyResource_Create_Happy(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 201, "{\"policy_id\":\"example-id\"}")}
	m := PolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPolicyResource_Create_NilClient exercises PolicyResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPolicyResource_Create_NilClient(t *testing.T) {
	r := &PolicyResource{}
	m := PolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPolicyResource_Create_BuildError exercises PolicyResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPolicyResource_Create_BuildError(t *testing.T) {
	r := &PolicyResource{client: newMalformedBaseURLClient(t)}
	m := PolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPolicyResource_Create_SendError exercises PolicyResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestPolicyResource_Create_SendError(t *testing.T) {
	r := &PolicyResource{client: newTransportErrorClient(t)}
	m := PolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPolicyResource_Create_APIError exercises PolicyResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPolicyResource_Create_APIError(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_policy")
}

// TestPolicyResource_Create_APIErrorReadBody exercises PolicyResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPolicyResource_Create_APIErrorReadBody(t *testing.T) {
	r := &PolicyResource{client: newMockClientReadErrorBody(t, 501)}
	m := PolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPolicyResource_Create_InvalidJSON exercises PolicyResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPolicyResource_Create_InvalidJSON(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 201, "{{")}
	m := PolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPolicyResource_Create_MapError exercises PolicyResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPolicyResource_Create_MapError(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 201, "{\"policy_id\":12345}")}
	m := PolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPolicyResource_Create_MissingID exercises PolicyResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestPolicyResource_Create_MissingID(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 201, "{}")}
	m := PolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestPolicyResource_Create_LocationFallback exercises PolicyResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestPolicyResource_Create_LocationFallback(t *testing.T) {
	r := &PolicyResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := PolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.PolicyId.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.PolicyId.ValueString(), "example-id")
	}
}

// TestPolicyResource_Read_Happy exercises PolicyResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestPolicyResource_Read_Happy(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 200, "{}")}
	m := PolicyResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPolicyResource_Read_NilClient exercises PolicyResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPolicyResource_Read_NilClient(t *testing.T) {
	r := &PolicyResource{}
	m := PolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPolicyResource_Read_BuildError exercises PolicyResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPolicyResource_Read_BuildError(t *testing.T) {
	r := &PolicyResource{client: newMalformedBaseURLClient(t)}
	m := PolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPolicyResource_Read_SendError exercises PolicyResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestPolicyResource_Read_SendError(t *testing.T) {
	r := &PolicyResource{client: newTransportErrorClient(t)}
	m := PolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPolicyResource_Read_NotFound exercises PolicyResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestPolicyResource_Read_NotFound(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 404, "")}
	m := PolicyResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPolicyResource_Read_APIError exercises PolicyResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPolicyResource_Read_APIError(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_policy")
}

// TestPolicyResource_Read_APIErrorReadBody exercises PolicyResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPolicyResource_Read_APIErrorReadBody(t *testing.T) {
	r := &PolicyResource{client: newMockClientReadErrorBody(t, 501)}
	m := PolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPolicyResource_Read_InvalidJSON exercises PolicyResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPolicyResource_Read_InvalidJSON(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 200, "{{")}
	m := PolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPolicyResource_Read_MapError exercises PolicyResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPolicyResource_Read_MapError(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 200, "{\"policy_id\":12345}")}
	m := PolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPolicyResource_Update_Happy exercises PolicyResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestPolicyResource_Update_Happy(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 200, "{}")}
	m := PolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPolicyResource_Update_NilClient exercises PolicyResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPolicyResource_Update_NilClient(t *testing.T) {
	r := &PolicyResource{}
	m := PolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPolicyResource_Update_BuildError exercises PolicyResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPolicyResource_Update_BuildError(t *testing.T) {
	r := &PolicyResource{client: newMalformedBaseURLClient(t)}
	m := PolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPolicyResource_Update_SendError exercises PolicyResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestPolicyResource_Update_SendError(t *testing.T) {
	r := &PolicyResource{client: newTransportErrorClient(t)}
	m := PolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPolicyResource_Update_APIError exercises PolicyResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPolicyResource_Update_APIError(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_policy")
}

// TestPolicyResource_Update_APIErrorReadBody exercises PolicyResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPolicyResource_Update_APIErrorReadBody(t *testing.T) {
	r := &PolicyResource{client: newMockClientReadErrorBody(t, 501)}
	m := PolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPolicyResource_Update_InvalidJSON exercises PolicyResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPolicyResource_Update_InvalidJSON(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 200, "{{")}
	m := PolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPolicyResource_Update_MapError exercises PolicyResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPolicyResource_Update_MapError(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 200, "{\"policy_id\":12345}")}
	m := PolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPolicyResource_Delete_Happy exercises PolicyResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestPolicyResource_Delete_Happy(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 204, "")}
	m := PolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPolicyResource_Delete_NilClient exercises PolicyResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPolicyResource_Delete_NilClient(t *testing.T) {
	r := &PolicyResource{}
	m := PolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPolicyResource_Delete_BuildError exercises PolicyResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPolicyResource_Delete_BuildError(t *testing.T) {
	r := &PolicyResource{client: newMalformedBaseURLClient(t)}
	m := PolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPolicyResource_Delete_SendError exercises PolicyResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestPolicyResource_Delete_SendError(t *testing.T) {
	r := &PolicyResource{client: newTransportErrorClient(t)}
	m := PolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPolicyResource_Delete_NotFoundSuccess exercises PolicyResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestPolicyResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 404, "")}
	m := PolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPolicyResource_Delete_APIError exercises PolicyResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPolicyResource_Delete_APIError(t *testing.T) {
	r := &PolicyResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_policy")
}

// TestPolicyResource_Delete_APIErrorReadBody exercises PolicyResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPolicyResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &PolicyResource{client: newMockClientReadErrorBody(t, 501)}
	m := PolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
