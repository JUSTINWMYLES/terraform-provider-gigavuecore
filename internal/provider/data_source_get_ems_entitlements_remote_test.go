package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEmsEntitlementsDataSource_Read_Happy exercises GetEmsEntitlementsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetEmsEntitlementsDataSource_Read_Happy(t *testing.T) {
	r := &GetEmsEntitlementsDataSource{client: newMockClientStatus(t, 200, "{\"entitlements\":[]}")}
	m := GetEmsEntitlementsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetEmsEntitlementsDataSource_Read_NilClient exercises GetEmsEntitlementsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetEmsEntitlementsDataSource_Read_NilClient(t *testing.T) {
	r := &GetEmsEntitlementsDataSource{}
	m := GetEmsEntitlementsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetEmsEntitlementsDataSource_Read_BuildError exercises GetEmsEntitlementsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetEmsEntitlementsDataSource_Read_BuildError(t *testing.T) {
	r := &GetEmsEntitlementsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetEmsEntitlementsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetEmsEntitlementsDataSource_Read_SendError exercises GetEmsEntitlementsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetEmsEntitlementsDataSource_Read_SendError(t *testing.T) {
	r := &GetEmsEntitlementsDataSource{client: newTransportErrorClient(t)}
	m := GetEmsEntitlementsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetEmsEntitlementsDataSource_Read_InvalidJSON exercises GetEmsEntitlementsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetEmsEntitlementsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetEmsEntitlementsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetEmsEntitlementsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
