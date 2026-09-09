package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllSystemAcmeCertificateDetailsDataSource_Read_Happy exercises LoadAllSystemAcmeCertificateDetailsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllSystemAcmeCertificateDetailsDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllSystemAcmeCertificateDetailsDataSource{client: newMockClientStatus(t, 200, "{\"certificateInfo\":[]}")}
	m := LoadAllSystemAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllSystemAcmeCertificateDetailsDataSource_Read_NilClient exercises LoadAllSystemAcmeCertificateDetailsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllSystemAcmeCertificateDetailsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllSystemAcmeCertificateDetailsDataSource{}
	m := LoadAllSystemAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllSystemAcmeCertificateDetailsDataSource_Read_BuildError exercises LoadAllSystemAcmeCertificateDetailsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllSystemAcmeCertificateDetailsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllSystemAcmeCertificateDetailsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllSystemAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllSystemAcmeCertificateDetailsDataSource_Read_SendError exercises LoadAllSystemAcmeCertificateDetailsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllSystemAcmeCertificateDetailsDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllSystemAcmeCertificateDetailsDataSource{client: newTransportErrorClient(t)}
	m := LoadAllSystemAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllSystemAcmeCertificateDetailsDataSource_Read_InvalidJSON exercises LoadAllSystemAcmeCertificateDetailsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllSystemAcmeCertificateDetailsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllSystemAcmeCertificateDetailsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllSystemAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
