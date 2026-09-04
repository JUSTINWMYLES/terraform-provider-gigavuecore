package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestFmTemplateResource_Create_Happy exercises FmTemplateResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestFmTemplateResource_Create_Happy(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 200, "{\"template_name\":\"example-id\"}")}
	m := FmTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmTemplateResource_Create_NilClient exercises FmTemplateResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFmTemplateResource_Create_NilClient(t *testing.T) {
	r := &FmTemplateResource{}
	m := FmTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFmTemplateResource_Create_BuildError exercises FmTemplateResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFmTemplateResource_Create_BuildError(t *testing.T) {
	r := &FmTemplateResource{client: newMalformedBaseURLClient(t)}
	m := FmTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFmTemplateResource_Create_SendError exercises FmTemplateResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestFmTemplateResource_Create_SendError(t *testing.T) {
	r := &FmTemplateResource{client: newTransportErrorClient(t)}
	m := FmTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFmTemplateResource_Create_APIError exercises FmTemplateResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFmTemplateResource_Create_APIError(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FmTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_fm_template")
}

// TestFmTemplateResource_Create_APIErrorReadBody exercises FmTemplateResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFmTemplateResource_Create_APIErrorReadBody(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientReadErrorBody(t, 501)}
	m := FmTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFmTemplateResource_Create_InvalidJSON exercises FmTemplateResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFmTemplateResource_Create_InvalidJSON(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 200, "{{")}
	m := FmTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFmTemplateResource_Create_MapError exercises FmTemplateResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFmTemplateResource_Create_MapError(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 200, "{\"template_name\":12345}")}
	m := FmTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFmTemplateResource_Create_MissingID exercises FmTemplateResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestFmTemplateResource_Create_MissingID(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 200, "{}")}
	m := FmTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestFmTemplateResource_Create_LocationFallback exercises FmTemplateResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestFmTemplateResource_Create_LocationFallback(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientWithLocation(t, 200, "http://example.test/folders/example-id", "{}")}
	m := FmTemplateResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.TemplateName.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.TemplateName.ValueString(), "example-id")
	}
}

// TestFmTemplateResource_Read_Happy exercises FmTemplateResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestFmTemplateResource_Read_Happy(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 200, "{}")}
	m := FmTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmTemplateResource_Read_NilClient exercises FmTemplateResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFmTemplateResource_Read_NilClient(t *testing.T) {
	r := &FmTemplateResource{}
	m := FmTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFmTemplateResource_Read_BuildError exercises FmTemplateResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFmTemplateResource_Read_BuildError(t *testing.T) {
	r := &FmTemplateResource{client: newMalformedBaseURLClient(t)}
	m := FmTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFmTemplateResource_Read_SendError exercises FmTemplateResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestFmTemplateResource_Read_SendError(t *testing.T) {
	r := &FmTemplateResource{client: newTransportErrorClient(t)}
	m := FmTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFmTemplateResource_Read_NotFound exercises FmTemplateResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestFmTemplateResource_Read_NotFound(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 404, "")}
	m := FmTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmTemplateResource_Read_APIError exercises FmTemplateResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFmTemplateResource_Read_APIError(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FmTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_fm_template")
}

// TestFmTemplateResource_Read_APIErrorReadBody exercises FmTemplateResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFmTemplateResource_Read_APIErrorReadBody(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientReadErrorBody(t, 501)}
	m := FmTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFmTemplateResource_Read_InvalidJSON exercises FmTemplateResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFmTemplateResource_Read_InvalidJSON(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 200, "{{")}
	m := FmTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFmTemplateResource_Read_MapError exercises FmTemplateResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFmTemplateResource_Read_MapError(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 200, "{\"template_name\":12345}")}
	m := FmTemplateResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFmTemplateResource_Update_Happy exercises FmTemplateResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestFmTemplateResource_Update_Happy(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 200, "{}")}
	m := FmTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmTemplateResource_Update_NilClient exercises FmTemplateResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFmTemplateResource_Update_NilClient(t *testing.T) {
	r := &FmTemplateResource{}
	m := FmTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFmTemplateResource_Update_BuildError exercises FmTemplateResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFmTemplateResource_Update_BuildError(t *testing.T) {
	r := &FmTemplateResource{client: newMalformedBaseURLClient(t)}
	m := FmTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFmTemplateResource_Update_SendError exercises FmTemplateResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestFmTemplateResource_Update_SendError(t *testing.T) {
	r := &FmTemplateResource{client: newTransportErrorClient(t)}
	m := FmTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFmTemplateResource_Update_APIError exercises FmTemplateResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFmTemplateResource_Update_APIError(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FmTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_fm_template")
}

// TestFmTemplateResource_Update_APIErrorReadBody exercises FmTemplateResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFmTemplateResource_Update_APIErrorReadBody(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientReadErrorBody(t, 501)}
	m := FmTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFmTemplateResource_Update_InvalidJSON exercises FmTemplateResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFmTemplateResource_Update_InvalidJSON(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 200, "{{")}
	m := FmTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestFmTemplateResource_Update_MapError exercises FmTemplateResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestFmTemplateResource_Update_MapError(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 200, "{\"template_name\":12345}")}
	m := FmTemplateResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestFmTemplateResource_Delete_Happy exercises FmTemplateResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestFmTemplateResource_Delete_Happy(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 204, "")}
	m := FmTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmTemplateResource_Delete_NilClient exercises FmTemplateResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFmTemplateResource_Delete_NilClient(t *testing.T) {
	r := &FmTemplateResource{}
	m := FmTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFmTemplateResource_Delete_BuildError exercises FmTemplateResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFmTemplateResource_Delete_BuildError(t *testing.T) {
	r := &FmTemplateResource{client: newMalformedBaseURLClient(t)}
	m := FmTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFmTemplateResource_Delete_SendError exercises FmTemplateResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestFmTemplateResource_Delete_SendError(t *testing.T) {
	r := &FmTemplateResource{client: newTransportErrorClient(t)}
	m := FmTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFmTemplateResource_Delete_NotFoundSuccess exercises FmTemplateResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestFmTemplateResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 404, "")}
	m := FmTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmTemplateResource_Delete_APIError exercises FmTemplateResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFmTemplateResource_Delete_APIError(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FmTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_fm_template")
}

// TestFmTemplateResource_Delete_APIErrorReadBody exercises FmTemplateResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFmTemplateResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &FmTemplateResource{client: newMockClientReadErrorBody(t, 501)}
	m := FmTemplateResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
