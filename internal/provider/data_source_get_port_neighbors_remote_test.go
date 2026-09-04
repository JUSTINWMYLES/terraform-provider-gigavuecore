package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPortNeighborsDataSource_Read_Happy exercises GetPortNeighborsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetPortNeighborsDataSource_Read_Happy(t *testing.T) {
	r := &GetPortNeighborsDataSource{client: newMockClientStatus(t, 200, "{\"portNeighbors\":[]}")}
	m := GetPortNeighborsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPortNeighborsDataSource_Read_NilClient exercises GetPortNeighborsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPortNeighborsDataSource_Read_NilClient(t *testing.T) {
	r := &GetPortNeighborsDataSource{}
	m := GetPortNeighborsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPortNeighborsDataSource_Read_BuildError exercises GetPortNeighborsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetPortNeighborsDataSource_Read_BuildError(t *testing.T) {
	r := &GetPortNeighborsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetPortNeighborsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetPortNeighborsDataSource_Read_SendError exercises GetPortNeighborsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetPortNeighborsDataSource_Read_SendError(t *testing.T) {
	r := &GetPortNeighborsDataSource{client: newTransportErrorClient(t)}
	m := GetPortNeighborsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetPortNeighborsDataSource_Read_InvalidJSON exercises GetPortNeighborsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetPortNeighborsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetPortNeighborsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetPortNeighborsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
