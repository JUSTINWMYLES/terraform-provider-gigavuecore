package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllNtpServerDataSource_Read_Happy exercises GetAllNtpServerDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetAllNtpServerDataSource_Read_Happy(t *testing.T) {
	r := &GetAllNtpServerDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetAllNtpServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllNtpServerDataSource_Read_NilClient exercises GetAllNtpServerDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllNtpServerDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllNtpServerDataSource{}
	m := GetAllNtpServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllNtpServerDataSource_Read_BuildError exercises GetAllNtpServerDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAllNtpServerDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllNtpServerDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllNtpServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAllNtpServerDataSource_Read_SendError exercises GetAllNtpServerDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAllNtpServerDataSource_Read_SendError(t *testing.T) {
	r := &GetAllNtpServerDataSource{client: newTransportErrorClient(t)}
	m := GetAllNtpServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAllNtpServerDataSource_Read_NotFound exercises GetAllNtpServerDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetAllNtpServerDataSource_Read_NotFound(t *testing.T) {
	r := &GetAllNtpServerDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetAllNtpServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetAllNtpServerDataSource_Read_APIError exercises GetAllNtpServerDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAllNtpServerDataSource_Read_APIError(t *testing.T) {
	r := &GetAllNtpServerDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetAllNtpServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_all_ntp_server")
}

// TestGetAllNtpServerDataSource_Read_APIErrorReadBody exercises GetAllNtpServerDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAllNtpServerDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetAllNtpServerDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetAllNtpServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetAllNtpServerDataSource_Read_InvalidJSON exercises GetAllNtpServerDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetAllNtpServerDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllNtpServerDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllNtpServerDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
