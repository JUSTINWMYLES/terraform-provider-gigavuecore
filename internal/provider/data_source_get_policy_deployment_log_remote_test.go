package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPolicyDeploymentLogDataSource_Read_Happy exercises GetPolicyDeploymentLogDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetPolicyDeploymentLogDataSource_Read_Happy(t *testing.T) {
	r := &GetPolicyDeploymentLogDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetPolicyDeploymentLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPolicyDeploymentLogDataSource_Read_NilClient exercises GetPolicyDeploymentLogDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPolicyDeploymentLogDataSource_Read_NilClient(t *testing.T) {
	r := &GetPolicyDeploymentLogDataSource{}
	m := GetPolicyDeploymentLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPolicyDeploymentLogDataSource_Read_BuildError exercises GetPolicyDeploymentLogDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetPolicyDeploymentLogDataSource_Read_BuildError(t *testing.T) {
	r := &GetPolicyDeploymentLogDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPolicyDeploymentLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetPolicyDeploymentLogDataSource_Read_SendError exercises GetPolicyDeploymentLogDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetPolicyDeploymentLogDataSource_Read_SendError(t *testing.T) {
	r := &GetPolicyDeploymentLogDataSource{client: newTransportErrorClient(t)}
	m := GetPolicyDeploymentLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetPolicyDeploymentLogDataSource_Read_NotFound exercises GetPolicyDeploymentLogDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetPolicyDeploymentLogDataSource_Read_NotFound(t *testing.T) {
	r := &GetPolicyDeploymentLogDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetPolicyDeploymentLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetPolicyDeploymentLogDataSource_Read_APIError exercises GetPolicyDeploymentLogDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetPolicyDeploymentLogDataSource_Read_APIError(t *testing.T) {
	r := &GetPolicyDeploymentLogDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetPolicyDeploymentLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_policy_deployment_log")
}

// TestGetPolicyDeploymentLogDataSource_Read_APIErrorReadBody exercises GetPolicyDeploymentLogDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetPolicyDeploymentLogDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetPolicyDeploymentLogDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetPolicyDeploymentLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetPolicyDeploymentLogDataSource_Read_InvalidJSON exercises GetPolicyDeploymentLogDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetPolicyDeploymentLogDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPolicyDeploymentLogDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPolicyDeploymentLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
