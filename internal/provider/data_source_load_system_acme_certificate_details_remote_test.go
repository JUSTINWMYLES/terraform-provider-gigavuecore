package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSystemAcmeCertificateDetailsDataSource_Read_Happy exercises LoadSystemAcmeCertificateDetailsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadSystemAcmeCertificateDetailsDataSource_Read_Happy(t *testing.T) {
	r := &LoadSystemAcmeCertificateDetailsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadSystemAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadSystemAcmeCertificateDetailsDataSource_Read_NilClient exercises LoadSystemAcmeCertificateDetailsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadSystemAcmeCertificateDetailsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadSystemAcmeCertificateDetailsDataSource{}
	m := LoadSystemAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadSystemAcmeCertificateDetailsDataSource_Read_BuildError exercises LoadSystemAcmeCertificateDetailsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadSystemAcmeCertificateDetailsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadSystemAcmeCertificateDetailsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadSystemAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadSystemAcmeCertificateDetailsDataSource_Read_SendError exercises LoadSystemAcmeCertificateDetailsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadSystemAcmeCertificateDetailsDataSource_Read_SendError(t *testing.T) {
	r := &LoadSystemAcmeCertificateDetailsDataSource{client: newTransportErrorClient(t)}
	m := LoadSystemAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadSystemAcmeCertificateDetailsDataSource_Read_NotFound exercises LoadSystemAcmeCertificateDetailsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadSystemAcmeCertificateDetailsDataSource_Read_NotFound(t *testing.T) {
	r := &LoadSystemAcmeCertificateDetailsDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadSystemAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadSystemAcmeCertificateDetailsDataSource_Read_APIError exercises LoadSystemAcmeCertificateDetailsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadSystemAcmeCertificateDetailsDataSource_Read_APIError(t *testing.T) {
	r := &LoadSystemAcmeCertificateDetailsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadSystemAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_system_acme_certificate_details")
}

// TestLoadSystemAcmeCertificateDetailsDataSource_Read_APIErrorReadBody exercises LoadSystemAcmeCertificateDetailsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadSystemAcmeCertificateDetailsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadSystemAcmeCertificateDetailsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadSystemAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadSystemAcmeCertificateDetailsDataSource_Read_InvalidJSON exercises LoadSystemAcmeCertificateDetailsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadSystemAcmeCertificateDetailsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadSystemAcmeCertificateDetailsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadSystemAcmeCertificateDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
