package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSslClientTrustStoreListDataSource_Read_Happy exercises GetSslClientTrustStoreListDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetSslClientTrustStoreListDataSource_Read_Happy(t *testing.T) {
	r := &GetSslClientTrustStoreListDataSource{client: newMockClientStatus(t, 200, "{\"clientTrustStores\":[]}")}
	m := GetSslClientTrustStoreListDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSslClientTrustStoreListDataSource_Read_NilClient exercises GetSslClientTrustStoreListDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSslClientTrustStoreListDataSource_Read_NilClient(t *testing.T) {
	r := &GetSslClientTrustStoreListDataSource{}
	m := GetSslClientTrustStoreListDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSslClientTrustStoreListDataSource_Read_BuildError exercises GetSslClientTrustStoreListDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetSslClientTrustStoreListDataSource_Read_BuildError(t *testing.T) {
	r := &GetSslClientTrustStoreListDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSslClientTrustStoreListDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetSslClientTrustStoreListDataSource_Read_SendError exercises GetSslClientTrustStoreListDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetSslClientTrustStoreListDataSource_Read_SendError(t *testing.T) {
	r := &GetSslClientTrustStoreListDataSource{client: newTransportErrorClient(t)}
	m := GetSslClientTrustStoreListDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetSslClientTrustStoreListDataSource_Read_InvalidJSON exercises GetSslClientTrustStoreListDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetSslClientTrustStoreListDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSslClientTrustStoreListDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSslClientTrustStoreListDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
