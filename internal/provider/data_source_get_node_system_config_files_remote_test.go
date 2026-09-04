package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNodeSystemConfigFilesDataSource_Read_Happy exercises GetNodeSystemConfigFilesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetNodeSystemConfigFilesDataSource_Read_Happy(t *testing.T) {
	r := &GetNodeSystemConfigFilesDataSource{client: newMockClientStatus(t, 200, "{\"files\":[]}")}
	m := GetNodeSystemConfigFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetNodeSystemConfigFilesDataSource_Read_NilClient exercises GetNodeSystemConfigFilesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetNodeSystemConfigFilesDataSource_Read_NilClient(t *testing.T) {
	r := &GetNodeSystemConfigFilesDataSource{}
	m := GetNodeSystemConfigFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetNodeSystemConfigFilesDataSource_Read_BuildError exercises GetNodeSystemConfigFilesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetNodeSystemConfigFilesDataSource_Read_BuildError(t *testing.T) {
	r := &GetNodeSystemConfigFilesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetNodeSystemConfigFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetNodeSystemConfigFilesDataSource_Read_SendError exercises GetNodeSystemConfigFilesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetNodeSystemConfigFilesDataSource_Read_SendError(t *testing.T) {
	r := &GetNodeSystemConfigFilesDataSource{client: newTransportErrorClient(t)}
	m := GetNodeSystemConfigFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetNodeSystemConfigFilesDataSource_Read_InvalidJSON exercises GetNodeSystemConfigFilesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetNodeSystemConfigFilesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetNodeSystemConfigFilesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetNodeSystemConfigFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
