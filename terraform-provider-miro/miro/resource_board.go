package miro

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceBoard() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the Board",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the Board",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
