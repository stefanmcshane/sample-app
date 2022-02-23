// Package config is used for defining structs of any config or secrets that the application may read in
package config

// Info is a sample struct for the informaton provided in ./sample-config.yml
type Info struct {
	ApplicationName string `json:"applicationName"`
}
