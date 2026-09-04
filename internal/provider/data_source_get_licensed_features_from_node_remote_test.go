package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetLicensedFeaturesFromNodeDataSource_Read_Happy exercises GetLicensedFeaturesFromNodeDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetLicensedFeaturesFromNodeDataSource_Read_Happy(t *testing.T) {
	r := &GetLicensedFeaturesFromNodeDataSource{client: newMockClientStatus(t, 200, "{\"licenses\":[]}")}
	m := GetLicensedFeaturesFromNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetLicensedFeaturesFromNodeDataSource_Read_NilClient exercises GetLicensedFeaturesFromNodeDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetLicensedFeaturesFromNodeDataSource_Read_NilClient(t *testing.T) {
	r := &GetLicensedFeaturesFromNodeDataSource{}
	m := GetLicensedFeaturesFromNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetLicensedFeaturesFromNodeDataSource_Read_BuildError exercises GetLicensedFeaturesFromNodeDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetLicensedFeaturesFromNodeDataSource_Read_BuildError(t *testing.T) {
	r := &GetLicensedFeaturesFromNodeDataSource{client: newMalformedBaseURLClient(t)}
	m := GetLicensedFeaturesFromNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetLicensedFeaturesFromNodeDataSource_Read_SendError exercises GetLicensedFeaturesFromNodeDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetLicensedFeaturesFromNodeDataSource_Read_SendError(t *testing.T) {
	r := &GetLicensedFeaturesFromNodeDataSource{client: newTransportErrorClient(t)}
	m := GetLicensedFeaturesFromNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetLicensedFeaturesFromNodeDataSource_Read_InvalidJSON exercises GetLicensedFeaturesFromNodeDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetLicensedFeaturesFromNodeDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetLicensedFeaturesFromNodeDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetLicensedFeaturesFromNodeDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
