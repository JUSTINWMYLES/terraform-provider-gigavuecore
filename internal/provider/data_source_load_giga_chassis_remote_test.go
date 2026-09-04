package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadGigaChassisDataSource_Read_Happy exercises LoadGigaChassisDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadGigaChassisDataSource_Read_Happy(t *testing.T) {
	r := &LoadGigaChassisDataSource{client: newMockClientStatus(t, 200, "{\"chassisList\":[]}")}
	m := LoadGigaChassisDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadGigaChassisDataSource_Read_NilClient exercises LoadGigaChassisDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadGigaChassisDataSource_Read_NilClient(t *testing.T) {
	r := &LoadGigaChassisDataSource{}
	m := LoadGigaChassisDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadGigaChassisDataSource_Read_BuildError exercises LoadGigaChassisDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadGigaChassisDataSource_Read_BuildError(t *testing.T) {
	r := &LoadGigaChassisDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadGigaChassisDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadGigaChassisDataSource_Read_SendError exercises LoadGigaChassisDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadGigaChassisDataSource_Read_SendError(t *testing.T) {
	r := &LoadGigaChassisDataSource{client: newTransportErrorClient(t)}
	m := LoadGigaChassisDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadGigaChassisDataSource_Read_InvalidJSON exercises LoadGigaChassisDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadGigaChassisDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadGigaChassisDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadGigaChassisDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
