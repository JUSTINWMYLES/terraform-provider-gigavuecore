package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestPcapProfileResource_Create_Happy exercises PcapProfileResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestPcapProfileResource_Create_Happy(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientStatus(t, 201, "{\"id\":\"example-id\"}")}
	m := PcapProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPcapProfileResource_Create_NilClient exercises PcapProfileResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPcapProfileResource_Create_NilClient(t *testing.T) {
	r := &PcapProfileResource{}
	m := PcapProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPcapProfileResource_Create_BuildError exercises PcapProfileResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPcapProfileResource_Create_BuildError(t *testing.T) {
	r := &PcapProfileResource{client: newMalformedBaseURLClient(t)}
	m := PcapProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPcapProfileResource_Create_SendError exercises PcapProfileResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestPcapProfileResource_Create_SendError(t *testing.T) {
	r := &PcapProfileResource{client: newTransportErrorClient(t)}
	m := PcapProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPcapProfileResource_Create_APIError exercises PcapProfileResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPcapProfileResource_Create_APIError(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PcapProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_pcap_profile")
}

// TestPcapProfileResource_Create_APIErrorReadBody exercises PcapProfileResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPcapProfileResource_Create_APIErrorReadBody(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := PcapProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPcapProfileResource_Create_InvalidJSON exercises PcapProfileResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPcapProfileResource_Create_InvalidJSON(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientStatus(t, 201, "{{")}
	m := PcapProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPcapProfileResource_Create_MapError exercises PcapProfileResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPcapProfileResource_Create_MapError(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientStatus(t, 201, "{\"id\":12345}")}
	m := PcapProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPcapProfileResource_Create_MissingID exercises PcapProfileResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestPcapProfileResource_Create_MissingID(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientStatus(t, 201, "{}")}
	m := PcapProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestPcapProfileResource_Create_LocationFallback exercises PcapProfileResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestPcapProfileResource_Create_LocationFallback(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := PcapProfileResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Id.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Id.ValueString(), "example-id")
	}
}

// TestPcapProfileResource_Read_Happy exercises PcapProfileResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestPcapProfileResource_Read_Happy(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientStatus(t, 200, "{}")}
	m := PcapProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPcapProfileResource_Read_NilClient exercises PcapProfileResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPcapProfileResource_Read_NilClient(t *testing.T) {
	r := &PcapProfileResource{}
	m := PcapProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPcapProfileResource_Read_BuildError exercises PcapProfileResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPcapProfileResource_Read_BuildError(t *testing.T) {
	r := &PcapProfileResource{client: newMalformedBaseURLClient(t)}
	m := PcapProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPcapProfileResource_Read_SendError exercises PcapProfileResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestPcapProfileResource_Read_SendError(t *testing.T) {
	r := &PcapProfileResource{client: newTransportErrorClient(t)}
	m := PcapProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPcapProfileResource_Read_NotFound exercises PcapProfileResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestPcapProfileResource_Read_NotFound(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientStatus(t, 404, "")}
	m := PcapProfileResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPcapProfileResource_Read_APIError exercises PcapProfileResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPcapProfileResource_Read_APIError(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PcapProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_pcap_profile")
}

// TestPcapProfileResource_Read_APIErrorReadBody exercises PcapProfileResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPcapProfileResource_Read_APIErrorReadBody(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := PcapProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPcapProfileResource_Read_InvalidJSON exercises PcapProfileResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPcapProfileResource_Read_InvalidJSON(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientStatus(t, 200, "{{")}
	m := PcapProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPcapProfileResource_Read_MapError exercises PcapProfileResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPcapProfileResource_Read_MapError(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientStatus(t, 200, "{\"id\":12345}")}
	m := PcapProfileResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPcapProfileResource_Delete_Happy exercises PcapProfileResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestPcapProfileResource_Delete_Happy(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientStatus(t, 204, "")}
	m := PcapProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPcapProfileResource_Delete_NilClient exercises PcapProfileResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPcapProfileResource_Delete_NilClient(t *testing.T) {
	r := &PcapProfileResource{}
	m := PcapProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPcapProfileResource_Delete_BuildError exercises PcapProfileResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPcapProfileResource_Delete_BuildError(t *testing.T) {
	r := &PcapProfileResource{client: newMalformedBaseURLClient(t)}
	m := PcapProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPcapProfileResource_Delete_SendError exercises PcapProfileResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestPcapProfileResource_Delete_SendError(t *testing.T) {
	r := &PcapProfileResource{client: newTransportErrorClient(t)}
	m := PcapProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPcapProfileResource_Delete_NotFoundSuccess exercises PcapProfileResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestPcapProfileResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientStatus(t, 404, "")}
	m := PcapProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPcapProfileResource_Delete_APIError exercises PcapProfileResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPcapProfileResource_Delete_APIError(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PcapProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_pcap_profile")
}

// TestPcapProfileResource_Delete_APIErrorReadBody exercises PcapProfileResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPcapProfileResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &PcapProfileResource{client: newMockClientReadErrorBody(t, 501)}
	m := PcapProfileResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
