package config

import "fmt"

type MySQL struct {
	User     string
	PassWord string
	Address  string
	Port     string
	Socket   string
	DataBase string
}

func (c *MySQL) Source() string {
	prot := "tcp"
	url := fmt.Sprintf("%s:%s", c.Address, c.Port)
	if c.Socket != "" {
		prot = "unix"
		url = c.Socket
	}

	return fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=true", c.User, c.PassWord, prot, url, c.DataBase)
}
