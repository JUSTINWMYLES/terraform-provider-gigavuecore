package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEventByIdDataSource_Read_Happy exercises GetEventByIdDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetEventByIdDataSource_Read_Happy(t *testing.T) {
	r := &GetEventByIdDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetEventByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetEventByIdDataSource_Read_NilClient exercises GetEventByIdDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetEventByIdDataSource_Read_NilClient(t *testing.T) {
	r := &GetEventByIdDataSource{}
	m := GetEventByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetEventByIdDataSource_Read_BuildError exercises GetEventByIdDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetEventByIdDataSource_Read_BuildError(t *testing.T) {
	r := &GetEventByIdDataSource{client: newMalformedBaseURLClient(t)}
	m := GetEventByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetEventByIdDataSource_Read_SendError exercises GetEventByIdDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetEventByIdDataSource_Read_SendError(t *testing.T) {
	r := &GetEventByIdDataSource{client: newTransportErrorClient(t)}
	m := GetEventByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetEventByIdDataSource_Read_NotFound exercises GetEventByIdDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetEventByIdDataSource_Read_NotFound(t *testing.T) {
	r := &GetEventByIdDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetEventByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetEventByIdDataSource_Read_APIError exercises GetEventByIdDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetEventByIdDataSource_Read_APIError(t *testing.T) {
	r := &GetEventByIdDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetEventByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_event_by_id")
}

// TestGetEventByIdDataSource_Read_APIErrorReadBody exercises GetEventByIdDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetEventByIdDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetEventByIdDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetEventByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetEventByIdDataSource_Read_InvalidJSON exercises GetEventByIdDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetEventByIdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetEventByIdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetEventByIdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
