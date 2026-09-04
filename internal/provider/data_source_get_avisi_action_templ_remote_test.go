package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAvisiActionTemplDataSource_Read_Happy exercises GetAvisiActionTemplDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAvisiActionTemplDataSource_Read_Happy(t *testing.T) {
	r := &GetAvisiActionTemplDataSource{client: newMockClientStatus(t, 200, "{\"avActionTemplates\":[]}")}
	m := GetAvisiActionTemplDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAvisiActionTemplDataSource_Read_NilClient exercises GetAvisiActionTemplDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAvisiActionTemplDataSource_Read_NilClient(t *testing.T) {
	r := &GetAvisiActionTemplDataSource{}
	m := GetAvisiActionTemplDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAvisiActionTemplDataSource_Read_BuildError exercises GetAvisiActionTemplDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAvisiActionTemplDataSource_Read_BuildError(t *testing.T) {
	r := &GetAvisiActionTemplDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAvisiActionTemplDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAvisiActionTemplDataSource_Read_SendError exercises GetAvisiActionTemplDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAvisiActionTemplDataSource_Read_SendError(t *testing.T) {
	r := &GetAvisiActionTemplDataSource{client: newTransportErrorClient(t)}
	m := GetAvisiActionTemplDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAvisiActionTemplDataSource_Read_InvalidJSON exercises GetAvisiActionTemplDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAvisiActionTemplDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAvisiActionTemplDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAvisiActionTemplDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
