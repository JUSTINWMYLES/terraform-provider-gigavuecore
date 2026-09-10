package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEmailServerDataSource_Read_Happy exercises GetEmailServerDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetEmailServerDataSource_Read_Happy(t *testing.T) {
	r := &GetEmailServerDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetEmailServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetEmailServerDataSource_Read_NilClient exercises GetEmailServerDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetEmailServerDataSource_Read_NilClient(t *testing.T) {
	r := &GetEmailServerDataSource{}
	m := GetEmailServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetEmailServerDataSource_Read_BuildError exercises GetEmailServerDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetEmailServerDataSource_Read_BuildError(t *testing.T) {
	r := &GetEmailServerDataSource{client: newMalformedBaseURLClient(t)}
	m := GetEmailServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetEmailServerDataSource_Read_SendError exercises GetEmailServerDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetEmailServerDataSource_Read_SendError(t *testing.T) {
	r := &GetEmailServerDataSource{client: newTransportErrorClient(t)}
	m := GetEmailServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetEmailServerDataSource_Read_NotFound exercises GetEmailServerDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetEmailServerDataSource_Read_NotFound(t *testing.T) {
	r := &GetEmailServerDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetEmailServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetEmailServerDataSource_Read_APIError exercises GetEmailServerDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetEmailServerDataSource_Read_APIError(t *testing.T) {
	r := &GetEmailServerDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetEmailServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_email_server")
}

// TestGetEmailServerDataSource_Read_APIErrorReadBody exercises GetEmailServerDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetEmailServerDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetEmailServerDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetEmailServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetEmailServerDataSource_Read_InvalidJSON exercises GetEmailServerDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetEmailServerDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetEmailServerDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetEmailServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
