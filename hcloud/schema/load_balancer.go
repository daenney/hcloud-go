package schema

import "time"

// LoadBalancer represents a Load Balancer in the Hetzner Cloud.
type LoadBalancer struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	PublicNet        LoadBalancerPublicNet    `json:"public_net"`
	PrivateNet       []LoadBalancerPrivateNet `json:"private_net"`
	Location         Location                 `json:"location"`
	LoadBalancerType LoadBalancerType         `json:"load_balancer_type"`
	Protection       LoadBalancerProtection   `json:"protection"`
	Labels           map[string]string        `json:"labels"`
	Created          time.Time                `json:"created"`
	Services         []LoadBalancerService    `json:"services"`
	Targets          []LoadBalancerTarget     `json:"targets"`
	Algorithm        LoadBalancerAlgorithm    `json:"algorithm"`
}

// LoadBalancerPublicNet defines the schema of a Load Balancers
// public network information.
type LoadBalancerPublicNet struct {
	Enabled bool                      `json:"enabled"`
	IPv4    LoadBalancerPublicNetIPv4 `json:"ipv4"`
	IPv6    LoadBalancerPublicNetIPv6 `json:"ipv6"`
}

// LoadBalancerPublicNetIPv4 defines the schema of a Load Balancers public
// network information for an IPv4.
type LoadBalancerPublicNetIPv4 struct {
	IP string `json:"ip"`
}

// LoadBalancerPublicNetIPv6 defines the schema of a Load Balancers public
// network information for an IPv6.
type LoadBalancerPublicNetIPv6 struct {
	IP string `json:"ip"`
}

// LoadBalancerPrivateNet defines the schema of a Load Balancers private network information.
type LoadBalancerPrivateNet struct {
	Network int    `json:"network"`
	IP      string `json:"ip"`
}

// LoadBalancerAlgorithm represents the algorithm of a Load Balancer.
type LoadBalancerAlgorithm struct {
	Type string `json:"type"`
}

// LoadBalancerProtection represents the protection level of a Load Balancer.
type LoadBalancerProtection struct {
	Delete bool `json:"delete"`
}

// LoadBalancerService represents a service of a Load Balancer.
type LoadBalancerService struct {
	Protocol        string                          `json:"protocol"`
	ListenPort      int                             `json:"listen_port,omitempty"`
	DestinationPort int                             `json:"destination_port,omitempty"`
	Proxyprotocol   bool                            `json:"proxyprotocol,omitempty"`
	HTTP            *LoadBalancerServiceHTTP        `json:"http,omitempty"`
	HealthCheck     *LoadBalancerServiceHealthCheck `json:"health_check,omitempty"`
}

// LoadBalancerServiceHTTP represents the http configuration for a LoadBalancerService.
type LoadBalancerServiceHTTP struct {
	CookieName     string `json:"cookie_name,omitempty"`
	CookieLifetime int    `json:"cookie_lifetime,omitempty"`
	Certificates   []int  `json:"certificates,omitempty"`
	RedirectHTTP   bool   `json:"redirect_http,omitempty"`
	StickySessions bool   `json:"sticky_sessions,omitempty"`
}

// LoadBalancerServiceHealthCheck represents a service health check configuration.
type LoadBalancerServiceHealthCheck struct {
	Protocol string                              `json:"protocol,omitempty"`
	Port     int                                 `json:"port,omitempty"`
	Interval int                                 `json:"interval,omitempty"`
	Timeout  int                                 `json:"timeout,omitempty"`
	Retries  int                                 `json:"retries,omitempty"`
	HTTP     *LoadBalancerServiceHealthCheckHTTP `json:"http,omitempty"`
}

// LoadBalancerServiceHealthCheckHTTP represents a http health check configuration.
type LoadBalancerServiceHealthCheckHTTP struct {
	Domain      string   `json:"domain,omitempty"`
	Path        string   `json:"path,omitempty"`
	Response    string   `json:"response,omitempty"`
	StatusCodes []string `json:"status_codes,omitempty"`
	TLS         bool     `json:"tls,omitempty"`
}

// LoadBalancerTarget represents a target of a Load Balancer.
type LoadBalancerTarget struct {
	Type         string                           `json:"type"`
	Server       *LoadBalancerTargetServer        `json:"server,omitempty"`
	HealthStatus []LoadBalancerTargetHealthStatus `json:"health_status,omitempty"`
	UsePrivateIP bool                             `json:"use_private_ip,omitempty"`
}

