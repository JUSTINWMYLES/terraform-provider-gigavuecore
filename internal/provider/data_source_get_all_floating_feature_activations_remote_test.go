package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllFloatingFeatureActivationsDataSource_Read_Happy exercises GetAllFloatingFeatureActivationsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllFloatingFeatureActivationsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllFloatingFeatureActivationsDataSource{client: newMockClientStatus(t, 200, "{\"activations\":[]}")}
	m := GetAllFloatingFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllFloatingFeatureActivationsDataSource_Read_NilClient exercises GetAllFloatingFeatureActivationsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllFloatingFeatureActivationsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllFloatingFeatureActivationsDataSource{}
	m := GetAllFloatingFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllFloatingFeatureActivationsDataSource_Read_BuildError exercises GetAllFloatingFeatureActivationsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllFloatingFeatureActivationsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllFloatingFeatureActivationsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllFloatingFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllFloatingFeatureActivationsDataSource_Read_SendError exercises GetAllFloatingFeatureActivationsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllFloatingFeatureActivationsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllFloatingFeatureActivationsDataSource{client: newTransportErrorClient(t)}
	m := GetAllFloatingFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllFloatingFeatureActivationsDataSource_Read_InvalidJSON exercises GetAllFloatingFeatureActivationsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllFloatingFeatureActivationsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllFloatingFeatureActivationsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllFloatingFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
