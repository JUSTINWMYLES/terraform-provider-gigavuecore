package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetDeactivableVblsDataSource_Read_Happy exercises GetDeactivableVblsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetDeactivableVblsDataSource_Read_Happy(t *testing.T) {
	r := &GetDeactivableVblsDataSource{client: newMockClientStatus(t, 200, "{\"vblActivations\":[]}")}
	m := GetDeactivableVblsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetDeactivableVblsDataSource_Read_NilClient exercises GetDeactivableVblsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetDeactivableVblsDataSource_Read_NilClient(t *testing.T) {
	r := &GetDeactivableVblsDataSource{}
	m := GetDeactivableVblsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetDeactivableVblsDataSource_Read_BuildError exercises GetDeactivableVblsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetDeactivableVblsDataSource_Read_BuildError(t *testing.T) {
	r := &GetDeactivableVblsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetDeactivableVblsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetDeactivableVblsDataSource_Read_SendError exercises GetDeactivableVblsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetDeactivableVblsDataSource_Read_SendError(t *testing.T) {
	r := &GetDeactivableVblsDataSource{client: newTransportErrorClient(t)}
	m := GetDeactivableVblsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetDeactivableVblsDataSource_Read_InvalidJSON exercises GetDeactivableVblsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetDeactivableVblsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetDeactivableVblsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetDeactivableVblsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
