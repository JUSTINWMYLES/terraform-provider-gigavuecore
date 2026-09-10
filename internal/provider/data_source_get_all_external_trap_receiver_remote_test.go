package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllExternalTrapReceiverDataSource_Read_Happy exercises GetAllExternalTrapReceiverDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAllExternalTrapReceiverDataSource_Read_Happy(t *testing.T) {
	r := &GetAllExternalTrapReceiverDataSource{client: newMockClientStatus(t, 200, "{\"externalTrapReceivers\":[]}")}
	m := GetAllExternalTrapReceiverDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAllExternalTrapReceiverDataSource_Read_NilClient exercises GetAllExternalTrapReceiverDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAllExternalTrapReceiverDataSource_Read_NilClient(t *testing.T) {
	r := &GetAllExternalTrapReceiverDataSource{}
	m := GetAllExternalTrapReceiverDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAllExternalTrapReceiverDataSource_Read_BuildError exercises GetAllExternalTrapReceiverDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAllExternalTrapReceiverDataSource_Read_BuildError(t *testing.T) {
	r := &GetAllExternalTrapReceiverDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAllExternalTrapReceiverDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllExternalTrapReceiverDataSource_Read_SendError exercises GetAllExternalTrapReceiverDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAllExternalTrapReceiverDataSource_Read_SendError(t *testing.T) {
	r := &GetAllExternalTrapReceiverDataSource{client: newTransportErrorClient(t)}
	m := GetAllExternalTrapReceiverDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAllExternalTrapReceiverDataSource_Read_InvalidJSON exercises GetAllExternalTrapReceiverDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAllExternalTrapReceiverDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAllExternalTrapReceiverDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAllExternalTrapReceiverDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
