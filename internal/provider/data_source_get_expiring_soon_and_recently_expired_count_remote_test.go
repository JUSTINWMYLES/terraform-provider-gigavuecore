package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_Happy exercises GetExpiringSoonAndRecentlyExpiredCountDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_Happy(t *testing.T) {
	r := &GetExpiringSoonAndRecentlyExpiredCountDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetExpiringSoonAndRecentlyExpiredCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_NilClient exercises GetExpiringSoonAndRecentlyExpiredCountDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_NilClient(t *testing.T) {
	r := &GetExpiringSoonAndRecentlyExpiredCountDataSource{}
	m := GetExpiringSoonAndRecentlyExpiredCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_BuildError exercises GetExpiringSoonAndRecentlyExpiredCountDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_BuildError(t *testing.T) {
	r := &GetExpiringSoonAndRecentlyExpiredCountDataSource{client: newMalformedBaseURLClient(t)}
	m := GetExpiringSoonAndRecentlyExpiredCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_SendError exercises GetExpiringSoonAndRecentlyExpiredCountDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_SendError(t *testing.T) {
	r := &GetExpiringSoonAndRecentlyExpiredCountDataSource{client: newTransportErrorClient(t)}
	m := GetExpiringSoonAndRecentlyExpiredCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_NotFound exercises GetExpiringSoonAndRecentlyExpiredCountDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_NotFound(t *testing.T) {
	r := &GetExpiringSoonAndRecentlyExpiredCountDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetExpiringSoonAndRecentlyExpiredCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_APIError exercises GetExpiringSoonAndRecentlyExpiredCountDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_APIError(t *testing.T) {
	r := &GetExpiringSoonAndRecentlyExpiredCountDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetExpiringSoonAndRecentlyExpiredCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_get_expiring_soon_and_recently_expired_count")
}

// TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_APIErrorReadBody exercises GetExpiringSoonAndRecentlyExpiredCountDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetExpiringSoonAndRecentlyExpiredCountDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := GetExpiringSoonAndRecentlyExpiredCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_InvalidJSON exercises GetExpiringSoonAndRecentlyExpiredCountDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetExpiringSoonAndRecentlyExpiredCountDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetExpiringSoonAndRecentlyExpiredCountDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetExpiringSoonAndRecentlyExpiredCountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
