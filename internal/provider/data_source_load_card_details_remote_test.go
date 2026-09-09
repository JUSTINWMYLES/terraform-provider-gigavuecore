package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadCardDetailsDataSource_Read_Happy exercises LoadCardDetailsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestLoadCardDetailsDataSource_Read_Happy(t *testing.T) {
	r := &LoadCardDetailsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := LoadCardDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadCardDetailsDataSource_Read_NilClient exercises LoadCardDetailsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadCardDetailsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadCardDetailsDataSource{}
	m := LoadCardDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadCardDetailsDataSource_Read_BuildError exercises LoadCardDetailsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestLoadCardDetailsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadCardDetailsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadCardDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestLoadCardDetailsDataSource_Read_SendError exercises LoadCardDetailsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestLoadCardDetailsDataSource_Read_SendError(t *testing.T) {
	r := &LoadCardDetailsDataSource{client: newTransportErrorClient(t)}
	m := LoadCardDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestLoadCardDetailsDataSource_Read_NotFound exercises LoadCardDetailsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestLoadCardDetailsDataSource_Read_NotFound(t *testing.T) {
	r := &LoadCardDetailsDataSource{client: newMockClientStatus(t, 404, "")}
	m := LoadCardDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestLoadCardDetailsDataSource_Read_APIError exercises LoadCardDetailsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestLoadCardDetailsDataSource_Read_APIError(t *testing.T) {
	r := &LoadCardDetailsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := LoadCardDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_load_card_details")
}

// TestLoadCardDetailsDataSource_Read_APIErrorReadBody exercises LoadCardDetailsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestLoadCardDetailsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &LoadCardDetailsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := LoadCardDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestLoadCardDetailsDataSource_Read_InvalidJSON exercises LoadCardDetailsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestLoadCardDetailsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadCardDetailsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadCardDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
