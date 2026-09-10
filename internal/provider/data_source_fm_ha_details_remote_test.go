package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestFmHaDetailsDataSource_Read_Happy exercises FmHaDetailsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestFmHaDetailsDataSource_Read_Happy(t *testing.T) {
	r := &FmHaDetailsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := FmHaDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmHaDetailsDataSource_Read_NilClient exercises FmHaDetailsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFmHaDetailsDataSource_Read_NilClient(t *testing.T) {
	r := &FmHaDetailsDataSource{}
	m := FmHaDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFmHaDetailsDataSource_Read_BuildError exercises FmHaDetailsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFmHaDetailsDataSource_Read_BuildError(t *testing.T) {
	r := &FmHaDetailsDataSource{client: newMalformedBaseURLClient(t)}
	m := FmHaDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFmHaDetailsDataSource_Read_SendError exercises FmHaDetailsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestFmHaDetailsDataSource_Read_SendError(t *testing.T) {
	r := &FmHaDetailsDataSource{client: newTransportErrorClient(t)}
	m := FmHaDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFmHaDetailsDataSource_Read_NotFound exercises FmHaDetailsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestFmHaDetailsDataSource_Read_NotFound(t *testing.T) {
	r := &FmHaDetailsDataSource{client: newMockClientStatus(t, 404, "")}
	m := FmHaDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestFmHaDetailsDataSource_Read_APIError exercises FmHaDetailsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFmHaDetailsDataSource_Read_APIError(t *testing.T) {
	r := &FmHaDetailsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FmHaDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_fm_ha_details")
}

// TestFmHaDetailsDataSource_Read_APIErrorReadBody exercises FmHaDetailsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFmHaDetailsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &FmHaDetailsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := FmHaDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestFmHaDetailsDataSource_Read_InvalidJSON exercises FmHaDetailsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestFmHaDetailsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &FmHaDetailsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := FmHaDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
