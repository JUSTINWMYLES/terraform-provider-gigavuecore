package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_Happy exercises GetLatestPeriodVolumeDetailedAndPeriodsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_Happy(t *testing.T) {
	r := &GetLatestPeriodVolumeDetailedAndPeriodsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetLatestPeriodVolumeDetailedAndPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_NilClient exercises GetLatestPeriodVolumeDetailedAndPeriodsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_NilClient(t *testing.T) {
	r := &GetLatestPeriodVolumeDetailedAndPeriodsDataSource{}
	m := GetLatestPeriodVolumeDetailedAndPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_BuildError exercises GetLatestPeriodVolumeDetailedAndPeriodsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_BuildError(t *testing.T) {
	r := &GetLatestPeriodVolumeDetailedAndPeriodsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetLatestPeriodVolumeDetailedAndPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_SendError exercises GetLatestPeriodVolumeDetailedAndPeriodsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_SendError(t *testing.T) {
	r := &GetLatestPeriodVolumeDetailedAndPeriodsDataSource{client: newTransportErrorClient(t)}
	m := GetLatestPeriodVolumeDetailedAndPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_NotFound exercises GetLatestPeriodVolumeDetailedAndPeriodsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_NotFound(t *testing.T) {
	r := &GetLatestPeriodVolumeDetailedAndPeriodsDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetLatestPeriodVolumeDetailedAndPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_APIError exercises GetLatestPeriodVolumeDetailedAndPeriodsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_APIError(t *testing.T) {
	r := &GetLatestPeriodVolumeDetailedAndPeriodsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetLatestPeriodVolumeDetailedAndPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_latest_period_volume_detailed_and_periods")
}

// TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_APIErrorReadBody exercises GetLatestPeriodVolumeDetailedAndPeriodsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetLatestPeriodVolumeDetailedAndPeriodsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetLatestPeriodVolumeDetailedAndPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_InvalidJSON exercises GetLatestPeriodVolumeDetailedAndPeriodsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetLatestPeriodVolumeDetailedAndPeriodsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetLatestPeriodVolumeDetailedAndPeriodsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetLatestPeriodVolumeDetailedAndPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
