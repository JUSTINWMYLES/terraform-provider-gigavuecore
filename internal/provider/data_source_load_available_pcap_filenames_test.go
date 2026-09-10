package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAvailablePcapFilenamesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAvailablePcapFilenamesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAvailablePcapFilenamesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAvailablePcapFilenamesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAvailablePcapFilenamesDataSourceMetadata(t *testing.T) {
	d := NewLoadAvailablePcapFilenamesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_available_pcap_filenames" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_available_pcap_filenames")
	}
}
