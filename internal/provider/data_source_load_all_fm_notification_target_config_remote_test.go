package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllFmNotificationTargetConfigDataSource_Read_Happy exercises LoadAllFmNotificationTargetConfigDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllFmNotificationTargetConfigDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllFmNotificationTargetConfigDataSource{client: newMockClientStatus(t, 200, "{\"fmNotificationTargetConfigs\":[]}")}
	m := LoadAllFmNotificationTargetConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllFmNotificationTargetConfigDataSource_Read_NilClient exercises LoadAllFmNotificationTargetConfigDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllFmNotificationTargetConfigDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllFmNotificationTargetConfigDataSource{}
	m := LoadAllFmNotificationTargetConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllFmNotificationTargetConfigDataSource_Read_BuildError exercises LoadAllFmNotificationTargetConfigDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllFmNotificationTargetConfigDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllFmNotificationTargetConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllFmNotificationTargetConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllFmNotificationTargetConfigDataSource_Read_SendError exercises LoadAllFmNotificationTargetConfigDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllFmNotificationTargetConfigDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllFmNotificationTargetConfigDataSource{client: newTransportErrorClient(t)}
	m := LoadAllFmNotificationTargetConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllFmNotificationTargetConfigDataSource_Read_InvalidJSON exercises LoadAllFmNotificationTargetConfigDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllFmNotificationTargetConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllFmNotificationTargetConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllFmNotificationTargetConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