// LoadBalancerTargetHealthStatus represents a health status of target of a Load Balancer.
type LoadBalancerTargetHealthStatus struct {
	ListenPort int    `json:"listen_port"`
	Status     string `json:"status"`
}

// LoadBalancerTargetServer represents a server target of a Load Balancer.
type LoadBalancerTargetServer struct {
	ID int `json:"id"`
}

// LoadBalancerListResponse defines the schema of the response when
// listing Load Balancer.
type LoadBalancerListResponse struct {
	LoadBalancers []LoadBalancer `json:"load_balancers"`
}

// LoadBalancerGetResponse defines the schema of the response when
// retrieving a single Load Balancer.
type LoadBalancerGetResponse struct {
	LoadBalancer LoadBalancer `json:"load_balancer"`
}

// LoadBalancerActionAddTargetRequest defines the schema of the request to
// add a target to a Load Balancer.
type LoadBalancerActionAddTargetRequest struct {
	Type         string                    `json:"type"`
	Server       *LoadBalancerTargetServer `json:"server,omitempty"`
	UsePrivateIP bool                      `json:"use_private_ip"`
}

// LoadBalancerActionRemoveTargetRequest defines the schema of the request to
// remove a target from a Load Balancer.
type LoadBalancerActionRemoveTargetRequest struct {
	Type   string                    `json:"type"`
	Server *LoadBalancerTargetServer `json:"server,omitempty"`
}

// LoadBalancerActionTargetResponse defines the schema of the response when
// adding or removing a target from a Load Balancer.
type LoadBalancerActionTargetResponse struct {
	Action Action `json:"action"`
}

// LoadBalancerActionAddServiceRequest defines the schema of the request to
// adding a service to a Load Balancer.
type LoadBalancerActionAddServiceRequest struct {
	Protocol        string                          `json:"protocol"`
	ListenPort      int                             `json:"listen_port,omitempty"`
	DestinationPort int                             `json:"destination_port,omitempty"`
	ProxyProtocol   *bool                           `json:"proxyprotocol,omitempty"`
	HTTP            *LoadBalancerServiceHTTP        `json:"http,omitempty"`
	HealthCheck     *LoadBalancerServiceHealthCheck `json:"health_check,omitempty"`
}

// LoadBalancerActionAddServiceResponse defines the schema of the response when
// creating a add service action.
type LoadBalancerActionAddServiceResponse struct {
	Action Action `json:"action"`
}

// LoadBalancerActionUpdateServiceRequest defines the schema of the request to
// update a service from a Load Balancer.
type LoadBalancerActionUpdateServiceRequest struct {
	Protocol        string                          `json:"protocol,omitempty"`
	ListenPort      int                             `json:"listen_port"`
	DestinationPort *int                            `json:"destination_port,omitempty"`
	ProxyProtocol   *bool                           `json:"proxyprotocol,omitempty"`
	HTTP            *LoadBalancerUpdateServiceHTTP  `json:"http,omitempty"`
	HealthCheck     *LoadBalancerServiceHealthCheck `json:"health_check,omitempty"`
}

// LoadBalancerUpdateServiceHTTP represents the http configuration for a LoadBalancerActionUpdateServiceRequest.
type LoadBalancerUpdateServiceHTTP struct {
	CookieName     string `json:"cookie_name,omitempty"`
	CookieLifetime int    `json:"cookie_lifetime,omitempty"`
	Certificates   []int  `json:"certificates,omitempty"`
	RedirectHTTP   *bool  `json:"redirect_http,omitempty"`
	StickySessions *bool  `json:"sticky_sessions,omitempty"`
}

// LoadBalancerActionUpdateServiceResponse defines the schema of the response when
// creating a update service action.
type LoadBalancerActionUpdateServiceResponse struct {
	Action Action `json:"action"`
}

// LoadBalancerDeleteServiceRequest defines the schema of the request to
// delete a service from a Load Balancer.
type LoadBalancerDeleteServiceRequest struct {
	ListenPort int `json:"listen_port"`
}

