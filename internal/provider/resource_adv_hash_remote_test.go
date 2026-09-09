package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestAdvHashResource_Create_Happy exercises AdvHashResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestAdvHashResource_Create_Happy(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 200, "{\"slot_id\":\"example-id\"}")}
	m := AdvHashResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAdvHashResource_Create_NilClient exercises AdvHashResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAdvHashResource_Create_NilClient(t *testing.T) {
	r := &AdvHashResource{}
	m := AdvHashResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAdvHashResource_Create_BuildError exercises AdvHashResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAdvHashResource_Create_BuildError(t *testing.T) {
	r := &AdvHashResource{client: newMalformedBaseURLClient(t)}
	m := AdvHashResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAdvHashResource_Create_SendError exercises AdvHashResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestAdvHashResource_Create_SendError(t *testing.T) {
	r := &AdvHashResource{client: newTransportErrorClient(t)}
	m := AdvHashResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAdvHashResource_Create_APIError exercises AdvHashResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAdvHashResource_Create_APIError(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AdvHashResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_adv_hash")
}

// TestAdvHashResource_Create_APIErrorReadBody exercises AdvHashResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAdvHashResource_Create_APIErrorReadBody(t *testing.T) {
	r := &AdvHashResource{client: newMockClientReadErrorBody(t, 501)}
	m := AdvHashResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestAdvHashResource_Create_InvalidJSON exercises AdvHashResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestAdvHashResource_Create_InvalidJSON(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 200, "{{")}
	m := AdvHashResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestAdvHashResource_Create_MapError exercises AdvHashResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestAdvHashResource_Create_MapError(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 200, "{\"slot_id\":12345}")}
	m := AdvHashResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestAdvHashResource_Create_MissingID exercises AdvHashResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestAdvHashResource_Create_MissingID(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 200, "{}")}
	m := AdvHashResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestAdvHashResource_Create_LocationFallback exercises AdvHashResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestAdvHashResource_Create_LocationFallback(t *testing.T) {
	r := &AdvHashResource{client: newMockClientWithLocation(t, 200, "http://example.test/folders/example-id", "{}")}
	m := AdvHashResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.SlotId.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.SlotId.ValueString(), "example-id")
	}
}

// TestAdvHashResource_Read_Happy exercises AdvHashResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestAdvHashResource_Read_Happy(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 200, "{}")}
	m := AdvHashResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestAdvHashResource_Read_NilClient exercises AdvHashResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAdvHashResource_Read_NilClient(t *testing.T) {
	r := &AdvHashResource{}
	m := AdvHashResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAdvHashResource_Read_BuildError exercises AdvHashResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAdvHashResource_Read_BuildError(t *testing.T) {
	r := &AdvHashResource{client: newMalformedBaseURLClient(t)}
	m := AdvHashResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAdvHashResource_Read_SendError exercises AdvHashResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestAdvHashResource_Read_SendError(t *testing.T) {
	r := &AdvHashResource{client: newTransportErrorClient(t)}
	m := AdvHashResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAdvHashResource_Read_NotFound exercises AdvHashResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestAdvHashResource_Read_NotFound(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 404, "")}
	m := AdvHashResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestAdvHashResource_Read_APIError exercises AdvHashResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAdvHashResource_Read_APIError(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AdvHashResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_adv_hash")
}

// TestAdvHashResource_Read_APIErrorReadBody exercises AdvHashResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAdvHashResource_Read_APIErrorReadBody(t *testing.T) {
	r := &AdvHashResource{client: newMockClientReadErrorBody(t, 501)}
	m := AdvHashResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestAdvHashResource_Read_InvalidJSON exercises AdvHashResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestAdvHashResource_Read_InvalidJSON(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 200, "{{")}
	m := AdvHashResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestAdvHashResource_Read_MapError exercises AdvHashResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestAdvHashResource_Read_MapError(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 200, "{\"slot_id\":12345}")}
	m := AdvHashResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestAdvHashResource_Update_Happy exercises AdvHashResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestAdvHashResource_Update_Happy(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 200, "{}")}
	m := AdvHashResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAdvHashResource_Update_NilClient exercises AdvHashResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAdvHashResource_Update_NilClient(t *testing.T) {
	r := &AdvHashResource{}
	m := AdvHashResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAdvHashResource_Update_BuildError exercises AdvHashResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAdvHashResource_Update_BuildError(t *testing.T) {
	r := &AdvHashResource{client: newMalformedBaseURLClient(t)}
	m := AdvHashResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAdvHashResource_Update_SendError exercises AdvHashResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestAdvHashResource_Update_SendError(t *testing.T) {
	r := &AdvHashResource{client: newTransportErrorClient(t)}
	m := AdvHashResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAdvHashResource_Update_APIError exercises AdvHashResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAdvHashResource_Update_APIError(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AdvHashResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_adv_hash")
}

// TestAdvHashResource_Update_APIErrorReadBody exercises AdvHashResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAdvHashResource_Update_APIErrorReadBody(t *testing.T) {
	r := &AdvHashResource{client: newMockClientReadErrorBody(t, 501)}
	m := AdvHashResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestAdvHashResource_Update_InvalidJSON exercises AdvHashResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestAdvHashResource_Update_InvalidJSON(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 200, "{{")}
	m := AdvHashResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestAdvHashResource_Update_MapError exercises AdvHashResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestAdvHashResource_Update_MapError(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 200, "{\"slot_id\":12345}")}
	m := AdvHashResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestAdvHashResource_Delete_Happy exercises AdvHashResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestAdvHashResource_Delete_Happy(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 204, "")}
	m := AdvHashResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAdvHashResource_Delete_NilClient exercises AdvHashResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAdvHashResource_Delete_NilClient(t *testing.T) {
	r := &AdvHashResource{}
	m := AdvHashResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAdvHashResource_Delete_BuildError exercises AdvHashResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAdvHashResource_Delete_BuildError(t *testing.T) {
	r := &AdvHashResource{client: newMalformedBaseURLClient(t)}
	m := AdvHashResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAdvHashResource_Delete_SendError exercises AdvHashResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestAdvHashResource_Delete_SendError(t *testing.T) {
	r := &AdvHashResource{client: newTransportErrorClient(t)}
	m := AdvHashResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAdvHashResource_Delete_NotFoundSuccess exercises AdvHashResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestAdvHashResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 404, "")}
	m := AdvHashResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAdvHashResource_Delete_APIError exercises AdvHashResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAdvHashResource_Delete_APIError(t *testing.T) {
	r := &AdvHashResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AdvHashResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_adv_hash")
}

// TestAdvHashResource_Delete_APIErrorReadBody exercises AdvHashResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAdvHashResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &AdvHashResource{client: newMockClientReadErrorBody(t, 501)}
	m := AdvHashResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
