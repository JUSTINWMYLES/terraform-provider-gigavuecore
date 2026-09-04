package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllUnlicensedChassisAndCardsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllUnlicensedChassisAndCardsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllUnlicensedChassisAndCardsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllUnlicensedChassisAndCardsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllUnlicensedChassisAndCardsDataSourceMetadata(t *testing.T) {
	d := NewGetAllUnlicensedChassisAndCardsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_unlicensed_chassis_and_cards" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_unlicensed_chassis_and_cards")
	}
}