// LoadBalancerDeleteServiceResponse defines the schema of the response when
// creating a delete_service action.
type LoadBalancerDeleteServiceResponse struct {
	Action Action `json:"action"`
}

// LoadBalancerCreateRequest defines the schema of the request to create a LoadBalancer.
type LoadBalancerCreateRequest struct {
	Name             string                 `json:"name"`
	LoadBalancerType interface{}            `json:"load_balancer_type"` // int or string
	Algorithm        *LoadBalancerAlgorithm `json:"algorithm,omitempty"`
	Location         string                 `json:"location,omitempty"`
	NetworkZone      string                 `json:"network_zone,omitempty"`
	Labels           *map[string]string     `json:"labels,omitempty"`
	Targets          []LoadBalancerTarget   `json:"targets,omitempty"`
	Services         []LoadBalancerService  `json:"services,omitempty"`
	PublicInterface  *bool                  `json:"public_interface,omitempty"`
	Network          *int                   `json:"network,omitempty"`
}

// LoadBalancerCreateResponse defines the schema of the response to
// create a LoadBalancer.
type LoadBalancerCreateResponse struct {
	LoadBalancer LoadBalancer `json:"load_balancer"`
	Action       Action       `json:"action"`
}

// LoadBalancerActionChangeProtectionRequest defines the schema of the request to
// change the resource protection of a load balancer.
type LoadBalancerActionChangeProtectionRequest struct {
	Delete *bool `json:"delete,omitempty"`
}

// LoadBalancerActionChangeProtectionResponse defines the schema of the response when
// changing the resource protection of a load balancer.
type LoadBalancerActionChangeProtectionResponse struct {
	Action Action `json:"action"`
}

// LoadBalancerUpdateRequest defines the schema of the request to update a load balancer.
type LoadBalancerUpdateRequest struct {
	Name   string             `json:"name,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
}

// LoadBalancerUpdateResponse defines the schema of the response when updating a load balancer.
type LoadBalancerUpdateResponse struct {
	LoadBalancer LoadBalancer `json:"load_balancer"`
}

// LoadBalancerActionChangeAlgorithmRequest defines the schema of the request to
// change the algorithm of a load balancer.
type LoadBalancerActionChangeAlgorithmRequest struct {
	Type string `json:"type"`
}

// LoadBalancerActionChangeAlgorithmResponse defines the schema of the response when
// changing the algorithm of a load balancer.
type LoadBalancerActionChangeAlgorithmResponse struct {
	Action Action `json:"action"`
}

// LoadBalancerActionAttachToNetworkRequest defines the schema for the request to
// attach a network to a Load Balancer.
type LoadBalancerActionAttachToNetworkRequest struct {
	Network int     `json:"network"`
	IP      *string `json:"ip,omitempty"`
}

// LoadBalancerActionAttachToNetworkResponse defines the schema of the response when
// creating an attach_to_network Load Balancer action.
type LoadBalancerActionAttachToNetworkResponse struct {
	Action Action `json:"action"`
}

// LoadBalancerActionDetachFromNetworkRequest defines the schema for the request to
// detach a network from a Load Balancer.
type LoadBalancerActionDetachFromNetworkRequest struct {
	Network int `json:"network"`
}

// LoadBalancerActionDetachFromNetworkResponse defines the schema of the response when
// creating a detach_from_network Load Balancer action.
type LoadBalancerActionDetachFromNetworkResponse struct {
	Action Action `json:"action"`
}

// LoadBalancerActionEnablePublicInterfaceRequest defines the schema for the request to
// enable the public interface of a Load Balancer.
type LoadBalancerActionEnablePublicInterfaceRequest struct{}

// LoadBalancerActionEnablePublicInterfaceResponse defines the schema of the response when
// creating a enable_public_interface Load Balancer action.
type LoadBalancerActionEnablePublicInterfaceResponse struct {
	Action Action `json:"action"`
}

// LoadBalancerActionDisablePublicInterfaceRequest defines the schema for the request to
// disable the public interface of a Load Balancer.
type LoadBalancerActionDisablePublicInterfaceRequest struct{}

// LoadBalancerActionDisablePublicInterfaceResponse defines the schema of the response when
// creating a disable_public_interface Load Balancer action.
type LoadBalancerActionDisablePublicInterfaceResponse struct {
	Action Action `json:"action"`
}