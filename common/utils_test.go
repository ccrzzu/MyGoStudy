package common

import (
	"fmt"
	"net"
	"testing"
)

func TestIP(t *testing.T) {
	ori := "192.168.78.123"
	fmt.Println("origin:", ori)
	i := InetAtoN(ori)
	fmt.Println(i)
	/*fmt.Println(byte(i>>24))
	fmt.Println(byte(i>>16))
	fmt.Println(byte(i>>8))
	fmt.Println(byte(i))*/
	fmt.Println(InetNtoA(i))

	fmt.Println(IPtoUInt32(net.IP{12,13,14,15}))
	fmt.Println(net.IP{12,13,14,15}.String())
	fmt.Println(UInt32toIP(202182159))
}
