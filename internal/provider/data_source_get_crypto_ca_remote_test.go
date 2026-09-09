package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCryptoCaDataSource_Read_Happy exercises GetCryptoCaDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetCryptoCaDataSource_Read_Happy(t *testing.T) {
	r := &GetCryptoCaDataSource{client: newMockClientStatus(t, 200, "{\"caList\":[]}")}
	m := GetCryptoCaDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetCryptoCaDataSource_Read_NilClient exercises GetCryptoCaDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetCryptoCaDataSource_Read_NilClient(t *testing.T) {
	r := &GetCryptoCaDataSource{}
	m := GetCryptoCaDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetCryptoCaDataSource_Read_BuildError exercises GetCryptoCaDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetCryptoCaDataSource_Read_BuildError(t *testing.T) {
	r := &GetCryptoCaDataSource{client: newMalformedBaseURLClient(t)}
	m := GetCryptoCaDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetCryptoCaDataSource_Read_SendError exercises GetCryptoCaDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetCryptoCaDataSource_Read_SendError(t *testing.T) {
	r := &GetCryptoCaDataSource{client: newTransportErrorClient(t)}
	m := GetCryptoCaDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetCryptoCaDataSource_Read_InvalidJSON exercises GetCryptoCaDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetCryptoCaDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetCryptoCaDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetCryptoCaDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
