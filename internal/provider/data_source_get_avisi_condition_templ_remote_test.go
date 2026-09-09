package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAvisiConditionTemplDataSource_Read_Happy exercises GetAvisiConditionTemplDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAvisiConditionTemplDataSource_Read_Happy(t *testing.T) {
	r := &GetAvisiConditionTemplDataSource{client: newMockClientStatus(t, 200, "{\"avConditionTemplates\":[]}")}
	m := GetAvisiConditionTemplDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAvisiConditionTemplDataSource_Read_NilClient exercises GetAvisiConditionTemplDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAvisiConditionTemplDataSource_Read_NilClient(t *testing.T) {
	r := &GetAvisiConditionTemplDataSource{}
	m := GetAvisiConditionTemplDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAvisiConditionTemplDataSource_Read_BuildError exercises GetAvisiConditionTemplDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAvisiConditionTemplDataSource_Read_BuildError(t *testing.T) {
	r := &GetAvisiConditionTemplDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAvisiConditionTemplDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAvisiConditionTemplDataSource_Read_SendError exercises GetAvisiConditionTemplDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAvisiConditionTemplDataSource_Read_SendError(t *testing.T) {
	r := &GetAvisiConditionTemplDataSource{client: newTransportErrorClient(t)}
	m := GetAvisiConditionTemplDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAvisiConditionTemplDataSource_Read_InvalidJSON exercises GetAvisiConditionTemplDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAvisiConditionTemplDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAvisiConditionTemplDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAvisiConditionTemplDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
