package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetRfsSyncDataSource_Read_Happy exercises GetRfsSyncDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetRfsSyncDataSource_Read_Happy(t *testing.T) {
	r := &GetRfsSyncDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetRfsSyncDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetRfsSyncDataSource_Read_NilClient exercises GetRfsSyncDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetRfsSyncDataSource_Read_NilClient(t *testing.T) {
	r := &GetRfsSyncDataSource{}
	m := GetRfsSyncDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetRfsSyncDataSource_Read_BuildError exercises GetRfsSyncDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetRfsSyncDataSource_Read_BuildError(t *testing.T) {
	r := &GetRfsSyncDataSource{client: newMalformedBaseURLClient(t)}
	m := GetRfsSyncDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetRfsSyncDataSource_Read_SendError exercises GetRfsSyncDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetRfsSyncDataSource_Read_SendError(t *testing.T) {
	r := &GetRfsSyncDataSource{client: newTransportErrorClient(t)}
	m := GetRfsSyncDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetRfsSyncDataSource_Read_NotFound exercises GetRfsSyncDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetRfsSyncDataSource_Read_NotFound(t *testing.T) {
	r := &GetRfsSyncDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetRfsSyncDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetRfsSyncDataSource_Read_APIError exercises GetRfsSyncDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetRfsSyncDataSource_Read_APIError(t *testing.T) {
	r := &GetRfsSyncDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetRfsSyncDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_rfs_sync")
}

// TestGetRfsSyncDataSource_Read_APIErrorReadBody exercises GetRfsSyncDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetRfsSyncDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetRfsSyncDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetRfsSyncDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetRfsSyncDataSource_Read_InvalidJSON exercises GetRfsSyncDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetRfsSyncDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetRfsSyncDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetRfsSyncDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
