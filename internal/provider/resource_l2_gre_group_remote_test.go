package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestL2GreGroupResource_Create_Happy exercises L2GreGroupResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestL2GreGroupResource_Create_Happy(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := L2GreGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestL2GreGroupResource_Create_NilClient exercises L2GreGroupResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestL2GreGroupResource_Create_NilClient(t *testing.T) {
	r := &L2GreGroupResource{}
	m := L2GreGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestL2GreGroupResource_Create_BuildError exercises L2GreGroupResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestL2GreGroupResource_Create_BuildError(t *testing.T) {
	r := &L2GreGroupResource{client: newMalformedBaseURLClient(t)}
	m := L2GreGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestL2GreGroupResource_Create_SendError exercises L2GreGroupResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestL2GreGroupResource_Create_SendError(t *testing.T) {
	r := &L2GreGroupResource{client: newTransportErrorClient(t)}
	m := L2GreGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestL2GreGroupResource_Create_APIError exercises L2GreGroupResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestL2GreGroupResource_Create_APIError(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := L2GreGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_l2_gre_group")
}

// TestL2GreGroupResource_Create_APIErrorReadBody exercises L2GreGroupResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestL2GreGroupResource_Create_APIErrorReadBody(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := L2GreGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestL2GreGroupResource_Create_InvalidJSON exercises L2GreGroupResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestL2GreGroupResource_Create_InvalidJSON(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 201, "{{")}
	m := L2GreGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestL2GreGroupResource_Create_MapError exercises L2GreGroupResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestL2GreGroupResource_Create_MapError(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := L2GreGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestL2GreGroupResource_Create_MissingID exercises L2GreGroupResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestL2GreGroupResource_Create_MissingID(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 201, "{}")}
	m := L2GreGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestL2GreGroupResource_Create_LocationFallback exercises L2GreGroupResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestL2GreGroupResource_Create_LocationFallback(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := L2GreGroupResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestL2GreGroupResource_Read_Happy exercises L2GreGroupResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestL2GreGroupResource_Read_Happy(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := L2GreGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestL2GreGroupResource_Read_NilClient exercises L2GreGroupResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestL2GreGroupResource_Read_NilClient(t *testing.T) {
	r := &L2GreGroupResource{}
	m := L2GreGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestL2GreGroupResource_Read_BuildError exercises L2GreGroupResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestL2GreGroupResource_Read_BuildError(t *testing.T) {
	r := &L2GreGroupResource{client: newMalformedBaseURLClient(t)}
	m := L2GreGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestL2GreGroupResource_Read_SendError exercises L2GreGroupResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestL2GreGroupResource_Read_SendError(t *testing.T) {
	r := &L2GreGroupResource{client: newTransportErrorClient(t)}
	m := L2GreGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestL2GreGroupResource_Read_NotFound exercises L2GreGroupResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestL2GreGroupResource_Read_NotFound(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 404, "")}
	m := L2GreGroupResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestL2GreGroupResource_Read_APIError exercises L2GreGroupResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestL2GreGroupResource_Read_APIError(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := L2GreGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_l2_gre_group")
}

// TestL2GreGroupResource_Read_APIErrorReadBody exercises L2GreGroupResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestL2GreGroupResource_Read_APIErrorReadBody(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := L2GreGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestL2GreGroupResource_Read_InvalidJSON exercises L2GreGroupResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestL2GreGroupResource_Read_InvalidJSON(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := L2GreGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestL2GreGroupResource_Read_MapError exercises L2GreGroupResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestL2GreGroupResource_Read_MapError(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := L2GreGroupResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestL2GreGroupResource_Update_Happy exercises L2GreGroupResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestL2GreGroupResource_Update_Happy(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 200, "{}")}
	m := L2GreGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestL2GreGroupResource_Update_NilClient exercises L2GreGroupResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestL2GreGroupResource_Update_NilClient(t *testing.T) {
	r := &L2GreGroupResource{}
	m := L2GreGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestL2GreGroupResource_Update_BuildError exercises L2GreGroupResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestL2GreGroupResource_Update_BuildError(t *testing.T) {
	r := &L2GreGroupResource{client: newMalformedBaseURLClient(t)}
	m := L2GreGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestL2GreGroupResource_Update_SendError exercises L2GreGroupResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestL2GreGroupResource_Update_SendError(t *testing.T) {
	r := &L2GreGroupResource{client: newTransportErrorClient(t)}
	m := L2GreGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestL2GreGroupResource_Update_APIError exercises L2GreGroupResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestL2GreGroupResource_Update_APIError(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := L2GreGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_l2_gre_group")
}

// TestL2GreGroupResource_Update_APIErrorReadBody exercises L2GreGroupResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestL2GreGroupResource_Update_APIErrorReadBody(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := L2GreGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestL2GreGroupResource_Update_InvalidJSON exercises L2GreGroupResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestL2GreGroupResource_Update_InvalidJSON(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 200, "{{")}
	m := L2GreGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestL2GreGroupResource_Update_MapError exercises L2GreGroupResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestL2GreGroupResource_Update_MapError(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := L2GreGroupResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestL2GreGroupResource_Delete_Happy exercises L2GreGroupResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestL2GreGroupResource_Delete_Happy(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 204, "")}
	m := L2GreGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestL2GreGroupResource_Delete_NilClient exercises L2GreGroupResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestL2GreGroupResource_Delete_NilClient(t *testing.T) {
	r := &L2GreGroupResource{}
	m := L2GreGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestL2GreGroupResource_Delete_BuildError exercises L2GreGroupResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestL2GreGroupResource_Delete_BuildError(t *testing.T) {
	r := &L2GreGroupResource{client: newMalformedBaseURLClient(t)}
	m := L2GreGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestL2GreGroupResource_Delete_SendError exercises L2GreGroupResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestL2GreGroupResource_Delete_SendError(t *testing.T) {
	r := &L2GreGroupResource{client: newTransportErrorClient(t)}
	m := L2GreGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestL2GreGroupResource_Delete_NotFoundSuccess exercises L2GreGroupResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestL2GreGroupResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 404, "")}
	m := L2GreGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestL2GreGroupResource_Delete_APIError exercises L2GreGroupResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestL2GreGroupResource_Delete_APIError(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := L2GreGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_l2_gre_group")
}

// TestL2GreGroupResource_Delete_APIErrorReadBody exercises L2GreGroupResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestL2GreGroupResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &L2GreGroupResource{client: newMockClientReadErrorBody(t, 501)}
	m := L2GreGroupResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
