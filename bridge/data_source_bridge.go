package bridge

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceBridgeValue() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBridgeValueRead,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceBridgeValueRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiURL := meta.(map[string]interface{})["api_url"].(string)
	namespace := d.Get("namespace").(string)
	name := d.Get("name").(string)

	url := fmt.Sprintf("%s/resource/%s/%s", apiURL, namespace, name)
	log.Printf("Fetching resource from URL: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error making HTTP request: %s", err)
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error response from API: %s", resp.Status)
		return diag.Errorf("Error fetching resource: %s", resp.Status)
	}

	var result OutputResource
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error decoding JSON response: %s", err)
		return diag.FromErr(err)
	}

	if err := d.Set("value", result.Value); err != nil {
		log.Printf("Error setting value: %s", err)
		return diag.FromErr(err)
	}

	if err := d.Set("arn", result.Arn); err != nil {
		log.Printf("Error setting ARN: %s", err)
		return diag.FromErr(err)
	}

	if err := d.Set("resource_type", result.ResourceType); err != nil {
		log.Printf("Error setting resource type: %s", err)
		return diag.FromErr(err)
	}

	d.SetId(name)
	log.Printf("Successfully fetched and set resource: %s", name)
	return nil
}
