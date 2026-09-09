package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMobilitySamNodeConfigsDataSource_Read_Happy exercises LoadMobilitySamNodeConfigsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadMobilitySamNodeConfigsDataSource_Read_Happy(t *testing.T) {
	r := &LoadMobilitySamNodeConfigsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadMobilitySamNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadMobilitySamNodeConfigsDataSource_Read_NilClient exercises LoadMobilitySamNodeConfigsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadMobilitySamNodeConfigsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadMobilitySamNodeConfigsDataSource{}
	m := LoadMobilitySamNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadMobilitySamNodeConfigsDataSource_Read_BuildError exercises LoadMobilitySamNodeConfigsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadMobilitySamNodeConfigsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadMobilitySamNodeConfigsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadMobilitySamNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadMobilitySamNodeConfigsDataSource_Read_SendError exercises LoadMobilitySamNodeConfigsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadMobilitySamNodeConfigsDataSource_Read_SendError(t *testing.T) {
	r := &LoadMobilitySamNodeConfigsDataSource{client: newTransportErrorClient(t)}
	m := LoadMobilitySamNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadMobilitySamNodeConfigsDataSource_Read_NotFound exercises LoadMobilitySamNodeConfigsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadMobilitySamNodeConfigsDataSource_Read_NotFound(t *testing.T) {
	r := &LoadMobilitySamNodeConfigsDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadMobilitySamNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadMobilitySamNodeConfigsDataSource_Read_APIError exercises LoadMobilitySamNodeConfigsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadMobilitySamNodeConfigsDataSource_Read_APIError(t *testing.T) {
	r := &LoadMobilitySamNodeConfigsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadMobilitySamNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_mobility_sam_node_configs")
}

// TestLoadMobilitySamNodeConfigsDataSource_Read_APIErrorReadBody exercises LoadMobilitySamNodeConfigsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadMobilitySamNodeConfigsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadMobilitySamNodeConfigsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadMobilitySamNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadMobilitySamNodeConfigsDataSource_Read_InvalidJSON exercises LoadMobilitySamNodeConfigsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadMobilitySamNodeConfigsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadMobilitySamNodeConfigsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadMobilitySamNodeConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
