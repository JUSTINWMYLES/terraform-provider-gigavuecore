package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetImageServerRepoByVersionDataSource_Read_Happy exercises GetImageServerRepoByVersionDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetImageServerRepoByVersionDataSource_Read_Happy(t *testing.T) {
	r := &GetImageServerRepoByVersionDataSource{client: newMockClientStatus(t, 200, "{\"imageRepoImagesByVersion\":[]}")}
	m := GetImageServerRepoByVersionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetImageServerRepoByVersionDataSource_Read_NilClient exercises GetImageServerRepoByVersionDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetImageServerRepoByVersionDataSource_Read_NilClient(t *testing.T) {
	r := &GetImageServerRepoByVersionDataSource{}
	m := GetImageServerRepoByVersionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetImageServerRepoByVersionDataSource_Read_BuildError exercises GetImageServerRepoByVersionDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetImageServerRepoByVersionDataSource_Read_BuildError(t *testing.T) {
	r := &GetImageServerRepoByVersionDataSource{client: newMalformedBaseURLClient(t)}
	m := GetImageServerRepoByVersionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetImageServerRepoByVersionDataSource_Read_SendError exercises GetImageServerRepoByVersionDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetImageServerRepoByVersionDataSource_Read_SendError(t *testing.T) {
	r := &GetImageServerRepoByVersionDataSource{client: newTransportErrorClient(t)}
	m := GetImageServerRepoByVersionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetImageServerRepoByVersionDataSource_Read_InvalidJSON exercises GetImageServerRepoByVersionDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetImageServerRepoByVersionDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetImageServerRepoByVersionDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetImageServerRepoByVersionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
