package provider

import "context"
import "github.com/hashicorp/terraform-plugin-framework/function"

var _ function.Function = (*LoadResultsBasedOnScrollIdFunction)(nil)

// LoadResultsBasedOnScrollIdFunction is the generated Terraform provider-defined function implementation.
type LoadResultsBasedOnScrollIdFunction struct {
	SourceOperation string
}

// Metadata returns the function name.
func (f *LoadResultsBasedOnScrollIdFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "load_results_based_on_scroll_id"
}

// Definition returns the function signature.
func (f *LoadResultsBasedOnScrollIdFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{Summary: "Load Results Based On Scroll Id", Description: "Load results  based on ScrollId", Parameters: []function.Parameter{function.StringParameter{Name: "scroll_id", Description: "scrollId which is used in the scroll API in order to retrieve the required batch of results"}, function.StringParameter{Name: "response_scroll_timeout", Description: "Scroll time out to mention how long it should keep the \"search context\" alive"}}, Return: function.StringReturn{}}
}

// Run executes the function logic.
func (f *LoadResultsBasedOnScrollIdFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	resp.Error = function.NewFuncError("Run is not wired to a remote API endpoint.")
}
