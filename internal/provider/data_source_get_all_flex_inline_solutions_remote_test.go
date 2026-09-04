package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllFlexInlineSolutionsDataSource_Read_Happy exercises GetAllFlexInlineSolutionsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllFlexInlineSolutionsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllFlexInlineSolutionsDataSource{client: newMockClientStatus(t, 200, "{\"flexInlineSolutions\":[]}")}
	m := GetAllFlexInlineSolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllFlexInlineSolutionsDataSource_Read_NilClient exercises GetAllFlexInlineSolutionsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllFlexInlineSolutionsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllFlexInlineSolutionsDataSource{}
	m := GetAllFlexInlineSolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllFlexInlineSolutionsDataSource_Read_BuildError exercises GetAllFlexInlineSolutionsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllFlexInlineSolutionsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllFlexInlineSolutionsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllFlexInlineSolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllFlexInlineSolutionsDataSource_Read_SendError exercises GetAllFlexInlineSolutionsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllFlexInlineSolutionsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllFlexInlineSolutionsDataSource{client: newTransportErrorClient(t)}
	m := GetAllFlexInlineSolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllFlexInlineSolutionsDataSource_Read_InvalidJSON exercises GetAllFlexInlineSolutionsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllFlexInlineSolutionsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllFlexInlineSolutionsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllFlexInlineSolutionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
