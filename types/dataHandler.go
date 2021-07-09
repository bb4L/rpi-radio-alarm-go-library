package types

// DataHandler the interface is there to abstract interaction with storage / the api
type DataHandler interface {
	GetAlarms(withWritePermission bool) ([]Alarm, error)
	GetAlarm(idx int, withWritePermission bool) (Alarm, error)
	AddAlarm(alarm Alarm) ([]Alarm, error)
	GetRadio(withWritePermission bool) (Radio, error)
	SaveRadio(radio Radio) (Radio, error)
	SaveAlarm(idx int, alarm Alarm) (Alarm, error)
	DeleteAlarm(idx int) ([]Alarm, error)
	SwitchRadio(running bool) (Radio, error)
}
