package example

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	var err error
	log.Printf("Creating user password\n")
	seed := meta.(int)
	log.Printf("Seed is %d\n", seed)
	if err = d.Set("password", "PAssword121"); err != nil {
		return fmt.Errorf("Error setting password: %s", err)
	}
	log.Println("Password set to ", "Password")
	return nil
}

func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Reading user password\n")
	return nil
}
func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Deleting user password\n")
	return nil
}
