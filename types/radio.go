package types

import (
	"os"
	"os/exec"

	"github.com/bb4L/rpi-radio-alarm-go-library/constants"
	"github.com/bb4L/rpi-radio-alarm-go-library/logging"
)

var logger = logging.GetLogger(os.Stdout, constants.DefaultPrefix, "radio")

// Radio representing the state of the radio
type Radio struct {
	Running bool `yaml:"running" json:"isPlaying"`
	Pid     int  `yaml:"pid"`
}

// StartRadio starts  the radio
func (r *Radio) StartRadio() {
	logger.Println("start radio")
	r.startRadioWithFunction(nil)
}

func (r *Radio) startRadioWithFunction(startRadioFunction func() (int, error)) {
	var err error

	if startRadioFunction == nil {
		logger.Println("using default start")
		startRadioFunction = defaultStartRadio
	}

	if r.Running {
		return
	}

	r.Pid, err = startRadioFunction()
	r.Running = true

	if err != nil {
		logger.Panicf("could not start radio %s", err)
	}

	logger.Printf("started radio process %d\n", r.Pid)
}

func defaultStartRadio() (int, error) {
	cmd := exec.Command("mplayer", "https://streamingp.shoutcast.com/hotmixradio-sunny-128.mp3", "volume 150")
	cmd.Stdout = logger.Writer()
	cmd.Stderr = logger.Writer()
	err := cmd.Start()
	return cmd.Process.Pid, err
}

// StopRadio stops the radio
func (r *Radio) StopRadio() error {
	logger.Printf("stop radio\n")
	return r.stopRadioWithFunction(nil)
}

func (r *Radio) stopRadioWithFunction(stopFunction func(int) error) (err error) {
	if stopFunction == nil {
		logger.Println("using default stop")
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
