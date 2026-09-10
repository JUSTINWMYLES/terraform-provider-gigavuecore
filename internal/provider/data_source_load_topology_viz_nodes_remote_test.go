package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadTopologyVizNodesDataSource_Read_Happy exercises LoadTopologyVizNodesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadTopologyVizNodesDataSource_Read_Happy(t *testing.T) {
	r := &LoadTopologyVizNodesDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := LoadTopologyVizNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadTopologyVizNodesDataSource_Read_NilClient exercises LoadTopologyVizNodesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadTopologyVizNodesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadTopologyVizNodesDataSource{}
	m := LoadTopologyVizNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadTopologyVizNodesDataSource_Read_BuildError exercises LoadTopologyVizNodesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadTopologyVizNodesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadTopologyVizNodesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadTopologyVizNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadTopologyVizNodesDataSource_Read_SendError exercises LoadTopologyVizNodesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadTopologyVizNodesDataSource_Read_SendError(t *testing.T) {
	r := &LoadTopologyVizNodesDataSource{client: newTransportErrorClient(t)}
	m := LoadTopologyVizNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadTopologyVizNodesDataSource_Read_InvalidJSON exercises LoadTopologyVizNodesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadTopologyVizNodesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadTopologyVizNodesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadTopologyVizNodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
