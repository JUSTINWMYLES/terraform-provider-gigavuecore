package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPtpParentByAliasDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetPtpParentByAliasDataSourceSchemaValidation(t *testing.T) {
	d := NewGetPtpParentByAliasDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetPtpParentByAliasDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetPtpParentByAliasDataSourceMetadata(t *testing.T) {
	d := NewGetPtpParentByAliasDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_ptp_parent_by_alias" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_ptp_parent_by_alias")
	}
}
