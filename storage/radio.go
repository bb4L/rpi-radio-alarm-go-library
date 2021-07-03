package storage

import "github.com/bb4L/rpi-radio-alarm-go-library/types"

// GetRadio returns the radio resource
func (storageHelper *Helper) GetRadio(withWritePermission bool) (types.Radio, error) {
	if withWritePermission {
		storageHelper.getMutex().Lock()
	} else {
		storageHelper.getMutex().RLock()
		defer storageHelper.getMutex().RUnlock()
	}
	data, err := storageHelper.getStoredData()
	if err != nil {
		return types.Radio{}, err
	}
	return data.Radio, nil

}

// SaveRadio saves given radio as the radio resource
func (storageHelper *Helper) SaveRadio(radio types.Radio) (types.Radio, error) {
	storageHelper.getMutex().Lock()
	data, err := storageHelper.getStoredData()

	if err != nil {
		return types.Radio{}, err
	}
	data.Radio = radio

	storageHelper.SaveStoredData(data)

	return storageHelper.GetRadio(false)

}

// SwitchRadio switches the radio to running if true is passed
func (storageHelper *Helper) SwitchRadio(running bool) (types.Radio, error) {
	radio, err := storageHelper.GetRadio(false)
	if err != nil {
		logger.Println("error switching radio")
		return types.Radio{}, err
	}

	if radio.Running == running {
		return radio, nil
	}

	if running {
		radio.StartRadio()
	} else {
		radio.StopRadio()
	}

	return storageHelper.SaveRadio(radio)
}
