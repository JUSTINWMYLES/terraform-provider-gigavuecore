package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteLdapUserGroupMappingAction_Invoke_Happy exercises DeleteLdapUserGroupMappingAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteLdapUserGroupMappingAction_Invoke_Happy(t *testing.T) {
	r := &DeleteLdapUserGroupMappingAction{client: newMockClientStatus(t, 200, "{}")}
	m := DeleteLdapUserGroupMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteLdapUserGroupMappingAction_Invoke_NilClient exercises DeleteLdapUserGroupMappingAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteLdapUserGroupMappingAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteLdapUserGroupMappingAction{}
	m := DeleteLdapUserGroupMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteLdapUserGroupMappingAction_Invoke_BuildError exercises DeleteLdapUserGroupMappingAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteLdapUserGroupMappingAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteLdapUserGroupMappingAction{client: newMalformedBaseURLClient(t)}
	m := DeleteLdapUserGroupMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteLdapUserGroupMappingAction_Invoke_SendError exercises DeleteLdapUserGroupMappingAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteLdapUserGroupMappingAction_Invoke_SendError(t *testing.T) {
	r := &DeleteLdapUserGroupMappingAction{client: newTransportErrorClient(t)}
	m := DeleteLdapUserGroupMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteLdapUserGroupMappingAction_Invoke_APIError exercises DeleteLdapUserGroupMappingAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteLdapUserGroupMappingAction_Invoke_APIError(t *testing.T) {
	r := &DeleteLdapUserGroupMappingAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteLdapUserGroupMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_ldap_user_group_mapping")
}

// TestDeleteLdapUserGroupMappingAction_Invoke_APIErrorReadBody exercises DeleteLdapUserGroupMappingAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteLdapUserGroupMappingAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteLdapUserGroupMappingAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteLdapUserGroupMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
