package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestPortPacketThresholdDataSource_Read_Happy exercises PortPacketThresholdDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestPortPacketThresholdDataSource_Read_Happy(t *testing.T) {
	r := &PortPacketThresholdDataSource{client: newMockClientStatus(t, 200, "{\"device\":[]}")}
	m := PortPacketThresholdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPortPacketThresholdDataSource_Read_NilClient exercises PortPacketThresholdDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPortPacketThresholdDataSource_Read_NilClient(t *testing.T) {
	r := &PortPacketThresholdDataSource{}
	m := PortPacketThresholdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPortPacketThresholdDataSource_Read_BuildError exercises PortPacketThresholdDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestPortPacketThresholdDataSource_Read_BuildError(t *testing.T) {
	r := &PortPacketThresholdDataSource{client: newMalformedBaseURLClient(t)}
	m := PortPacketThresholdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestPortPacketThresholdDataSource_Read_SendError exercises PortPacketThresholdDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestPortPacketThresholdDataSource_Read_SendError(t *testing.T) {
	r := &PortPacketThresholdDataSource{client: newTransportErrorClient(t)}
	m := PortPacketThresholdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestPortPacketThresholdDataSource_Read_InvalidJSON exercises PortPacketThresholdDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestPortPacketThresholdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &PortPacketThresholdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := PortPacketThresholdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
