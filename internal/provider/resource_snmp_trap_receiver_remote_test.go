package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSnmpTrapReceiverResource_Create_Happy exercises SnmpTrapReceiverResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestSnmpTrapReceiverResource_Create_Happy(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 200, "{\"alias\":\"example-id\"}")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSnmpTrapReceiverResource_Create_NilClient exercises SnmpTrapReceiverResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSnmpTrapReceiverResource_Create_NilClient(t *testing.T) {
	r := &SnmpTrapReceiverResource{}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSnmpTrapReceiverResource_Create_BuildError exercises SnmpTrapReceiverResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSnmpTrapReceiverResource_Create_BuildError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMalformedBaseURLClient(t)}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSnmpTrapReceiverResource_Create_SendError exercises SnmpTrapReceiverResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestSnmpTrapReceiverResource_Create_SendError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newTransportErrorClient(t)}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSnmpTrapReceiverResource_Create_APIError exercises SnmpTrapReceiverResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSnmpTrapReceiverResource_Create_APIError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_snmp_trap_receiver")
}

// TestSnmpTrapReceiverResource_Create_APIErrorReadBody exercises SnmpTrapReceiverResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSnmpTrapReceiverResource_Create_APIErrorReadBody(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientReadErrorBody(t, 501)}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSnmpTrapReceiverResource_Create_InvalidJSON exercises SnmpTrapReceiverResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSnmpTrapReceiverResource_Create_InvalidJSON(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 200, "{{")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSnmpTrapReceiverResource_Create_MapError exercises SnmpTrapReceiverResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSnmpTrapReceiverResource_Create_MapError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSnmpTrapReceiverResource_Create_MissingID exercises SnmpTrapReceiverResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestSnmpTrapReceiverResource_Create_MissingID(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 200, "{}")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestSnmpTrapReceiverResource_Create_LocationFallback exercises SnmpTrapReceiverResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestSnmpTrapReceiverResource_Create_LocationFallback(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientWithLocation(t, 200, "http://example.test/folders/example-id", "{}")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestSnmpTrapReceiverResource_Read_Happy exercises SnmpTrapReceiverResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestSnmpTrapReceiverResource_Read_Happy(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 200, "{}")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSnmpTrapReceiverResource_Read_NilClient exercises SnmpTrapReceiverResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSnmpTrapReceiverResource_Read_NilClient(t *testing.T) {
	r := &SnmpTrapReceiverResource{}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSnmpTrapReceiverResource_Read_BuildError exercises SnmpTrapReceiverResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSnmpTrapReceiverResource_Read_BuildError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMalformedBaseURLClient(t)}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSnmpTrapReceiverResource_Read_SendError exercises SnmpTrapReceiverResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestSnmpTrapReceiverResource_Read_SendError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newTransportErrorClient(t)}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSnmpTrapReceiverResource_Read_NotFound exercises SnmpTrapReceiverResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestSnmpTrapReceiverResource_Read_NotFound(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 404, "")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestSnmpTrapReceiverResource_Read_APIError exercises SnmpTrapReceiverResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSnmpTrapReceiverResource_Read_APIError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_snmp_trap_receiver")
}

// TestSnmpTrapReceiverResource_Read_APIErrorReadBody exercises SnmpTrapReceiverResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSnmpTrapReceiverResource_Read_APIErrorReadBody(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientReadErrorBody(t, 501)}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSnmpTrapReceiverResource_Read_InvalidJSON exercises SnmpTrapReceiverResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSnmpTrapReceiverResource_Read_InvalidJSON(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 200, "{{")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSnmpTrapReceiverResource_Read_MapError exercises SnmpTrapReceiverResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSnmpTrapReceiverResource_Read_MapError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSnmpTrapReceiverResource_Update_Happy exercises SnmpTrapReceiverResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestSnmpTrapReceiverResource_Update_Happy(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 200, "{}")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSnmpTrapReceiverResource_Update_NilClient exercises SnmpTrapReceiverResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSnmpTrapReceiverResource_Update_NilClient(t *testing.T) {
	r := &SnmpTrapReceiverResource{}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSnmpTrapReceiverResource_Update_BuildError exercises SnmpTrapReceiverResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSnmpTrapReceiverResource_Update_BuildError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMalformedBaseURLClient(t)}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSnmpTrapReceiverResource_Update_SendError exercises SnmpTrapReceiverResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestSnmpTrapReceiverResource_Update_SendError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newTransportErrorClient(t)}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSnmpTrapReceiverResource_Update_APIError exercises SnmpTrapReceiverResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSnmpTrapReceiverResource_Update_APIError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_snmp_trap_receiver")
}

// TestSnmpTrapReceiverResource_Update_APIErrorReadBody exercises SnmpTrapReceiverResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSnmpTrapReceiverResource_Update_APIErrorReadBody(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientReadErrorBody(t, 501)}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestSnmpTrapReceiverResource_Update_InvalidJSON exercises SnmpTrapReceiverResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestSnmpTrapReceiverResource_Update_InvalidJSON(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 200, "{{")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestSnmpTrapReceiverResource_Update_MapError exercises SnmpTrapReceiverResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestSnmpTrapReceiverResource_Update_MapError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestSnmpTrapReceiverResource_Delete_Happy exercises SnmpTrapReceiverResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestSnmpTrapReceiverResource_Delete_Happy(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 200, "")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSnmpTrapReceiverResource_Delete_NilClient exercises SnmpTrapReceiverResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSnmpTrapReceiverResource_Delete_NilClient(t *testing.T) {
	r := &SnmpTrapReceiverResource{}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSnmpTrapReceiverResource_Delete_BuildError exercises SnmpTrapReceiverResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSnmpTrapReceiverResource_Delete_BuildError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMalformedBaseURLClient(t)}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSnmpTrapReceiverResource_Delete_SendError exercises SnmpTrapReceiverResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestSnmpTrapReceiverResource_Delete_SendError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newTransportErrorClient(t)}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSnmpTrapReceiverResource_Delete_NotFoundSuccess exercises SnmpTrapReceiverResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestSnmpTrapReceiverResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 404, "")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSnmpTrapReceiverResource_Delete_APIError exercises SnmpTrapReceiverResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSnmpTrapReceiverResource_Delete_APIError(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_snmp_trap_receiver")
}

// TestSnmpTrapReceiverResource_Delete_APIErrorReadBody exercises SnmpTrapReceiverResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSnmpTrapReceiverResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &SnmpTrapReceiverResource{client: newMockClientReadErrorBody(t, 501)}
	m := SnmpTrapReceiverResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
