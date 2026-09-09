package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestResourceSelectionDataSource_Read_Happy exercises ResourceSelectionDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestResourceSelectionDataSource_Read_Happy(t *testing.T) {
	r := &ResourceSelectionDataSource{client: newMockClientStatus(t, 200, "{\"resources\":[]}")}
	m := ResourceSelectionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestResourceSelectionDataSource_Read_NilClient exercises ResourceSelectionDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestResourceSelectionDataSource_Read_NilClient(t *testing.T) {
	r := &ResourceSelectionDataSource{}
	m := ResourceSelectionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestResourceSelectionDataSource_Read_BuildError exercises ResourceSelectionDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestResourceSelectionDataSource_Read_BuildError(t *testing.T) {
	r := &ResourceSelectionDataSource{client: newMalformedBaseURLClient(t)}
	m := ResourceSelectionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestResourceSelectionDataSource_Read_SendError exercises ResourceSelectionDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestResourceSelectionDataSource_Read_SendError(t *testing.T) {
	r := &ResourceSelectionDataSource{client: newTransportErrorClient(t)}
	m := ResourceSelectionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestResourceSelectionDataSource_Read_InvalidJSON exercises ResourceSelectionDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestResourceSelectionDataSource_Read_InvalidJSON(t *testing.T) {
	r := &ResourceSelectionDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := ResourceSelectionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
