package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGsCardsDataSource_Read_Happy exercises GsCardsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGsCardsDataSource_Read_Happy(t *testing.T) {
	r := &GsCardsDataSource{client: newMockClientStatus(t, 200, "{\"gsCards\":[]}")}
	m := GsCardsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGsCardsDataSource_Read_NilClient exercises GsCardsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGsCardsDataSource_Read_NilClient(t *testing.T) {
	r := &GsCardsDataSource{}
	m := GsCardsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGsCardsDataSource_Read_BuildError exercises GsCardsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGsCardsDataSource_Read_BuildError(t *testing.T) {
	r := &GsCardsDataSource{client: newMalformedBaseURLClient(t)}
	m := GsCardsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGsCardsDataSource_Read_SendError exercises GsCardsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGsCardsDataSource_Read_SendError(t *testing.T) {
	r := &GsCardsDataSource{client: newTransportErrorClient(t)}
	m := GsCardsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGsCardsDataSource_Read_InvalidJSON exercises GsCardsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGsCardsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GsCardsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GsCardsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
