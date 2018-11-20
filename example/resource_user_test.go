package example

import (
	"fmt"
	r "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestUser(t *testing.T) {
	r.Test(t, r.TestCase{
		Providers: testProviders,
		Steps: []r.TestStep{
			r.TestStep{
				Config: fmt.Sprintf(`
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
                `),
				Check: func(s *terraform.State) error {
					gotUntyped := s.RootModule().Outputs["user_user1_username"].Value
					got, ok := gotUntyped.(string)
					if !ok {
						return fmt.Errorf("output for \"user_user1\" is not a string")
					}

					if !strings.Contains(got, "user1") {
						return fmt.Errorf("user1 is missing")
					}
					if expected, got := "user1", gotUntyped; got != expected {
						return fmt.Errorf("incorrect user name: expected %v, got %v", expected, got)
					}

					return nil

				},
			},
		},
	})
}
