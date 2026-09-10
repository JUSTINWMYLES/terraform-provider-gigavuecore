package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestIntentPolicyResource_Create_Happy exercises IntentPolicyResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestIntentPolicyResource_Create_Happy(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 201, "{\"name\":\"example-id\"}")}
	m := IntentPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestIntentPolicyResource_Create_NilClient exercises IntentPolicyResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIntentPolicyResource_Create_NilClient(t *testing.T) {
	r := &IntentPolicyResource{}
	m := IntentPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestIntentPolicyResource_Create_BuildError exercises IntentPolicyResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestIntentPolicyResource_Create_BuildError(t *testing.T) {
	r := &IntentPolicyResource{client: newMalformedBaseURLClient(t)}
	m := IntentPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestIntentPolicyResource_Create_SendError exercises IntentPolicyResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestIntentPolicyResource_Create_SendError(t *testing.T) {
	r := &IntentPolicyResource{client: newTransportErrorClient(t)}
	m := IntentPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestIntentPolicyResource_Create_APIError exercises IntentPolicyResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestIntentPolicyResource_Create_APIError(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := IntentPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_intent_policy")
}

// TestIntentPolicyResource_Create_APIErrorReadBody exercises IntentPolicyResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestIntentPolicyResource_Create_APIErrorReadBody(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientReadErrorBody(t, 500)}
	m := IntentPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestIntentPolicyResource_Create_InvalidJSON exercises IntentPolicyResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestIntentPolicyResource_Create_InvalidJSON(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 201, "{{")}
	m := IntentPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestIntentPolicyResource_Create_MapError exercises IntentPolicyResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestIntentPolicyResource_Create_MapError(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 201, "{\"name\":12345}")}
	m := IntentPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestIntentPolicyResource_Create_MissingID exercises IntentPolicyResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestIntentPolicyResource_Create_MissingID(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 201, "{}")}
	m := IntentPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestIntentPolicyResource_Create_LocationFallback exercises IntentPolicyResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestIntentPolicyResource_Create_LocationFallback(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := IntentPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Name.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.Name.ValueString(), "example-id")
	}
}

// TestIntentPolicyResource_Read_Happy exercises IntentPolicyResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestIntentPolicyResource_Read_Happy(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 200, "{}")}
	m := IntentPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestIntentPolicyResource_Read_NilClient exercises IntentPolicyResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIntentPolicyResource_Read_NilClient(t *testing.T) {
	r := &IntentPolicyResource{}
	m := IntentPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestIntentPolicyResource_Read_BuildError exercises IntentPolicyResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestIntentPolicyResource_Read_BuildError(t *testing.T) {
	r := &IntentPolicyResource{client: newMalformedBaseURLClient(t)}
	m := IntentPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestIntentPolicyResource_Read_SendError exercises IntentPolicyResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestIntentPolicyResource_Read_SendError(t *testing.T) {
	r := &IntentPolicyResource{client: newTransportErrorClient(t)}
	m := IntentPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestIntentPolicyResource_Read_NotFound exercises IntentPolicyResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestIntentPolicyResource_Read_NotFound(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 404, "")}
	m := IntentPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestIntentPolicyResource_Read_APIError exercises IntentPolicyResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestIntentPolicyResource_Read_APIError(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := IntentPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_intent_policy")
}

// TestIntentPolicyResource_Read_APIErrorReadBody exercises IntentPolicyResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestIntentPolicyResource_Read_APIErrorReadBody(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientReadErrorBody(t, 500)}
	m := IntentPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestIntentPolicyResource_Read_InvalidJSON exercises IntentPolicyResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestIntentPolicyResource_Read_InvalidJSON(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 200, "{{")}
	m := IntentPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestIntentPolicyResource_Read_MapError exercises IntentPolicyResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestIntentPolicyResource_Read_MapError(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 200, "{\"name\":12345}")}
	m := IntentPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestIntentPolicyResource_Update_Happy exercises IntentPolicyResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestIntentPolicyResource_Update_Happy(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 200, "{}")}
	m := IntentPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestIntentPolicyResource_Update_NilClient exercises IntentPolicyResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIntentPolicyResource_Update_NilClient(t *testing.T) {
	r := &IntentPolicyResource{}
	m := IntentPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestIntentPolicyResource_Update_BuildError exercises IntentPolicyResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestIntentPolicyResource_Update_BuildError(t *testing.T) {
	r := &IntentPolicyResource{client: newMalformedBaseURLClient(t)}
	m := IntentPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestIntentPolicyResource_Update_SendError exercises IntentPolicyResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestIntentPolicyResource_Update_SendError(t *testing.T) {
	r := &IntentPolicyResource{client: newTransportErrorClient(t)}
	m := IntentPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestIntentPolicyResource_Update_APIError exercises IntentPolicyResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestIntentPolicyResource_Update_APIError(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := IntentPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_intent_policy")
}

// TestIntentPolicyResource_Update_APIErrorReadBody exercises IntentPolicyResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestIntentPolicyResource_Update_APIErrorReadBody(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientReadErrorBody(t, 500)}
	m := IntentPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestIntentPolicyResource_Update_InvalidJSON exercises IntentPolicyResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestIntentPolicyResource_Update_InvalidJSON(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 200, "{{")}
	m := IntentPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestIntentPolicyResource_Update_MapError exercises IntentPolicyResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestIntentPolicyResource_Update_MapError(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 200, "{\"name\":12345}")}
	m := IntentPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestIntentPolicyResource_Delete_Happy exercises IntentPolicyResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestIntentPolicyResource_Delete_Happy(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 204, "")}
	m := IntentPolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestIntentPolicyResource_Delete_NilClient exercises IntentPolicyResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestIntentPolicyResource_Delete_NilClient(t *testing.T) {
	r := &IntentPolicyResource{}
	m := IntentPolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestIntentPolicyResource_Delete_BuildError exercises IntentPolicyResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestIntentPolicyResource_Delete_BuildError(t *testing.T) {
	r := &IntentPolicyResource{client: newMalformedBaseURLClient(t)}
	m := IntentPolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestIntentPolicyResource_Delete_SendError exercises IntentPolicyResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestIntentPolicyResource_Delete_SendError(t *testing.T) {
	r := &IntentPolicyResource{client: newTransportErrorClient(t)}
	m := IntentPolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestIntentPolicyResource_Delete_NotFoundSuccess exercises IntentPolicyResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestIntentPolicyResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 404, "")}
	m := IntentPolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestIntentPolicyResource_Delete_APIError exercises IntentPolicyResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestIntentPolicyResource_Delete_APIError(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := IntentPolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_intent_policy")
}

// TestIntentPolicyResource_Delete_APIErrorReadBody exercises IntentPolicyResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestIntentPolicyResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &IntentPolicyResource{client: newMockClientReadErrorBody(t, 500)}
	m := IntentPolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
