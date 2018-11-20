package example

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"log"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	//TODO: provide backwards compability
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"seed": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SEED", nil),
				Description: `seed for generating password`,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"user": resourceUser(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	log.Printf("Configuring provider\n")
	genSeed := d.Get("seed").(int)
	return genSeed, nil
}
