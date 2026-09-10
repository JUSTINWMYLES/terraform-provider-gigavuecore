package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllFloatingFeatureActivationsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllFloatingFeatureActivationsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllFloatingFeatureActivationsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllFloatingFeatureActivationsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllFloatingFeatureActivationsDataSourceMetadata(t *testing.T) {
	d := NewGetAllFloatingFeatureActivationsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_floating_feature_activations" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_floating_feature_activations")
	}
}
