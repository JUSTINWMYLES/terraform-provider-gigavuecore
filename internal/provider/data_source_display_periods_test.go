package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestDisplayPeriodsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestDisplayPeriodsDataSourceSchemaValidation(t *testing.T) {
	d := NewDisplayPeriodsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestDisplayPeriodsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestDisplayPeriodsDataSourceMetadata(t *testing.T) {
	d := NewDisplayPeriodsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_display_periods" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_display_periods")
	}
}
