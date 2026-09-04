package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetForeignMastersDataSource_Read_Happy exercises GetForeignMastersDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetForeignMastersDataSource_Read_Happy(t *testing.T) {
	r := &GetForeignMastersDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetForeignMastersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetForeignMastersDataSource_Read_NilClient exercises GetForeignMastersDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetForeignMastersDataSource_Read_NilClient(t *testing.T) {
	r := &GetForeignMastersDataSource{}
	m := GetForeignMastersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetForeignMastersDataSource_Read_BuildError exercises GetForeignMastersDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetForeignMastersDataSource_Read_BuildError(t *testing.T) {
	r := &GetForeignMastersDataSource{client: newMalformedBaseURLClient(t)}
	m := GetForeignMastersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetForeignMastersDataSource_Read_SendError exercises GetForeignMastersDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetForeignMastersDataSource_Read_SendError(t *testing.T) {
	r := &GetForeignMastersDataSource{client: newTransportErrorClient(t)}
	m := GetForeignMastersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetForeignMastersDataSource_Read_NotFound exercises GetForeignMastersDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetForeignMastersDataSource_Read_NotFound(t *testing.T) {
	r := &GetForeignMastersDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetForeignMastersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetForeignMastersDataSource_Read_APIError exercises GetForeignMastersDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetForeignMastersDataSource_Read_APIError(t *testing.T) {
	r := &GetForeignMastersDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetForeignMastersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_foreign_masters")
}

// TestGetForeignMastersDataSource_Read_APIErrorReadBody exercises GetForeignMastersDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetForeignMastersDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetForeignMastersDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetForeignMastersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetForeignMastersDataSource_Read_InvalidJSON exercises GetForeignMastersDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetForeignMastersDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetForeignMastersDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetForeignMastersDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
