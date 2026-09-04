package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestPortFilterRuleResource_Create_Happy exercises PortFilterRuleResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestPortFilterRuleResource_Create_Happy(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 201, "{\"rule_id\":1}")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortFilterRuleResource_Create_NilClient exercises PortFilterRuleResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortFilterRuleResource_Create_NilClient(t *testing.T) {
	r := &PortFilterRuleResource{}
	m := PortFilterRuleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortFilterRuleResource_Create_BuildError exercises PortFilterRuleResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortFilterRuleResource_Create_BuildError(t *testing.T) {
	r := &PortFilterRuleResource{client: newMalformedBaseURLClient(t)}
	m := PortFilterRuleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortFilterRuleResource_Create_SendError exercises PortFilterRuleResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortFilterRuleResource_Create_SendError(t *testing.T) {
	r := &PortFilterRuleResource{client: newTransportErrorClient(t)}
	m := PortFilterRuleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortFilterRuleResource_Create_APIError exercises PortFilterRuleResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortFilterRuleResource_Create_APIError(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating gigavuecore_port_filter_rule")
}

// TestPortFilterRuleResource_Create_APIErrorReadBody exercises PortFilterRuleResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortFilterRuleResource_Create_APIErrorReadBody(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortFilterRuleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPortFilterRuleResource_Create_InvalidJSON exercises PortFilterRuleResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPortFilterRuleResource_Create_InvalidJSON(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 201, "{{")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPortFilterRuleResource_Create_MapError exercises PortFilterRuleResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPortFilterRuleResource_Create_MapError(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 201, "{\"rule_id\":\"not-valid\"}")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPortFilterRuleResource_Create_MissingID exercises PortFilterRuleResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestPortFilterRuleResource_Create_MissingID(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 201, "{}")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestPortFilterRuleResource_Read_Happy exercises PortFilterRuleResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestPortFilterRuleResource_Read_Happy(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 200, "{}")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortFilterRuleResource_Read_NilClient exercises PortFilterRuleResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortFilterRuleResource_Read_NilClient(t *testing.T) {
	r := &PortFilterRuleResource{}
	m := PortFilterRuleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortFilterRuleResource_Read_BuildError exercises PortFilterRuleResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortFilterRuleResource_Read_BuildError(t *testing.T) {
	r := &PortFilterRuleResource{client: newMalformedBaseURLClient(t)}
	m := PortFilterRuleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortFilterRuleResource_Read_SendError exercises PortFilterRuleResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortFilterRuleResource_Read_SendError(t *testing.T) {
	r := &PortFilterRuleResource{client: newTransportErrorClient(t)}
	m := PortFilterRuleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortFilterRuleResource_Read_NotFound exercises PortFilterRuleResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestPortFilterRuleResource_Read_NotFound(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 404, "")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortFilterRuleResource_Read_APIError exercises PortFilterRuleResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortFilterRuleResource_Read_APIError(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_port_filter_rule")
}

// TestPortFilterRuleResource_Read_APIErrorReadBody exercises PortFilterRuleResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortFilterRuleResource_Read_APIErrorReadBody(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientReadErrorBody(t, 500)}
	m := PortFilterRuleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPortFilterRuleResource_Read_InvalidJSON exercises PortFilterRuleResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPortFilterRuleResource_Read_InvalidJSON(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 200, "{{")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPortFilterRuleResource_Read_MapError exercises PortFilterRuleResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPortFilterRuleResource_Read_MapError(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 200, "{\"rule_id\":\"not-valid\"}")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPortFilterRuleResource_Update_Happy exercises PortFilterRuleResource.updateRemote against an httptest mock: happy path returns the success status with no errors.
func TestPortFilterRuleResource_Update_Happy(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 200, "{}")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortFilterRuleResource_Update_NilClient exercises PortFilterRuleResource.updateRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortFilterRuleResource_Update_NilClient(t *testing.T) {
	r := &PortFilterRuleResource{}
	m := PortFilterRuleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortFilterRuleResource_Update_BuildError exercises PortFilterRuleResource.updateRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortFilterRuleResource_Update_BuildError(t *testing.T) {
	r := &PortFilterRuleResource{client: newMalformedBaseURLClient(t)}
	m := PortFilterRuleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortFilterRuleResource_Update_SendError exercises PortFilterRuleResource.updateRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortFilterRuleResource_Update_SendError(t *testing.T) {
	r := &PortFilterRuleResource{client: newTransportErrorClient(t)}
	m := PortFilterRuleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortFilterRuleResource_Update_APIError exercises PortFilterRuleResource.updateRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortFilterRuleResource_Update_APIError(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error updating gigavuecore_port_filter_rule")
}

// TestPortFilterRuleResource_Update_APIErrorReadBody exercises PortFilterRuleResource.updateRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortFilterRuleResource_Update_APIErrorReadBody(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientReadErrorBody(t, 500)}
	m := PortFilterRuleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestPortFilterRuleResource_Update_InvalidJSON exercises PortFilterRuleResource.updateRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestPortFilterRuleResource_Update_InvalidJSON(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 200, "{{")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestPortFilterRuleResource_Update_MapError exercises PortFilterRuleResource.updateRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestPortFilterRuleResource_Update_MapError(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 200, "{\"rule_id\":\"not-valid\"}")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.UpdateResponse{}
	r.updateRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestPortFilterRuleResource_Delete_Happy exercises PortFilterRuleResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestPortFilterRuleResource_Delete_Happy(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 204, "")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortFilterRuleResource_Delete_NilClient exercises PortFilterRuleResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortFilterRuleResource_Delete_NilClient(t *testing.T) {
	r := &PortFilterRuleResource{}
	m := PortFilterRuleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortFilterRuleResource_Delete_BuildError exercises PortFilterRuleResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPortFilterRuleResource_Delete_BuildError(t *testing.T) {
	r := &PortFilterRuleResource{client: newMalformedBaseURLClient(t)}
	m := PortFilterRuleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPortFilterRuleResource_Delete_SendError exercises PortFilterRuleResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestPortFilterRuleResource_Delete_SendError(t *testing.T) {
	r := &PortFilterRuleResource{client: newTransportErrorClient(t)}
	m := PortFilterRuleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPortFilterRuleResource_Delete_NotFoundSuccess exercises PortFilterRuleResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestPortFilterRuleResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 404, "")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortFilterRuleResource_Delete_APIError exercises PortFilterRuleResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPortFilterRuleResource_Delete_APIError(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PortFilterRuleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting gigavuecore_port_filter_rule")
}

// TestPortFilterRuleResource_Delete_APIErrorReadBody exercises PortFilterRuleResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPortFilterRuleResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &PortFilterRuleResource{client: newMockClientReadErrorBody(t, 501)}
	m := PortFilterRuleResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
