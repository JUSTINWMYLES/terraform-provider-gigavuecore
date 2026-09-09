package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMobilitySamNodeDataSource_Read_Happy exercises LoadMobilitySamNodeDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadMobilitySamNodeDataSource_Read_Happy(t *testing.T) {
	r := &LoadMobilitySamNodeDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadMobilitySamNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadMobilitySamNodeDataSource_Read_NilClient exercises LoadMobilitySamNodeDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadMobilitySamNodeDataSource_Read_NilClient(t *testing.T) {
	r := &LoadMobilitySamNodeDataSource{}
	m := LoadMobilitySamNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadMobilitySamNodeDataSource_Read_BuildError exercises LoadMobilitySamNodeDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadMobilitySamNodeDataSource_Read_BuildError(t *testing.T) {
	r := &LoadMobilitySamNodeDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadMobilitySamNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadMobilitySamNodeDataSource_Read_SendError exercises LoadMobilitySamNodeDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadMobilitySamNodeDataSource_Read_SendError(t *testing.T) {
	r := &LoadMobilitySamNodeDataSource{client: newTransportErrorClient(t)}
	m := LoadMobilitySamNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadMobilitySamNodeDataSource_Read_NotFound exercises LoadMobilitySamNodeDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadMobilitySamNodeDataSource_Read_NotFound(t *testing.T) {
	r := &LoadMobilitySamNodeDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadMobilitySamNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadMobilitySamNodeDataSource_Read_APIError exercises LoadMobilitySamNodeDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadMobilitySamNodeDataSource_Read_APIError(t *testing.T) {
	r := &LoadMobilitySamNodeDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadMobilitySamNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_mobility_sam_node")
}

// TestLoadMobilitySamNodeDataSource_Read_APIErrorReadBody exercises LoadMobilitySamNodeDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadMobilitySamNodeDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadMobilitySamNodeDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadMobilitySamNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadMobilitySamNodeDataSource_Read_InvalidJSON exercises LoadMobilitySamNodeDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadMobilitySamNodeDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadMobilitySamNodeDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadMobilitySamNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
