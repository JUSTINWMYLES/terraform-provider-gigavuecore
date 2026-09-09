package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetGsopInfoDataSource_Read_Happy exercises GetGsopInfoDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetGsopInfoDataSource_Read_Happy(t *testing.T) {
	r := &GetGsopInfoDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetGsopInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetGsopInfoDataSource_Read_NilClient exercises GetGsopInfoDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetGsopInfoDataSource_Read_NilClient(t *testing.T) {
	r := &GetGsopInfoDataSource{}
	m := GetGsopInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetGsopInfoDataSource_Read_BuildError exercises GetGsopInfoDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetGsopInfoDataSource_Read_BuildError(t *testing.T) {
	r := &GetGsopInfoDataSource{client: newMalformedBaseURLClient(t)}
	m := GetGsopInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetGsopInfoDataSource_Read_SendError exercises GetGsopInfoDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetGsopInfoDataSource_Read_SendError(t *testing.T) {
	r := &GetGsopInfoDataSource{client: newTransportErrorClient(t)}
	m := GetGsopInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetGsopInfoDataSource_Read_NotFound exercises GetGsopInfoDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetGsopInfoDataSource_Read_NotFound(t *testing.T) {
	r := &GetGsopInfoDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetGsopInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetGsopInfoDataSource_Read_APIError exercises GetGsopInfoDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetGsopInfoDataSource_Read_APIError(t *testing.T) {
	r := &GetGsopInfoDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetGsopInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_gsop_info")
}

// TestGetGsopInfoDataSource_Read_APIErrorReadBody exercises GetGsopInfoDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetGsopInfoDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetGsopInfoDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetGsopInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetGsopInfoDataSource_Read_InvalidJSON exercises GetGsopInfoDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetGsopInfoDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetGsopInfoDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetGsopInfoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
