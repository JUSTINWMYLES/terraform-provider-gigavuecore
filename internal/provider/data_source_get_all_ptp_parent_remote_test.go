package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllPtpParentDataSource_Read_Happy exercises GetAllPtpParentDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllPtpParentDataSource_Read_Happy(t *testing.T) {
	r := &GetAllPtpParentDataSource{client: newMockClientStatus(t, 200, "{\"parents\":[]}")}
	m := GetAllPtpParentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllPtpParentDataSource_Read_NilClient exercises GetAllPtpParentDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllPtpParentDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllPtpParentDataSource{}
	m := GetAllPtpParentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllPtpParentDataSource_Read_BuildError exercises GetAllPtpParentDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllPtpParentDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllPtpParentDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllPtpParentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllPtpParentDataSource_Read_SendError exercises GetAllPtpParentDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllPtpParentDataSource_Read_SendError(t *testing.T) {
	r := &GetAllPtpParentDataSource{client: newTransportErrorClient(t)}
	m := GetAllPtpParentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllPtpParentDataSource_Read_InvalidJSON exercises GetAllPtpParentDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllPtpParentDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllPtpParentDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllPtpParentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
