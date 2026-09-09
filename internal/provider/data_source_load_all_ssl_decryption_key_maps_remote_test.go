package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllSslDecryptionKeyMapsDataSource_Read_Happy exercises LoadAllSslDecryptionKeyMapsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllSslDecryptionKeyMapsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllSslDecryptionKeyMapsDataSource{client: newMockClientStatus(t, 200, "{\"keyMaps\":[]}")}
	m := LoadAllSslDecryptionKeyMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllSslDecryptionKeyMapsDataSource_Read_NilClient exercises LoadAllSslDecryptionKeyMapsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllSslDecryptionKeyMapsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllSslDecryptionKeyMapsDataSource{}
	m := LoadAllSslDecryptionKeyMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllSslDecryptionKeyMapsDataSource_Read_BuildError exercises LoadAllSslDecryptionKeyMapsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllSslDecryptionKeyMapsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllSslDecryptionKeyMapsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllSslDecryptionKeyMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllSslDecryptionKeyMapsDataSource_Read_SendError exercises LoadAllSslDecryptionKeyMapsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllSslDecryptionKeyMapsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllSslDecryptionKeyMapsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllSslDecryptionKeyMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllSslDecryptionKeyMapsDataSource_Read_InvalidJSON exercises LoadAllSslDecryptionKeyMapsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllSslDecryptionKeyMapsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllSslDecryptionKeyMapsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllSslDecryptionKeyMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
