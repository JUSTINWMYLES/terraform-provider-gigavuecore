package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTimePropertiesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllTimePropertiesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllTimePropertiesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllTimePropertiesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllTimePropertiesDataSourceMetadata(t *testing.T) {
	d := NewGetAllTimePropertiesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_time_properties" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_time_properties")
	}
}
