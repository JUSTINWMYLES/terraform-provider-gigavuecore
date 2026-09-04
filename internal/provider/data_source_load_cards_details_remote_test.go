package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadCardsDetailsDataSource_Read_Happy exercises LoadCardsDetailsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadCardsDetailsDataSource_Read_Happy(t *testing.T) {
	r := &LoadCardsDetailsDataSource{client: newMockClientStatus(t, 200, "{\"cards\":[]}")}
	m := LoadCardsDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadCardsDetailsDataSource_Read_NilClient exercises LoadCardsDetailsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadCardsDetailsDataSource_Read_NilClient(t *testing.T) {
	r := &LoadCardsDetailsDataSource{}
	m := LoadCardsDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadCardsDetailsDataSource_Read_BuildError exercises LoadCardsDetailsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadCardsDetailsDataSource_Read_BuildError(t *testing.T) {
	r := &LoadCardsDetailsDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadCardsDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadCardsDetailsDataSource_Read_SendError exercises LoadCardsDetailsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadCardsDetailsDataSource_Read_SendError(t *testing.T) {
	r := &LoadCardsDetailsDataSource{client: newTransportErrorClient(t)}
	m := LoadCardsDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadCardsDetailsDataSource_Read_InvalidJSON exercises LoadCardsDetailsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadCardsDetailsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadCardsDetailsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadCardsDetailsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
