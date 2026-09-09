package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllImageRepoImagesDataSource_Read_Happy exercises LoadAllImageRepoImagesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestLoadAllImageRepoImagesDataSource_Read_Happy(t *testing.T) {
	r := &LoadAllImageRepoImagesDataSource{client: newMockClientStatus(t, 200, "{\"ImageRepoImagesQueryResponse\":[]}")}
	m := LoadAllImageRepoImagesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestLoadAllImageRepoImagesDataSource_Read_NilClient exercises LoadAllImageRepoImagesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestLoadAllImageRepoImagesDataSource_Read_NilClient(t *testing.T) {
	r := &LoadAllImageRepoImagesDataSource{}
	m := LoadAllImageRepoImagesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestLoadAllImageRepoImagesDataSource_Read_BuildError exercises LoadAllImageRepoImagesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestLoadAllImageRepoImagesDataSource_Read_BuildError(t *testing.T) {
	r := &LoadAllImageRepoImagesDataSource{client: newMalformedBaseURLClient(t)}
	m := LoadAllImageRepoImagesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllImageRepoImagesDataSource_Read_SendError exercises LoadAllImageRepoImagesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestLoadAllImageRepoImagesDataSource_Read_SendError(t *testing.T) {
	r := &LoadAllImageRepoImagesDataSource{client: newTransportErrorClient(t)}
	m := LoadAllImageRepoImagesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestLoadAllImageRepoImagesDataSource_Read_InvalidJSON exercises LoadAllImageRepoImagesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestLoadAllImageRepoImagesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &LoadAllImageRepoImagesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := LoadAllImageRepoImagesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
