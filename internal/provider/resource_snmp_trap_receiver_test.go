package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSnmpTrapReceiverResourceSchemaValidation verifies that the generated resource schema is valid.
func TestSnmpTrapReceiverResourceSchemaValidation(t *testing.T) {
	r := &SnmpTrapReceiverResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestSnmpTrapReceiverResourceMetadata verifies that the generated resource reports the expected type name.
func TestSnmpTrapReceiverResourceMetadata(t *testing.T) {
	r := &SnmpTrapReceiverResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_snmp_trap_receiver" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_snmp_trap_receiver")
	}
}
