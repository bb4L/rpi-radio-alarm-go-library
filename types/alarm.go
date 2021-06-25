// Package types contains all the relevant structs/types for the library
package types

// Represent a alarm
type Alarm struct {
	Name   string `yaml:"name" json:"name"`
	Hour   int    `yaml:"hour" json:"hour"`
	Minute int    `yaml:"minute" json:"min"`
	Days   []int  `yaml:"days" json:"days"`
	Active bool   `yaml:"active" json:"on"`
}
