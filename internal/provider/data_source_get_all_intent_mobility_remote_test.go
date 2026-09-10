package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllIntentMobilityDataSource_Read_Happy exercises GetAllIntentMobilityDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllIntentMobilityDataSource_Read_Happy(t *testing.T) {
	r := &GetAllIntentMobilityDataSource{client: newMockClientStatus(t, 200, "[]")}
	m := GetAllIntentMobilityDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllIntentMobilityDataSource_Read_NilClient exercises GetAllIntentMobilityDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllIntentMobilityDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllIntentMobilityDataSource{}
	m := GetAllIntentMobilityDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllIntentMobilityDataSource_Read_BuildError exercises GetAllIntentMobilityDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllIntentMobilityDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllIntentMobilityDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllIntentMobilityDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllIntentMobilityDataSource_Read_SendError exercises GetAllIntentMobilityDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllIntentMobilityDataSource_Read_SendError(t *testing.T) {
	r := &GetAllIntentMobilityDataSource{client: newTransportErrorClient(t)}
	m := GetAllIntentMobilityDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllIntentMobilityDataSource_Read_InvalidJSON exercises GetAllIntentMobilityDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllIntentMobilityDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllIntentMobilityDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllIntentMobilityDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
