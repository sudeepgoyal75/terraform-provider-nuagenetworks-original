package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.2.1"
)

func resourceDSCPRemarkingPolicyTable() *schema.Resource {
    return &schema.Resource{
        Create: resourceDSCPRemarkingPolicyTableCreate,
        Read:   resourceDSCPRemarkingPolicyTableRead,
        Update: resourceDSCPRemarkingPolicyTableUpdate,
        Delete: resourceDSCPRemarkingPolicyTableDelete,
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
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}

func resourceDSCPRemarkingPolicyTableCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize DSCPRemarkingPolicyTable object
    o := &vspk.DSCPRemarkingPolicyTable{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_me"); ok {
        parent := &vspk.Me{ID: attr.(string)}
        err := parent.CreateDSCPRemarkingPolicyTable(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        err := parent.CreateDSCPRemarkingPolicyTable(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceDSCPRemarkingPolicyTableRead(d, m)
}

func resourceDSCPRemarkingPolicyTableRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DSCPRemarkingPolicyTable{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("description", o.Description)
    d.Set("entity_scope", o.EntityScope)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceDSCPRemarkingPolicyTableUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DSCPRemarkingPolicyTable{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceDSCPRemarkingPolicyTableDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DSCPRemarkingPolicyTable{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}