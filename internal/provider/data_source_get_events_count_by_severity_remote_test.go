package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEventsCountBySeverityDataSource_Read_Happy exercises GetEventsCountBySeverityDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetEventsCountBySeverityDataSource_Read_Happy(t *testing.T) {
	r := &GetEventsCountBySeverityDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetEventsCountBySeverityDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetEventsCountBySeverityDataSource_Read_NilClient exercises GetEventsCountBySeverityDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetEventsCountBySeverityDataSource_Read_NilClient(t *testing.T) {
	r := &GetEventsCountBySeverityDataSource{}
	m := GetEventsCountBySeverityDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetEventsCountBySeverityDataSource_Read_BuildError exercises GetEventsCountBySeverityDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetEventsCountBySeverityDataSource_Read_BuildError(t *testing.T) {
	r := &GetEventsCountBySeverityDataSource{client: newMalformedBaseURLClient(t)}
	m := GetEventsCountBySeverityDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetEventsCountBySeverityDataSource_Read_SendError exercises GetEventsCountBySeverityDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetEventsCountBySeverityDataSource_Read_SendError(t *testing.T) {
	r := &GetEventsCountBySeverityDataSource{client: newTransportErrorClient(t)}
	m := GetEventsCountBySeverityDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetEventsCountBySeverityDataSource_Read_NotFound exercises GetEventsCountBySeverityDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetEventsCountBySeverityDataSource_Read_NotFound(t *testing.T) {
	r := &GetEventsCountBySeverityDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetEventsCountBySeverityDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetEventsCountBySeverityDataSource_Read_APIError exercises GetEventsCountBySeverityDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetEventsCountBySeverityDataSource_Read_APIError(t *testing.T) {
	r := &GetEventsCountBySeverityDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetEventsCountBySeverityDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_events_count_by_severity")
}

// TestGetEventsCountBySeverityDataSource_Read_APIErrorReadBody exercises GetEventsCountBySeverityDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetEventsCountBySeverityDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetEventsCountBySeverityDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetEventsCountBySeverityDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetEventsCountBySeverityDataSource_Read_InvalidJSON exercises GetEventsCountBySeverityDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetEventsCountBySeverityDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetEventsCountBySeverityDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetEventsCountBySeverityDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
