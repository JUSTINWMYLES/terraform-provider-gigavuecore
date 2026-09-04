package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadCardDetailsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadCardDetailsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadCardDetailsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadCardDetailsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadCardDetailsDataSourceMetadata(t *testing.T) {
	d := NewLoadCardDetailsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_card_details" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_card_details")
	}
}
