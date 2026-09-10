package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllVblFeatureActivationsDataSource_Read_Happy exercises GetAllVblFeatureActivationsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllVblFeatureActivationsDataSource_Read_Happy(t *testing.T) {
	r := &GetAllVblFeatureActivationsDataSource{client: newMockClientStatus(t, 200, "{\"activations\":[]}")}
	m := GetAllVblFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllVblFeatureActivationsDataSource_Read_NilClient exercises GetAllVblFeatureActivationsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllVblFeatureActivationsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllVblFeatureActivationsDataSource{}
	m := GetAllVblFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllVblFeatureActivationsDataSource_Read_BuildError exercises GetAllVblFeatureActivationsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllVblFeatureActivationsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllVblFeatureActivationsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllVblFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllVblFeatureActivationsDataSource_Read_SendError exercises GetAllVblFeatureActivationsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllVblFeatureActivationsDataSource_Read_SendError(t *testing.T) {
	r := &GetAllVblFeatureActivationsDataSource{client: newTransportErrorClient(t)}
	m := GetAllVblFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllVblFeatureActivationsDataSource_Read_InvalidJSON exercises GetAllVblFeatureActivationsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllVblFeatureActivationsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllVblFeatureActivationsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllVblFeatureActivationsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
