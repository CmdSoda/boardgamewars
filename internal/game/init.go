package game

import (
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Data struct {
	Startup
	AllAirbases           AirbasesMap
	AllAircraftParameters AircraftParametersMap
	Air2AirWeapons        Air2AirWeaponLibrary
	AllWarParties         WarPartyMap
	AllAircrafts          AircraftsMap
	AllPilots             PilotsMap
	AllCounters           CounterList
	CountryDataMap
	World
	Statistics
	Settings
}

var Globals = Data{}

// Shortcuts to the warparties
var WarPartyIdUSA = WarPartyId(uuid.MustParse("92432884-3974-11ec-8d3d-0242ac130003"))
var WarPartyIdRussia = WarPartyId(uuid.MustParse("a261b7c6-3974-11ec-8d3d-0242ac130003"))
var WarPartyIdUK = WarPartyId(uuid.MustParse("5a6dffaa-3975-11ec-8d3d-0242ac130003"))
var WarPartyIdGermany = WarPartyId(uuid.MustParse("5e737c4c-3975-11ec-8d3d-0242ac130003"))

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
	Globals.AllWarParties = map[WarPartyId]*WarParty{}
	Globals.AllAircrafts = map[AircraftId]*Aircraft{}
	Globals.AllPilots = map[PilotId]*Pilot{}
	if Globals.AllWarParties, err = LoadWarParties(); err != nil {
		return err
	}
	Globals.AllAirbases = map[AirbaseId]Airbase{}
	Globals.Statistics = NewStatistics()
	randomizer.Init(seed)
	Globals.World = NewWorld()
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
