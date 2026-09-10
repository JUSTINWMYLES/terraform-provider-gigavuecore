package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetKeystoreKeysDataSource_Read_Happy exercises GetKeystoreKeysDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetKeystoreKeysDataSource_Read_Happy(t *testing.T) {
	r := &GetKeystoreKeysDataSource{client: newMockClientStatus(t, 200, "{\"keys\":[]}")}
	m := GetKeystoreKeysDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetKeystoreKeysDataSource_Read_NilClient exercises GetKeystoreKeysDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetKeystoreKeysDataSource_Read_NilClient(t *testing.T) {
	r := &GetKeystoreKeysDataSource{}
	m := GetKeystoreKeysDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetKeystoreKeysDataSource_Read_BuildError exercises GetKeystoreKeysDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetKeystoreKeysDataSource_Read_BuildError(t *testing.T) {
	r := &GetKeystoreKeysDataSource{client: newMalformedBaseURLClient(t)}
	m := GetKeystoreKeysDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetKeystoreKeysDataSource_Read_SendError exercises GetKeystoreKeysDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetKeystoreKeysDataSource_Read_SendError(t *testing.T) {
	r := &GetKeystoreKeysDataSource{client: newTransportErrorClient(t)}
	m := GetKeystoreKeysDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetKeystoreKeysDataSource_Read_InvalidJSON exercises GetKeystoreKeysDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetKeystoreKeysDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetKeystoreKeysDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetKeystoreKeysDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
