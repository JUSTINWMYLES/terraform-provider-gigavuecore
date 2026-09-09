package provider

import (
	"context"
	"testing"
)

// TestTrafficPolicyGraphListResource_List_Happy exercises TrafficPolicyGraphListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestTrafficPolicyGraphListResource_List_Happy(t *testing.T) {
	r := &TrafficPolicyGraphListResource{client: newMockClientStatus(t, 200, "{\"trafficPolicyGraphs\":[]}")}
	m := TrafficPolicyGraphListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestTrafficPolicyGraphListResource_List_NilClient exercises TrafficPolicyGraphListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTrafficPolicyGraphListResource_List_NilClient(t *testing.T) {
	r := &TrafficPolicyGraphListResource{}
	m := TrafficPolicyGraphListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestTrafficPolicyGraphListResource_List_BuildError exercises TrafficPolicyGraphListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestTrafficPolicyGraphListResource_List_BuildError(t *testing.T) {
	r := &TrafficPolicyGraphListResource{client: newMalformedBaseURLClient(t)}
	m := TrafficPolicyGraphListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestTrafficPolicyGraphListResource_List_SendError exercises TrafficPolicyGraphListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestTrafficPolicyGraphListResource_List_SendError(t *testing.T) {
	r := &TrafficPolicyGraphListResource{client: newTransportErrorClient(t)}
	m := TrafficPolicyGraphListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestTrafficPolicyGraphListResource_List_InvalidJSON exercises TrafficPolicyGraphListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestTrafficPolicyGraphListResource_List_InvalidJSON(t *testing.T) {
	r := &TrafficPolicyGraphListResource{client: newMockClientStatus(t, 200, "{{")}
	m := TrafficPolicyGraphListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
