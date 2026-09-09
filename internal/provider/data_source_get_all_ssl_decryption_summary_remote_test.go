package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllSslDecryptionSummaryDataSource_Read_Happy exercises GetAllSslDecryptionSummaryDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllSslDecryptionSummaryDataSource_Read_Happy(t *testing.T) {
	r := &GetAllSslDecryptionSummaryDataSource{client: newMockClientStatus(t, 200, "{\"sslDecryptionReportsSummary\":[]}")}
	m := GetAllSslDecryptionSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllSslDecryptionSummaryDataSource_Read_NilClient exercises GetAllSslDecryptionSummaryDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllSslDecryptionSummaryDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllSslDecryptionSummaryDataSource{}
	m := GetAllSslDecryptionSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllSslDecryptionSummaryDataSource_Read_BuildError exercises GetAllSslDecryptionSummaryDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllSslDecryptionSummaryDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllSslDecryptionSummaryDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllSslDecryptionSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllSslDecryptionSummaryDataSource_Read_SendError exercises GetAllSslDecryptionSummaryDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllSslDecryptionSummaryDataSource_Read_SendError(t *testing.T) {
	r := &GetAllSslDecryptionSummaryDataSource{client: newTransportErrorClient(t)}
	m := GetAllSslDecryptionSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllSslDecryptionSummaryDataSource_Read_InvalidJSON exercises GetAllSslDecryptionSummaryDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllSslDecryptionSummaryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllSslDecryptionSummaryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllSslDecryptionSummaryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
