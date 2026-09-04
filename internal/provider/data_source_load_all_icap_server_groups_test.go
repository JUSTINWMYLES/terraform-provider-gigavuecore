package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllIcapServerGroupsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllIcapServerGroupsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllIcapServerGroupsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllIcapServerGroupsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllIcapServerGroupsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllIcapServerGroupsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_icap_server_groups" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_icap_server_groups")
	}
}
