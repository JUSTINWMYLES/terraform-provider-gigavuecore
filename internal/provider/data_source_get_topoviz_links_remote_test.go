package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTopovizLinksDataSource_Read_Happy exercises GetTopovizLinksDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetTopovizLinksDataSource_Read_Happy(t *testing.T) {
	r := &GetTopovizLinksDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := GetTopovizLinksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetTopovizLinksDataSource_Read_NilClient exercises GetTopovizLinksDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetTopovizLinksDataSource_Read_NilClient(t *testing.T) {
	r := &GetTopovizLinksDataSource{}
	m := GetTopovizLinksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetTopovizLinksDataSource_Read_BuildError exercises GetTopovizLinksDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetTopovizLinksDataSource_Read_BuildError(t *testing.T) {
	r := &GetTopovizLinksDataSource{client: newMalformedBaseURLClient(t)}
	m := GetTopovizLinksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetTopovizLinksDataSource_Read_SendError exercises GetTopovizLinksDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetTopovizLinksDataSource_Read_SendError(t *testing.T) {
	r := &GetTopovizLinksDataSource{client: newTransportErrorClient(t)}
	m := GetTopovizLinksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetTopovizLinksDataSource_Read_InvalidJSON exercises GetTopovizLinksDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetTopovizLinksDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetTopovizLinksDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetTopovizLinksDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
