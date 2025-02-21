package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.1.1"
)

func resourceOverlayMirrorDestination() *schema.Resource {
    return &schema.Resource{
        Create: resourceOverlayMirrorDestinationCreate,
        Read:   resourceOverlayMirrorDestinationRead,
        Update: resourceOverlayMirrorDestinationUpdate,
        Delete: resourceOverlayMirrorDestinationDelete,
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
            "esi": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "redundancy_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "template_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "virtual_network_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "end_point_type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "trigger_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceOverlayMirrorDestinationCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize OverlayMirrorDestination object
    o := &vspk.OverlayMirrorDestination{
        EndPointType: d.Get("end_point_type").(string),
    }
    if attr, ok := d.GetOk("esi"); ok {
        o.ESI = attr.(string)
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("redundancy_enabled"); ok {
        o.RedundancyEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("template_id"); ok {
        o.TemplateID = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("virtual_network_id"); ok {
        o.VirtualNetworkID = attr.(string)
    }
    if attr, ok := d.GetOk("trigger_type"); ok {
        o.TriggerType = attr.(string)
    }
    parent := &vspk.L2Domain{ID: d.Get("parent_l2_domain").(string)}
    err := parent.CreateOverlayMirrorDestination(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceOverlayMirrorDestinationRead(d, m)
}

func resourceOverlayMirrorDestinationRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayMirrorDestination{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("esi", o.ESI)
    d.Set("name", o.Name)
    d.Set("redundancy_enabled", o.RedundancyEnabled)
    d.Set("template_id", o.TemplateID)
    d.Set("description", o.Description)
    d.Set("virtual_network_id", o.VirtualNetworkID)
    d.Set("end_point_type", o.EndPointType)
    d.Set("trigger_type", o.TriggerType)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceOverlayMirrorDestinationUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayMirrorDestination{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.EndPointType = d.Get("end_point_type").(string)
    
    if attr, ok := d.GetOk("esi"); ok {
        o.ESI = attr.(string)
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("redundancy_enabled"); ok {
        o.RedundancyEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("template_id"); ok {
        o.TemplateID = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("virtual_network_id"); ok {
        o.VirtualNetworkID = attr.(string)
    }
    if attr, ok := d.GetOk("trigger_type"); ok {
        o.TriggerType = attr.(string)
    }

    o.Save()

    return nil
}

func resourceOverlayMirrorDestinationDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayMirrorDestination{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}