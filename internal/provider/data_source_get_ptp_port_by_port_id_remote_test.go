package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPtpPortByPortIdDataSource_Read_Happy exercises GetPtpPortByPortIdDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetPtpPortByPortIdDataSource_Read_Happy(t *testing.T) {
	r := &GetPtpPortByPortIdDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetPtpPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPtpPortByPortIdDataSource_Read_NilClient exercises GetPtpPortByPortIdDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPtpPortByPortIdDataSource_Read_NilClient(t *testing.T) {
	r := &GetPtpPortByPortIdDataSource{}
	m := GetPtpPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPtpPortByPortIdDataSource_Read_BuildError exercises GetPtpPortByPortIdDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetPtpPortByPortIdDataSource_Read_BuildError(t *testing.T) {
	r := &GetPtpPortByPortIdDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPtpPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetPtpPortByPortIdDataSource_Read_SendError exercises GetPtpPortByPortIdDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetPtpPortByPortIdDataSource_Read_SendError(t *testing.T) {
	r := &GetPtpPortByPortIdDataSource{client: newTransportErrorClient(t)}
	m := GetPtpPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetPtpPortByPortIdDataSource_Read_NotFound exercises GetPtpPortByPortIdDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetPtpPortByPortIdDataSource_Read_NotFound(t *testing.T) {
	r := &GetPtpPortByPortIdDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetPtpPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetPtpPortByPortIdDataSource_Read_APIError exercises GetPtpPortByPortIdDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetPtpPortByPortIdDataSource_Read_APIError(t *testing.T) {
	r := &GetPtpPortByPortIdDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetPtpPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_ptp_port_by_port_id")
}

// TestGetPtpPortByPortIdDataSource_Read_APIErrorReadBody exercises GetPtpPortByPortIdDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetPtpPortByPortIdDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetPtpPortByPortIdDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetPtpPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetPtpPortByPortIdDataSource_Read_InvalidJSON exercises GetPtpPortByPortIdDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetPtpPortByPortIdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPtpPortByPortIdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPtpPortByPortIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
