package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListDeviceConfigFilesDataSource_Read_Happy exercises ListDeviceConfigFilesDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestListDeviceConfigFilesDataSource_Read_Happy(t *testing.T) {
	r := &ListDeviceConfigFilesDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := ListDeviceConfigFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestListDeviceConfigFilesDataSource_Read_NilClient exercises ListDeviceConfigFilesDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListDeviceConfigFilesDataSource_Read_NilClient(t *testing.T) {
	r := &ListDeviceConfigFilesDataSource{}
	m := ListDeviceConfigFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestListDeviceConfigFilesDataSource_Read_BuildError exercises ListDeviceConfigFilesDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestListDeviceConfigFilesDataSource_Read_BuildError(t *testing.T) {
	r := &ListDeviceConfigFilesDataSource{client: newMalformedBaseURLClient(t)}
	m := ListDeviceConfigFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestListDeviceConfigFilesDataSource_Read_SendError exercises ListDeviceConfigFilesDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestListDeviceConfigFilesDataSource_Read_SendError(t *testing.T) {
	r := &ListDeviceConfigFilesDataSource{client: newTransportErrorClient(t)}
	m := ListDeviceConfigFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestListDeviceConfigFilesDataSource_Read_NotFound exercises ListDeviceConfigFilesDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestListDeviceConfigFilesDataSource_Read_NotFound(t *testing.T) {
	r := &ListDeviceConfigFilesDataSource{client: newMockClientStatus(t, 404, "")}
	m := ListDeviceConfigFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestListDeviceConfigFilesDataSource_Read_APIError exercises ListDeviceConfigFilesDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestListDeviceConfigFilesDataSource_Read_APIError(t *testing.T) {
	r := &ListDeviceConfigFilesDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ListDeviceConfigFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_list_device_config_files")
}

// TestListDeviceConfigFilesDataSource_Read_APIErrorReadBody exercises ListDeviceConfigFilesDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestListDeviceConfigFilesDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &ListDeviceConfigFilesDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := ListDeviceConfigFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestListDeviceConfigFilesDataSource_Read_InvalidJSON exercises ListDeviceConfigFilesDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestListDeviceConfigFilesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &ListDeviceConfigFilesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := ListDeviceConfigFilesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
