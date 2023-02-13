package clients

import "github.com/greenhell1337/shopForDB/pkg/models/address"

type Client struct {
	Name             string
	Surname          string
	TimeRegistration string
	Date             map[string]string
	address.Address
}
