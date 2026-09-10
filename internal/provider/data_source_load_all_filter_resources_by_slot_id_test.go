package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllFilterResourcesBySlotIdDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllFilterResourcesBySlotIdDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllFilterResourcesBySlotIdDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllFilterResourcesBySlotIdDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllFilterResourcesBySlotIdDataSourceMetadata(t *testing.T) {
	d := NewLoadAllFilterResourcesBySlotIdDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_filter_resources_by_slot_id" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_filter_resources_by_slot_id")
	}
}
