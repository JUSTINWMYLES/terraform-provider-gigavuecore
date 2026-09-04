package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllSslDecryptionEndpointsDataSource_Read_Happy exercises LoadAllSslDecryptionEndpointsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllSslDecryptionEndpointsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllSslDecryptionEndpointsDataSource{client: newMockClientStatus(t, 200, "{\"endpoints\":[]}")}
	m := LoadAllSslDecryptionEndpointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllSslDecryptionEndpointsDataSource_Read_NilClient exercises LoadAllSslDecryptionEndpointsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllSslDecryptionEndpointsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllSslDecryptionEndpointsDataSource{}
	m := LoadAllSslDecryptionEndpointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllSslDecryptionEndpointsDataSource_Read_BuildError exercises LoadAllSslDecryptionEndpointsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllSslDecryptionEndpointsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllSslDecryptionEndpointsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllSslDecryptionEndpointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllSslDecryptionEndpointsDataSource_Read_SendError exercises LoadAllSslDecryptionEndpointsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllSslDecryptionEndpointsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllSslDecryptionEndpointsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllSslDecryptionEndpointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllSslDecryptionEndpointsDataSource_Read_InvalidJSON exercises LoadAllSslDecryptionEndpointsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllSslDecryptionEndpointsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllSslDecryptionEndpointsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllSslDecryptionEndpointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
