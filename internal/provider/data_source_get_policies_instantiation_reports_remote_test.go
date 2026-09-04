package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPoliciesInstantiationReportsDataSource_Read_Happy exercises GetPoliciesInstantiationReportsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetPoliciesInstantiationReportsDataSource_Read_Happy(t *testing.T) {
	r := &GetPoliciesInstantiationReportsDataSource{client: newMockClientStatus(t, 200, "{\"avPolicyTriggerReports\":[]}")}
	m := GetPoliciesInstantiationReportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPoliciesInstantiationReportsDataSource_Read_NilClient exercises GetPoliciesInstantiationReportsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPoliciesInstantiationReportsDataSource_Read_NilClient(t *testing.T) {
	r := &GetPoliciesInstantiationReportsDataSource{}
	m := GetPoliciesInstantiationReportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPoliciesInstantiationReportsDataSource_Read_BuildError exercises GetPoliciesInstantiationReportsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetPoliciesInstantiationReportsDataSource_Read_BuildError(t *testing.T) {
	r := &GetPoliciesInstantiationReportsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPoliciesInstantiationReportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetPoliciesInstantiationReportsDataSource_Read_SendError exercises GetPoliciesInstantiationReportsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetPoliciesInstantiationReportsDataSource_Read_SendError(t *testing.T) {
	r := &GetPoliciesInstantiationReportsDataSource{client: newTransportErrorClient(t)}
	m := GetPoliciesInstantiationReportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetPoliciesInstantiationReportsDataSource_Read_InvalidJSON exercises GetPoliciesInstantiationReportsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetPoliciesInstantiationReportsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPoliciesInstantiationReportsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPoliciesInstantiationReportsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
