package common

type SessionType uint8

const (
	UnknownSessionType = SessionType(0)
)

// Weather represents the weather present on the track
type Weather uint8

const (
	Clear      = Weather(0)
	LightCloud = Weather(1)
	Overcast   = Weather(2)
	LightRain  = Weather(3)
	HeavyRain  = Weather(4)
	Storm      = Weather(5)
)

// ZoneFlag represents the flag being flown in a marshal zone
type ZoneFlag int8

const (
	UnknownZoneFlag = ZoneFlag(-1)
	None            = ZoneFlag(0)
	Green           = ZoneFlag(1)
	Blue            = ZoneFlag(2)
	Yellow          = ZoneFlag(3)
	Red             = ZoneFlag(4)
)
