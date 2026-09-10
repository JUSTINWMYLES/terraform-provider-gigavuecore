package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllFeatureActivationsDataSource_Read_Happy exercises GetAllFeatureActivationsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllFeatureActivationsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllFeatureActivationsDataSource{client: newMockClientStatus(t, 200, "{\"activations\":[]}")}
	m := GetAllFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllFeatureActivationsDataSource_Read_NilClient exercises GetAllFeatureActivationsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllFeatureActivationsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllFeatureActivationsDataSource{}
	m := GetAllFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllFeatureActivationsDataSource_Read_BuildError exercises GetAllFeatureActivationsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllFeatureActivationsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllFeatureActivationsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllFeatureActivationsDataSource_Read_SendError exercises GetAllFeatureActivationsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllFeatureActivationsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllFeatureActivationsDataSource{client: newTransportErrorClient(t)}
	m := GetAllFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllFeatureActivationsDataSource_Read_InvalidJSON exercises GetAllFeatureActivationsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllFeatureActivationsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllFeatureActivationsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
