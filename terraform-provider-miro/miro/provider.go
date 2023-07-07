package miro

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_token": {
				Type:        schema.TypeString,
				Description: "Access key for Miro API",
				Required:    true,
				Sensitive:   true,
			},
		},
		ResourcesMap:   map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
