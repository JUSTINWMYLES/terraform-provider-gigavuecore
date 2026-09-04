package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllAlarmSuppressionMetadataDataSource_Read_Happy exercises GetAllAlarmSuppressionMetadataDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllAlarmSuppressionMetadataDataSource_Read_Happy(t *testing.T) {
	r := &GetAllAlarmSuppressionMetadataDataSource{client: newMockClientStatus(t, 200, "{\"suppressedEntities\":[]}")}
	m := GetAllAlarmSuppressionMetadataDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllAlarmSuppressionMetadataDataSource_Read_NilClient exercises GetAllAlarmSuppressionMetadataDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllAlarmSuppressionMetadataDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllAlarmSuppressionMetadataDataSource{}
	m := GetAllAlarmSuppressionMetadataDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllAlarmSuppressionMetadataDataSource_Read_BuildError exercises GetAllAlarmSuppressionMetadataDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllAlarmSuppressionMetadataDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllAlarmSuppressionMetadataDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllAlarmSuppressionMetadataDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllAlarmSuppressionMetadataDataSource_Read_SendError exercises GetAllAlarmSuppressionMetadataDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllAlarmSuppressionMetadataDataSource_Read_SendError(t *testing.T) {
	r := &GetAllAlarmSuppressionMetadataDataSource{client: newTransportErrorClient(t)}
	m := GetAllAlarmSuppressionMetadataDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllAlarmSuppressionMetadataDataSource_Read_InvalidJSON exercises GetAllAlarmSuppressionMetadataDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllAlarmSuppressionMetadataDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllAlarmSuppressionMetadataDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllAlarmSuppressionMetadataDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
