package common

// DriverStatus represents where the driver is on track and what kind of lap they may be on
type DriverStatus uint8

const (
	InGarage  = DriverStatus(0)
	FlyingLap = DriverStatus(1)
	InLap     = DriverStatus(2)
	OutLap    = DriverStatus(3)
	OnTrack   = DriverStatus(4)
)

// Formula is the type of formula racing that is occuring
type Formula uint8

const (
	F1        = Formula(0)
	F1Classic = Formula(1)
	F2        = Formula(2)
	F1Generic = Formula(3)
)

// PitStatus represents the cars position in relation to the pit lane
type PitStatus uint8

const (
	NoPitStatus = PitStatus(0)
	Pitting     = PitStatus(1)
	InPitArea   = PitStatus(2)
)

// ResultStatus represents the status of the driver for the given lap
type ResultStatus uint8

const (
	InvalidResultStatus = ResultStatus(0)
	Inactive            = ResultStatus(1)
	Active              = ResultStatus(2)
	Finished            = ResultStatus(3)
	Disqualified        = ResultStatus(4)
	NotClassified       = ResultStatus(5)
	Retired             = ResultStatus(6)
)

// SafetyCarStatus represents the status of the safety car on the circuit
type SafetyCarStatus uint8

const (
	NoSafetyCar      = SafetyCarStatus(0)
	FullSafetyCar    = SafetyCarStatus(1)
	VirtualSafetyCar = SafetyCarStatus(2)
)

// Sector represents which sector on the track the car is in
type Sector uint8

const (
	Sector1 = Sector(0)
	Sector2 = Sector(1)
	Sector3 = Sector(2)
)

// SessionType represents the type of session currently happening
type SessionType uint8

const (
	UnknownSessionType = SessionType(0)
	Practice1          = SessionType(1)
	Practice2          = SessionType(2)
	Practice3          = SessionType(3)
	ShortPractice      = SessionType(4)
	Qualifying1        = SessionType(5)
	Qualifying2        = SessionType(6)
	Qualifying3        = SessionType(7)
	ShortQualifying    = SessionType(8)
	OneShotQualifying  = SessionType(9)
	Race               = SessionType(10)
	Race2              = SessionType(11)
	TimeTrial          = SessionType(12)
)

// Track represents the track being raced on
type Track int8

const (
	UnknownTrack     = Track(-1)
	Melbourne        = Track(0)
	PaulRicard       = Track(1)
	Shanghai         = Track(2)
	Bahrain          = Track(3)
	Catalunya        = Track(4)
	Monaco           = Track(5)
	Montreal         = Track(6)
	Silverstone      = Track(7)
	Hockenheim       = Track(8)
	Hungaroring      = Track(9)
	Spa              = Track(10)
	Monza            = Track(11)
	Singapore        = Track(12)
	Suzuka           = Track(13)
	AbuDhabi         = Track(14)
	Texas            = Track(15)
	Brazil           = Track(16)
	Austria          = Track(17)
	Sochi            = Track(18)
	Mexico           = Track(19)
	Azerbaijan       = Track(20)
	SakhirShort      = Track(21)
	SilverstoneShort = Track(22)
	TexasShort       = Track(23)
	SuzukaShort      = Track(24)
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
