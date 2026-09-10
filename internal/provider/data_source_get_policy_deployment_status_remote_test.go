package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPolicyDeploymentStatusDataSource_Read_Happy exercises GetPolicyDeploymentStatusDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetPolicyDeploymentStatusDataSource_Read_Happy(t *testing.T) {
	r := &GetPolicyDeploymentStatusDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetPolicyDeploymentStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPolicyDeploymentStatusDataSource_Read_NilClient exercises GetPolicyDeploymentStatusDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPolicyDeploymentStatusDataSource_Read_NilClient(t *testing.T) {
	r := &GetPolicyDeploymentStatusDataSource{}
	m := GetPolicyDeploymentStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPolicyDeploymentStatusDataSource_Read_BuildError exercises GetPolicyDeploymentStatusDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetPolicyDeploymentStatusDataSource_Read_BuildError(t *testing.T) {
	r := &GetPolicyDeploymentStatusDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPolicyDeploymentStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetPolicyDeploymentStatusDataSource_Read_SendError exercises GetPolicyDeploymentStatusDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetPolicyDeploymentStatusDataSource_Read_SendError(t *testing.T) {
	r := &GetPolicyDeploymentStatusDataSource{client: newTransportErrorClient(t)}
	m := GetPolicyDeploymentStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetPolicyDeploymentStatusDataSource_Read_NotFound exercises GetPolicyDeploymentStatusDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetPolicyDeploymentStatusDataSource_Read_NotFound(t *testing.T) {
	r := &GetPolicyDeploymentStatusDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetPolicyDeploymentStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetPolicyDeploymentStatusDataSource_Read_APIError exercises GetPolicyDeploymentStatusDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetPolicyDeploymentStatusDataSource_Read_APIError(t *testing.T) {
	r := &GetPolicyDeploymentStatusDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetPolicyDeploymentStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_policy_deployment_status")
}

// TestGetPolicyDeploymentStatusDataSource_Read_APIErrorReadBody exercises GetPolicyDeploymentStatusDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetPolicyDeploymentStatusDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetPolicyDeploymentStatusDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetPolicyDeploymentStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetPolicyDeploymentStatusDataSource_Read_InvalidJSON exercises GetPolicyDeploymentStatusDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetPolicyDeploymentStatusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPolicyDeploymentStatusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPolicyDeploymentStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
