package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllVlanResourcesDataSource_Read_Happy exercises LoadAllVlanResourcesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadAllVlanResourcesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllVlanResourcesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadAllVlanResourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllVlanResourcesDataSource_Read_NilClient exercises LoadAllVlanResourcesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllVlanResourcesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllVlanResourcesDataSource{}
	m := LoadAllVlanResourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllVlanResourcesDataSource_Read_BuildError exercises LoadAllVlanResourcesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadAllVlanResourcesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllVlanResourcesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllVlanResourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadAllVlanResourcesDataSource_Read_SendError exercises LoadAllVlanResourcesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadAllVlanResourcesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllVlanResourcesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllVlanResourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadAllVlanResourcesDataSource_Read_NotFound exercises LoadAllVlanResourcesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadAllVlanResourcesDataSource_Read_NotFound(t *testing.T) {
	r := &LoadAllVlanResourcesDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadAllVlanResourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadAllVlanResourcesDataSource_Read_APIError exercises LoadAllVlanResourcesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadAllVlanResourcesDataSource_Read_APIError(t *testing.T) {
	r := &LoadAllVlanResourcesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadAllVlanResourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_all_vlan_resources")
}

// TestLoadAllVlanResourcesDataSource_Read_APIErrorReadBody exercises LoadAllVlanResourcesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadAllVlanResourcesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadAllVlanResourcesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadAllVlanResourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadAllVlanResourcesDataSource_Read_InvalidJSON exercises LoadAllVlanResourcesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadAllVlanResourcesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllVlanResourcesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllVlanResourcesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
