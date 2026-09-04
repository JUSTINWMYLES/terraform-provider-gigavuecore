package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestExpiryCountFlAndNllDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestExpiryCountFlAndNllDataSourceSchemaValidation(t *testing.T) {
	d := NewExpiryCountFlAndNllDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestExpiryCountFlAndNllDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestExpiryCountFlAndNllDataSourceMetadata(t *testing.T) {
	d := NewExpiryCountFlAndNllDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_expiry_count_fl_and_nll" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_expiry_count_fl_and_nll")
	}
}
