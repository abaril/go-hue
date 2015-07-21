package configuration

import (
	"testing"
	"fmt"
	"github.com/abaril/go-hue/src/portal"
	"github.com/abaril/go-hue/src/key"
)

var (
	username_api_key_filename = "../../conf/example.conf"
	username_api_key = ""
)

func init() {
	k := key.New(username_api_key_filename)
	username_api_key = k.Username
}

func TestCreateUser(t *testing.T) {
	fmt.Println("configuration.CreateUser()")
	portal := portal.GetPortal()
	ccc := NewConfiguration(portal[0].InternalIpAddress)
	response := ccc.CreateUser("go-hue-user", "go-lang-test")
	fmt.Println("\t", response[0])
}

func TestDeleteUser(t *testing.T) {
	fmt.Println("configuration.DeleteUser()")
	portal := portal.GetPortal()
	ccc := NewConfiguration(portal[0].InternalIpAddress)
	response := ccc.DeleteUser(username_api_key, "ae2b1fca515949e5d54fb22b8ed95575")
	fmt.Println("\t", response[0])
}

func TestGetFullState(t *testing.T) {
	fmt.Println("configuration.GetFullState()")
	portal := portal.GetPortal()
	ccc := NewConfiguration(portal[0].InternalIpAddress)
	response := ccc.GetFullState(username_api_key)
	fmt.Println(response)
}

func TestGetConfiguration(t *testing.T) {
	portal := portal.GetPortal()
	fmt.Println("configuration.GetConfiguration()")
	for i := range portal {
		p := portal[i]
		h := NewConfiguration(p.InternalIpAddress)
		c := h.GetConfiguration(username_api_key)
		fmt.Printf("\tName:         %s\n", c.Name)
		fmt.Printf("\tUtc:          %s\n", c.Utc)
		fmt.Printf("\tSwVersion:    %s\n", c.SwVersion)
		fmt.Printf("\tProxyAddress: %s\n", c.ProxyAddress)
		fmt.Printf("\tProxyPort:    %d\n", c.ProxyPort)
		fmt.Printf("\tMac:          %s\n", c.Mac)
		fmt.Printf("\tLinkButton:   %t\n", c.LinkButton)
		fmt.Printf("\tIpAddress:    %s\n", c.IpAddress)
		fmt.Printf("\tNetMask:      %s\n", c.NetMask)
		fmt.Printf("\tGateway:      %s\n", c.Gateway)
		fmt.Printf("\tDhcp:         %t\n", c.Dhcp)
		fmt.Printf("\tSwUpdate:\n")
		fmt.Printf("\t\tUpdateState: %d\n", c.SwUpdate.UpdateState)
		fmt.Printf("\t\tUrl:         %s\n", c.SwUpdate.Url)
		fmt.Printf("\t\tText:        %s\n", c.SwUpdate.Text)
		fmt.Printf("\t\tNotify:      %t\n", c.SwUpdate.Notify)
		fmt.Printf("\tWhitelist:\n")
		for j := range c.Whitelist {
			fmt.Printf("\t\tKey: %s\n", j)
			fmt.Printf("\t\t\tLastUseDate: %s\n", c.Whitelist[j].LastUseDate)
			fmt.Printf("\t\t\tCreateDate:  %s\n", c.Whitelist[j].CreateDate)
			fmt.Printf("\t\t\tName:        %s\n", c.Whitelist[j].Name)
		}
		t.Log(p.InternalIpAddress)
	}
}
