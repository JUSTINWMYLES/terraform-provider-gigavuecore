package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAcmeServerDetailsDataSource_Read_Happy exercises GetAcmeServerDetailsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAcmeServerDetailsDataSource_Read_Happy(t *testing.T) {
	r := &GetAcmeServerDetailsDataSource{client: newMockClientStatus(t, 200, "{\"fmAcmeServerConfigDetails\":[]}")}
	m := GetAcmeServerDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAcmeServerDetailsDataSource_Read_NilClient exercises GetAcmeServerDetailsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAcmeServerDetailsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAcmeServerDetailsDataSource{}
	m := GetAcmeServerDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAcmeServerDetailsDataSource_Read_BuildError exercises GetAcmeServerDetailsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAcmeServerDetailsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAcmeServerDetailsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAcmeServerDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAcmeServerDetailsDataSource_Read_SendError exercises GetAcmeServerDetailsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAcmeServerDetailsDataSource_Read_SendError(t *testing.T) {
	r := &GetAcmeServerDetailsDataSource{client: newTransportErrorClient(t)}
	m := GetAcmeServerDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAcmeServerDetailsDataSource_Read_InvalidJSON exercises GetAcmeServerDetailsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAcmeServerDetailsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAcmeServerDetailsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAcmeServerDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
