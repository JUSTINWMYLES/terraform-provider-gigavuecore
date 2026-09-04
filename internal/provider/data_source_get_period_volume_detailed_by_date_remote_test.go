package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPeriodVolumeDetailedByDateDataSource_Read_Happy exercises GetPeriodVolumeDetailedByDateDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetPeriodVolumeDetailedByDateDataSource_Read_Happy(t *testing.T) {
	r := &GetPeriodVolumeDetailedByDateDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetPeriodVolumeDetailedByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPeriodVolumeDetailedByDateDataSource_Read_NilClient exercises GetPeriodVolumeDetailedByDateDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPeriodVolumeDetailedByDateDataSource_Read_NilClient(t *testing.T) {
	r := &GetPeriodVolumeDetailedByDateDataSource{}
	m := GetPeriodVolumeDetailedByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPeriodVolumeDetailedByDateDataSource_Read_BuildError exercises GetPeriodVolumeDetailedByDateDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetPeriodVolumeDetailedByDateDataSource_Read_BuildError(t *testing.T) {
	r := &GetPeriodVolumeDetailedByDateDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPeriodVolumeDetailedByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetPeriodVolumeDetailedByDateDataSource_Read_SendError exercises GetPeriodVolumeDetailedByDateDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetPeriodVolumeDetailedByDateDataSource_Read_SendError(t *testing.T) {
	r := &GetPeriodVolumeDetailedByDateDataSource{client: newTransportErrorClient(t)}
	m := GetPeriodVolumeDetailedByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetPeriodVolumeDetailedByDateDataSource_Read_NotFound exercises GetPeriodVolumeDetailedByDateDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetPeriodVolumeDetailedByDateDataSource_Read_NotFound(t *testing.T) {
	r := &GetPeriodVolumeDetailedByDateDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetPeriodVolumeDetailedByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetPeriodVolumeDetailedByDateDataSource_Read_APIError exercises GetPeriodVolumeDetailedByDateDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetPeriodVolumeDetailedByDateDataSource_Read_APIError(t *testing.T) {
	r := &GetPeriodVolumeDetailedByDateDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetPeriodVolumeDetailedByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_period_volume_detailed_by_date")
}

// TestGetPeriodVolumeDetailedByDateDataSource_Read_APIErrorReadBody exercises GetPeriodVolumeDetailedByDateDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetPeriodVolumeDetailedByDateDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetPeriodVolumeDetailedByDateDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetPeriodVolumeDetailedByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetPeriodVolumeDetailedByDateDataSource_Read_InvalidJSON exercises GetPeriodVolumeDetailedByDateDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetPeriodVolumeDetailedByDateDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPeriodVolumeDetailedByDateDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPeriodVolumeDetailedByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
