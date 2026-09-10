package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestExpiryCountFlAndNllDataSource_Read_Happy exercises ExpiryCountFlAndNllDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestExpiryCountFlAndNllDataSource_Read_Happy(t *testing.T) {
	r := &ExpiryCountFlAndNllDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := ExpiryCountFlAndNllDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExpiryCountFlAndNllDataSource_Read_NilClient exercises ExpiryCountFlAndNllDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExpiryCountFlAndNllDataSource_Read_NilClient(t *testing.T) {
	r := &ExpiryCountFlAndNllDataSource{}
	m := ExpiryCountFlAndNllDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExpiryCountFlAndNllDataSource_Read_BuildError exercises ExpiryCountFlAndNllDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestExpiryCountFlAndNllDataSource_Read_BuildError(t *testing.T) {
	r := &ExpiryCountFlAndNllDataSource{client: newMalformedBaseURLClient(t)}
	m := ExpiryCountFlAndNllDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestExpiryCountFlAndNllDataSource_Read_SendError exercises ExpiryCountFlAndNllDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestExpiryCountFlAndNllDataSource_Read_SendError(t *testing.T) {
	r := &ExpiryCountFlAndNllDataSource{client: newTransportErrorClient(t)}
	m := ExpiryCountFlAndNllDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestExpiryCountFlAndNllDataSource_Read_InvalidJSON exercises ExpiryCountFlAndNllDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestExpiryCountFlAndNllDataSource_Read_InvalidJSON(t *testing.T) {
	r := &ExpiryCountFlAndNllDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := ExpiryCountFlAndNllDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
