package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAcmeCertificateDetailsDataSource_Read_Happy exercises GetAcmeCertificateDetailsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAcmeCertificateDetailsDataSource_Read_Happy(t *testing.T) {
	r := &GetAcmeCertificateDetailsDataSource{client: newMockClientStatus(t, 200, "{\"acmeCertificate\":[]}")}
	m := GetAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAcmeCertificateDetailsDataSource_Read_NilClient exercises GetAcmeCertificateDetailsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAcmeCertificateDetailsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAcmeCertificateDetailsDataSource{}
	m := GetAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAcmeCertificateDetailsDataSource_Read_BuildError exercises GetAcmeCertificateDetailsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAcmeCertificateDetailsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAcmeCertificateDetailsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAcmeCertificateDetailsDataSource_Read_SendError exercises GetAcmeCertificateDetailsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAcmeCertificateDetailsDataSource_Read_SendError(t *testing.T) {
	r := &GetAcmeCertificateDetailsDataSource{client: newTransportErrorClient(t)}
	m := GetAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAcmeCertificateDetailsDataSource_Read_InvalidJSON exercises GetAcmeCertificateDetailsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAcmeCertificateDetailsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAcmeCertificateDetailsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
