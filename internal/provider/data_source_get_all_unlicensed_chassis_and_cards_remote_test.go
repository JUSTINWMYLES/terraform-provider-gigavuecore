package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllUnlicensedChassisAndCardsDataSource_Read_Happy exercises GetAllUnlicensedChassisAndCardsDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetAllUnlicensedChassisAndCardsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllUnlicensedChassisAndCardsDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetAllUnlicensedChassisAndCardsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllUnlicensedChassisAndCardsDataSource_Read_NilClient exercises GetAllUnlicensedChassisAndCardsDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllUnlicensedChassisAndCardsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllUnlicensedChassisAndCardsDataSource{}
	m := GetAllUnlicensedChassisAndCardsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllUnlicensedChassisAndCardsDataSource_Read_BuildError exercises GetAllUnlicensedChassisAndCardsDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAllUnlicensedChassisAndCardsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllUnlicensedChassisAndCardsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllUnlicensedChassisAndCardsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAllUnlicensedChassisAndCardsDataSource_Read_SendError exercises GetAllUnlicensedChassisAndCardsDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAllUnlicensedChassisAndCardsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllUnlicensedChassisAndCardsDataSource{client: newTransportErrorClient(t)}
	m := GetAllUnlicensedChassisAndCardsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAllUnlicensedChassisAndCardsDataSource_Read_NotFound exercises GetAllUnlicensedChassisAndCardsDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetAllUnlicensedChassisAndCardsDataSource_Read_NotFound(t *testing.T) {
	r := &GetAllUnlicensedChassisAndCardsDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetAllUnlicensedChassisAndCardsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetAllUnlicensedChassisAndCardsDataSource_Read_APIError exercises GetAllUnlicensedChassisAndCardsDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAllUnlicensedChassisAndCardsDataSource_Read_APIError(t *testing.T) {
	r := &GetAllUnlicensedChassisAndCardsDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetAllUnlicensedChassisAndCardsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_all_unlicensed_chassis_and_cards")
}

// TestGetAllUnlicensedChassisAndCardsDataSource_Read_APIErrorReadBody exercises GetAllUnlicensedChassisAndCardsDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAllUnlicensedChassisAndCardsDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetAllUnlicensedChassisAndCardsDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetAllUnlicensedChassisAndCardsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetAllUnlicensedChassisAndCardsDataSource_Read_InvalidJSON exercises GetAllUnlicensedChassisAndCardsDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetAllUnlicensedChassisAndCardsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllUnlicensedChassisAndCardsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllUnlicensedChassisAndCardsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
