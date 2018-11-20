provider "example" {
  seed = 12345
}

resource "user" "user1" {
  provider = "example"
  username = "user1"
}
output "user_user1_username" {
  value = "${user.user1.username}"
}
output "user_user1_password" {
  value = "${user.user1.password}"
}