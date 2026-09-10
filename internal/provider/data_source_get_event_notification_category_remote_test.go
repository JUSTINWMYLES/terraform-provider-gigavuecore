package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEventNotificationCategoryDataSource_Read_Happy exercises GetEventNotificationCategoryDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetEventNotificationCategoryDataSource_Read_Happy(t *testing.T) {
	r := &GetEventNotificationCategoryDataSource{client: newMockClientStatus(t, 200, "{\"eventScopes\":[]}")}
	m := GetEventNotificationCategoryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetEventNotificationCategoryDataSource_Read_NilClient exercises GetEventNotificationCategoryDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetEventNotificationCategoryDataSource_Read_NilClient(t *testing.T) {
	r := &GetEventNotificationCategoryDataSource{}
	m := GetEventNotificationCategoryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetEventNotificationCategoryDataSource_Read_BuildError exercises GetEventNotificationCategoryDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetEventNotificationCategoryDataSource_Read_BuildError(t *testing.T) {
	r := &GetEventNotificationCategoryDataSource{client: newMalformedBaseURLClient(t)}
	m := GetEventNotificationCategoryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetEventNotificationCategoryDataSource_Read_SendError exercises GetEventNotificationCategoryDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetEventNotificationCategoryDataSource_Read_SendError(t *testing.T) {
	r := &GetEventNotificationCategoryDataSource{client: newTransportErrorClient(t)}
	m := GetEventNotificationCategoryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetEventNotificationCategoryDataSource_Read_InvalidJSON exercises GetEventNotificationCategoryDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetEventNotificationCategoryDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetEventNotificationCategoryDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetEventNotificationCategoryDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
