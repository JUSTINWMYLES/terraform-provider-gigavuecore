package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPtpPortsCountersByPortIdDataSource_Read_Happy exercises GetPtpPortsCountersByPortIdDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetPtpPortsCountersByPortIdDataSource_Read_Happy(t *testing.T) {
	r := &GetPtpPortsCountersByPortIdDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetPtpPortsCountersByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPtpPortsCountersByPortIdDataSource_Read_NilClient exercises GetPtpPortsCountersByPortIdDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPtpPortsCountersByPortIdDataSource_Read_NilClient(t *testing.T) {
	r := &GetPtpPortsCountersByPortIdDataSource{}
	m := GetPtpPortsCountersByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPtpPortsCountersByPortIdDataSource_Read_BuildError exercises GetPtpPortsCountersByPortIdDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetPtpPortsCountersByPortIdDataSource_Read_BuildError(t *testing.T) {
	r := &GetPtpPortsCountersByPortIdDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPtpPortsCountersByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetPtpPortsCountersByPortIdDataSource_Read_SendError exercises GetPtpPortsCountersByPortIdDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetPtpPortsCountersByPortIdDataSource_Read_SendError(t *testing.T) {
	r := &GetPtpPortsCountersByPortIdDataSource{client: newTransportErrorClient(t)}
	m := GetPtpPortsCountersByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetPtpPortsCountersByPortIdDataSource_Read_NotFound exercises GetPtpPortsCountersByPortIdDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetPtpPortsCountersByPortIdDataSource_Read_NotFound(t *testing.T) {
	r := &GetPtpPortsCountersByPortIdDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetPtpPortsCountersByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetPtpPortsCountersByPortIdDataSource_Read_APIError exercises GetPtpPortsCountersByPortIdDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetPtpPortsCountersByPortIdDataSource_Read_APIError(t *testing.T) {
	r := &GetPtpPortsCountersByPortIdDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetPtpPortsCountersByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_ptp_ports_counters_by_port_id")
}

// TestGetPtpPortsCountersByPortIdDataSource_Read_APIErrorReadBody exercises GetPtpPortsCountersByPortIdDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetPtpPortsCountersByPortIdDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetPtpPortsCountersByPortIdDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetPtpPortsCountersByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetPtpPortsCountersByPortIdDataSource_Read_InvalidJSON exercises GetPtpPortsCountersByPortIdDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetPtpPortsCountersByPortIdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPtpPortsCountersByPortIdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPtpPortsCountersByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
