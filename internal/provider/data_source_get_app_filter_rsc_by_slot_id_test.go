package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAppFilterRscBySlotIdDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAppFilterRscBySlotIdDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAppFilterRscBySlotIdDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAppFilterRscBySlotIdDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAppFilterRscBySlotIdDataSourceMetadata(t *testing.T) {
	d := NewGetAppFilterRscBySlotIdDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_app_filter_rsc_by_slot_id" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_app_filter_rsc_by_slot_id")
	}
}
