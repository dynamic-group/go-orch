package orch

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	// "log"
	"net"
	"time"
)

const (
	HeartBeatInterval int = 5
)

type Orch struct {
	client *etcd.Client
}

type Service struct {
	Created time.Time
	Name    string
	Host    string
	Port    int
	Addr    net.IP
}

func (o *Orch) Register(name string, host string, ip net.IP, port int) (*Service, error) {
	_, err := o.client.Set(fmt.Sprintf("services/%s/nodes/%s", name, host), fmt.Sprintf("%s:%d", ip.String(), port), 0)
	if err != nil {
		return &Service{}, err
	}
	return &Service{
		Created: time.Now(),
		Name:    name,
		Host:    host,
		Port:    port,
		Addr:    ip,
	}, nil
}

func (o *Orch) Unregister(s *Service) error {
	_, err := o.client.Delete(fmt.Sprintf("services/%s/nodes/%s", s.Name, s.Host), false)
	return err
}

func NewOrch(etcdClient *etcd.Client) *Orch {
	return &Orch{
		etcdClient,
	}
}
