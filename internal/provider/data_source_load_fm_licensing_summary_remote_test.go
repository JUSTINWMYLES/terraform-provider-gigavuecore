package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadFmLicensingSummaryDataSource_Read_Happy exercises LoadFmLicensingSummaryDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadFmLicensingSummaryDataSource_Read_Happy(t *testing.T) {
	r := &LoadFmLicensingSummaryDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadFmLicensingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadFmLicensingSummaryDataSource_Read_NilClient exercises LoadFmLicensingSummaryDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadFmLicensingSummaryDataSource_Read_NilClient(t *testing.T) {
	r := &LoadFmLicensingSummaryDataSource{}
	m := LoadFmLicensingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadFmLicensingSummaryDataSource_Read_BuildError exercises LoadFmLicensingSummaryDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadFmLicensingSummaryDataSource_Read_BuildError(t *testing.T) {
	r := &LoadFmLicensingSummaryDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadFmLicensingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadFmLicensingSummaryDataSource_Read_SendError exercises LoadFmLicensingSummaryDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadFmLicensingSummaryDataSource_Read_SendError(t *testing.T) {
	r := &LoadFmLicensingSummaryDataSource{client: newTransportErrorClient(t)}
	m := LoadFmLicensingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadFmLicensingSummaryDataSource_Read_NotFound exercises LoadFmLicensingSummaryDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadFmLicensingSummaryDataSource_Read_NotFound(t *testing.T) {
	r := &LoadFmLicensingSummaryDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadFmLicensingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadFmLicensingSummaryDataSource_Read_APIError exercises LoadFmLicensingSummaryDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadFmLicensingSummaryDataSource_Read_APIError(t *testing.T) {
	r := &LoadFmLicensingSummaryDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadFmLicensingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_fm_licensing_summary")
}

// TestLoadFmLicensingSummaryDataSource_Read_APIErrorReadBody exercises LoadFmLicensingSummaryDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadFmLicensingSummaryDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadFmLicensingSummaryDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadFmLicensingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadFmLicensingSummaryDataSource_Read_InvalidJSON exercises LoadFmLicensingSummaryDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadFmLicensingSummaryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadFmLicensingSummaryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadFmLicensingSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
