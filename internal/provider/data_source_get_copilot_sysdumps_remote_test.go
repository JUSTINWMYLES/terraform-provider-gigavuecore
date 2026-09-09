package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCopilotSysdumpsDataSource_Read_Happy exercises GetCopilotSysdumpsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetCopilotSysdumpsDataSource_Read_Happy(t *testing.T) {
	r := &GetCopilotSysdumpsDataSource{client: newMockClientStatus(t, 200, "{\"sysdumpFiles\":[]}")}
	m := GetCopilotSysdumpsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetCopilotSysdumpsDataSource_Read_NilClient exercises GetCopilotSysdumpsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetCopilotSysdumpsDataSource_Read_NilClient(t *testing.T) {
	r := &GetCopilotSysdumpsDataSource{}
	m := GetCopilotSysdumpsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetCopilotSysdumpsDataSource_Read_BuildError exercises GetCopilotSysdumpsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetCopilotSysdumpsDataSource_Read_BuildError(t *testing.T) {
	r := &GetCopilotSysdumpsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetCopilotSysdumpsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetCopilotSysdumpsDataSource_Read_SendError exercises GetCopilotSysdumpsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetCopilotSysdumpsDataSource_Read_SendError(t *testing.T) {
	r := &GetCopilotSysdumpsDataSource{client: newTransportErrorClient(t)}
	m := GetCopilotSysdumpsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetCopilotSysdumpsDataSource_Read_InvalidJSON exercises GetCopilotSysdumpsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetCopilotSysdumpsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetCopilotSysdumpsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetCopilotSysdumpsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
