package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
)

func resourceWirelessPort() *schema.Resource {
    return &schema.Resource{
        Create: resourceWirelessPortCreate,
        Read:   resourceWirelessPortRead,
        Update: resourceWirelessPortUpdate,
        Delete: resourceWirelessPortDelete,
        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },
        Schema: map[string]*schema.Schema{
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vlan_range": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "generic_config": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "permitted_action": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "physical_name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "wifi_frequency_band": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "wifi_mode": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "port_type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "country_code": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "frequency_channel": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "use_user_mnemonic": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "user_mnemonic": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_egress_qos_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceWirelessPortCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize WirelessPort object
    o := &vspk.WirelessPort{
        Name: d.Get("name").(string),
        PhysicalName: d.Get("physical_name").(string),
        WifiFrequencyBand: d.Get("wifi_frequency_band").(string),
        WifiMode: d.Get("wifi_mode").(string),
        PortType: d.Get("port_type").(string),
        CountryCode: d.Get("country_code").(string),
        FrequencyChannel: d.Get("frequency_channel").(string),
    }
    if attr, ok := d.GetOk("generic_config"); ok {
        o.GenericConfig = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("use_user_mnemonic"); ok {
        o.UseUserMnemonic = attr.(bool)
    }
    if attr, ok := d.GetOk("user_mnemonic"); ok {
        o.UserMnemonic = attr.(string)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.NSGateway{ID: d.Get("parent_ns_gateway").(string)}
    err := parent.CreateWirelessPort(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceWirelessPortRead(d, m)
}

func resourceWirelessPortRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.WirelessPort{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("vlan_range", o.VLANRange)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("generic_config", o.GenericConfig)
    d.Set("permitted_action", o.PermittedAction)
    d.Set("description", o.Description)
    d.Set("physical_name", o.PhysicalName)
    d.Set("wifi_frequency_band", o.WifiFrequencyBand)
    d.Set("wifi_mode", o.WifiMode)
    d.Set("entity_scope", o.EntityScope)
    d.Set("port_type", o.PortType)
    d.Set("country_code", o.CountryCode)
    d.Set("frequency_channel", o.FrequencyChannel)
    d.Set("use_user_mnemonic", o.UseUserMnemonic)
    d.Set("user_mnemonic", o.UserMnemonic)
    d.Set("associated_egress_qos_policy_id", o.AssociatedEgressQOSPolicyID)
    d.Set("status", o.Status)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceWirelessPortUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.WirelessPort{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.PhysicalName = d.Get("physical_name").(string)
    o.WifiFrequencyBand = d.Get("wifi_frequency_band").(string)
    o.WifiMode = d.Get("wifi_mode").(string)
    o.PortType = d.Get("port_type").(string)
    o.CountryCode = d.Get("country_code").(string)
    o.FrequencyChannel = d.Get("frequency_channel").(string)
    
    if attr, ok := d.GetOk("generic_config"); ok {
        o.GenericConfig = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("use_user_mnemonic"); ok {
        o.UseUserMnemonic = attr.(bool)
    }
    if attr, ok := d.GetOk("user_mnemonic"); ok {
        o.UserMnemonic = attr.(string)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceWirelessPortDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.WirelessPort{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}