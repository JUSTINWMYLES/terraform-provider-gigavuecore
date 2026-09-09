package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllIntentMobilityDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllIntentMobilityDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllIntentMobilityDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllIntentMobilityDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllIntentMobilityDataSourceMetadata(t *testing.T) {
	d := NewGetAllIntentMobilityDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_intent_mobility" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_intent_mobility")
	}
}
