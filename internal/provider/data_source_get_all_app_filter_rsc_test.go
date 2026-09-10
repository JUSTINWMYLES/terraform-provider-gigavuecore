package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllAppFilterRscDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllAppFilterRscDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllAppFilterRscDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllAppFilterRscDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllAppFilterRscDataSourceMetadata(t *testing.T) {
	d := NewGetAllAppFilterRscDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_app_filter_rsc" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_app_filter_rsc")
	}
}
