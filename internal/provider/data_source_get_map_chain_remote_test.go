package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMapChainDataSource_Read_Happy exercises GetMapChainDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetMapChainDataSource_Read_Happy(t *testing.T) {
	r := &GetMapChainDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetMapChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetMapChainDataSource_Read_NilClient exercises GetMapChainDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetMapChainDataSource_Read_NilClient(t *testing.T) {
	r := &GetMapChainDataSource{}
	m := GetMapChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetMapChainDataSource_Read_BuildError exercises GetMapChainDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetMapChainDataSource_Read_BuildError(t *testing.T) {
	r := &GetMapChainDataSource{client: newMalformedBaseURLClient(t)}
	m := GetMapChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetMapChainDataSource_Read_SendError exercises GetMapChainDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetMapChainDataSource_Read_SendError(t *testing.T) {
	r := &GetMapChainDataSource{client: newTransportErrorClient(t)}
	m := GetMapChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetMapChainDataSource_Read_NotFound exercises GetMapChainDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetMapChainDataSource_Read_NotFound(t *testing.T) {
	r := &GetMapChainDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetMapChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetMapChainDataSource_Read_APIError exercises GetMapChainDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetMapChainDataSource_Read_APIError(t *testing.T) {
	r := &GetMapChainDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetMapChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_map_chain")
}

// TestGetMapChainDataSource_Read_APIErrorReadBody exercises GetMapChainDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetMapChainDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetMapChainDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetMapChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetMapChainDataSource_Read_InvalidJSON exercises GetMapChainDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetMapChainDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetMapChainDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetMapChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
