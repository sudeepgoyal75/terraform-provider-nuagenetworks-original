package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.10.2"
)

func resourceVCenterHypervisor() *schema.Resource {
    return &schema.Resource{
        Create: resourceVCenterHypervisorCreate,
        Read:   resourceVCenterHypervisorRead,
        Update: resourceVCenterHypervisorUpdate,
        Delete: resourceVCenterHypervisorDelete,
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
            "vcenter_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vcenter_password": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vcenter_user": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vrs_configuration_time_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "vrs_metrics_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vrs_state": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "NOT_DEPLOYED",
            },
            "v_require_nuage_metadata": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "managed_object_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_vrs_deployed_date": &schema.Schema{
                Type:     schema.TypeFloat,
                Optional: true,
                Computed: true,
            },
            "data_dns1": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "data_dns2": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "data_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "data_ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "data_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "data_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "datapath_sync_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "scope": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "secondary_nuage_controller": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "removed_from_vcenter_inventory": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "generic_split_activation": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "separate_data_network": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "deployment_count": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "personality": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "destination_mirror_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "no_mirror",
            },
            "metadata_server_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "metadata_server_listen_port": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "metadata_server_port": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "metadata_service_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "network_uplink_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "network_uplink_interface_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "network_uplink_interface_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "network_uplink_interface_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nfs_log_server": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nfs_mount_path": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mgmt_dns1": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mgmt_dns2": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mgmt_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mgmt_ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mgmt_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mgmt_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "dhcp_relay_server": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mirror_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "site_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "allow_data_dhcp": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "allow_mgmt_dhcp": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "flow_eviction_threshold": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "vm_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "toolbox_deployment_mode": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "toolbox_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "toolbox_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "toolbox_password": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "toolbox_user_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "portgroup_metadata": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "nova_client_version": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "nova_metadata_service_auth_url": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nova_metadata_service_endpoint": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nova_metadata_service_password": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nova_metadata_service_tenant": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nova_metadata_service_username": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nova_metadata_shared_secret": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "nova_region_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "upgrade_package_password": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "upgrade_package_url": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "upgrade_package_username": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "upgrade_script_time_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "upgrade_status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "upgrade_timedout": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "primary_nuage_controller": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vrs_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vrs_password": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vrs_user_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "static_route": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "static_route_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "static_route_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "ntp_server1": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "ntp_server2": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mtu": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "successfully_applied_version": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multi_vmssupport": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "multicast_receive_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast_receive_interface_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast_receive_interface_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast_receive_range": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast_send_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast_send_interface_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast_send_interface_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast_source_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "customized_script_url": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "available_networks": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "ovf_url": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "hypervisor_ip": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "hypervisor_password": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "hypervisor_user": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "parent_vcenter_data_center": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_cluster"},
            },
            "parent_vcenter_cluster": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_data_center"},
            },
        },
    }
}

func resourceVCenterHypervisorCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize VCenterHypervisor object
    o := &vspk.VCenterHypervisor{
        Name: d.Get("name").(string),
        DataNetworkPortgroup: d.Get("data_network_portgroup").(string),
        Description: d.Get("description").(string),
        MgmtNetworkPortgroup: d.Get("mgmt_network_portgroup").(string),
        VmNetworkPortgroup: d.Get("vm_network_portgroup").(string),
        HypervisorIP: d.Get("hypervisor_ip").(string),
        HypervisorPassword: d.Get("hypervisor_password").(string),
        HypervisorUser: d.Get("hypervisor_user").(string),
    }
    if attr, ok := d.GetOk("vcenter_ip"); ok {
        o.VCenterIP = attr.(string)
    }
    if attr, ok := d.GetOk("vcenter_password"); ok {
        o.VCenterPassword = attr.(string)
    }
    if attr, ok := d.GetOk("vcenter_user"); ok {
        o.VCenterUser = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_configuration_time_limit"); ok {
        o.VRSConfigurationTimeLimit = attr.(int)
    }
    if attr, ok := d.GetOk("vrs_metrics_id"); ok {
        o.VRSMetricsID = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_state"); ok {
        o.VRSState = attr.(string)
    }
    if attr, ok := d.GetOk("v_require_nuage_metadata"); ok {
        o.VRequireNuageMetadata = attr.(bool)
    }
    if attr, ok := d.GetOk("managed_object_id"); ok {
        o.ManagedObjectID = attr.(string)
    }
    if attr, ok := d.GetOk("last_vrs_deployed_date"); ok {
        o.LastVRSDeployedDate = attr.(float64)
    }
    if attr, ok := d.GetOk("data_dns1"); ok {
        o.DataDNS1 = attr.(string)
    }
    if attr, ok := d.GetOk("data_dns2"); ok {
        o.DataDNS2 = attr.(string)
    }
    if attr, ok := d.GetOk("data_gateway"); ok {
        o.DataGateway = attr.(string)
    }
    if attr, ok := d.GetOk("data_ip_address"); ok {
        o.DataIPAddress = attr.(string)
    }
    if attr, ok := d.GetOk("data_netmask"); ok {
        o.DataNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("datapath_sync_timeout"); ok {
        o.DatapathSyncTimeout = attr.(int)
    }
    if attr, ok := d.GetOk("scope"); ok {
        o.Scope = attr.(bool)
    }
    if attr, ok := d.GetOk("secondary_nuage_controller"); ok {
        o.SecondaryNuageController = attr.(string)
    }
    if attr, ok := d.GetOk("removed_from_vcenter_inventory"); ok {
        o.RemovedFromVCenterInventory = attr.(bool)
    }
    if attr, ok := d.GetOk("generic_split_activation"); ok {
        o.GenericSplitActivation = attr.(bool)
    }
    if attr, ok := d.GetOk("separate_data_network"); ok {
        o.SeparateDataNetwork = attr.(bool)
    }
    if attr, ok := d.GetOk("deployment_count"); ok {
        o.DeploymentCount = attr.(int)
    }
    if attr, ok := d.GetOk("personality"); ok {
        o.Personality = attr.(string)
    }
    if attr, ok := d.GetOk("destination_mirror_port"); ok {
        o.DestinationMirrorPort = attr.(string)
    }
    if attr, ok := d.GetOk("metadata_server_ip"); ok {
        o.MetadataServerIP = attr.(string)
    }
    if attr, ok := d.GetOk("metadata_server_listen_port"); ok {
        o.MetadataServerListenPort = attr.(int)
    }
    if attr, ok := d.GetOk("metadata_server_port"); ok {
        o.MetadataServerPort = attr.(int)
    }
    if attr, ok := d.GetOk("metadata_service_enabled"); ok {
        o.MetadataServiceEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("network_uplink_interface"); ok {
        o.NetworkUplinkInterface = attr.(string)
    }
    if attr, ok := d.GetOk("network_uplink_interface_gateway"); ok {
        o.NetworkUplinkInterfaceGateway = attr.(string)
    }
    if attr, ok := d.GetOk("network_uplink_interface_ip"); ok {
        o.NetworkUplinkInterfaceIp = attr.(string)
    }
    if attr, ok := d.GetOk("network_uplink_interface_netmask"); ok {
        o.NetworkUplinkInterfaceNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("nfs_log_server"); ok {
        o.NfsLogServer = attr.(string)
    }
    if attr, ok := d.GetOk("nfs_mount_path"); ok {
        o.NfsMountPath = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_dns1"); ok {
        o.MgmtDNS1 = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_dns2"); ok {
        o.MgmtDNS2 = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_gateway"); ok {
        o.MgmtGateway = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_ip_address"); ok {
        o.MgmtIPAddress = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_netmask"); ok {
        o.MgmtNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("dhcp_relay_server"); ok {
        o.DhcpRelayServer = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_network_portgroup"); ok {
        o.MirrorNetworkPortgroup = attr.(string)
    }
    if attr, ok := d.GetOk("site_id"); ok {
        o.SiteId = attr.(string)
    }
    if attr, ok := d.GetOk("allow_data_dhcp"); ok {
        o.AllowDataDHCP = attr.(bool)
    }
    if attr, ok := d.GetOk("allow_mgmt_dhcp"); ok {
        o.AllowMgmtDHCP = attr.(bool)
    }
    if attr, ok := d.GetOk("flow_eviction_threshold"); ok {
        o.FlowEvictionThreshold = attr.(int)
    }
    if attr, ok := d.GetOk("toolbox_deployment_mode"); ok {
        o.ToolboxDeploymentMode = attr.(bool)
    }
    if attr, ok := d.GetOk("toolbox_group"); ok {
        o.ToolboxGroup = attr.(string)
    }
    if attr, ok := d.GetOk("toolbox_ip"); ok {
        o.ToolboxIP = attr.(string)
    }
    if attr, ok := d.GetOk("toolbox_password"); ok {
        o.ToolboxPassword = attr.(string)
    }
    if attr, ok := d.GetOk("toolbox_user_name"); ok {
        o.ToolboxUserName = attr.(string)
    }
    if attr, ok := d.GetOk("portgroup_metadata"); ok {
        o.PortgroupMetadata = attr.(bool)
    }
    if attr, ok := d.GetOk("nova_client_version"); ok {
        o.NovaClientVersion = attr.(int)
    }
    if attr, ok := d.GetOk("nova_metadata_service_auth_url"); ok {
        o.NovaMetadataServiceAuthUrl = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_endpoint"); ok {
        o.NovaMetadataServiceEndpoint = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_password"); ok {
        o.NovaMetadataServicePassword = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_tenant"); ok {
        o.NovaMetadataServiceTenant = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_username"); ok {
        o.NovaMetadataServiceUsername = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_shared_secret"); ok {
        o.NovaMetadataSharedSecret = attr.(string)
    }
    if attr, ok := d.GetOk("nova_region_name"); ok {
        o.NovaRegionName = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_package_password"); ok {
        o.UpgradePackagePassword = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_package_url"); ok {
        o.UpgradePackageURL = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_package_username"); ok {
        o.UpgradePackageUsername = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_script_time_limit"); ok {
        o.UpgradeScriptTimeLimit = attr.(int)
    }
    if attr, ok := d.GetOk("upgrade_status"); ok {
        o.UpgradeStatus = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_timedout"); ok {
        o.UpgradeTimedout = attr.(bool)
    }
    if attr, ok := d.GetOk("primary_nuage_controller"); ok {
        o.PrimaryNuageController = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_id"); ok {
        o.VrsId = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_password"); ok {
        o.VrsPassword = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_user_name"); ok {
        o.VrsUserName = attr.(string)
    }
    if attr, ok := d.GetOk("static_route"); ok {
        o.StaticRoute = attr.(string)
    }
    if attr, ok := d.GetOk("static_route_gateway"); ok {
        o.StaticRouteGateway = attr.(string)
    }
    if attr, ok := d.GetOk("static_route_netmask"); ok {
        o.StaticRouteNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("ntp_server1"); ok {
        o.NtpServer1 = attr.(string)
    }
    if attr, ok := d.GetOk("ntp_server2"); ok {
        o.NtpServer2 = attr.(string)
    }
    if attr, ok := d.GetOk("mtu"); ok {
        o.Mtu = attr.(int)
    }
    if attr, ok := d.GetOk("successfully_applied_version"); ok {
        o.SuccessfullyAppliedVersion = attr.(string)
    }
    if attr, ok := d.GetOk("multi_vmssupport"); ok {
        o.MultiVMSsupport = attr.(bool)
    }
    if attr, ok := d.GetOk("multicast_receive_interface"); ok {
        o.MulticastReceiveInterface = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_receive_interface_ip"); ok {
        o.MulticastReceiveInterfaceIP = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_receive_interface_netmask"); ok {
        o.MulticastReceiveInterfaceNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_receive_range"); ok {
        o.MulticastReceiveRange = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_send_interface"); ok {
        o.MulticastSendInterface = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_send_interface_ip"); ok {
        o.MulticastSendInterfaceIP = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_send_interface_netmask"); ok {
        o.MulticastSendInterfaceNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_source_portgroup"); ok {
        o.MulticastSourcePortgroup = attr.(string)
    }
    if attr, ok := d.GetOk("customized_script_url"); ok {
        o.CustomizedScriptURL = attr.(string)
    }
    if attr, ok := d.GetOk("available_networks"); ok {
        o.AvailableNetworks = attr.([]interface{})
    }
    if attr, ok := d.GetOk("ovf_url"); ok {
        o.OvfURL = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_vcenter_data_center"); ok {
        parent := &vspk.VCenterDataCenter{ID: attr.(string)}
        err := parent.CreateVCenterHypervisor(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vcenter_cluster"); ok {
        parent := &vspk.VCenterCluster{ID: attr.(string)}
        err := parent.CreateVCenterHypervisor(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceVCenterHypervisorRead(d, m)
}

func resourceVCenterHypervisorRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VCenterHypervisor{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("vcenter_ip", o.VCenterIP)
    d.Set("vcenter_password", o.VCenterPassword)
    d.Set("vcenter_user", o.VCenterUser)
    d.Set("vrs_configuration_time_limit", o.VRSConfigurationTimeLimit)
    d.Set("vrs_metrics_id", o.VRSMetricsID)
    d.Set("vrs_state", o.VRSState)
    d.Set("v_require_nuage_metadata", o.VRequireNuageMetadata)
    d.Set("name", o.Name)
    d.Set("managed_object_id", o.ManagedObjectID)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("last_vrs_deployed_date", o.LastVRSDeployedDate)
    d.Set("data_dns1", o.DataDNS1)
    d.Set("data_dns2", o.DataDNS2)
    d.Set("data_gateway", o.DataGateway)
    d.Set("data_ip_address", o.DataIPAddress)
    d.Set("data_netmask", o.DataNetmask)
    d.Set("data_network_portgroup", o.DataNetworkPortgroup)
    d.Set("datapath_sync_timeout", o.DatapathSyncTimeout)
    d.Set("scope", o.Scope)
    d.Set("secondary_nuage_controller", o.SecondaryNuageController)
    d.Set("removed_from_vcenter_inventory", o.RemovedFromVCenterInventory)
    d.Set("generic_split_activation", o.GenericSplitActivation)
    d.Set("separate_data_network", o.SeparateDataNetwork)
    d.Set("deployment_count", o.DeploymentCount)
    d.Set("personality", o.Personality)
    d.Set("description", o.Description)
    d.Set("destination_mirror_port", o.DestinationMirrorPort)
    d.Set("metadata_server_ip", o.MetadataServerIP)
    d.Set("metadata_server_listen_port", o.MetadataServerListenPort)
    d.Set("metadata_server_port", o.MetadataServerPort)
    d.Set("metadata_service_enabled", o.MetadataServiceEnabled)
    d.Set("network_uplink_interface", o.NetworkUplinkInterface)
    d.Set("network_uplink_interface_gateway", o.NetworkUplinkInterfaceGateway)
    d.Set("network_uplink_interface_ip", o.NetworkUplinkInterfaceIp)
    d.Set("network_uplink_interface_netmask", o.NetworkUplinkInterfaceNetmask)
    d.Set("nfs_log_server", o.NfsLogServer)
    d.Set("nfs_mount_path", o.NfsMountPath)
    d.Set("mgmt_dns1", o.MgmtDNS1)
    d.Set("mgmt_dns2", o.MgmtDNS2)
    d.Set("mgmt_gateway", o.MgmtGateway)
    d.Set("mgmt_ip_address", o.MgmtIPAddress)
    d.Set("mgmt_netmask", o.MgmtNetmask)
    d.Set("mgmt_network_portgroup", o.MgmtNetworkPortgroup)
    d.Set("dhcp_relay_server", o.DhcpRelayServer)
    d.Set("mirror_network_portgroup", o.MirrorNetworkPortgroup)
    d.Set("site_id", o.SiteId)
    d.Set("allow_data_dhcp", o.AllowDataDHCP)
    d.Set("allow_mgmt_dhcp", o.AllowMgmtDHCP)
    d.Set("flow_eviction_threshold", o.FlowEvictionThreshold)
    d.Set("vm_network_portgroup", o.VmNetworkPortgroup)
    d.Set("entity_scope", o.EntityScope)
    d.Set("toolbox_deployment_mode", o.ToolboxDeploymentMode)
    d.Set("toolbox_group", o.ToolboxGroup)
    d.Set("toolbox_ip", o.ToolboxIP)
    d.Set("toolbox_password", o.ToolboxPassword)
    d.Set("toolbox_user_name", o.ToolboxUserName)
    d.Set("portgroup_metadata", o.PortgroupMetadata)
    d.Set("nova_client_version", o.NovaClientVersion)
    d.Set("nova_metadata_service_auth_url", o.NovaMetadataServiceAuthUrl)
    d.Set("nova_metadata_service_endpoint", o.NovaMetadataServiceEndpoint)
    d.Set("nova_metadata_service_password", o.NovaMetadataServicePassword)
    d.Set("nova_metadata_service_tenant", o.NovaMetadataServiceTenant)
    d.Set("nova_metadata_service_username", o.NovaMetadataServiceUsername)
    d.Set("nova_metadata_shared_secret", o.NovaMetadataSharedSecret)
    d.Set("nova_region_name", o.NovaRegionName)
    d.Set("upgrade_package_password", o.UpgradePackagePassword)
    d.Set("upgrade_package_url", o.UpgradePackageURL)
    d.Set("upgrade_package_username", o.UpgradePackageUsername)
    d.Set("upgrade_script_time_limit", o.UpgradeScriptTimeLimit)
    d.Set("upgrade_status", o.UpgradeStatus)
    d.Set("upgrade_timedout", o.UpgradeTimedout)
    d.Set("primary_nuage_controller", o.PrimaryNuageController)
    d.Set("vrs_id", o.VrsId)
    d.Set("vrs_password", o.VrsPassword)
    d.Set("vrs_user_name", o.VrsUserName)
    d.Set("static_route", o.StaticRoute)
    d.Set("static_route_gateway", o.StaticRouteGateway)
    d.Set("static_route_netmask", o.StaticRouteNetmask)
    d.Set("ntp_server1", o.NtpServer1)
    d.Set("ntp_server2", o.NtpServer2)
    d.Set("mtu", o.Mtu)
    d.Set("successfully_applied_version", o.SuccessfullyAppliedVersion)
    d.Set("multi_vmssupport", o.MultiVMSsupport)
    d.Set("multicast_receive_interface", o.MulticastReceiveInterface)
    d.Set("multicast_receive_interface_ip", o.MulticastReceiveInterfaceIP)
    d.Set("multicast_receive_interface_netmask", o.MulticastReceiveInterfaceNetmask)
    d.Set("multicast_receive_range", o.MulticastReceiveRange)
    d.Set("multicast_send_interface", o.MulticastSendInterface)
    d.Set("multicast_send_interface_ip", o.MulticastSendInterfaceIP)
    d.Set("multicast_send_interface_netmask", o.MulticastSendInterfaceNetmask)
    d.Set("multicast_source_portgroup", o.MulticastSourcePortgroup)
    d.Set("customized_script_url", o.CustomizedScriptURL)
    d.Set("available_networks", o.AvailableNetworks)
    d.Set("ovf_url", o.OvfURL)
    d.Set("external_id", o.ExternalID)
    d.Set("hypervisor_ip", o.HypervisorIP)
    d.Set("hypervisor_password", o.HypervisorPassword)
    d.Set("hypervisor_user", o.HypervisorUser)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceVCenterHypervisorUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VCenterHypervisor{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.DataNetworkPortgroup = d.Get("data_network_portgroup").(string)
    o.Description = d.Get("description").(string)
    o.MgmtNetworkPortgroup = d.Get("mgmt_network_portgroup").(string)
    o.VmNetworkPortgroup = d.Get("vm_network_portgroup").(string)
    o.HypervisorIP = d.Get("hypervisor_ip").(string)
    o.HypervisorPassword = d.Get("hypervisor_password").(string)
    o.HypervisorUser = d.Get("hypervisor_user").(string)
    
    if attr, ok := d.GetOk("vcenter_ip"); ok {
        o.VCenterIP = attr.(string)
    }
    if attr, ok := d.GetOk("vcenter_password"); ok {
        o.VCenterPassword = attr.(string)
    }
    if attr, ok := d.GetOk("vcenter_user"); ok {
        o.VCenterUser = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_configuration_time_limit"); ok {
        o.VRSConfigurationTimeLimit = attr.(int)
    }
    if attr, ok := d.GetOk("vrs_metrics_id"); ok {
        o.VRSMetricsID = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_state"); ok {
        o.VRSState = attr.(string)
    }
    if attr, ok := d.GetOk("v_require_nuage_metadata"); ok {
        o.VRequireNuageMetadata = attr.(bool)
    }
    if attr, ok := d.GetOk("managed_object_id"); ok {
        o.ManagedObjectID = attr.(string)
    }
    if attr, ok := d.GetOk("last_vrs_deployed_date"); ok {
        o.LastVRSDeployedDate = attr.(float64)
    }
    if attr, ok := d.GetOk("data_dns1"); ok {
        o.DataDNS1 = attr.(string)
    }
    if attr, ok := d.GetOk("data_dns2"); ok {
        o.DataDNS2 = attr.(string)
    }
    if attr, ok := d.GetOk("data_gateway"); ok {
        o.DataGateway = attr.(string)
    }
    if attr, ok := d.GetOk("data_ip_address"); ok {
        o.DataIPAddress = attr.(string)
    }
    if attr, ok := d.GetOk("data_netmask"); ok {
        o.DataNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("datapath_sync_timeout"); ok {
        o.DatapathSyncTimeout = attr.(int)
    }
    if attr, ok := d.GetOk("scope"); ok {
        o.Scope = attr.(bool)
    }
    if attr, ok := d.GetOk("secondary_nuage_controller"); ok {
        o.SecondaryNuageController = attr.(string)
    }
    if attr, ok := d.GetOk("removed_from_vcenter_inventory"); ok {
        o.RemovedFromVCenterInventory = attr.(bool)
    }
    if attr, ok := d.GetOk("generic_split_activation"); ok {
        o.GenericSplitActivation = attr.(bool)
    }
    if attr, ok := d.GetOk("separate_data_network"); ok {
        o.SeparateDataNetwork = attr.(bool)
    }
    if attr, ok := d.GetOk("deployment_count"); ok {
        o.DeploymentCount = attr.(int)
    }
    if attr, ok := d.GetOk("personality"); ok {
        o.Personality = attr.(string)
    }
    if attr, ok := d.GetOk("destination_mirror_port"); ok {
        o.DestinationMirrorPort = attr.(string)
    }
    if attr, ok := d.GetOk("metadata_server_ip"); ok {
        o.MetadataServerIP = attr.(string)
    }
    if attr, ok := d.GetOk("metadata_server_listen_port"); ok {
        o.MetadataServerListenPort = attr.(int)
    }
    if attr, ok := d.GetOk("metadata_server_port"); ok {
        o.MetadataServerPort = attr.(int)
    }
    if attr, ok := d.GetOk("metadata_service_enabled"); ok {
        o.MetadataServiceEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("network_uplink_interface"); ok {
        o.NetworkUplinkInterface = attr.(string)
    }
    if attr, ok := d.GetOk("network_uplink_interface_gateway"); ok {
        o.NetworkUplinkInterfaceGateway = attr.(string)
    }
    if attr, ok := d.GetOk("network_uplink_interface_ip"); ok {
        o.NetworkUplinkInterfaceIp = attr.(string)
    }
    if attr, ok := d.GetOk("network_uplink_interface_netmask"); ok {
        o.NetworkUplinkInterfaceNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("nfs_log_server"); ok {
        o.NfsLogServer = attr.(string)
    }
    if attr, ok := d.GetOk("nfs_mount_path"); ok {
        o.NfsMountPath = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_dns1"); ok {
        o.MgmtDNS1 = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_dns2"); ok {
        o.MgmtDNS2 = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_gateway"); ok {
        o.MgmtGateway = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_ip_address"); ok {
        o.MgmtIPAddress = attr.(string)
    }
    if attr, ok := d.GetOk("mgmt_netmask"); ok {
        o.MgmtNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("dhcp_relay_server"); ok {
        o.DhcpRelayServer = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_network_portgroup"); ok {
        o.MirrorNetworkPortgroup = attr.(string)
    }
    if attr, ok := d.GetOk("site_id"); ok {
        o.SiteId = attr.(string)
    }
    if attr, ok := d.GetOk("allow_data_dhcp"); ok {
        o.AllowDataDHCP = attr.(bool)
    }
    if attr, ok := d.GetOk("allow_mgmt_dhcp"); ok {
        o.AllowMgmtDHCP = attr.(bool)
    }
    if attr, ok := d.GetOk("flow_eviction_threshold"); ok {
        o.FlowEvictionThreshold = attr.(int)
    }
    if attr, ok := d.GetOk("toolbox_deployment_mode"); ok {
        o.ToolboxDeploymentMode = attr.(bool)
    }
    if attr, ok := d.GetOk("toolbox_group"); ok {
        o.ToolboxGroup = attr.(string)
    }
    if attr, ok := d.GetOk("toolbox_ip"); ok {
        o.ToolboxIP = attr.(string)
    }
    if attr, ok := d.GetOk("toolbox_password"); ok {
        o.ToolboxPassword = attr.(string)
    }
    if attr, ok := d.GetOk("toolbox_user_name"); ok {
        o.ToolboxUserName = attr.(string)
    }
    if attr, ok := d.GetOk("portgroup_metadata"); ok {
        o.PortgroupMetadata = attr.(bool)
    }
    if attr, ok := d.GetOk("nova_client_version"); ok {
        o.NovaClientVersion = attr.(int)
    }
    if attr, ok := d.GetOk("nova_metadata_service_auth_url"); ok {
        o.NovaMetadataServiceAuthUrl = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_endpoint"); ok {
        o.NovaMetadataServiceEndpoint = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_password"); ok {
        o.NovaMetadataServicePassword = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_tenant"); ok {
        o.NovaMetadataServiceTenant = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_service_username"); ok {
        o.NovaMetadataServiceUsername = attr.(string)
    }
    if attr, ok := d.GetOk("nova_metadata_shared_secret"); ok {
        o.NovaMetadataSharedSecret = attr.(string)
    }
    if attr, ok := d.GetOk("nova_region_name"); ok {
        o.NovaRegionName = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_package_password"); ok {
        o.UpgradePackagePassword = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_package_url"); ok {
        o.UpgradePackageURL = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_package_username"); ok {
        o.UpgradePackageUsername = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_script_time_limit"); ok {
        o.UpgradeScriptTimeLimit = attr.(int)
    }
    if attr, ok := d.GetOk("upgrade_status"); ok {
        o.UpgradeStatus = attr.(string)
    }
    if attr, ok := d.GetOk("upgrade_timedout"); ok {
        o.UpgradeTimedout = attr.(bool)
    }
    if attr, ok := d.GetOk("primary_nuage_controller"); ok {
        o.PrimaryNuageController = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_id"); ok {
        o.VrsId = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_password"); ok {
        o.VrsPassword = attr.(string)
    }
    if attr, ok := d.GetOk("vrs_user_name"); ok {
        o.VrsUserName = attr.(string)
    }
    if attr, ok := d.GetOk("static_route"); ok {
        o.StaticRoute = attr.(string)
    }
    if attr, ok := d.GetOk("static_route_gateway"); ok {
        o.StaticRouteGateway = attr.(string)
    }
    if attr, ok := d.GetOk("static_route_netmask"); ok {
        o.StaticRouteNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("ntp_server1"); ok {
        o.NtpServer1 = attr.(string)
    }
    if attr, ok := d.GetOk("ntp_server2"); ok {
        o.NtpServer2 = attr.(string)
    }
    if attr, ok := d.GetOk("mtu"); ok {
        o.Mtu = attr.(int)
    }
    if attr, ok := d.GetOk("successfully_applied_version"); ok {
        o.SuccessfullyAppliedVersion = attr.(string)
    }
    if attr, ok := d.GetOk("multi_vmssupport"); ok {
        o.MultiVMSsupport = attr.(bool)
    }
    if attr, ok := d.GetOk("multicast_receive_interface"); ok {
        o.MulticastReceiveInterface = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_receive_interface_ip"); ok {
        o.MulticastReceiveInterfaceIP = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_receive_interface_netmask"); ok {
        o.MulticastReceiveInterfaceNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_receive_range"); ok {
        o.MulticastReceiveRange = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_send_interface"); ok {
        o.MulticastSendInterface = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_send_interface_ip"); ok {
        o.MulticastSendInterfaceIP = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_send_interface_netmask"); ok {
        o.MulticastSendInterfaceNetmask = attr.(string)
    }
    if attr, ok := d.GetOk("multicast_source_portgroup"); ok {
        o.MulticastSourcePortgroup = attr.(string)
    }
    if attr, ok := d.GetOk("customized_script_url"); ok {
        o.CustomizedScriptURL = attr.(string)
    }
    if attr, ok := d.GetOk("available_networks"); ok {
        o.AvailableNetworks = attr.([]interface{})
    }
    if attr, ok := d.GetOk("ovf_url"); ok {
        o.OvfURL = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceVCenterHypervisorDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VCenterHypervisor{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}