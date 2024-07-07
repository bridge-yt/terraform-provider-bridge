package bridge

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BRIDGE_API_URL", nil),
				Description: "URL of the Bridge API",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"bridge_output": resourceBridgeOutput(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"bridge_value": dataSourceBridgeValue(),
		},
		ConfigureFunc: providerConfigure,
	}
}

// providerConfigure configures the provider with the API URL.
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return map[string]interface{}{
		"api_url": d.Get("api_url").(string),
	}, nil
}
