package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNotifConfigResourceSchemaValidation verifies that the generated resource schema is valid.
func TestNotifConfigResourceSchemaValidation(t *testing.T) {
	r := &NotifConfigResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestNotifConfigResourceMetadata verifies that the generated resource reports the expected type name.
func TestNotifConfigResourceMetadata(t *testing.T) {
	r := &NotifConfigResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_notif_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_notif_config")
	}
}
