package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestInterfaceResource_Create_Happy exercises InterfaceResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestInterfaceResource_Create_Happy(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := InterfaceResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInterfaceResource_Create_NilClient exercises InterfaceResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInterfaceResource_Create_NilClient(t *testing.T) {
	r := &InterfaceResource{}
	m := InterfaceResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInterfaceResource_Create_BuildError exercises InterfaceResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInterfaceResource_Create_BuildError(t *testing.T) {
	r := &InterfaceResource{client: newMalformedBaseURLClient(t)}
	m := InterfaceResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInterfaceResource_Create_SendError exercises InterfaceResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestInterfaceResource_Create_SendError(t *testing.T) {
	r := &InterfaceResource{client: newTransportErrorClient(t)}
	m := InterfaceResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInterfaceResource_Create_APIError exercises InterfaceResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInterfaceResource_Create_APIError(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InterfaceResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_interface")
}

// TestInterfaceResource_Create_APIErrorReadBody exercises InterfaceResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInterfaceResource_Create_APIErrorReadBody(t *testing.T) {
	r := &InterfaceResource{client: newMockClientReadErrorBody(t, 501)}
	m := InterfaceResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestInterfaceResource_Create_InvalidJSON exercises InterfaceResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestInterfaceResource_Create_InvalidJSON(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 201, "{{")}
	m := InterfaceResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestInterfaceResource_Create_MapError exercises InterfaceResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestInterfaceResource_Create_MapError(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := InterfaceResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestInterfaceResource_Create_MissingID exercises InterfaceResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestInterfaceResource_Create_MissingID(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 201, "{}")}
	m := InterfaceResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestInterfaceResource_Create_LocationFallback exercises InterfaceResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestInterfaceResource_Create_LocationFallback(t *testing.T) {
	r := &InterfaceResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := InterfaceResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestInterfaceResource_Read_Happy exercises InterfaceResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestInterfaceResource_Read_Happy(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 200, "{}")}
	m := InterfaceResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestInterfaceResource_Read_NilClient exercises InterfaceResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInterfaceResource_Read_NilClient(t *testing.T) {
	r := &InterfaceResource{}
	m := InterfaceResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInterfaceResource_Read_BuildError exercises InterfaceResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInterfaceResource_Read_BuildError(t *testing.T) {
	r := &InterfaceResource{client: newMalformedBaseURLClient(t)}
	m := InterfaceResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInterfaceResource_Read_SendError exercises InterfaceResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestInterfaceResource_Read_SendError(t *testing.T) {
	r := &InterfaceResource{client: newTransportErrorClient(t)}
	m := InterfaceResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInterfaceResource_Read_NotFound exercises InterfaceResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestInterfaceResource_Read_NotFound(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 404, "")}
	m := InterfaceResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestInterfaceResource_Read_APIError exercises InterfaceResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInterfaceResource_Read_APIError(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InterfaceResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_interface")
}

// TestInterfaceResource_Read_APIErrorReadBody exercises InterfaceResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInterfaceResource_Read_APIErrorReadBody(t *testing.T) {
	r := &InterfaceResource{client: newMockClientReadErrorBody(t, 501)}
	m := InterfaceResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestInterfaceResource_Read_InvalidJSON exercises InterfaceResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestInterfaceResource_Read_InvalidJSON(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 200, "{{")}
	m := InterfaceResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestInterfaceResource_Read_MapError exercises InterfaceResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestInterfaceResource_Read_MapError(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := InterfaceResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestInterfaceResource_Update_Happy exercises InterfaceResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestInterfaceResource_Update_Happy(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 200, "{}")}
	m := InterfaceResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInterfaceResource_Update_NilClient exercises InterfaceResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInterfaceResource_Update_NilClient(t *testing.T) {
	r := &InterfaceResource{}
	m := InterfaceResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInterfaceResource_Update_BuildError exercises InterfaceResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInterfaceResource_Update_BuildError(t *testing.T) {
	r := &InterfaceResource{client: newMalformedBaseURLClient(t)}
	m := InterfaceResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInterfaceResource_Update_SendError exercises InterfaceResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestInterfaceResource_Update_SendError(t *testing.T) {
	r := &InterfaceResource{client: newTransportErrorClient(t)}
	m := InterfaceResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInterfaceResource_Update_APIError exercises InterfaceResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInterfaceResource_Update_APIError(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InterfaceResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_interface")
}

// TestInterfaceResource_Update_APIErrorReadBody exercises InterfaceResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInterfaceResource_Update_APIErrorReadBody(t *testing.T) {
	r := &InterfaceResource{client: newMockClientReadErrorBody(t, 501)}
	m := InterfaceResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestInterfaceResource_Update_InvalidJSON exercises InterfaceResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestInterfaceResource_Update_InvalidJSON(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 200, "{{")}
	m := InterfaceResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestInterfaceResource_Update_MapError exercises InterfaceResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestInterfaceResource_Update_MapError(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := InterfaceResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestInterfaceResource_Delete_Happy exercises InterfaceResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestInterfaceResource_Delete_Happy(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 204, "")}
	m := InterfaceResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInterfaceResource_Delete_NilClient exercises InterfaceResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInterfaceResource_Delete_NilClient(t *testing.T) {
	r := &InterfaceResource{}
	m := InterfaceResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInterfaceResource_Delete_BuildError exercises InterfaceResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInterfaceResource_Delete_BuildError(t *testing.T) {
	r := &InterfaceResource{client: newMalformedBaseURLClient(t)}
	m := InterfaceResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInterfaceResource_Delete_SendError exercises InterfaceResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestInterfaceResource_Delete_SendError(t *testing.T) {
	r := &InterfaceResource{client: newTransportErrorClient(t)}
	m := InterfaceResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInterfaceResource_Delete_NotFoundSuccess exercises InterfaceResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestInterfaceResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 404, "")}
	m := InterfaceResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInterfaceResource_Delete_APIError exercises InterfaceResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInterfaceResource_Delete_APIError(t *testing.T) {
	r := &InterfaceResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InterfaceResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_interface")
}

// TestInterfaceResource_Delete_APIErrorReadBody exercises InterfaceResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInterfaceResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &InterfaceResource{client: newMockClientReadErrorBody(t, 501)}
	m := InterfaceResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
