package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource_Read_Happy exercises GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource_Read_Happy(t *testing.T) {
	r := &GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource_Read_NilClient exercises GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource_Read_NilClient(t *testing.T) {
	r := &GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource{}
	m := GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource_Read_BuildError exercises GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource_Read_BuildError(t *testing.T) {
	r := &GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource{client: newMalformedBaseURLClient(t)}
	m := GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource_Read_SendError exercises GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource_Read_SendError(t *testing.T) {
	r := &GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource{client: newTransportErrorClient(t)}
	m := GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource_Read_InvalidJSON exercises GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
