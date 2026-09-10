package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetHsmKeyMapsDataSource_Read_Happy exercises GetHsmKeyMapsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetHsmKeyMapsDataSource_Read_Happy(t *testing.T) {
	r := &GetHsmKeyMapsDataSource{client: newMockClientStatus(t, 200, "{\"hsmKeyMaps\":[]}")}
	m := GetHsmKeyMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetHsmKeyMapsDataSource_Read_NilClient exercises GetHsmKeyMapsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetHsmKeyMapsDataSource_Read_NilClient(t *testing.T) {
	r := &GetHsmKeyMapsDataSource{}
	m := GetHsmKeyMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetHsmKeyMapsDataSource_Read_BuildError exercises GetHsmKeyMapsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetHsmKeyMapsDataSource_Read_BuildError(t *testing.T) {
	r := &GetHsmKeyMapsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetHsmKeyMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetHsmKeyMapsDataSource_Read_SendError exercises GetHsmKeyMapsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetHsmKeyMapsDataSource_Read_SendError(t *testing.T) {
	r := &GetHsmKeyMapsDataSource{client: newTransportErrorClient(t)}
	m := GetHsmKeyMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetHsmKeyMapsDataSource_Read_InvalidJSON exercises GetHsmKeyMapsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetHsmKeyMapsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetHsmKeyMapsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetHsmKeyMapsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
