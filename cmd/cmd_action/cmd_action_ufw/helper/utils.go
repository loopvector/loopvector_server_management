package helper

import (
	"errors"
	"fmt"
	"loopvector_server_management/controller"
	"loopvector_server_management/model"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// type UfwTrafficPolicy struct {
// 	Policy string
// }

// const (
// 	UfwTrafficPolicyAllow = "allow"
// 	UfwTrafficPolicyDeny  = "deny"
// )

var (
	ip          string
	ips         []string
	ipWithPort  string
	ipsWithPort []string
	port        string
	ports       []string
	protocol    string
)

func RunUfwTrafficPolicyCommandE(cmd *cobra.Command, args []string, trafficPolicy controller.UfwTrafficPolicy) error {
	serverName := model.ServerNameModel{Name: args[0]}
	if ip != "" {
		controller.SetUfwIpAddresses(
			serverName,
			GetUfwIpAddressesTrafficPolicyRequest([]string{ip}, protocol),
			trafficPolicy,
		)
	} else if len(ips) != 0 {
		controller.SetUfwIpAddresses(
			serverName,
			GetUfwIpAddressesTrafficPolicyRequest(ips, protocol),
			trafficPolicy,
		)
	} else if ipWithPort != "" {
		controller.SetUfwIpAddressesWithPort(
			serverName,
			GetUfwIpAddressesWithPortTrafficPolicyRequest([]string{ipWithPort}, protocol),
			trafficPolicy,
		)
	} else if len(ipsWithPort) != 0 {
		controller.SetUfwIpAddressesWithPort(
			serverName,
			GetUfwIpAddressesWithPortTrafficPolicyRequest(ipsWithPort, protocol),
			trafficPolicy,
		)
	} else if port != "" {
		controller.SetUfwPorts(
			serverName,
			GetUfwPortsTrafficPolicyRequest([]string{port}, protocol),
			trafficPolicy,
		)
	} else if len(ports) != 0 {
		controller.SetUfwPorts(
			serverName,
			GetUfwPortsTrafficPolicyRequest(ports, protocol),
			trafficPolicy,
		)
	} else {
		return errors.New("no ip/ip:port/port provided")
	}
	return nil
}

func InitUfwTrafficPolicyCommandFlags(command *cobra.Command) {
	command.Flags().StringVar(&ip, "ip", "", "ip address to allow in ufw")
	command.Flags().StringVar(&ipWithPort, "ipWithPort", "", "ip address with ports to allow in ufw")
	command.Flags().StringVar(&port, "port", "", "port to allow in ufw")
	command.Flags().StringVar(&protocol, "protocol", "", "protocol of the ip/ip:port/port to allow in ufw")

	command.Flags().StringSliceVar(&ips, "ips", []string{}, "ip addresses to allow in ufw")
	command.Flags().StringSliceVar(&ipsWithPort, "ipsWithPort", []string{}, "ip addresses with port to allow in ufw")
	command.Flags().StringSliceVar(&ports, "ports", []string{}, "ports to allow in ufw")

	command.MarkFlagsMutuallyExclusive("ip", "ips", "ipWithPort", "ipsWithPort", "port", "ports")
}

func GetUfwIpAddressesTrafficPolicyRequest(ips []string, protocol string) []controller.UfwIpAddressesTrafficPolicyRequest {
	allIps := []controller.UfwIpAddressesTrafficPolicyRequest{}
	for _, oneIp := range ips {
		allIps = append(
			allIps,
			controller.UfwIpAddressesTrafficPolicyRequest{
				Ip:       oneIp,
				Protocol: protocol,
			},
		)
	}
	return allIps
}

func GetUfwIpAddressesWithPortTrafficPolicyRequest(ipsWithPort []string, protocol string) []controller.UfwIpAddressesWithPortTrafficPolicyRequest {
	allIpsWithPort := []controller.UfwIpAddressesWithPortTrafficPolicyRequest{}
	for _, oneIpWithPort := range ipsWithPort {
		parts := strings.Split(oneIpWithPort, ":")
		numPort, err := strconv.ParseUint(parts[1], 10, 32)
		if err != nil {
			fmt.Println("Error:", err)
			panic(err)
		}
		numUint32Port := uint32(numPort)
		allIpsWithPort = append(
			allIpsWithPort,
			controller.UfwIpAddressesWithPortTrafficPolicyRequest{
				Ip:       parts[0],
				Port:     numUint32Port,
				Protocol: protocol,
			},
		)
	}
	return allIpsWithPort
}

func GetUfwPortsTrafficPolicyRequest(ports []string, protocol string) []controller.UfwPortsTrafficPolicyRequest {
	allPorts := []controller.UfwPortsTrafficPolicyRequest{}
	for _, onePort := range ports {
		numPort, err := strconv.ParseUint(onePort, 10, 32)
		if err != nil {
			panic(err)
		}
		numUint32Port := uint32(numPort)
		allPorts = append(
			allPorts,
			controller.UfwPortsTrafficPolicyRequest{
				Port:     numUint32Port,
				Protocol: protocol,
			},
		)
	}
	return allPorts
}
