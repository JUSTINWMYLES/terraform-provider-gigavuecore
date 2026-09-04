package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetDeployDataSource_Read_Happy exercises GetDeployDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetDeployDataSource_Read_Happy(t *testing.T) {
	r := &GetDeployDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetDeployDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetDeployDataSource_Read_NilClient exercises GetDeployDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetDeployDataSource_Read_NilClient(t *testing.T) {
	r := &GetDeployDataSource{}
	m := GetDeployDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetDeployDataSource_Read_BuildError exercises GetDeployDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetDeployDataSource_Read_BuildError(t *testing.T) {
	r := &GetDeployDataSource{client: newMalformedBaseURLClient(t)}
	m := GetDeployDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetDeployDataSource_Read_SendError exercises GetDeployDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetDeployDataSource_Read_SendError(t *testing.T) {
	r := &GetDeployDataSource{client: newTransportErrorClient(t)}
	m := GetDeployDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetDeployDataSource_Read_NotFound exercises GetDeployDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetDeployDataSource_Read_NotFound(t *testing.T) {
	r := &GetDeployDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetDeployDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetDeployDataSource_Read_APIError exercises GetDeployDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetDeployDataSource_Read_APIError(t *testing.T) {
	r := &GetDeployDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetDeployDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_deploy")
}

// TestGetDeployDataSource_Read_APIErrorReadBody exercises GetDeployDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetDeployDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetDeployDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetDeployDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetDeployDataSource_Read_InvalidJSON exercises GetDeployDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetDeployDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetDeployDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetDeployDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
