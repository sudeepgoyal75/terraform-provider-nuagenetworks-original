package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.6.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceInfrastructureaccessprofile() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceInfrastructureaccessprofileRead,
        Schema: map[string]*schema.Schema{
            "filter": dataSourceFiltersSchema(),
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ssh_auth_mode": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "password": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_ip_filter": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "user_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}


func dataSourceInfrastructureaccessprofileRead(d *schema.ResourceData, m interface{}) error {
    filteredInfrastructureaccessprofiles := vspk.InfrastructureaccessprofilesList{}
    err := &bambou.Error{}
    fetchFilter := &bambou.FetchingInfo{}
    
    filters, filtersOk := d.GetOk("filter")
    if filtersOk {
        fetchFilter = bambou.NewFetchingInfo()
        for _, v := range filters.(*schema.Set).List() {
            m := v.(map[string]interface{})
            if fetchFilter.Filter != "" {
                fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string),  m["operator"].(string),  m["value"].(string))
            } else {
                fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
            }
           
        }
    }

    Infrastructureaccessprofile := &vspk.Infrastructureaccessprofile{}

    if len(filteredInfrastructureaccessprofiles) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredInfrastructureaccessprofiles) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Infrastructureaccessprofile = filteredInfrastructureaccessprofiles[0]

    d.Set("ssh_auth_mode", Infrastructureaccessprofile.SSHAuthMode)
    d.Set("name", Infrastructureaccessprofile.Name)
    d.Set("password", Infrastructureaccessprofile.Password)
    d.Set("last_updated_by", Infrastructureaccessprofile.LastUpdatedBy)
    d.Set("description", Infrastructureaccessprofile.Description)
    d.Set("enterprise_id", Infrastructureaccessprofile.EnterpriseID)
    d.Set("entity_scope", Infrastructureaccessprofile.EntityScope)
    d.Set("source_ip_filter", Infrastructureaccessprofile.SourceIPFilter)
    d.Set("user_name", Infrastructureaccessprofile.UserName)
    d.Set("external_id", Infrastructureaccessprofile.ExternalID)
    
    d.Set("id", Infrastructureaccessprofile.Identifier())
    d.Set("parent_id", Infrastructureaccessprofile.ParentID)
    d.Set("parent_type", Infrastructureaccessprofile.ParentType)
    d.Set("owner", Infrastructureaccessprofile.Owner)

    d.SetId(Infrastructureaccessprofile.Identifier())
    
    return nil
}