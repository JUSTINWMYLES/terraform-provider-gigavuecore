package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestElbResource_Create_Happy exercises ElbResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestElbResource_Create_Happy(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := ElbResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestElbResource_Create_NilClient exercises ElbResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestElbResource_Create_NilClient(t *testing.T) {
	r := &ElbResource{}
	m := ElbResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestElbResource_Create_BuildError exercises ElbResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestElbResource_Create_BuildError(t *testing.T) {
	r := &ElbResource{client: newMalformedBaseURLClient(t)}
	m := ElbResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestElbResource_Create_SendError exercises ElbResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestElbResource_Create_SendError(t *testing.T) {
	r := &ElbResource{client: newTransportErrorClient(t)}
	m := ElbResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestElbResource_Create_APIError exercises ElbResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestElbResource_Create_APIError(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ElbResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_elb")
}

// TestElbResource_Create_APIErrorReadBody exercises ElbResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestElbResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ElbResource{client: newMockClientReadErrorBody(t, 501)}
	m := ElbResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestElbResource_Create_InvalidJSON exercises ElbResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestElbResource_Create_InvalidJSON(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 201, "{{")}
	m := ElbResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestElbResource_Create_MapError exercises ElbResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestElbResource_Create_MapError(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := ElbResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestElbResource_Create_MissingID exercises ElbResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestElbResource_Create_MissingID(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 201, "{}")}
	m := ElbResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestElbResource_Create_LocationFallback exercises ElbResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestElbResource_Create_LocationFallback(t *testing.T) {
	r := &ElbResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ElbResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestElbResource_Read_Happy exercises ElbResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestElbResource_Read_Happy(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 200, "{}")}
	m := ElbResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestElbResource_Read_NilClient exercises ElbResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestElbResource_Read_NilClient(t *testing.T) {
	r := &ElbResource{}
	m := ElbResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestElbResource_Read_BuildError exercises ElbResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestElbResource_Read_BuildError(t *testing.T) {
	r := &ElbResource{client: newMalformedBaseURLClient(t)}
	m := ElbResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestElbResource_Read_SendError exercises ElbResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestElbResource_Read_SendError(t *testing.T) {
	r := &ElbResource{client: newTransportErrorClient(t)}
	m := ElbResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestElbResource_Read_NotFound exercises ElbResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestElbResource_Read_NotFound(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 404, "")}
	m := ElbResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestElbResource_Read_APIError exercises ElbResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestElbResource_Read_APIError(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ElbResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_elb")
}

// TestElbResource_Read_APIErrorReadBody exercises ElbResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestElbResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ElbResource{client: newMockClientReadErrorBody(t, 501)}
	m := ElbResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestElbResource_Read_InvalidJSON exercises ElbResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestElbResource_Read_InvalidJSON(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 200, "{{")}
	m := ElbResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestElbResource_Read_MapError exercises ElbResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestElbResource_Read_MapError(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ElbResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestElbResource_Update_Happy exercises ElbResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestElbResource_Update_Happy(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 200, "{}")}
	m := ElbResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestElbResource_Update_NilClient exercises ElbResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestElbResource_Update_NilClient(t *testing.T) {
	r := &ElbResource{}
	m := ElbResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestElbResource_Update_BuildError exercises ElbResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestElbResource_Update_BuildError(t *testing.T) {
	r := &ElbResource{client: newMalformedBaseURLClient(t)}
	m := ElbResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestElbResource_Update_SendError exercises ElbResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestElbResource_Update_SendError(t *testing.T) {
	r := &ElbResource{client: newTransportErrorClient(t)}
	m := ElbResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestElbResource_Update_APIError exercises ElbResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestElbResource_Update_APIError(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ElbResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_elb")
}

// TestElbResource_Update_APIErrorReadBody exercises ElbResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestElbResource_Update_APIErrorReadBody(t *testing.T) {
	r := &ElbResource{client: newMockClientReadErrorBody(t, 501)}
	m := ElbResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestElbResource_Update_InvalidJSON exercises ElbResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestElbResource_Update_InvalidJSON(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 200, "{{")}
	m := ElbResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestElbResource_Update_MapError exercises ElbResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestElbResource_Update_MapError(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := ElbResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestElbResource_Delete_Happy exercises ElbResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestElbResource_Delete_Happy(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 204, "")}
	m := ElbResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestElbResource_Delete_NilClient exercises ElbResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestElbResource_Delete_NilClient(t *testing.T) {
	r := &ElbResource{}
	m := ElbResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestElbResource_Delete_BuildError exercises ElbResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestElbResource_Delete_BuildError(t *testing.T) {
	r := &ElbResource{client: newMalformedBaseURLClient(t)}
	m := ElbResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestElbResource_Delete_SendError exercises ElbResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestElbResource_Delete_SendError(t *testing.T) {
	r := &ElbResource{client: newTransportErrorClient(t)}
	m := ElbResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestElbResource_Delete_NotFoundSuccess exercises ElbResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestElbResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 404, "")}
	m := ElbResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestElbResource_Delete_APIError exercises ElbResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestElbResource_Delete_APIError(t *testing.T) {
	r := &ElbResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ElbResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_elb")
}

// TestElbResource_Delete_APIErrorReadBody exercises ElbResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestElbResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ElbResource{client: newMockClientReadErrorBody(t, 501)}
	m := ElbResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
