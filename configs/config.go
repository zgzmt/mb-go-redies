package configs

import "fmt"

type ServerConfig struct {
	Address string
}

func (s *ServerConfig)SetAddress(ip string, port int){
	s.Address = fmt.Sprint("%s:%d", ip, port)
}