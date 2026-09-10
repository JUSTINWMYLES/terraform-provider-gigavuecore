package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestDisplayPeriodsDataSource_Read_Happy exercises DisplayPeriodsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestDisplayPeriodsDataSource_Read_Happy(t *testing.T) {
	r := &DisplayPeriodsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := DisplayPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDisplayPeriodsDataSource_Read_NilClient exercises DisplayPeriodsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDisplayPeriodsDataSource_Read_NilClient(t *testing.T) {
	r := &DisplayPeriodsDataSource{}
	m := DisplayPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDisplayPeriodsDataSource_Read_BuildError exercises DisplayPeriodsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDisplayPeriodsDataSource_Read_BuildError(t *testing.T) {
	r := &DisplayPeriodsDataSource{client: newMalformedBaseURLClient(t)}
	m := DisplayPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDisplayPeriodsDataSource_Read_SendError exercises DisplayPeriodsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestDisplayPeriodsDataSource_Read_SendError(t *testing.T) {
	r := &DisplayPeriodsDataSource{client: newTransportErrorClient(t)}
	m := DisplayPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDisplayPeriodsDataSource_Read_NotFound exercises DisplayPeriodsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestDisplayPeriodsDataSource_Read_NotFound(t *testing.T) {
	r := &DisplayPeriodsDataSource{client: newMockClientStatus(t, 404, "")}
	m := DisplayPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestDisplayPeriodsDataSource_Read_APIError exercises DisplayPeriodsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDisplayPeriodsDataSource_Read_APIError(t *testing.T) {
	r := &DisplayPeriodsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DisplayPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_display_periods")
}

// TestDisplayPeriodsDataSource_Read_APIErrorReadBody exercises DisplayPeriodsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDisplayPeriodsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &DisplayPeriodsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := DisplayPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestDisplayPeriodsDataSource_Read_InvalidJSON exercises DisplayPeriodsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestDisplayPeriodsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &DisplayPeriodsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := DisplayPeriodsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
