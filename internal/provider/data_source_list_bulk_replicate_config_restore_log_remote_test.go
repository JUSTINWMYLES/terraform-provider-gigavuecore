package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListBulkReplicateConfigRestoreLogDataSource_Read_Happy exercises ListBulkReplicateConfigRestoreLogDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestListBulkReplicateConfigRestoreLogDataSource_Read_Happy(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := ListBulkReplicateConfigRestoreLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestListBulkReplicateConfigRestoreLogDataSource_Read_NilClient exercises ListBulkReplicateConfigRestoreLogDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestListBulkReplicateConfigRestoreLogDataSource_Read_NilClient(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogDataSource{}
	m := ListBulkReplicateConfigRestoreLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestListBulkReplicateConfigRestoreLogDataSource_Read_BuildError exercises ListBulkReplicateConfigRestoreLogDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestListBulkReplicateConfigRestoreLogDataSource_Read_BuildError(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogDataSource{client: newMalformedBaseURLClient(t)}
	m := ListBulkReplicateConfigRestoreLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestListBulkReplicateConfigRestoreLogDataSource_Read_SendError exercises ListBulkReplicateConfigRestoreLogDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestListBulkReplicateConfigRestoreLogDataSource_Read_SendError(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogDataSource{client: newTransportErrorClient(t)}
	m := ListBulkReplicateConfigRestoreLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestListBulkReplicateConfigRestoreLogDataSource_Read_NotFound exercises ListBulkReplicateConfigRestoreLogDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestListBulkReplicateConfigRestoreLogDataSource_Read_NotFound(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogDataSource{client: newMockClientStatus(t, 404, "")}
	m := ListBulkReplicateConfigRestoreLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestListBulkReplicateConfigRestoreLogDataSource_Read_APIError exercises ListBulkReplicateConfigRestoreLogDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestListBulkReplicateConfigRestoreLogDataSource_Read_APIError(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogDataSource{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ListBulkReplicateConfigRestoreLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading gigavuecore_list_bulk_replicate_config_restore_log")
}

// TestListBulkReplicateConfigRestoreLogDataSource_Read_APIErrorReadBody exercises ListBulkReplicateConfigRestoreLogDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestListBulkReplicateConfigRestoreLogDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogDataSource{client: newMockClientReadErrorBody(t, 501)}
	m := ListBulkReplicateConfigRestoreLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestListBulkReplicateConfigRestoreLogDataSource_Read_InvalidJSON exercises ListBulkReplicateConfigRestoreLogDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestListBulkReplicateConfigRestoreLogDataSource_Read_InvalidJSON(t *testing.T) {
	r := &ListBulkReplicateConfigRestoreLogDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := ListBulkReplicateConfigRestoreLogDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
