package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlexInlineMapConfigStatusDataSource_Read_Happy exercises GetFlexInlineMapConfigStatusDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetFlexInlineMapConfigStatusDataSource_Read_Happy(t *testing.T) {
	r := &GetFlexInlineMapConfigStatusDataSource{client: newMockClientStatus(t, 200, "{\"flexInlineDeployStatus\":[]}")}
	m := GetFlexInlineMapConfigStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFlexInlineMapConfigStatusDataSource_Read_NilClient exercises GetFlexInlineMapConfigStatusDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFlexInlineMapConfigStatusDataSource_Read_NilClient(t *testing.T) {
	r := &GetFlexInlineMapConfigStatusDataSource{}
	m := GetFlexInlineMapConfigStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFlexInlineMapConfigStatusDataSource_Read_BuildError exercises GetFlexInlineMapConfigStatusDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetFlexInlineMapConfigStatusDataSource_Read_BuildError(t *testing.T) {
	r := &GetFlexInlineMapConfigStatusDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFlexInlineMapConfigStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetFlexInlineMapConfigStatusDataSource_Read_SendError exercises GetFlexInlineMapConfigStatusDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetFlexInlineMapConfigStatusDataSource_Read_SendError(t *testing.T) {
	r := &GetFlexInlineMapConfigStatusDataSource{client: newTransportErrorClient(t)}
	m := GetFlexInlineMapConfigStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetFlexInlineMapConfigStatusDataSource_Read_InvalidJSON exercises GetFlexInlineMapConfigStatusDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetFlexInlineMapConfigStatusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFlexInlineMapConfigStatusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFlexInlineMapConfigStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
