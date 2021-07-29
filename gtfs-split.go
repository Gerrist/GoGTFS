package main

import (
	"github.com/Gerrist/GoGTFS/GTFS"
	"github.com/Gerrist/GoGTFS/util"
	"log"
)

func main() {
	filterAgency := "TRANSDEV"

	gtfs := GTFS.Store{}

	log.Println("[Import]", "Importing agency.txt")
	gtfs.ReadFile("agency", "./gtfs-data/agency.txt")

	log.Println("[Import]", "Importing calendar_dates.txt")
	gtfs.ReadFile("calendar_dates", "./gtfs-data/calendar_dates.txt")

	log.Println("[Import]", "Importing routes.txt")
	gtfs.ReadFile("routes", "./gtfs-data/routes.txt")

	log.Println("[Import]", "Importing shapes.txt")
	gtfs.ReadFile("shapes", "./gtfs-data/shapes.txt")

	log.Println("[Import]", "Importing stop_times.txt")
	gtfs.ReadFile("stop_times", "./gtfs-data/stop_times.txt")

	log.Println("[Import]", "Importing stops.txt")
	gtfs.ReadFile("stops", "./gtfs-data/stops.txt")

	log.Println("[Import]", "Importing transfers.txt")
	gtfs.ReadFile("transfers", "./gtfs-data/transfers.txt")

	log.Println("[Import]", "Importing trips.txt")
	gtfs.ReadFile("trips", "./gtfs-data/trips.txt")

	log.Println("[Filter]", "Filtering GTFS with", filterAgency, "data")

	newGtfs := GTFS.Store{}

	for _, agency := range gtfs.Agency {
		if agency.Id == filterAgency {
			newGtfs.Agency = append(newGtfs.Agency, agency)
		}
	}

	routeIds := make([]int, 0)
	for _, route := range gtfs.Route {
		if route.AgencyId == filterAgency {
			newGtfs.Route = append(newGtfs.Route, route)
			routeIds = append(routeIds, route.RouteId)
		}
	}

	tripIds := make([]int, 0)
	shapeIds := make([]int, 0)
	serviceIds := make([]int, 0)
	for _, trip := range gtfs.Trip {
		if util.IndexOfInt(trip.RouteId, routeIds) > -1 {
			newGtfs.Trip = append(newGtfs.Trip, trip)
			tripIds = append(tripIds, trip.TripId)
			serviceIds = append(serviceIds, trip.ServiceId)
			shapeIds = append(shapeIds, trip.ShapeId)
		}
	}

	for _, calendarDate := range gtfs.CalendarDates {
		if util.IndexOfInt(calendarDate.ServiceId, serviceIds) > -1 {
			newGtfs.CalendarDates = append(newGtfs.CalendarDates, calendarDate)
		}
	}

	stopIds := make([]int, 0)
	for _, stopTime := range gtfs.StopTime {
		if util.IndexOfInt(stopTime.TripId, tripIds) > -1 {
			newGtfs.StopTime = append(newGtfs.StopTime, stopTime)
			stopIds = append(stopIds, stopTime.StopId)
		}
	}

	for _, stop := range gtfs.Stop {
		if util.IndexOfInt(stop.Id, stopIds) > -1 {
			newGtfs.Stop = append(newGtfs.Stop, stop)
		}
	}

	for _, shape := range gtfs.Shape {
		if util.IndexOfInt(shape.Id, shapeIds) > -1 {
			newGtfs.Shape = append(newGtfs.Shape, shape)
		}
	}

	log.Println("[Export]", "Exporting new GTFS with", filterAgency, "data")

	newGtfs.Export("new-gtfs")
}
