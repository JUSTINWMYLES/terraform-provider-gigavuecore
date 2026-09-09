package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestPtpConfigResource_Create_Happy exercises PtpConfigResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestPtpConfigResource_Create_Happy(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 201, "{\"alias\":\"example-id\"}")}
	m := PtpConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPtpConfigResource_Create_NilClient exercises PtpConfigResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPtpConfigResource_Create_NilClient(t *testing.T) {
	r := &PtpConfigResource{}
	m := PtpConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPtpConfigResource_Create_BuildError exercises PtpConfigResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPtpConfigResource_Create_BuildError(t *testing.T) {
	r := &PtpConfigResource{client: newMalformedBaseURLClient(t)}
	m := PtpConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPtpConfigResource_Create_SendError exercises PtpConfigResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestPtpConfigResource_Create_SendError(t *testing.T) {
	r := &PtpConfigResource{client: newTransportErrorClient(t)}
	m := PtpConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPtpConfigResource_Create_APIError exercises PtpConfigResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPtpConfigResource_Create_APIError(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PtpConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_ptp_config")
}

// TestPtpConfigResource_Create_APIErrorReadBody exercises PtpConfigResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPtpConfigResource_Create_APIErrorReadBody(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientReadErrorBody(t, 501)}
	m := PtpConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPtpConfigResource_Create_InvalidJSON exercises PtpConfigResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPtpConfigResource_Create_InvalidJSON(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 201, "{{")}
	m := PtpConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPtpConfigResource_Create_MapError exercises PtpConfigResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPtpConfigResource_Create_MapError(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 201, "{\"alias\":12345}")}
	m := PtpConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPtpConfigResource_Create_MissingID exercises PtpConfigResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestPtpConfigResource_Create_MissingID(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 201, "{}")}
	m := PtpConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestPtpConfigResource_Create_LocationFallback exercises PtpConfigResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestPtpConfigResource_Create_LocationFallback(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := PtpConfigResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Alias.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Alias.ValueString(), "example-id")
	}
}

// TestPtpConfigResource_Read_Happy exercises PtpConfigResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestPtpConfigResource_Read_Happy(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 200, "{}")}
	m := PtpConfigResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPtpConfigResource_Read_NilClient exercises PtpConfigResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPtpConfigResource_Read_NilClient(t *testing.T) {
	r := &PtpConfigResource{}
	m := PtpConfigResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPtpConfigResource_Read_BuildError exercises PtpConfigResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPtpConfigResource_Read_BuildError(t *testing.T) {
	r := &PtpConfigResource{client: newMalformedBaseURLClient(t)}
	m := PtpConfigResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPtpConfigResource_Read_SendError exercises PtpConfigResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestPtpConfigResource_Read_SendError(t *testing.T) {
	r := &PtpConfigResource{client: newTransportErrorClient(t)}
	m := PtpConfigResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPtpConfigResource_Read_NotFound exercises PtpConfigResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestPtpConfigResource_Read_NotFound(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 404, "")}
	m := PtpConfigResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPtpConfigResource_Read_APIError exercises PtpConfigResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPtpConfigResource_Read_APIError(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PtpConfigResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_ptp_config")
}

// TestPtpConfigResource_Read_APIErrorReadBody exercises PtpConfigResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPtpConfigResource_Read_APIErrorReadBody(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientReadErrorBody(t, 501)}
	m := PtpConfigResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPtpConfigResource_Read_InvalidJSON exercises PtpConfigResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPtpConfigResource_Read_InvalidJSON(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 200, "{{")}
	m := PtpConfigResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPtpConfigResource_Read_MapError exercises PtpConfigResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPtpConfigResource_Read_MapError(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := PtpConfigResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPtpConfigResource_Update_Happy exercises PtpConfigResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestPtpConfigResource_Update_Happy(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 200, "{}")}
	m := PtpConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPtpConfigResource_Update_NilClient exercises PtpConfigResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPtpConfigResource_Update_NilClient(t *testing.T) {
	r := &PtpConfigResource{}
	m := PtpConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPtpConfigResource_Update_BuildError exercises PtpConfigResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPtpConfigResource_Update_BuildError(t *testing.T) {
	r := &PtpConfigResource{client: newMalformedBaseURLClient(t)}
	m := PtpConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPtpConfigResource_Update_SendError exercises PtpConfigResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestPtpConfigResource_Update_SendError(t *testing.T) {
	r := &PtpConfigResource{client: newTransportErrorClient(t)}
	m := PtpConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPtpConfigResource_Update_APIError exercises PtpConfigResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPtpConfigResource_Update_APIError(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PtpConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_ptp_config")
}

// TestPtpConfigResource_Update_APIErrorReadBody exercises PtpConfigResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPtpConfigResource_Update_APIErrorReadBody(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientReadErrorBody(t, 501)}
	m := PtpConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPtpConfigResource_Update_InvalidJSON exercises PtpConfigResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPtpConfigResource_Update_InvalidJSON(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 200, "{{")}
	m := PtpConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPtpConfigResource_Update_MapError exercises PtpConfigResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPtpConfigResource_Update_MapError(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 200, "{\"alias\":12345}")}
	m := PtpConfigResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPtpConfigResource_Delete_Happy exercises PtpConfigResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestPtpConfigResource_Delete_Happy(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 204, "")}
	m := PtpConfigResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPtpConfigResource_Delete_NilClient exercises PtpConfigResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPtpConfigResource_Delete_NilClient(t *testing.T) {
	r := &PtpConfigResource{}
	m := PtpConfigResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPtpConfigResource_Delete_BuildError exercises PtpConfigResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPtpConfigResource_Delete_BuildError(t *testing.T) {
	r := &PtpConfigResource{client: newMalformedBaseURLClient(t)}
	m := PtpConfigResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPtpConfigResource_Delete_SendError exercises PtpConfigResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestPtpConfigResource_Delete_SendError(t *testing.T) {
	r := &PtpConfigResource{client: newTransportErrorClient(t)}
	m := PtpConfigResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPtpConfigResource_Delete_NotFoundSuccess exercises PtpConfigResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestPtpConfigResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 404, "")}
	m := PtpConfigResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPtpConfigResource_Delete_APIError exercises PtpConfigResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPtpConfigResource_Delete_APIError(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PtpConfigResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_ptp_config")
}

// TestPtpConfigResource_Delete_APIErrorReadBody exercises PtpConfigResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPtpConfigResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &PtpConfigResource{client: newMockClientReadErrorBody(t, 501)}
	m := PtpConfigResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
