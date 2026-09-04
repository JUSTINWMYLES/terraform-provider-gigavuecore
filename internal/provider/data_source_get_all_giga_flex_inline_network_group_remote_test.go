package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllGigaFlexInlineNetworkGroupDataSource_Read_Happy exercises GetAllGigaFlexInlineNetworkGroupDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllGigaFlexInlineNetworkGroupDataSource_Read_Happy(t *testing.T) {
	r := &GetAllGigaFlexInlineNetworkGroupDataSource{client: newMockClientStatus(t, 200, "{\"gigaFlexInlineNetworkGroups\":[]}")}
	m := GetAllGigaFlexInlineNetworkGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllGigaFlexInlineNetworkGroupDataSource_Read_NilClient exercises GetAllGigaFlexInlineNetworkGroupDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllGigaFlexInlineNetworkGroupDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllGigaFlexInlineNetworkGroupDataSource{}
	m := GetAllGigaFlexInlineNetworkGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllGigaFlexInlineNetworkGroupDataSource_Read_BuildError exercises GetAllGigaFlexInlineNetworkGroupDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllGigaFlexInlineNetworkGroupDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllGigaFlexInlineNetworkGroupDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllGigaFlexInlineNetworkGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllGigaFlexInlineNetworkGroupDataSource_Read_SendError exercises GetAllGigaFlexInlineNetworkGroupDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllGigaFlexInlineNetworkGroupDataSource_Read_SendError(t *testing.T) {
	r := &GetAllGigaFlexInlineNetworkGroupDataSource{client: newTransportErrorClient(t)}
	m := GetAllGigaFlexInlineNetworkGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllGigaFlexInlineNetworkGroupDataSource_Read_InvalidJSON exercises GetAllGigaFlexInlineNetworkGroupDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllGigaFlexInlineNetworkGroupDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllGigaFlexInlineNetworkGroupDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllGigaFlexInlineNetworkGroupDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
