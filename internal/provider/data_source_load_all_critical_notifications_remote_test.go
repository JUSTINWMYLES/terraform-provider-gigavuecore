package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllCriticalNotificationsDataSource_Read_Happy exercises LoadAllCriticalNotificationsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllCriticalNotificationsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllCriticalNotificationsDataSource{client: newMockClientStatus(t, 200, "{\"message\":[]}")}
	m := LoadAllCriticalNotificationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllCriticalNotificationsDataSource_Read_NilClient exercises LoadAllCriticalNotificationsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllCriticalNotificationsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllCriticalNotificationsDataSource{}
	m := LoadAllCriticalNotificationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllCriticalNotificationsDataSource_Read_BuildError exercises LoadAllCriticalNotificationsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllCriticalNotificationsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllCriticalNotificationsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllCriticalNotificationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllCriticalNotificationsDataSource_Read_SendError exercises LoadAllCriticalNotificationsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllCriticalNotificationsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllCriticalNotificationsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllCriticalNotificationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllCriticalNotificationsDataSource_Read_InvalidJSON exercises LoadAllCriticalNotificationsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllCriticalNotificationsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllCriticalNotificationsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllCriticalNotificationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
