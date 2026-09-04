package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPeriodVolumeByDateDataSource_Read_Happy exercises GetPeriodVolumeByDateDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetPeriodVolumeByDateDataSource_Read_Happy(t *testing.T) {
	r := &GetPeriodVolumeByDateDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetPeriodVolumeByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPeriodVolumeByDateDataSource_Read_NilClient exercises GetPeriodVolumeByDateDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPeriodVolumeByDateDataSource_Read_NilClient(t *testing.T) {
	r := &GetPeriodVolumeByDateDataSource{}
	m := GetPeriodVolumeByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPeriodVolumeByDateDataSource_Read_BuildError exercises GetPeriodVolumeByDateDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetPeriodVolumeByDateDataSource_Read_BuildError(t *testing.T) {
	r := &GetPeriodVolumeByDateDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPeriodVolumeByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetPeriodVolumeByDateDataSource_Read_SendError exercises GetPeriodVolumeByDateDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetPeriodVolumeByDateDataSource_Read_SendError(t *testing.T) {
	r := &GetPeriodVolumeByDateDataSource{client: newTransportErrorClient(t)}
	m := GetPeriodVolumeByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetPeriodVolumeByDateDataSource_Read_NotFound exercises GetPeriodVolumeByDateDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetPeriodVolumeByDateDataSource_Read_NotFound(t *testing.T) {
	r := &GetPeriodVolumeByDateDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetPeriodVolumeByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetPeriodVolumeByDateDataSource_Read_APIError exercises GetPeriodVolumeByDateDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetPeriodVolumeByDateDataSource_Read_APIError(t *testing.T) {
	r := &GetPeriodVolumeByDateDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetPeriodVolumeByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_period_volume_by_date")
}

// TestGetPeriodVolumeByDateDataSource_Read_APIErrorReadBody exercises GetPeriodVolumeByDateDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetPeriodVolumeByDateDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetPeriodVolumeByDateDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetPeriodVolumeByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetPeriodVolumeByDateDataSource_Read_InvalidJSON exercises GetPeriodVolumeByDateDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetPeriodVolumeByDateDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPeriodVolumeByDateDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPeriodVolumeByDateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
