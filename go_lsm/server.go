package main

type CreateServerRequest struct {
	DisplayName       string `json:"serverDisplayName"`
	IPv4Address       string `json:"ipv4Address"`
	IPv4AddressSubnet string `json:"ipv4AddressSubnet"`
	IPv6Address       string `json:"ipv6Address"`
	IPv6AddressSubnet string `json:"ipv6AddressSubnet"`
	RootPassword      string `json:"rootPassword"`
	AdminUsername     string `json:"adminUsername"`
	AdminUserPassword string `json:"adminUserPassword"`
	AdminUserSSHKey   string `json:"adminUserSSHKey"`
	RootUserSSHKey    string `json:"rootUserSSHKey"`
}

func (csr *CreateServerRequest) Create() {
	
}
