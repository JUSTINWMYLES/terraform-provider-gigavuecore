package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadLicensingConfigDataSource_Read_Happy exercises LoadLicensingConfigDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadLicensingConfigDataSource_Read_Happy(t *testing.T) {
	r := &LoadLicensingConfigDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadLicensingConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadLicensingConfigDataSource_Read_NilClient exercises LoadLicensingConfigDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadLicensingConfigDataSource_Read_NilClient(t *testing.T) {
	r := &LoadLicensingConfigDataSource{}
	m := LoadLicensingConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadLicensingConfigDataSource_Read_BuildError exercises LoadLicensingConfigDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadLicensingConfigDataSource_Read_BuildError(t *testing.T) {
	r := &LoadLicensingConfigDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadLicensingConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadLicensingConfigDataSource_Read_SendError exercises LoadLicensingConfigDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadLicensingConfigDataSource_Read_SendError(t *testing.T) {
	r := &LoadLicensingConfigDataSource{client: newTransportErrorClient(t)}
	m := LoadLicensingConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadLicensingConfigDataSource_Read_NotFound exercises LoadLicensingConfigDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadLicensingConfigDataSource_Read_NotFound(t *testing.T) {
	r := &LoadLicensingConfigDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadLicensingConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadLicensingConfigDataSource_Read_APIError exercises LoadLicensingConfigDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadLicensingConfigDataSource_Read_APIError(t *testing.T) {
	r := &LoadLicensingConfigDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadLicensingConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_licensing_config")
}

// TestLoadLicensingConfigDataSource_Read_APIErrorReadBody exercises LoadLicensingConfigDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadLicensingConfigDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadLicensingConfigDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadLicensingConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadLicensingConfigDataSource_Read_InvalidJSON exercises LoadLicensingConfigDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadLicensingConfigDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadLicensingConfigDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadLicensingConfigDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
