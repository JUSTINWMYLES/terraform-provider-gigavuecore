package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGsCardsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGsCardsDataSourceSchemaValidation(t *testing.T) {
	d := NewGsCardsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGsCardsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGsCardsDataSourceMetadata(t *testing.T) {
	d := NewGsCardsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_gs_cards" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_gs_cards")
	}
}
