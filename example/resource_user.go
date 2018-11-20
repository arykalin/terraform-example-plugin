package example

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"math/rand"
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
	username := d.Get("username").(string)
	log.Printf("Creating user password for user %s\n", username)
	seed := meta.(int)
	log.Printf("Seed is %d\n", seed)
	password := randSeq(6, seed)
	if err = d.Set("password", password); err != nil {
		return fmt.Errorf("Error setting password: %s", err)
	}
	log.Println("Password set to ", password)
	d.SetId(username)
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

func randSeq(n, s int) string {
	rand.Seed(int64(s))
	var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
