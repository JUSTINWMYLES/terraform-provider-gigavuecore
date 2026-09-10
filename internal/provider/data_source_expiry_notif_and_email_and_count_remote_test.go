package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestExpiryNotifAndEmailAndCountDataSource_Read_Happy exercises ExpiryNotifAndEmailAndCountDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestExpiryNotifAndEmailAndCountDataSource_Read_Happy(t *testing.T) {
	r := &ExpiryNotifAndEmailAndCountDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := ExpiryNotifAndEmailAndCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExpiryNotifAndEmailAndCountDataSource_Read_NilClient exercises ExpiryNotifAndEmailAndCountDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExpiryNotifAndEmailAndCountDataSource_Read_NilClient(t *testing.T) {
	r := &ExpiryNotifAndEmailAndCountDataSource{}
	m := ExpiryNotifAndEmailAndCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExpiryNotifAndEmailAndCountDataSource_Read_BuildError exercises ExpiryNotifAndEmailAndCountDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestExpiryNotifAndEmailAndCountDataSource_Read_BuildError(t *testing.T) {
	r := &ExpiryNotifAndEmailAndCountDataSource{client: newMalformedBaseURLClient(t)}
	m := ExpiryNotifAndEmailAndCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestExpiryNotifAndEmailAndCountDataSource_Read_SendError exercises ExpiryNotifAndEmailAndCountDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestExpiryNotifAndEmailAndCountDataSource_Read_SendError(t *testing.T) {
	r := &ExpiryNotifAndEmailAndCountDataSource{client: newTransportErrorClient(t)}
	m := ExpiryNotifAndEmailAndCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestExpiryNotifAndEmailAndCountDataSource_Read_InvalidJSON exercises ExpiryNotifAndEmailAndCountDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestExpiryNotifAndEmailAndCountDataSource_Read_InvalidJSON(t *testing.T) {
	r := &ExpiryNotifAndEmailAndCountDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := ExpiryNotifAndEmailAndCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
