package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadFilterTemplateLimitsDataSource_Read_Happy exercises LoadFilterTemplateLimitsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadFilterTemplateLimitsDataSource_Read_Happy(t *testing.T) {
	r := &LoadFilterTemplateLimitsDataSource{client: newMockClientStatus(t, 200, "{\"filterTemplatesLimits\":[]}")}
	m := LoadFilterTemplateLimitsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadFilterTemplateLimitsDataSource_Read_NilClient exercises LoadFilterTemplateLimitsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadFilterTemplateLimitsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadFilterTemplateLimitsDataSource{}
	m := LoadFilterTemplateLimitsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadFilterTemplateLimitsDataSource_Read_BuildError exercises LoadFilterTemplateLimitsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadFilterTemplateLimitsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadFilterTemplateLimitsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadFilterTemplateLimitsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadFilterTemplateLimitsDataSource_Read_SendError exercises LoadFilterTemplateLimitsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadFilterTemplateLimitsDataSource_Read_SendError(t *testing.T) {
	r := &LoadFilterTemplateLimitsDataSource{client: newTransportErrorClient(t)}
	m := LoadFilterTemplateLimitsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadFilterTemplateLimitsDataSource_Read_InvalidJSON exercises LoadFilterTemplateLimitsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadFilterTemplateLimitsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadFilterTemplateLimitsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadFilterTemplateLimitsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
