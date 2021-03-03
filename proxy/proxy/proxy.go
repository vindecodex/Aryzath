package proxy

import (
	"Aryzath/proxy/app"
	"fmt"
)

type Proxy struct {
	DbManager *app.DatabaseManager
	Mode      string
}

func NewProxy() *Proxy {
	return &Proxy{
		DbManager: &app.DatabaseManager{
			Data: make([]string, 0),
		},
		Mode: "None",
	}
}

func (p *Proxy) GetData(password string) ([]string, error) {
	if p.Mode == "Get" {
		return p.DbManager.GetData(password)
	}
	return nil, fmt.Errorf("Invalid Proxy mode")
}

func (p *Proxy) SetData(password string, data string) error {
	if p.Mode == "Set" {
		return p.DbManager.SetData(password, data)
	}
	return fmt.Errorf("Invalid Proxy mode")
}

func (p *Proxy) SetProxyMode(mode string) {
	p.Mode = mode
}
