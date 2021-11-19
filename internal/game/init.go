package game

import (
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var WarPartyIdUSA = WarPartyId(uuid.MustParse("92432884-3974-11ec-8d3d-0242ac130003"))
var WarPartyIdRussia = WarPartyId(uuid.MustParse("a261b7c6-3974-11ec-8d3d-0242ac130003"))
var WarPartyIdUK = WarPartyId(uuid.MustParse("5a6dffaa-3975-11ec-8d3d-0242ac130003"))
var WarPartyIdGermany = WarPartyId(uuid.MustParse("5e737c4c-3975-11ec-8d3d-0242ac130003"))

func CreateWarParties() {
	// USA
	wpUSA := WarParty{
		WarPartyId:    WarPartyIdUSA,
		WarPartyColor: Blue,
		Name:          "USA",
		Country:       countrycodes.USA,
		Pilots:        []PilotId{},
		Aircrafts:     []AircraftId{},
	}
	Globals.AllWarParties[WarPartyIdUSA] = &wpUSA

	// Russia
	wpRussia := WarParty{
		WarPartyId:    WarPartyIdRussia,
		WarPartyColor: Red,
		Name:          "Russia",
		Country:       countrycodes.Russia,
		Pilots:        []PilotId{},
		Aircrafts:     []AircraftId{},
	}
	Globals.AllWarParties[WarPartyIdRussia] = &wpRussia

	// UK
	wpUK := WarParty{
		WarPartyId:    WarPartyIdUK,
		WarPartyColor: Red,
		Name:          "UK",
		Country:       countrycodes.UK,
		Pilots:        []PilotId{},
		Aircrafts:     []AircraftId{},
	}
	Globals.AllWarParties[WarPartyIdUK] = &wpUK

	// Germany
	wpGermany := WarParty{
		WarPartyId:    WarPartyIdGermany,
		WarPartyColor: Blue,
		Name:          "Germany",
		Country:       countrycodes.Germany,
		Pilots:        []PilotId{},
		Aircrafts:     []AircraftId{},
	}
	Globals.AllWarParties[WarPartyIdGermany] = &wpGermany
}

func loadconfig() error {
	var err error
	if Globals.Config, err = LoadConfig("gameconfig.json"); err != nil {
		return err
	}
	return nil
}

func initgame(seed int64) error {
	var err error
	Log.Info("game engine is starting...\n")
	Globals.AllWarParties = map[WarPartyId]*WarParty{}
	Globals.AllAircrafts = map[AircraftId]*Aircraft{}
	Globals.AllPilots = map[PilotId]*Pilot{}
	CreateWarParties()
	Globals.AllAirbases = map[AirbaseId]Airbase{}
	Globals.Statistics = NewStatistics()
	randomizer.Init(seed)
	if err = LoadAircraftParameters(); err != nil {
		return err
	}
	if err = LoadAir2AirWeapons(); err != nil {
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
	if err = loadconfig(); err != nil {
		return err
	}
	lvl, err = logrus.ParseLevel(Globals.Config.LogLevel)
	if err != nil {
		Log.Errorf("error while parsing log level: %s", err.Error())
	}
	Log.SetLevel(lvl)
	return initgame(seed)
}

func InitGameWithLogLevel(seed int64, level logrus.Level) error {
	Log = logrus.New()
	Log.SetLevel(level)
	if err := loadconfig(); err != nil {
		return err
	}
	return initgame(seed)
}