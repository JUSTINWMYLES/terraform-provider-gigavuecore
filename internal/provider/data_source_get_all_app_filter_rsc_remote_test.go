package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllAppFilterRscDataSource_Read_Happy exercises GetAllAppFilterRscDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllAppFilterRscDataSource_Read_Happy(t *testing.T) {
	r := &GetAllAppFilterRscDataSource{client: newMockClientStatus(t, 200, "{\"appFilterRscs\":[]}")}
	m := GetAllAppFilterRscDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllAppFilterRscDataSource_Read_NilClient exercises GetAllAppFilterRscDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllAppFilterRscDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllAppFilterRscDataSource{}
	m := GetAllAppFilterRscDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllAppFilterRscDataSource_Read_BuildError exercises GetAllAppFilterRscDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllAppFilterRscDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllAppFilterRscDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllAppFilterRscDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllAppFilterRscDataSource_Read_SendError exercises GetAllAppFilterRscDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllAppFilterRscDataSource_Read_SendError(t *testing.T) {
	r := &GetAllAppFilterRscDataSource{client: newTransportErrorClient(t)}
	m := GetAllAppFilterRscDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllAppFilterRscDataSource_Read_InvalidJSON exercises GetAllAppFilterRscDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllAppFilterRscDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllAppFilterRscDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllAppFilterRscDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
