package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadNrtStatsFlexInlineSolutionDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadNrtStatsFlexInlineSolutionDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadNrtStatsFlexInlineSolutionDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadNrtStatsFlexInlineSolutionDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadNrtStatsFlexInlineSolutionDataSourceMetadata(t *testing.T) {
	d := NewLoadNrtStatsFlexInlineSolutionDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_nrt_stats_flex_inline_solution" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_nrt_stats_flex_inline_solution")
	}
}
