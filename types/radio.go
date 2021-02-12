package types

// Radio struct to represent a radio
type Radio struct {
	Running bool `yaml:"running" json:"isPlaying"`
	Pid     int  `yaml:"pid"`
}
