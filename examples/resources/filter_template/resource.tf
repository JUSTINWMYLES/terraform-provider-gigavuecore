resource "gigavuecore_filter_template" "example" {
  alias         = "example"
  cluster_id    = "example"
  qualifier_set = ["mplsLabelTtl"]
  qualifiers    = ["ipsrc"]
}
