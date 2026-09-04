package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadFilterTemplateLimitBySlotIdDataSource_Read_Happy exercises LoadFilterTemplateLimitBySlotIdDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadFilterTemplateLimitBySlotIdDataSource_Read_Happy(t *testing.T) {
	r := &LoadFilterTemplateLimitBySlotIdDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadFilterTemplateLimitBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadFilterTemplateLimitBySlotIdDataSource_Read_NilClient exercises LoadFilterTemplateLimitBySlotIdDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadFilterTemplateLimitBySlotIdDataSource_Read_NilClient(t *testing.T) {
	r := &LoadFilterTemplateLimitBySlotIdDataSource{}
	m := LoadFilterTemplateLimitBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadFilterTemplateLimitBySlotIdDataSource_Read_BuildError exercises LoadFilterTemplateLimitBySlotIdDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadFilterTemplateLimitBySlotIdDataSource_Read_BuildError(t *testing.T) {
	r := &LoadFilterTemplateLimitBySlotIdDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadFilterTemplateLimitBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadFilterTemplateLimitBySlotIdDataSource_Read_SendError exercises LoadFilterTemplateLimitBySlotIdDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadFilterTemplateLimitBySlotIdDataSource_Read_SendError(t *testing.T) {
	r := &LoadFilterTemplateLimitBySlotIdDataSource{client: newTransportErrorClient(t)}
	m := LoadFilterTemplateLimitBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadFilterTemplateLimitBySlotIdDataSource_Read_NotFound exercises LoadFilterTemplateLimitBySlotIdDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadFilterTemplateLimitBySlotIdDataSource_Read_NotFound(t *testing.T) {
	r := &LoadFilterTemplateLimitBySlotIdDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadFilterTemplateLimitBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadFilterTemplateLimitBySlotIdDataSource_Read_APIError exercises LoadFilterTemplateLimitBySlotIdDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadFilterTemplateLimitBySlotIdDataSource_Read_APIError(t *testing.T) {
	r := &LoadFilterTemplateLimitBySlotIdDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadFilterTemplateLimitBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_filter_template_limit_by_slot_id")
}

// TestLoadFilterTemplateLimitBySlotIdDataSource_Read_APIErrorReadBody exercises LoadFilterTemplateLimitBySlotIdDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadFilterTemplateLimitBySlotIdDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadFilterTemplateLimitBySlotIdDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadFilterTemplateLimitBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadFilterTemplateLimitBySlotIdDataSource_Read_InvalidJSON exercises LoadFilterTemplateLimitBySlotIdDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadFilterTemplateLimitBySlotIdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadFilterTemplateLimitBySlotIdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadFilterTemplateLimitBySlotIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
