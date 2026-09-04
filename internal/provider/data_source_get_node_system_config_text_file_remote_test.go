package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNodeSystemConfigTextFileDataSource_Read_Happy exercises GetNodeSystemConfigTextFileDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetNodeSystemConfigTextFileDataSource_Read_Happy(t *testing.T) {
	r := &GetNodeSystemConfigTextFileDataSource{client: newMockClientStatus(t, 200, "{\"configFiles\":[]}")}
	m := GetNodeSystemConfigTextFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetNodeSystemConfigTextFileDataSource_Read_NilClient exercises GetNodeSystemConfigTextFileDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetNodeSystemConfigTextFileDataSource_Read_NilClient(t *testing.T) {
	r := &GetNodeSystemConfigTextFileDataSource{}
	m := GetNodeSystemConfigTextFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetNodeSystemConfigTextFileDataSource_Read_BuildError exercises GetNodeSystemConfigTextFileDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetNodeSystemConfigTextFileDataSource_Read_BuildError(t *testing.T) {
	r := &GetNodeSystemConfigTextFileDataSource{client: newMalformedBaseURLClient(t)}
	m := GetNodeSystemConfigTextFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetNodeSystemConfigTextFileDataSource_Read_SendError exercises GetNodeSystemConfigTextFileDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetNodeSystemConfigTextFileDataSource_Read_SendError(t *testing.T) {
	r := &GetNodeSystemConfigTextFileDataSource{client: newTransportErrorClient(t)}
	m := GetNodeSystemConfigTextFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetNodeSystemConfigTextFileDataSource_Read_InvalidJSON exercises GetNodeSystemConfigTextFileDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetNodeSystemConfigTextFileDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetNodeSystemConfigTextFileDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetNodeSystemConfigTextFileDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
