package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUserTokensDataSource_Read_Happy exercises GetUserTokensDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetUserTokensDataSource_Read_Happy(t *testing.T) {
	r := &GetUserTokensDataSource{client: newMockClientStatus(t, 200, "{\"fmUserTokenEntities\":[]}")}
	m := GetUserTokensDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetUserTokensDataSource_Read_NilClient exercises GetUserTokensDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetUserTokensDataSource_Read_NilClient(t *testing.T) {
	r := &GetUserTokensDataSource{}
	m := GetUserTokensDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetUserTokensDataSource_Read_BuildError exercises GetUserTokensDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetUserTokensDataSource_Read_BuildError(t *testing.T) {
	r := &GetUserTokensDataSource{client: newMalformedBaseURLClient(t)}
	m := GetUserTokensDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetUserTokensDataSource_Read_SendError exercises GetUserTokensDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetUserTokensDataSource_Read_SendError(t *testing.T) {
	r := &GetUserTokensDataSource{client: newTransportErrorClient(t)}
	m := GetUserTokensDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetUserTokensDataSource_Read_InvalidJSON exercises GetUserTokensDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetUserTokensDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetUserTokensDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetUserTokensDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
