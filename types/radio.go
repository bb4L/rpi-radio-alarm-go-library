package types

import (
	"os"
	"os/exec"

	"github.com/bb4L/rpi-radio-alarm-go-library/logging"
)

// Struct to represent a radio
type Radio struct {
	Running bool `yaml:"running" json:"isPlaying"`
	Pid     int  `yaml:"pid"`
}

// Start the radio
func (r *Radio) StartRadio() {
	logging.GetInfoLogger().Println("start radio")
	r.startRadioWithFunction(nil)
}

func (r *Radio) startRadioWithFunction(startRadioFunction func() (int, error)) {
	var err error = nil
	if startRadioFunction == nil {
		logging.GetInfoLogger().Println("using default start")
		startRadioFunction = defaultStartRadio
	}

	if r.Running {
		return
	}

	r.Pid, err = startRadioFunction()
	r.Running = true

	if err != nil {
		logging.GetErrorLogger().Panicf("could not start radio %s", err)
	}

	logging.GetInfoLogger().Printf("started radio process %d", r.Pid)
}

func defaultStartRadio() (int, error) {
	cmd := exec.Command("mplayer", "https://streamingp.shoutcast.com/hotmixradio-sunny-128.mp3", "volume 150")
	cmd.Stdout = logging.GetInfoLogger().Writer()
	cmd.Stderr = logging.GetErrorLogger().Writer()
	err := cmd.Start()
	return cmd.Process.Pid, err
}

// Stop the radio
func (r *Radio) StopRadio() error {
	logging.GetInfoLogger().Printf("stop radio")
	return r.stopRadioWithFunction(nil)
}

func (r *Radio) stopRadioWithFunction(stopFunction func(int) error) (err error) {
	if stopFunction == nil {
		logging.GetInfoLogger().Println("using default stop")
		stopFunction = defaultStopRadio
	}

	if !r.Running || r.Pid == -1 {
		return
	}

	err = stopFunction(r.Pid)
	if err != nil {
		return
	}

	r.Running = false
	r.Pid = -1

	return
}

func defaultStopRadio(pid int) error {
	var err error

	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	err = process.Kill()

	if err != nil {
		return err
	}
	return nil
}
