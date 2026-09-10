package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTrustStoreDataSource_Read_Happy exercises GetTrustStoreDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetTrustStoreDataSource_Read_Happy(t *testing.T) {
	r := &GetTrustStoreDataSource{client: newMockClientStatus(t, 200, "{\"trustStoreEntries\":[]}")}
	m := GetTrustStoreDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetTrustStoreDataSource_Read_NilClient exercises GetTrustStoreDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetTrustStoreDataSource_Read_NilClient(t *testing.T) {
	r := &GetTrustStoreDataSource{}
	m := GetTrustStoreDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetTrustStoreDataSource_Read_BuildError exercises GetTrustStoreDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetTrustStoreDataSource_Read_BuildError(t *testing.T) {
	r := &GetTrustStoreDataSource{client: newMalformedBaseURLClient(t)}
	m := GetTrustStoreDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetTrustStoreDataSource_Read_SendError exercises GetTrustStoreDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetTrustStoreDataSource_Read_SendError(t *testing.T) {
	r := &GetTrustStoreDataSource{client: newTransportErrorClient(t)}
	m := GetTrustStoreDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetTrustStoreDataSource_Read_InvalidJSON exercises GetTrustStoreDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetTrustStoreDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetTrustStoreDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetTrustStoreDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
