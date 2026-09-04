package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllFeatureActivationsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllFeatureActivationsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllFeatureActivationsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllFeatureActivationsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllFeatureActivationsDataSourceMetadata(t *testing.T) {
	d := NewGetAllFeatureActivationsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_feature_activations" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_feature_activations")
	}
}
