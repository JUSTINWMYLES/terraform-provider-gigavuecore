package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadTopologyVizEndPointsDataSource_Read_Happy exercises LoadTopologyVizEndPointsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadTopologyVizEndPointsDataSource_Read_Happy(t *testing.T) {
	r := &LoadTopologyVizEndPointsDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := LoadTopologyVizEndPointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadTopologyVizEndPointsDataSource_Read_NilClient exercises LoadTopologyVizEndPointsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadTopologyVizEndPointsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadTopologyVizEndPointsDataSource{}
	m := LoadTopologyVizEndPointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadTopologyVizEndPointsDataSource_Read_BuildError exercises LoadTopologyVizEndPointsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadTopologyVizEndPointsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadTopologyVizEndPointsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadTopologyVizEndPointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadTopologyVizEndPointsDataSource_Read_SendError exercises LoadTopologyVizEndPointsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadTopologyVizEndPointsDataSource_Read_SendError(t *testing.T) {
	r := &LoadTopologyVizEndPointsDataSource{client: newTransportErrorClient(t)}
	m := LoadTopologyVizEndPointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadTopologyVizEndPointsDataSource_Read_InvalidJSON exercises LoadTopologyVizEndPointsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadTopologyVizEndPointsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadTopologyVizEndPointsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadTopologyVizEndPointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
