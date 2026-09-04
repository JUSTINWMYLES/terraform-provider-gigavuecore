package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetGigaPortByPortIdDataSource_Read_Happy exercises GetGigaPortByPortIdDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetGigaPortByPortIdDataSource_Read_Happy(t *testing.T) {
	r := &GetGigaPortByPortIdDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetGigaPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetGigaPortByPortIdDataSource_Read_NilClient exercises GetGigaPortByPortIdDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetGigaPortByPortIdDataSource_Read_NilClient(t *testing.T) {
	r := &GetGigaPortByPortIdDataSource{}
	m := GetGigaPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetGigaPortByPortIdDataSource_Read_BuildError exercises GetGigaPortByPortIdDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetGigaPortByPortIdDataSource_Read_BuildError(t *testing.T) {
	r := &GetGigaPortByPortIdDataSource{client: newMalformedBaseURLClient(t)}
	m := GetGigaPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetGigaPortByPortIdDataSource_Read_SendError exercises GetGigaPortByPortIdDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetGigaPortByPortIdDataSource_Read_SendError(t *testing.T) {
	r := &GetGigaPortByPortIdDataSource{client: newTransportErrorClient(t)}
	m := GetGigaPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetGigaPortByPortIdDataSource_Read_NotFound exercises GetGigaPortByPortIdDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetGigaPortByPortIdDataSource_Read_NotFound(t *testing.T) {
	r := &GetGigaPortByPortIdDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetGigaPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetGigaPortByPortIdDataSource_Read_APIError exercises GetGigaPortByPortIdDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetGigaPortByPortIdDataSource_Read_APIError(t *testing.T) {
	r := &GetGigaPortByPortIdDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetGigaPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_giga_port_by_port_id")
}

// TestGetGigaPortByPortIdDataSource_Read_APIErrorReadBody exercises GetGigaPortByPortIdDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetGigaPortByPortIdDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetGigaPortByPortIdDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetGigaPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetGigaPortByPortIdDataSource_Read_InvalidJSON exercises GetGigaPortByPortIdDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetGigaPortByPortIdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetGigaPortByPortIdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetGigaPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
