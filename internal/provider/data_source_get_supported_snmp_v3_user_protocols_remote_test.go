package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSupportedSnmpV3UserProtocolsDataSource_Read_Happy exercises GetSupportedSnmpV3UserProtocolsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSupportedSnmpV3UserProtocolsDataSource_Read_Happy(t *testing.T) {
	r := &GetSupportedSnmpV3UserProtocolsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSupportedSnmpV3UserProtocolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSupportedSnmpV3UserProtocolsDataSource_Read_NilClient exercises GetSupportedSnmpV3UserProtocolsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSupportedSnmpV3UserProtocolsDataSource_Read_NilClient(t *testing.T) {
	r := &GetSupportedSnmpV3UserProtocolsDataSource{}
	m := GetSupportedSnmpV3UserProtocolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSupportedSnmpV3UserProtocolsDataSource_Read_BuildError exercises GetSupportedSnmpV3UserProtocolsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSupportedSnmpV3UserProtocolsDataSource_Read_BuildError(t *testing.T) {
	r := &GetSupportedSnmpV3UserProtocolsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSupportedSnmpV3UserProtocolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSupportedSnmpV3UserProtocolsDataSource_Read_SendError exercises GetSupportedSnmpV3UserProtocolsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSupportedSnmpV3UserProtocolsDataSource_Read_SendError(t *testing.T) {
	r := &GetSupportedSnmpV3UserProtocolsDataSource{client: newTransportErrorClient(t)}
	m := GetSupportedSnmpV3UserProtocolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSupportedSnmpV3UserProtocolsDataSource_Read_NotFound exercises GetSupportedSnmpV3UserProtocolsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSupportedSnmpV3UserProtocolsDataSource_Read_NotFound(t *testing.T) {
	r := &GetSupportedSnmpV3UserProtocolsDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSupportedSnmpV3UserProtocolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSupportedSnmpV3UserProtocolsDataSource_Read_APIError exercises GetSupportedSnmpV3UserProtocolsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSupportedSnmpV3UserProtocolsDataSource_Read_APIError(t *testing.T) {
	r := &GetSupportedSnmpV3UserProtocolsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetSupportedSnmpV3UserProtocolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_supported_snmp_v3_user_protocols")
}

// TestGetSupportedSnmpV3UserProtocolsDataSource_Read_APIErrorReadBody exercises GetSupportedSnmpV3UserProtocolsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSupportedSnmpV3UserProtocolsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSupportedSnmpV3UserProtocolsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetSupportedSnmpV3UserProtocolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSupportedSnmpV3UserProtocolsDataSource_Read_InvalidJSON exercises GetSupportedSnmpV3UserProtocolsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSupportedSnmpV3UserProtocolsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSupportedSnmpV3UserProtocolsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSupportedSnmpV3UserProtocolsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
