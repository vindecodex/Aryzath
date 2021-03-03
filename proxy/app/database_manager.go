package app

import "fmt"

type DatabaseManager struct {
	Data []string
}

func (d *DatabaseManager) GetData(password string) ([]string, error) {
	if password == "admin" || password == "user" {
		fmt.Printf("Data accessed by %s", password)
		return d.Data, nil
	}
	return nil, fmt.Errorf("Invalid password")
}

func (d *DatabaseManager) SetData(password string, data string) error {
	if password == "admin" {
		d.Data = append(d.Data, data)
		return nil
	}
	return fmt.Errorf("Requres Admin Password")
}
