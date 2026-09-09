package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadEventsDataSource_Read_Happy exercises LoadEventsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadEventsDataSource_Read_Happy(t *testing.T) {
	r := &LoadEventsDataSource{client: newMockClientStatus(t, 200, "{\"events\":[]}")}
	m := LoadEventsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadEventsDataSource_Read_NilClient exercises LoadEventsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadEventsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadEventsDataSource{}
	m := LoadEventsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadEventsDataSource_Read_BuildError exercises LoadEventsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadEventsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadEventsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadEventsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadEventsDataSource_Read_SendError exercises LoadEventsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadEventsDataSource_Read_SendError(t *testing.T) {
	r := &LoadEventsDataSource{client: newTransportErrorClient(t)}
	m := LoadEventsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadEventsDataSource_Read_InvalidJSON exercises LoadEventsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadEventsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadEventsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadEventsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
