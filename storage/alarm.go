package storage

import "github.com/bb4L/rpi-radio-alarm-go-library/types"

// DeleteAlarm delets the alarm on the given index
func (storageHelper *Helper) DeleteAlarm(alarmIdx int) ([]types.Alarm, error) {
	storageHelper.getMutex().Lock()
	data, err := storageHelper.getStoredData()
	if err != nil {
		return []types.Alarm{}, err
	}
	if alarmIdx == len(data.Alarms)-1 {
		data.Alarms = data.Alarms[:alarmIdx]

	} else {
		data.Alarms = append(data.Alarms[:alarmIdx], data.Alarms[alarmIdx+1:]...)

	}

	storageHelper.SaveStoredData(data)
	return storageHelper.GetAlarms(false)
}

// GetAlarms returns all the alarms
func (storageHelper *Helper) GetAlarms(withWritePermission bool) ([]types.Alarm, error) {
	if withWritePermission {
		storageHelper.getMutex().Lock()
	} else {
		storageHelper.getMutex().RLock()
		defer storageHelper.getMutex().RUnlock()
	}

	data, err := storageHelper.getStoredData()
	if err != nil {
		return []types.Alarm{}, err
	}
	return data.Alarms, nil

}

// GetAlarm returns the alarn on the given index
func (storageHelper *Helper) GetAlarm(idx int, withWritePermission bool) (types.Alarm, error) {
	if withWritePermission {
		storageHelper.getMutex().Lock()
	} else {
		storageHelper.getMutex().RLock()
		defer storageHelper.getMutex().RUnlock()
	}

	data, err := storageHelper.getStoredData()
	if err != nil {
		return types.Alarm{}, err
	}
	return data.Alarms[idx], nil
}

// AddAlarm adds a given alarm
func (storageHelper *Helper) AddAlarm(alarm types.Alarm) ([]types.Alarm, error) {
	storageHelper.getMutex().Lock()
	defer storageHelper.getMutex().Unlock()

	data, err := storageHelper.getStoredData()
	if err != nil {
		return []types.Alarm{}, err
	}

	data.Alarms = append(data.Alarms, alarm)

	storageHelper.SaveStoredData(data)
	return storageHelper.GetAlarms(false)
}

// SaveAlarm savest tha given alarm at the given index to the storage
func (storageHelper *Helper) SaveAlarm(idx int, alarm types.Alarm) (types.Alarm, error) {
	storageHelper.getMutex().Lock()
	data, err := storageHelper.getStoredData()
	if err != nil {
		return types.Alarm{}, err
	}

	data.Alarms[idx] = alarm
	storageHelper.SaveStoredData(data)

	return storageHelper.GetAlarm(idx, false)
}
