package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestEmailRecipientResource_Create_Happy exercises EmailRecipientResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestEmailRecipientResource_Create_Happy(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientStatus(t, 201, "{\"email_address\":\"example-id\"}")}
	m := EmailRecipientResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEmailRecipientResource_Create_NilClient exercises EmailRecipientResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEmailRecipientResource_Create_NilClient(t *testing.T) {
	r := &EmailRecipientResource{}
	m := EmailRecipientResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEmailRecipientResource_Create_BuildError exercises EmailRecipientResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEmailRecipientResource_Create_BuildError(t *testing.T) {
	r := &EmailRecipientResource{client: newMalformedBaseURLClient(t)}
	m := EmailRecipientResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEmailRecipientResource_Create_SendError exercises EmailRecipientResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestEmailRecipientResource_Create_SendError(t *testing.T) {
	r := &EmailRecipientResource{client: newTransportErrorClient(t)}
	m := EmailRecipientResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEmailRecipientResource_Create_APIError exercises EmailRecipientResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEmailRecipientResource_Create_APIError(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EmailRecipientResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_email_recipient")
}

// TestEmailRecipientResource_Create_APIErrorReadBody exercises EmailRecipientResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEmailRecipientResource_Create_APIErrorReadBody(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientReadErrorBody(t, 501)}
	m := EmailRecipientResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestEmailRecipientResource_Create_InvalidJSON exercises EmailRecipientResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestEmailRecipientResource_Create_InvalidJSON(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientStatus(t, 201, "{{")}
	m := EmailRecipientResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestEmailRecipientResource_Create_MapError exercises EmailRecipientResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestEmailRecipientResource_Create_MapError(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientStatus(t, 201, "{\"email_address\":12345}")}
	m := EmailRecipientResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestEmailRecipientResource_Create_MissingID exercises EmailRecipientResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestEmailRecipientResource_Create_MissingID(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientStatus(t, 201, "{}")}
	m := EmailRecipientResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestEmailRecipientResource_Create_LocationFallback exercises EmailRecipientResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestEmailRecipientResource_Create_LocationFallback(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := EmailRecipientResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.EmailAddress.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.EmailAddress.ValueString(), "example-id")
	}
}

// TestEmailRecipientResource_Read_Happy exercises EmailRecipientResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestEmailRecipientResource_Read_Happy(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientStatus(t, 200, "{}")}
	m := EmailRecipientResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestEmailRecipientResource_Read_NilClient exercises EmailRecipientResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEmailRecipientResource_Read_NilClient(t *testing.T) {
	r := &EmailRecipientResource{}
	m := EmailRecipientResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEmailRecipientResource_Read_BuildError exercises EmailRecipientResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEmailRecipientResource_Read_BuildError(t *testing.T) {
	r := &EmailRecipientResource{client: newMalformedBaseURLClient(t)}
	m := EmailRecipientResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEmailRecipientResource_Read_SendError exercises EmailRecipientResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestEmailRecipientResource_Read_SendError(t *testing.T) {
	r := &EmailRecipientResource{client: newTransportErrorClient(t)}
	m := EmailRecipientResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEmailRecipientResource_Read_NotFound exercises EmailRecipientResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestEmailRecipientResource_Read_NotFound(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientStatus(t, 404, "")}
	m := EmailRecipientResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestEmailRecipientResource_Read_APIError exercises EmailRecipientResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEmailRecipientResource_Read_APIError(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EmailRecipientResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_email_recipient")
}

// TestEmailRecipientResource_Read_APIErrorReadBody exercises EmailRecipientResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEmailRecipientResource_Read_APIErrorReadBody(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientReadErrorBody(t, 501)}
	m := EmailRecipientResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestEmailRecipientResource_Read_InvalidJSON exercises EmailRecipientResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestEmailRecipientResource_Read_InvalidJSON(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientStatus(t, 200, "{{")}
	m := EmailRecipientResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestEmailRecipientResource_Read_MapError exercises EmailRecipientResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestEmailRecipientResource_Read_MapError(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientStatus(t, 200, "{\"email_address\":12345}")}
	m := EmailRecipientResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestEmailRecipientResource_Delete_Happy exercises EmailRecipientResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestEmailRecipientResource_Delete_Happy(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientStatus(t, 204, "")}
	m := EmailRecipientResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEmailRecipientResource_Delete_NilClient exercises EmailRecipientResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEmailRecipientResource_Delete_NilClient(t *testing.T) {
	r := &EmailRecipientResource{}
	m := EmailRecipientResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEmailRecipientResource_Delete_BuildError exercises EmailRecipientResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEmailRecipientResource_Delete_BuildError(t *testing.T) {
	r := &EmailRecipientResource{client: newMalformedBaseURLClient(t)}
	m := EmailRecipientResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEmailRecipientResource_Delete_SendError exercises EmailRecipientResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestEmailRecipientResource_Delete_SendError(t *testing.T) {
	r := &EmailRecipientResource{client: newTransportErrorClient(t)}
	m := EmailRecipientResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEmailRecipientResource_Delete_NotFoundSuccess exercises EmailRecipientResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestEmailRecipientResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientStatus(t, 404, "")}
	m := EmailRecipientResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEmailRecipientResource_Delete_APIError exercises EmailRecipientResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEmailRecipientResource_Delete_APIError(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EmailRecipientResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_email_recipient")
}

// TestEmailRecipientResource_Delete_APIErrorReadBody exercises EmailRecipientResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEmailRecipientResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &EmailRecipientResource{client: newMockClientReadErrorBody(t, 501)}
	m := EmailRecipientResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
