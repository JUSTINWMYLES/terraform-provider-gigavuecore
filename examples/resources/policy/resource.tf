resource "gigavuecore_policy" "example" {
  description = "example"
  enabled = true
  name = "example"
  then_do = [{
    action = "example"
    params = [{
      key = "example"
      value = "example"
    }]
  }]
  when_condition = [{
    condition = "example"
    params = [{
      key = "example"
      value = "example"
    }]
  }]
}
