package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllArpEntriesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllArpEntriesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllArpEntriesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllArpEntriesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllArpEntriesDataSourceMetadata(t *testing.T) {
	d := NewGetAllArpEntriesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_arp_entries" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_arp_entries")
	}
}
