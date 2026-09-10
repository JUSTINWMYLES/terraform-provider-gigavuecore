package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestFmHaDetailsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestFmHaDetailsDataSourceSchemaValidation(t *testing.T) {
	d := NewFmHaDetailsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestFmHaDetailsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestFmHaDetailsDataSourceMetadata(t *testing.T) {
	d := NewFmHaDetailsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_fm_ha_details" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_fm_ha_details")
	}
}
