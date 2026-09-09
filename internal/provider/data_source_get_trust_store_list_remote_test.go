package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTrustStoreListDataSource_Read_Happy exercises GetTrustStoreListDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetTrustStoreListDataSource_Read_Happy(t *testing.T) {
	r := &GetTrustStoreListDataSource{client: newMockClientStatus(t, 200, "{\"trustStoreList\":[]}")}
	m := GetTrustStoreListDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetTrustStoreListDataSource_Read_NilClient exercises GetTrustStoreListDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetTrustStoreListDataSource_Read_NilClient(t *testing.T) {
	r := &GetTrustStoreListDataSource{}
	m := GetTrustStoreListDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetTrustStoreListDataSource_Read_BuildError exercises GetTrustStoreListDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetTrustStoreListDataSource_Read_BuildError(t *testing.T) {
	r := &GetTrustStoreListDataSource{client: newMalformedBaseURLClient(t)}
	m := GetTrustStoreListDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetTrustStoreListDataSource_Read_SendError exercises GetTrustStoreListDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetTrustStoreListDataSource_Read_SendError(t *testing.T) {
	r := &GetTrustStoreListDataSource{client: newTransportErrorClient(t)}
	m := GetTrustStoreListDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetTrustStoreListDataSource_Read_InvalidJSON exercises GetTrustStoreListDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetTrustStoreListDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetTrustStoreListDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetTrustStoreListDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
