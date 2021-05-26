package gps

import (
	"errors"

	"github.com/julienlevasseur/goLibs/logging"
	gpsd "github.com/stratoberry/go-gpsd"
	"github.com/umahmood/haversine"
)

// TPV Time Position Velocity
var TPV *gpsd.TPVReport

// Init initialize the gps.
func Init() {
	logging.Log.Debug("Initializing GPS")
	var gps *gpsd.Session
	var err error

	if gps, err = gpsd.Dial(gpsd.DefaultAddress); err != nil {
		logging.Log.Fatalf("Failed to connect to GPSD: %s", err)
	}

	gps.AddFilter("TPV", func(r interface{}) {
		TPV = r.(*gpsd.TPVReport)
	})

	done := gps.Watch()
	<-done
	logging.Log.Debug("GPS Initialized")
}

// Coordinates return the current gps coordinates.
func Coordinates() (haversine.Coord, error) {
	if TPV != nil {
		return haversine.Coord{Lat: TPV.Lat, Lon: TPV.Lon}, nil
	}

	return haversine.Coord{
		Lat: 0,
		Lon: 0,
	}, errors.New("coordinates not aquired yet")
}

// Heading return the current Heading (Cap).
func Heading() (heading float64, err error) {
	if TPV != nil {
		return TPV.Track, nil
	}

	return 0, errors.New("heading not available")
}
