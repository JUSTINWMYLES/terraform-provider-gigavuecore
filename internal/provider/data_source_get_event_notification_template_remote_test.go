package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEventNotificationTemplateDataSource_Read_Happy exercises GetEventNotificationTemplateDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetEventNotificationTemplateDataSource_Read_Happy(t *testing.T) {
	r := &GetEventNotificationTemplateDataSource{client: newMockClientStatus(t, 200, "{\"templates\":[]}")}
	m := GetEventNotificationTemplateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetEventNotificationTemplateDataSource_Read_NilClient exercises GetEventNotificationTemplateDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetEventNotificationTemplateDataSource_Read_NilClient(t *testing.T) {
	r := &GetEventNotificationTemplateDataSource{}
	m := GetEventNotificationTemplateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetEventNotificationTemplateDataSource_Read_BuildError exercises GetEventNotificationTemplateDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetEventNotificationTemplateDataSource_Read_BuildError(t *testing.T) {
	r := &GetEventNotificationTemplateDataSource{client: newMalformedBaseURLClient(t)}
	m := GetEventNotificationTemplateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetEventNotificationTemplateDataSource_Read_SendError exercises GetEventNotificationTemplateDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetEventNotificationTemplateDataSource_Read_SendError(t *testing.T) {
	r := &GetEventNotificationTemplateDataSource{client: newTransportErrorClient(t)}
	m := GetEventNotificationTemplateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetEventNotificationTemplateDataSource_Read_InvalidJSON exercises GetEventNotificationTemplateDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetEventNotificationTemplateDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetEventNotificationTemplateDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetEventNotificationTemplateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
