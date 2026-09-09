package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEligibleRenewersDataSource_Read_Happy exercises GetEligibleRenewersDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetEligibleRenewersDataSource_Read_Happy(t *testing.T) {
	r := &GetEligibleRenewersDataSource{client: newMockClientStatus(t, 200, "{\"activations\":[]}")}
	m := GetEligibleRenewersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetEligibleRenewersDataSource_Read_NilClient exercises GetEligibleRenewersDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetEligibleRenewersDataSource_Read_NilClient(t *testing.T) {
	r := &GetEligibleRenewersDataSource{}
	m := GetEligibleRenewersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetEligibleRenewersDataSource_Read_BuildError exercises GetEligibleRenewersDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetEligibleRenewersDataSource_Read_BuildError(t *testing.T) {
	r := &GetEligibleRenewersDataSource{client: newMalformedBaseURLClient(t)}
	m := GetEligibleRenewersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetEligibleRenewersDataSource_Read_SendError exercises GetEligibleRenewersDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetEligibleRenewersDataSource_Read_SendError(t *testing.T) {
	r := &GetEligibleRenewersDataSource{client: newTransportErrorClient(t)}
	m := GetEligibleRenewersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetEligibleRenewersDataSource_Read_InvalidJSON exercises GetEligibleRenewersDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetEligibleRenewersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetEligibleRenewersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetEligibleRenewersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
