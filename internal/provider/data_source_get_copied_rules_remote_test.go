package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCopiedRulesDataSource_Read_Happy exercises GetCopiedRulesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetCopiedRulesDataSource_Read_Happy(t *testing.T) {
	r := &GetCopiedRulesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetCopiedRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetCopiedRulesDataSource_Read_NilClient exercises GetCopiedRulesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetCopiedRulesDataSource_Read_NilClient(t *testing.T) {
	r := &GetCopiedRulesDataSource{}
	m := GetCopiedRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetCopiedRulesDataSource_Read_BuildError exercises GetCopiedRulesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetCopiedRulesDataSource_Read_BuildError(t *testing.T) {
	r := &GetCopiedRulesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetCopiedRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetCopiedRulesDataSource_Read_SendError exercises GetCopiedRulesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetCopiedRulesDataSource_Read_SendError(t *testing.T) {
	r := &GetCopiedRulesDataSource{client: newTransportErrorClient(t)}
	m := GetCopiedRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetCopiedRulesDataSource_Read_NotFound exercises GetCopiedRulesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetCopiedRulesDataSource_Read_NotFound(t *testing.T) {
	r := &GetCopiedRulesDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetCopiedRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetCopiedRulesDataSource_Read_APIError exercises GetCopiedRulesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetCopiedRulesDataSource_Read_APIError(t *testing.T) {
	r := &GetCopiedRulesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetCopiedRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_copied_rules")
}

// TestGetCopiedRulesDataSource_Read_APIErrorReadBody exercises GetCopiedRulesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetCopiedRulesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetCopiedRulesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetCopiedRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetCopiedRulesDataSource_Read_InvalidJSON exercises GetCopiedRulesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetCopiedRulesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetCopiedRulesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetCopiedRulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
