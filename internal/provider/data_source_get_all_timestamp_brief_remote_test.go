package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTimestampBriefDataSource_Read_Happy exercises GetAllTimestampBriefDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllTimestampBriefDataSource_Read_Happy(t *testing.T) {
	r := &GetAllTimestampBriefDataSource{client: newMockClientStatus(t, 200, "{\"timestamps\":[]}")}
	m := GetAllTimestampBriefDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllTimestampBriefDataSource_Read_NilClient exercises GetAllTimestampBriefDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllTimestampBriefDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllTimestampBriefDataSource{}
	m := GetAllTimestampBriefDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllTimestampBriefDataSource_Read_BuildError exercises GetAllTimestampBriefDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllTimestampBriefDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllTimestampBriefDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllTimestampBriefDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllTimestampBriefDataSource_Read_SendError exercises GetAllTimestampBriefDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllTimestampBriefDataSource_Read_SendError(t *testing.T) {
	r := &GetAllTimestampBriefDataSource{client: newTransportErrorClient(t)}
	m := GetAllTimestampBriefDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllTimestampBriefDataSource_Read_InvalidJSON exercises GetAllTimestampBriefDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllTimestampBriefDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllTimestampBriefDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllTimestampBriefDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
