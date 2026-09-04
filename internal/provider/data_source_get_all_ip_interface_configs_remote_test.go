package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllIpInterfaceConfigsDataSource_Read_Happy exercises GetAllIpInterfaceConfigsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetAllIpInterfaceConfigsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllIpInterfaceConfigsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetAllIpInterfaceConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllIpInterfaceConfigsDataSource_Read_NilClient exercises GetAllIpInterfaceConfigsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllIpInterfaceConfigsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllIpInterfaceConfigsDataSource{}
	m := GetAllIpInterfaceConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllIpInterfaceConfigsDataSource_Read_BuildError exercises GetAllIpInterfaceConfigsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAllIpInterfaceConfigsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllIpInterfaceConfigsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllIpInterfaceConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAllIpInterfaceConfigsDataSource_Read_SendError exercises GetAllIpInterfaceConfigsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAllIpInterfaceConfigsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllIpInterfaceConfigsDataSource{client: newTransportErrorClient(t)}
	m := GetAllIpInterfaceConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAllIpInterfaceConfigsDataSource_Read_NotFound exercises GetAllIpInterfaceConfigsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetAllIpInterfaceConfigsDataSource_Read_NotFound(t *testing.T) {
	r := &GetAllIpInterfaceConfigsDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetAllIpInterfaceConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetAllIpInterfaceConfigsDataSource_Read_APIError exercises GetAllIpInterfaceConfigsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAllIpInterfaceConfigsDataSource_Read_APIError(t *testing.T) {
	r := &GetAllIpInterfaceConfigsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetAllIpInterfaceConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_all_ip_interface_configs")
}

// TestGetAllIpInterfaceConfigsDataSource_Read_APIErrorReadBody exercises GetAllIpInterfaceConfigsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAllIpInterfaceConfigsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetAllIpInterfaceConfigsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetAllIpInterfaceConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetAllIpInterfaceConfigsDataSource_Read_InvalidJSON exercises GetAllIpInterfaceConfigsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetAllIpInterfaceConfigsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllIpInterfaceConfigsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllIpInterfaceConfigsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
