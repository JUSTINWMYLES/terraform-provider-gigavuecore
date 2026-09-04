package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadCardsDetailsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadCardsDetailsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadCardsDetailsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadCardsDetailsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadCardsDetailsDataSourceMetadata(t *testing.T) {
	d := NewLoadCardsDetailsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_cards_details" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_cards_details")
	}
}
