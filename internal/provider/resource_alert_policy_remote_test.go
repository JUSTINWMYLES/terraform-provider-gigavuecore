package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestAlertPolicyResource_Create_Happy exercises AlertPolicyResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestAlertPolicyResource_Create_Happy(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 200, "{\"policy_name\":\"example-id\"}")}
	m := AlertPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAlertPolicyResource_Create_NilClient exercises AlertPolicyResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAlertPolicyResource_Create_NilClient(t *testing.T) {
	r := &AlertPolicyResource{}
	m := AlertPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAlertPolicyResource_Create_BuildError exercises AlertPolicyResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAlertPolicyResource_Create_BuildError(t *testing.T) {
	r := &AlertPolicyResource{client: newMalformedBaseURLClient(t)}
	m := AlertPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAlertPolicyResource_Create_SendError exercises AlertPolicyResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestAlertPolicyResource_Create_SendError(t *testing.T) {
	r := &AlertPolicyResource{client: newTransportErrorClient(t)}
	m := AlertPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAlertPolicyResource_Create_APIError exercises AlertPolicyResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAlertPolicyResource_Create_APIError(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AlertPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_alert_policy")
}

// TestAlertPolicyResource_Create_APIErrorReadBody exercises AlertPolicyResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAlertPolicyResource_Create_APIErrorReadBody(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientReadErrorBody(t, 501)}
	m := AlertPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestAlertPolicyResource_Create_InvalidJSON exercises AlertPolicyResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestAlertPolicyResource_Create_InvalidJSON(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 200, "{{")}
	m := AlertPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestAlertPolicyResource_Create_MapError exercises AlertPolicyResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestAlertPolicyResource_Create_MapError(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 200, "{\"policy_name\":12345}")}
	m := AlertPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestAlertPolicyResource_Create_MissingID exercises AlertPolicyResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestAlertPolicyResource_Create_MissingID(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 200, "{}")}
	m := AlertPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestAlertPolicyResource_Create_LocationFallback exercises AlertPolicyResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header's trailing path segment.
func TestAlertPolicyResource_Create_LocationFallback(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientWithLocation(t, 200, "http://example.test/folders/example-id", "{}")}
	m := AlertPolicyResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.PolicyName.ValueString() != "example-id" {
		t.Fatalf("identifier = %q, want %q", m.PolicyName.ValueString(), "example-id")
	}
}

// TestAlertPolicyResource_Read_Happy exercises AlertPolicyResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestAlertPolicyResource_Read_Happy(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 200, "{}")}
	m := AlertPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestAlertPolicyResource_Read_NilClient exercises AlertPolicyResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAlertPolicyResource_Read_NilClient(t *testing.T) {
	r := &AlertPolicyResource{}
	m := AlertPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAlertPolicyResource_Read_BuildError exercises AlertPolicyResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAlertPolicyResource_Read_BuildError(t *testing.T) {
	r := &AlertPolicyResource{client: newMalformedBaseURLClient(t)}
	m := AlertPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAlertPolicyResource_Read_SendError exercises AlertPolicyResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestAlertPolicyResource_Read_SendError(t *testing.T) {
	r := &AlertPolicyResource{client: newTransportErrorClient(t)}
	m := AlertPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAlertPolicyResource_Read_NotFound exercises AlertPolicyResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestAlertPolicyResource_Read_NotFound(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 404, "")}
	m := AlertPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestAlertPolicyResource_Read_APIError exercises AlertPolicyResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAlertPolicyResource_Read_APIError(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AlertPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_alert_policy")
}

// TestAlertPolicyResource_Read_APIErrorReadBody exercises AlertPolicyResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAlertPolicyResource_Read_APIErrorReadBody(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientReadErrorBody(t, 501)}
	m := AlertPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestAlertPolicyResource_Read_InvalidJSON exercises AlertPolicyResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestAlertPolicyResource_Read_InvalidJSON(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 200, "{{")}
	m := AlertPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestAlertPolicyResource_Read_MapError exercises AlertPolicyResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestAlertPolicyResource_Read_MapError(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 200, "{\"policy_name\":12345}")}
	m := AlertPolicyResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestAlertPolicyResource_Update_Happy exercises AlertPolicyResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestAlertPolicyResource_Update_Happy(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 200, "{}")}
	m := AlertPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAlertPolicyResource_Update_NilClient exercises AlertPolicyResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAlertPolicyResource_Update_NilClient(t *testing.T) {
	r := &AlertPolicyResource{}
	m := AlertPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAlertPolicyResource_Update_BuildError exercises AlertPolicyResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAlertPolicyResource_Update_BuildError(t *testing.T) {
	r := &AlertPolicyResource{client: newMalformedBaseURLClient(t)}
	m := AlertPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAlertPolicyResource_Update_SendError exercises AlertPolicyResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestAlertPolicyResource_Update_SendError(t *testing.T) {
	r := &AlertPolicyResource{client: newTransportErrorClient(t)}
	m := AlertPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAlertPolicyResource_Update_APIError exercises AlertPolicyResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAlertPolicyResource_Update_APIError(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AlertPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_alert_policy")
}

// TestAlertPolicyResource_Update_APIErrorReadBody exercises AlertPolicyResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAlertPolicyResource_Update_APIErrorReadBody(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientReadErrorBody(t, 501)}
	m := AlertPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestAlertPolicyResource_Update_InvalidJSON exercises AlertPolicyResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestAlertPolicyResource_Update_InvalidJSON(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 200, "{{")}
	m := AlertPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestAlertPolicyResource_Update_MapError exercises AlertPolicyResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestAlertPolicyResource_Update_MapError(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 200, "{\"policy_name\":12345}")}
	m := AlertPolicyResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestAlertPolicyResource_Delete_Happy exercises AlertPolicyResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestAlertPolicyResource_Delete_Happy(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 204, "")}
	m := AlertPolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAlertPolicyResource_Delete_NilClient exercises AlertPolicyResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAlertPolicyResource_Delete_NilClient(t *testing.T) {
	r := &AlertPolicyResource{}
	m := AlertPolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAlertPolicyResource_Delete_BuildError exercises AlertPolicyResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAlertPolicyResource_Delete_BuildError(t *testing.T) {
	r := &AlertPolicyResource{client: newMalformedBaseURLClient(t)}
	m := AlertPolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAlertPolicyResource_Delete_SendError exercises AlertPolicyResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestAlertPolicyResource_Delete_SendError(t *testing.T) {
	r := &AlertPolicyResource{client: newTransportErrorClient(t)}
	m := AlertPolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAlertPolicyResource_Delete_NotFoundSuccess exercises AlertPolicyResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestAlertPolicyResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 404, "")}
	m := AlertPolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAlertPolicyResource_Delete_APIError exercises AlertPolicyResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAlertPolicyResource_Delete_APIError(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AlertPolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_alert_policy")
}

// TestAlertPolicyResource_Delete_APIErrorReadBody exercises AlertPolicyResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAlertPolicyResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &AlertPolicyResource{client: newMockClientReadErrorBody(t, 501)}
	m := AlertPolicyResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
