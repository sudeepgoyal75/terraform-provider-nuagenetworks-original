package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.2.1"
)

func resourceVNFDescriptor() *schema.Resource {
    return &schema.Resource{
        Create: resourceVNFDescriptorCreate,
        Read:   resourceVNFDescriptorRead,
        Update: resourceVNFDescriptorUpdate,
        Delete: resourceVNFDescriptorDelete,
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
            "cpu_count": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "memory_mb": &schema.Schema{
                Type:     schema.TypeInt,
                Required: true,
            },
            "vendor": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "metadata_id": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "visible": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: true,
            },
            "associated_vnf_threshold_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "storage_gb": &schema.Schema{
                Type:     schema.TypeInt,
                Required: true,
            },
            "parent_vnf_catalog": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceVNFDescriptorCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize VNFDescriptor object
    o := &vspk.VNFDescriptor{
        Name: d.Get("name").(string),
        MemoryMB: d.Get("memory_mb").(int),
        MetadataID: d.Get("metadata_id").(string),
        StorageGB: d.Get("storage_gb").(int),
    }
    if attr, ok := d.GetOk("cpu_count"); ok {
        o.CPUCount = attr.(int)
    }
    if attr, ok := d.GetOk("vendor"); ok {
        o.Vendor = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("visible"); ok {
        o.Visible = attr.(bool)
    }
    if attr, ok := d.GetOk("associated_vnf_threshold_policy_id"); ok {
        o.AssociatedVNFThresholdPolicyID = attr.(string)
    }
    parent := &vspk.VNFCatalog{ID: d.Get("parent_vnf_catalog").(string)}
    err := parent.CreateVNFDescriptor(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceVNFDescriptorRead(d, m)
}

func resourceVNFDescriptorRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VNFDescriptor{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("cpu_count", o.CPUCount)
    d.Set("name", o.Name)
    d.Set("memory_mb", o.MemoryMB)
    d.Set("vendor", o.Vendor)
    d.Set("description", o.Description)
    d.Set("metadata_id", o.MetadataID)
    d.Set("visible", o.Visible)
    d.Set("associated_vnf_threshold_policy_id", o.AssociatedVNFThresholdPolicyID)
    d.Set("storage_gb", o.StorageGB)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceVNFDescriptorUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VNFDescriptor{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.MemoryMB = d.Get("memory_mb").(int)
    o.MetadataID = d.Get("metadata_id").(string)
    o.StorageGB = d.Get("storage_gb").(int)
    
    if attr, ok := d.GetOk("cpu_count"); ok {
        o.CPUCount = attr.(int)
    }
    if attr, ok := d.GetOk("vendor"); ok {
        o.Vendor = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("visible"); ok {
        o.Visible = attr.(bool)
    }
    if attr, ok := d.GetOk("associated_vnf_threshold_policy_id"); ok {
        o.AssociatedVNFThresholdPolicyID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceVNFDescriptorDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VNFDescriptor{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}