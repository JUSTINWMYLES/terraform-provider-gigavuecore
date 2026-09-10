package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAvisiPoliciesDataSource_Read_Happy exercises GetAvisiPoliciesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAvisiPoliciesDataSource_Read_Happy(t *testing.T) {
	r := &GetAvisiPoliciesDataSource{client: newMockClientStatus(t, 200, "{\"avPolicies\":[]}")}
	m := GetAvisiPoliciesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAvisiPoliciesDataSource_Read_NilClient exercises GetAvisiPoliciesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAvisiPoliciesDataSource_Read_NilClient(t *testing.T) {
	r := &GetAvisiPoliciesDataSource{}
	m := GetAvisiPoliciesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAvisiPoliciesDataSource_Read_BuildError exercises GetAvisiPoliciesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAvisiPoliciesDataSource_Read_BuildError(t *testing.T) {
	r := &GetAvisiPoliciesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAvisiPoliciesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAvisiPoliciesDataSource_Read_SendError exercises GetAvisiPoliciesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAvisiPoliciesDataSource_Read_SendError(t *testing.T) {
	r := &GetAvisiPoliciesDataSource{client: newTransportErrorClient(t)}
	m := GetAvisiPoliciesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAvisiPoliciesDataSource_Read_InvalidJSON exercises GetAvisiPoliciesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAvisiPoliciesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAvisiPoliciesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAvisiPoliciesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
