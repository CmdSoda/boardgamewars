package game

import (
	"github.com/CmdSoda/boardgamewars/internal/countrycodes"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"github.com/google/uuid"
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
	Globals.AllWarParties[WarPartyIdUSA] = wpUSA

	// Russia
	wpRussia := WarParty{
		WarPartyId:    WarPartyIdRussia,
		WarPartyColor: Red,
		Name:          "Russia",
		Country:       countrycodes.Russia,
		Pilots:        []PilotId{},
		Aircrafts:     []AircraftId{},
	}
	Globals.AllWarParties[WarPartyIdRussia] = wpRussia

	// UK
	wpUK := WarParty{
		WarPartyId:    WarPartyIdUK,
		WarPartyColor: Red,
		Name:          "UK",
		Country:       countrycodes.UK,
		Pilots:        []PilotId{},
		Aircrafts:     []AircraftId{},
	}
	Globals.AllWarParties[WarPartyIdUK] = wpUK

	// Germany
	wpGermany := WarParty{
		WarPartyId:    WarPartyIdGermany,
		WarPartyColor: Blue,
		Name:          "Germany",
		Country:       countrycodes.Germany,
		Pilots:        []PilotId{},
		Aircrafts:     []AircraftId{},
	}
	Globals.AllWarParties[WarPartyIdGermany] = wpGermany
}

func InitGame() error {
	var err error
	Globals.AllWarParties = map[WarPartyId]WarParty{}
	Globals.AllAircrafts = map[AircraftId]Aircraft{}
	CreateWarParties()
	Globals.AirbaseList = map[AirbaseId]Airbase{}
	randomizer.Init()
	if err = LoadAircrafts(); err != nil {
		return err
	}
	if err = LoadAir2AirWeapons(); err != nil {
		return err
	}
	return nil
}
