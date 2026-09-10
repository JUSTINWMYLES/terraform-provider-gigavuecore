package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllFilterResourcesBySlotIdDataSource_Read_Happy exercises LoadAllFilterResourcesBySlotIdDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadAllFilterResourcesBySlotIdDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllFilterResourcesBySlotIdDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadAllFilterResourcesBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllFilterResourcesBySlotIdDataSource_Read_NilClient exercises LoadAllFilterResourcesBySlotIdDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllFilterResourcesBySlotIdDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllFilterResourcesBySlotIdDataSource{}
	m := LoadAllFilterResourcesBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllFilterResourcesBySlotIdDataSource_Read_BuildError exercises LoadAllFilterResourcesBySlotIdDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadAllFilterResourcesBySlotIdDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllFilterResourcesBySlotIdDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllFilterResourcesBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadAllFilterResourcesBySlotIdDataSource_Read_SendError exercises LoadAllFilterResourcesBySlotIdDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadAllFilterResourcesBySlotIdDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllFilterResourcesBySlotIdDataSource{client: newTransportErrorClient(t)}
	m := LoadAllFilterResourcesBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadAllFilterResourcesBySlotIdDataSource_Read_NotFound exercises LoadAllFilterResourcesBySlotIdDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadAllFilterResourcesBySlotIdDataSource_Read_NotFound(t *testing.T) {
	r := &LoadAllFilterResourcesBySlotIdDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadAllFilterResourcesBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadAllFilterResourcesBySlotIdDataSource_Read_APIError exercises LoadAllFilterResourcesBySlotIdDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadAllFilterResourcesBySlotIdDataSource_Read_APIError(t *testing.T) {
	r := &LoadAllFilterResourcesBySlotIdDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadAllFilterResourcesBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_all_filter_resources_by_slot_id")
}

// TestLoadAllFilterResourcesBySlotIdDataSource_Read_APIErrorReadBody exercises LoadAllFilterResourcesBySlotIdDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadAllFilterResourcesBySlotIdDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadAllFilterResourcesBySlotIdDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadAllFilterResourcesBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadAllFilterResourcesBySlotIdDataSource_Read_InvalidJSON exercises LoadAllFilterResourcesBySlotIdDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadAllFilterResourcesBySlotIdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllFilterResourcesBySlotIdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllFilterResourcesBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
