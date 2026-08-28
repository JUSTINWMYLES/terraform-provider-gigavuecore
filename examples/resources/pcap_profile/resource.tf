resource "gigavuecore_pcap_profile" "example" {
  alias_list = "example"
  cluster_id = "example"
  pcap_configs = [{
    alias        = "example"
    channel_port = "example"
    direction    = "example"
    packet_limit = 1
    pcap_rules_list = [{
      dscp = "example"
      dst_mac = {
        address = "example"
        mask    = "example"
      }
      dstipv4_addrandmask = {
        address = "example"
        mask    = "example"
      }
      ether_type       = "example"
      inner_vlan       = 1
      ip4_frag         = "example"
      ip4_ttl          = 1
      ip_ver           = "example"
      packet_hit_count = 1
      portdst          = 1
      portsrc          = 1
      protocol         = 1
      rule_id          = 1
      src_mac = {
        address = "example"
        mask    = "example"
      }
      srcipv4_addrandmask = {
        address = "example"
        mask    = "example"
      }
      tcpctl = 1
      vlan   = 1
    }]
    port = "example"
  }]
  port_ids = "example"
}
