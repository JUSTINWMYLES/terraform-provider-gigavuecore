package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemArpEntriesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSystemArpEntriesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSystemArpEntriesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSystemArpEntriesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSystemArpEntriesDataSourceMetadata(t *testing.T) {
	d := NewGetSystemArpEntriesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_system_arp_entries" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_system_arp_entries")
	}
}
