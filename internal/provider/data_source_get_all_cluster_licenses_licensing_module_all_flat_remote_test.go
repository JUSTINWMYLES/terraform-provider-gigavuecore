package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_Happy exercises GetAllClusterLicensesLicensingModuleAllFlatDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_Happy(t *testing.T) {
	r := &GetAllClusterLicensesLicensingModuleAllFlatDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetAllClusterLicensesLicensingModuleAllFlatDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_NilClient exercises GetAllClusterLicensesLicensingModuleAllFlatDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllClusterLicensesLicensingModuleAllFlatDataSource{}
	m := GetAllClusterLicensesLicensingModuleAllFlatDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_BuildError exercises GetAllClusterLicensesLicensingModuleAllFlatDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllClusterLicensesLicensingModuleAllFlatDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllClusterLicensesLicensingModuleAllFlatDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_SendError exercises GetAllClusterLicensesLicensingModuleAllFlatDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_SendError(t *testing.T) {
	r := &GetAllClusterLicensesLicensingModuleAllFlatDataSource{client: newTransportErrorClient(t)}
	m := GetAllClusterLicensesLicensingModuleAllFlatDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_NotFound exercises GetAllClusterLicensesLicensingModuleAllFlatDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_NotFound(t *testing.T) {
	r := &GetAllClusterLicensesLicensingModuleAllFlatDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetAllClusterLicensesLicensingModuleAllFlatDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_APIError exercises GetAllClusterLicensesLicensingModuleAllFlatDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_APIError(t *testing.T) {
	r := &GetAllClusterLicensesLicensingModuleAllFlatDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetAllClusterLicensesLicensingModuleAllFlatDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat")
}

// TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_APIErrorReadBody exercises GetAllClusterLicensesLicensingModuleAllFlatDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetAllClusterLicensesLicensingModuleAllFlatDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetAllClusterLicensesLicensingModuleAllFlatDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_InvalidJSON exercises GetAllClusterLicensesLicensingModuleAllFlatDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetAllClusterLicensesLicensingModuleAllFlatDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllClusterLicensesLicensingModuleAllFlatDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllClusterLicensesLicensingModuleAllFlatDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
