package bridge

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// OutputResource represents the structure of the data sent to/received from the Bridge API.
type OutputResource struct {
	Name         string `json:"name"`
	Arn          string `json:"arn"`
	Value        string `json:"value"`
	ResourceType string `json:"resource_type"`
}

// resourceBridgeOutput defines the schema and operations for the "bridge_output" resource.
func resourceBridgeOutput() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bridge_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"arn": {
				Type:     schema.TypeString,
				Required: false,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: false,
			},
			"bridge_register": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
		CreateContext: resourceBridgeOutputCreate,
		ReadContext:   resourceBridgeOutputRead,
		UpdateContext: resourceBridgeOutputUpdate,
		DeleteContext: resourceBridgeOutputDelete,
	}
}

// resourceBridgeOutputCreate handles the creation of the resource.
func resourceBridgeOutputCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	providerMeta := meta.(map[string]interface{})
	apiURL := providerMeta["api_url"].(string)

	if !d.Get("bridge_register").(bool) {
		return nil // Exit early if registration is disabled
	}

	name := d.Get("bridge_name").(string)
	value := d.Get("value").(string)
	arn := d.Get("arn").(string)
	resourceType := d.Get("resource_type").(string)

	outputResource := OutputResource{Name: name, Value: value, Arn: arn, ResourceType: resourceType}
	resourceData, err := json.Marshal(outputResource)
	if err != nil {
		return diag.FromErr(err)
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/resource", apiURL), bytes.NewBuffer(resourceData))
	if err != nil {
		return diag.FromErr(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		var apiError map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&apiError); err != nil {
			return diag.FromErr(err)
		}
		return diag.Errorf("Bridge API Error (%d): %s", resp.StatusCode, apiError["message"])
	}

	d.SetId(name)
	return nil
}

// resourceBridgeOutputRead reads the current state of the resource from the Bridge API.
func resourceBridgeOutputRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	providerMeta := meta.(map[string]interface{})
	apiURL := providerMeta["api_url"].(string)

	name := d.Id() // Get the resource name from the ID

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/resource/%s", apiURL, name), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return diag.Errorf("Failed to read output resource: %s", resp.Status)
	}

	var outputResource OutputResource
	if err := json.NewDecoder(resp.Body).Decode(&outputResource); err != nil {
		return diag.FromErr(err)
	}

	// Update the Terraform state with the values from the API
	d.Set("bridge_name", outputResource.Name)
	d.Set("value", outputResource.Value)
	d.Set("arn", outputResource.Arn)
	d.Set("resource_type", outputResource.ResourceType)
	d.Set("bridge_register", true) // Assuming the resource is registered if it exists in the API

	return nil
}

// resourceBridgeOutputUpdate handles the update of the resource.
func resourceBridgeOutputUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	providerMeta := meta.(map[string]interface{})
	apiURL := providerMeta["api_url"].(string)

	name := d.Get("bridge_name").(string)
	value := d.Get("value").(string)
	arn := d.Get("arn").(string)
	resourceType := d.Get("resource_type").(string)

	outputResource := OutputResource{Name: name, Value: value, Arn: arn, ResourceType: resourceType}
	resourceData, err := json.Marshal(outputResource)
	if err != nil {
		return diag.FromErr(err)
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/resource/%s", apiURL, name), bytes.NewBuffer(resourceData))
	if err != nil {
		return diag.FromErr(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiError map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&apiError); err != nil {
			return diag.FromErr(err)
		}
		return diag.Errorf("Bridge API Error (%d): %s", resp.StatusCode, apiError["message"])
	}

	d.SetId(name)
	return nil
}

// resourceBridgeOutputDelete deletes the resource from the Bridge API.
func resourceBridgeOutputDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	providerMeta := meta.(map[string]interface{})
	apiURL := providerMeta["api_url"].(string)
	name := d.Id()

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/resource/%s", apiURL, name), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return diag.Errorf("failed to delete output resource: %s", resp.Status)
	}

	d.SetId("")

	return nil
}
