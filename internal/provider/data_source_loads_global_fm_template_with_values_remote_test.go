package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadsGlobalFmTemplateWithValuesDataSource_Read_Happy exercises LoadsGlobalFmTemplateWithValuesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadsGlobalFmTemplateWithValuesDataSource_Read_Happy(t *testing.T) {
	r := &LoadsGlobalFmTemplateWithValuesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadsGlobalFmTemplateWithValuesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadsGlobalFmTemplateWithValuesDataSource_Read_NilClient exercises LoadsGlobalFmTemplateWithValuesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadsGlobalFmTemplateWithValuesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadsGlobalFmTemplateWithValuesDataSource{}
	m := LoadsGlobalFmTemplateWithValuesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadsGlobalFmTemplateWithValuesDataSource_Read_BuildError exercises LoadsGlobalFmTemplateWithValuesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadsGlobalFmTemplateWithValuesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadsGlobalFmTemplateWithValuesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadsGlobalFmTemplateWithValuesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadsGlobalFmTemplateWithValuesDataSource_Read_SendError exercises LoadsGlobalFmTemplateWithValuesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadsGlobalFmTemplateWithValuesDataSource_Read_SendError(t *testing.T) {
	r := &LoadsGlobalFmTemplateWithValuesDataSource{client: newTransportErrorClient(t)}
	m := LoadsGlobalFmTemplateWithValuesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadsGlobalFmTemplateWithValuesDataSource_Read_NotFound exercises LoadsGlobalFmTemplateWithValuesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadsGlobalFmTemplateWithValuesDataSource_Read_NotFound(t *testing.T) {
	r := &LoadsGlobalFmTemplateWithValuesDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadsGlobalFmTemplateWithValuesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadsGlobalFmTemplateWithValuesDataSource_Read_APIError exercises LoadsGlobalFmTemplateWithValuesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadsGlobalFmTemplateWithValuesDataSource_Read_APIError(t *testing.T) {
	r := &LoadsGlobalFmTemplateWithValuesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadsGlobalFmTemplateWithValuesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_loads_global_fm_template_with_values")
}

// TestLoadsGlobalFmTemplateWithValuesDataSource_Read_APIErrorReadBody exercises LoadsGlobalFmTemplateWithValuesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadsGlobalFmTemplateWithValuesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadsGlobalFmTemplateWithValuesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadsGlobalFmTemplateWithValuesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadsGlobalFmTemplateWithValuesDataSource_Read_InvalidJSON exercises LoadsGlobalFmTemplateWithValuesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadsGlobalFmTemplateWithValuesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadsGlobalFmTemplateWithValuesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadsGlobalFmTemplateWithValuesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
