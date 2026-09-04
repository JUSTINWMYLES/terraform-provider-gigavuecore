package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGigaStreamThresholdDataSource_Read_Happy exercises GigaStreamThresholdDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGigaStreamThresholdDataSource_Read_Happy(t *testing.T) {
	r := &GigaStreamThresholdDataSource{client: newMockClientStatus(t, 200, "{\"device\":[]}")}
	m := GigaStreamThresholdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGigaStreamThresholdDataSource_Read_NilClient exercises GigaStreamThresholdDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGigaStreamThresholdDataSource_Read_NilClient(t *testing.T) {
	r := &GigaStreamThresholdDataSource{}
	m := GigaStreamThresholdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGigaStreamThresholdDataSource_Read_BuildError exercises GigaStreamThresholdDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGigaStreamThresholdDataSource_Read_BuildError(t *testing.T) {
	r := &GigaStreamThresholdDataSource{client: newMalformedBaseURLClient(t)}
	m := GigaStreamThresholdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGigaStreamThresholdDataSource_Read_SendError exercises GigaStreamThresholdDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGigaStreamThresholdDataSource_Read_SendError(t *testing.T) {
	r := &GigaStreamThresholdDataSource{client: newTransportErrorClient(t)}
	m := GigaStreamThresholdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGigaStreamThresholdDataSource_Read_InvalidJSON exercises GigaStreamThresholdDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGigaStreamThresholdDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GigaStreamThresholdDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GigaStreamThresholdDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
