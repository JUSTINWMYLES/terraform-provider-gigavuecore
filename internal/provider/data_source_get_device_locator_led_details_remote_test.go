package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetDeviceLocatorLedDetailsDataSource_Read_Happy exercises GetDeviceLocatorLedDetailsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetDeviceLocatorLedDetailsDataSource_Read_Happy(t *testing.T) {
	r := &GetDeviceLocatorLedDetailsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetDeviceLocatorLedDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetDeviceLocatorLedDetailsDataSource_Read_NilClient exercises GetDeviceLocatorLedDetailsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetDeviceLocatorLedDetailsDataSource_Read_NilClient(t *testing.T) {
	r := &GetDeviceLocatorLedDetailsDataSource{}
	m := GetDeviceLocatorLedDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetDeviceLocatorLedDetailsDataSource_Read_BuildError exercises GetDeviceLocatorLedDetailsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetDeviceLocatorLedDetailsDataSource_Read_BuildError(t *testing.T) {
	r := &GetDeviceLocatorLedDetailsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetDeviceLocatorLedDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetDeviceLocatorLedDetailsDataSource_Read_SendError exercises GetDeviceLocatorLedDetailsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetDeviceLocatorLedDetailsDataSource_Read_SendError(t *testing.T) {
	r := &GetDeviceLocatorLedDetailsDataSource{client: newTransportErrorClient(t)}
	m := GetDeviceLocatorLedDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetDeviceLocatorLedDetailsDataSource_Read_NotFound exercises GetDeviceLocatorLedDetailsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetDeviceLocatorLedDetailsDataSource_Read_NotFound(t *testing.T) {
	r := &GetDeviceLocatorLedDetailsDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetDeviceLocatorLedDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetDeviceLocatorLedDetailsDataSource_Read_APIError exercises GetDeviceLocatorLedDetailsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetDeviceLocatorLedDetailsDataSource_Read_APIError(t *testing.T) {
	r := &GetDeviceLocatorLedDetailsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetDeviceLocatorLedDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_device_locator_led_details")
}

// TestGetDeviceLocatorLedDetailsDataSource_Read_APIErrorReadBody exercises GetDeviceLocatorLedDetailsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetDeviceLocatorLedDetailsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetDeviceLocatorLedDetailsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetDeviceLocatorLedDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetDeviceLocatorLedDetailsDataSource_Read_InvalidJSON exercises GetDeviceLocatorLedDetailsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetDeviceLocatorLedDetailsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetDeviceLocatorLedDetailsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetDeviceLocatorLedDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
