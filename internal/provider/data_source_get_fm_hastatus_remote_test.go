package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFmHastatusDataSource_Read_Happy exercises GetFmHastatusDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetFmHastatusDataSource_Read_Happy(t *testing.T) {
	r := &GetFmHastatusDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetFmHastatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFmHastatusDataSource_Read_NilClient exercises GetFmHastatusDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFmHastatusDataSource_Read_NilClient(t *testing.T) {
	r := &GetFmHastatusDataSource{}
	m := GetFmHastatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFmHastatusDataSource_Read_BuildError exercises GetFmHastatusDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetFmHastatusDataSource_Read_BuildError(t *testing.T) {
	r := &GetFmHastatusDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFmHastatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetFmHastatusDataSource_Read_SendError exercises GetFmHastatusDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetFmHastatusDataSource_Read_SendError(t *testing.T) {
	r := &GetFmHastatusDataSource{client: newTransportErrorClient(t)}
	m := GetFmHastatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetFmHastatusDataSource_Read_NotFound exercises GetFmHastatusDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetFmHastatusDataSource_Read_NotFound(t *testing.T) {
	r := &GetFmHastatusDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetFmHastatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetFmHastatusDataSource_Read_APIError exercises GetFmHastatusDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetFmHastatusDataSource_Read_APIError(t *testing.T) {
	r := &GetFmHastatusDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetFmHastatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_fm_hastatus")
}

// TestGetFmHastatusDataSource_Read_APIErrorReadBody exercises GetFmHastatusDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetFmHastatusDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetFmHastatusDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetFmHastatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetFmHastatusDataSource_Read_InvalidJSON exercises GetFmHastatusDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetFmHastatusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFmHastatusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFmHastatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
