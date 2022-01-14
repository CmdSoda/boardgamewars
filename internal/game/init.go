package game

import (
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"github.com/sirupsen/logrus"
)

type Data struct {
	Startup
	AllAirbases           AirbasesMap
	AllAircraftParameters AircraftParametersMap
	Air2AirWeapons        Air2AirWeaponLibrary
	AllAircrafts          AircraftsMap
	AllPilots             PilotsMap
	AllCounters           CounterList
	CountryDataMap
	World
	Statistics
	Settings
	EventList
}

var Globals = Data{}

func loadConfig() error {
	var err error
	if Globals.Startup, err = LoadStartup("startup.json"); err != nil {
		return err
	}
	return nil
}

func initGame(seed int64) error {
	var err error
	Log.Info("game engine is starting...\n")
	Globals.AllAircrafts = map[AircraftId]*Aircraft{}
	Globals.AllPilots = map[PilotId]*Pilot{}
	Globals.AllAirbases = map[AirbaseId]*Airbase{}
	Globals.Statistics = NewStatistics()
	randomizer.Init(seed)
	Globals.World = NewWorld()
	Globals.EventList = NewEventList()
	if err = LoadSettings("settings.json"); err != nil {
		return err
	}
	if err = LoadAircraftParameters(); err != nil {
		return err
	}
	if err = LoadAircraftParameters(); err != nil {
		return err
	}
	if err = LoadAir2AirWeapons(); err != nil {
		return err
	}
	if err = LoadCountryData(); err != nil {
		return err
	}
	return nil
}

// InitGame initialisiert das Spiel mit dem angegebenen seed. Wird 0 angegeben, wird ein seed basierend auf
// der Systemzeit benutzt.
func InitGame(seed int64) error {
	var err error
	var lvl logrus.Level
	Log = logrus.New()
	if err = loadConfig(); err != nil {
		return err
	}
	lvl, err = logrus.ParseLevel(Globals.Startup.LogLevel)
	if err != nil {
		Log.Errorf("error while parsing log level: %s", err.Error())
	}
	Log.SetLevel(lvl)
	return initGame(seed)
}

func InitGameWithLogLevel(seed int64, level logrus.Level) error {
	Log = logrus.New()
	Log.SetLevel(level)
	if err := loadConfig(); err != nil {
		return err
	}
	return initGame(seed)
}
