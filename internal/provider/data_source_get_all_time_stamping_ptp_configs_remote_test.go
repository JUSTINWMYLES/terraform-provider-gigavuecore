package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTimeStampingPtpConfigsDataSource_Read_Happy exercises GetAllTimeStampingPtpConfigsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllTimeStampingPtpConfigsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllTimeStampingPtpConfigsDataSource{client: newMockClientStatus(t, 200, "{\"ptpConfigs\":[]}")}
	m := GetAllTimeStampingPtpConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllTimeStampingPtpConfigsDataSource_Read_NilClient exercises GetAllTimeStampingPtpConfigsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllTimeStampingPtpConfigsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllTimeStampingPtpConfigsDataSource{}
	m := GetAllTimeStampingPtpConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllTimeStampingPtpConfigsDataSource_Read_BuildError exercises GetAllTimeStampingPtpConfigsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllTimeStampingPtpConfigsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllTimeStampingPtpConfigsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllTimeStampingPtpConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllTimeStampingPtpConfigsDataSource_Read_SendError exercises GetAllTimeStampingPtpConfigsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllTimeStampingPtpConfigsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllTimeStampingPtpConfigsDataSource{client: newTransportErrorClient(t)}
	m := GetAllTimeStampingPtpConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllTimeStampingPtpConfigsDataSource_Read_InvalidJSON exercises GetAllTimeStampingPtpConfigsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllTimeStampingPtpConfigsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllTimeStampingPtpConfigsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllTimeStampingPtpConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
