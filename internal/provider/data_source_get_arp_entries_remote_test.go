package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetArpEntriesDataSource_Read_Happy exercises GetArpEntriesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetArpEntriesDataSource_Read_Happy(t *testing.T) {
	r := &GetArpEntriesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetArpEntriesDataSource_Read_NilClient exercises GetArpEntriesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetArpEntriesDataSource_Read_NilClient(t *testing.T) {
	r := &GetArpEntriesDataSource{}
	m := GetArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetArpEntriesDataSource_Read_BuildError exercises GetArpEntriesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetArpEntriesDataSource_Read_BuildError(t *testing.T) {
	r := &GetArpEntriesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetArpEntriesDataSource_Read_SendError exercises GetArpEntriesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetArpEntriesDataSource_Read_SendError(t *testing.T) {
	r := &GetArpEntriesDataSource{client: newTransportErrorClient(t)}
	m := GetArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetArpEntriesDataSource_Read_NotFound exercises GetArpEntriesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetArpEntriesDataSource_Read_NotFound(t *testing.T) {
	r := &GetArpEntriesDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetArpEntriesDataSource_Read_APIError exercises GetArpEntriesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetArpEntriesDataSource_Read_APIError(t *testing.T) {
	r := &GetArpEntriesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_arp_entries")
}

// TestGetArpEntriesDataSource_Read_APIErrorReadBody exercises GetArpEntriesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetArpEntriesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetArpEntriesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetArpEntriesDataSource_Read_InvalidJSON exercises GetArpEntriesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetArpEntriesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetArpEntriesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetArpEntriesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
