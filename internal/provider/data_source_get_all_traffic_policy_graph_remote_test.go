package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTrafficPolicyGraphDataSource_Read_Happy exercises GetAllTrafficPolicyGraphDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllTrafficPolicyGraphDataSource_Read_Happy(t *testing.T) {
	r := &GetAllTrafficPolicyGraphDataSource{client: newMockClientStatus(t, 200, "{\"trafficPolicyGraphs\":[]}")}
	m := GetAllTrafficPolicyGraphDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllTrafficPolicyGraphDataSource_Read_NilClient exercises GetAllTrafficPolicyGraphDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllTrafficPolicyGraphDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllTrafficPolicyGraphDataSource{}
	m := GetAllTrafficPolicyGraphDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllTrafficPolicyGraphDataSource_Read_BuildError exercises GetAllTrafficPolicyGraphDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllTrafficPolicyGraphDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllTrafficPolicyGraphDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllTrafficPolicyGraphDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllTrafficPolicyGraphDataSource_Read_SendError exercises GetAllTrafficPolicyGraphDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllTrafficPolicyGraphDataSource_Read_SendError(t *testing.T) {
	r := &GetAllTrafficPolicyGraphDataSource{client: newTransportErrorClient(t)}
	m := GetAllTrafficPolicyGraphDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllTrafficPolicyGraphDataSource_Read_InvalidJSON exercises GetAllTrafficPolicyGraphDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllTrafficPolicyGraphDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllTrafficPolicyGraphDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllTrafficPolicyGraphDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
