package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetArpEntriesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetArpEntriesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetArpEntriesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetArpEntriesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetArpEntriesDataSourceMetadata(t *testing.T) {
	d := NewGetArpEntriesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_arp_entries" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_arp_entries")
	}
}
