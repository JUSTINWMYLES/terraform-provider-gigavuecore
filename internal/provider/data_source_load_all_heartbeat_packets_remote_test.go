package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllHeartbeatPacketsDataSource_Read_Happy exercises LoadAllHeartbeatPacketsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllHeartbeatPacketsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllHeartbeatPacketsDataSource{client: newMockClientStatus(t, 200, "{\"hbCustomPackets\":[]}")}
	m := LoadAllHeartbeatPacketsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllHeartbeatPacketsDataSource_Read_NilClient exercises LoadAllHeartbeatPacketsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllHeartbeatPacketsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllHeartbeatPacketsDataSource{}
	m := LoadAllHeartbeatPacketsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllHeartbeatPacketsDataSource_Read_BuildError exercises LoadAllHeartbeatPacketsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllHeartbeatPacketsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllHeartbeatPacketsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllHeartbeatPacketsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllHeartbeatPacketsDataSource_Read_SendError exercises LoadAllHeartbeatPacketsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllHeartbeatPacketsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllHeartbeatPacketsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllHeartbeatPacketsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllHeartbeatPacketsDataSource_Read_InvalidJSON exercises LoadAllHeartbeatPacketsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllHeartbeatPacketsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllHeartbeatPacketsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllHeartbeatPacketsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
