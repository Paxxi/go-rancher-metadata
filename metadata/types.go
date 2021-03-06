package metadata

type Stack struct {
	EnvironmentName string    `json:"environment_name"`
	EnvironmentUUID string    `json:"environment_uuid"`
	Name            string    `json:"name"`
	Services        []Service `json:"services"`
}

type HealthCheck struct {
	HealthyThreshold   int    `json:"healthy_threshold"`
	Interval           int    `json:"interval"`
	Port               int    `json:"port"`
	RequestLine        string `json:"request_line"`
	ResponseTimeout    int    `json:"response_timeout"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
}

type Service struct {
	Scale              int                    `json:"scale"`
	Name               string                 `json:"name"`
	StackName          string                 `json:"stack_name"`
	StackUUID          string                 `json:"stack_uuid"`
	Kind               string                 `json:"kind"`
	Hostname           string                 `json:"hostname"`
	Vip                string                 `json:"vip"`
	CreateIndex        int                    `json:"create_index"`
	UUID               string                 `json:"uuid"`
	ExternalIps        []string               `json:"external_ips"`
	Sidekicks          []string               `json:"sidekicks"`
	Containers         []Container            `json:"containers"`
	Ports              []string               `json:"ports"`
	Labels             map[string]string      `json:"labels"`
	Links              map[string]string      `json:"links"`
	Metadata           map[string]interface{} `json:"metadata"`
	Token              string                 `json:"token"`
	Fqdn               string                 `json:"fqdn"`
	HealthCheck        HealthCheck            `json:"health_check"`
	PrimaryServiceName string                 `json:"primary_service_name"`
	LBConfig           LBConfig               `json:"lb_config"`
}

type Container struct {
	Name                     string            `json:"name"`
	PrimaryIp                string            `json:"primary_ip"`
	PrimaryMacAddress        string            `json:"primary_mac_address"`
	Ips                      []string          `json:"ips"`
	Ports                    []string          `json:"ports"`
	ServiceName              string            `json:"service_name"`
	StackName                string            `json:"stack_name"`
	Labels                   map[string]string `json:"labels"`
	CreateIndex              int               `json:"create_index"`
	HostUUID                 string            `json:"host_uuid"`
	UUID                     string            `json:"uuid"`
	State                    string            `json:"state"`
	HealthState              string            `json:"health_state"`
	ExternalId               string            `json:"external_id"`
	StartCount               int               `json:"start_count"`
	MemoryReservation        int64             `json:"memory_reservation"`
	MilliCPUReservation      int64             `json:"milli_cpu_reservation"`
	Dns                      []string          `json:"dns"`
	DnsSearch                []string          `json:"dns_search"`
	HealthCheckHosts         []string          `json:"health_check_hosts"`
	NetworkFromContainerUUID string            `json:"network_from_container_uuid"`
	NetworkUUID              string            `json:"network_uuid"`
}

type Network struct {
	Name      string                 `json:"name"`
	UUID      string                 `json:"uuid"`
	Metadata  map[string]interface{} `json:"metadata"`
	HostPorts bool                   `json:"host_ports"`
	Default   bool                   `json:"is_default"`
}

type Host struct {
	Name           string            `json:"name"`
	AgentIP        string            `json:"agent_ip"`
	HostId         int               `json:"host_id"`
	Labels         map[string]string `json:"labels"`
	UUID           string            `json:"uuid"`
	Hostname       string            `json:"hostname"`
	Memory         int64             `json:"memory"`
	MilliCPU       int64             `json:"milli_cpu"`
	LocalStorageMb int64             `json:"local_storage_mb"`
}

type PortRule struct {
	SourcePort  int    `json:"source_port"`
	Protocol    string `json:"protocol"`
	Path        string `json:"path"`
	Hostname    string `json:"hostname"`
	Service     string `json:"service"`
	TargetPort  int    `json:"target_port"`
	Priority    int    `json:"priority"`
	BackendName string `json:"backend_name"`
	Selector    string `json:"selector"`
}

type LBConfig struct {
	Certs            []string           `json:"certs"`
	DefaultCert      string             `json:"default_cert"`
	PortRules        []PortRule         `json:"port_rules"`
	Config           string             `json:"config"`
	StickinessPolicy LBStickinessPolicy `json:"stickiness_policy"`
}

type LBStickinessPolicy struct {
	Name     string `json:"name"`
	Cookie   string `json:"cookie"`
	Domain   string `json:"domain"`
	Indirect bool   `json:"indirect"`
	Nocache  bool   `json:"nocache"`
	Postonly bool   `json:"postonly"`
	Mode     string `json:"mode"`
}
