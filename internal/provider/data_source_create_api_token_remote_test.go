package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestCreateApiTokenDataSource_Read_Happy exercises CreateApiTokenDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestCreateApiTokenDataSource_Read_Happy(t *testing.T) {
	r := &CreateApiTokenDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := CreateApiTokenDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateApiTokenDataSource_Read_NilClient exercises CreateApiTokenDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateApiTokenDataSource_Read_NilClient(t *testing.T) {
	r := &CreateApiTokenDataSource{}
	m := CreateApiTokenDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateApiTokenDataSource_Read_BuildError exercises CreateApiTokenDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateApiTokenDataSource_Read_BuildError(t *testing.T) {
	r := &CreateApiTokenDataSource{client: newMalformedBaseURLClient(t)}
	m := CreateApiTokenDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateApiTokenDataSource_Read_SendError exercises CreateApiTokenDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateApiTokenDataSource_Read_SendError(t *testing.T) {
	r := &CreateApiTokenDataSource{client: newTransportErrorClient(t)}
	m := CreateApiTokenDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateApiTokenDataSource_Read_NotFound exercises CreateApiTokenDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestCreateApiTokenDataSource_Read_NotFound(t *testing.T) {
	r := &CreateApiTokenDataSource{client: newMockClientStatus(t, 404, "")}
	m := CreateApiTokenDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestCreateApiTokenDataSource_Read_APIError exercises CreateApiTokenDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateApiTokenDataSource_Read_APIError(t *testing.T) {
	r := &CreateApiTokenDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateApiTokenDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_create_api_token")
}

// TestCreateApiTokenDataSource_Read_APIErrorReadBody exercises CreateApiTokenDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateApiTokenDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &CreateApiTokenDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := CreateApiTokenDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestCreateApiTokenDataSource_Read_InvalidJSON exercises CreateApiTokenDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestCreateApiTokenDataSource_Read_InvalidJSON(t *testing.T) {
	r := &CreateApiTokenDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := CreateApiTokenDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
