package GoGTFS

import (
	"encoding/csv"
	"log"
	"os"
)

type Agency struct {
	Id       string
	Name     string
	URL      string
	Timezone string
	Phone    string
}

func AgencyIndex(value string) int {
	indices := []string{"agency_id", "agency_name", "agency_url", "agency_timezone", "agency_phone"}
	return IndexOf(value, indices)
}

type CalendarDate struct {
	ServiceId     string
	Date          string
	ExceptionType int
}

func CalendarDateIndex(value string) int {
	indices := []string{"service_id", "date", "exception_type"}
	return IndexOf(value, indices)
}

type Route struct {
	RouteId        string
	AgencyId       string
	ExternalCode   string
	RouteShortName string
	RouteLongName  string
	RouteDesc      string
	RouteType      string
	RouteColor     string
	RouteTextColor string
	RouteURL       string
}

func RouteIndex(value string) int {
	indices := []string{"route_id", "agency_id", "external_code", "route_short_name", "route_long_name", "route_desc", "route_type", "route_color", "route_text_color", "route_url"}
	return IndexOf(value, indices)
}

type Shape struct {
	Id           string
	PTSequence   int
	Lat          float64
	Lon          float64
	DistTraveled int
}

func ShapeIndex(value string) int {
	indices := []string{"shape_id", "shape_pt_sequence", "shape_pt_lat", "shape_pt_lon", "shape_dist_traveled"}
	return IndexOf(value, indices)
}

type StopTime struct {
	TripId            string
	Sequence          int
	StopId            string
	StopHeadsign      string
	ArrivalTime       string
	DepartureTime     string
	PickUpType        int
	DropOffType       int
	Timepoint         int
	ShapeDistTraveled int
	FareUnitsTraveled int
}

func StopTimeIndex(value string) int {
	indices := []string{"trip_id", "stop_sequence", "stop_id", "stop_headsign", "arrival_time", "departure_time", "pickup_type", "drop_off_type", "timepoint", "shape_dist_traveled", "fare_units_traveled"}
	return IndexOf(value, indices)
}

type Stop struct {
	Id                 string
	Code               string
	Name               string
	Lat                float64
	Lon                float64
	LocationType       int
	ParentStation      string
	StopTimezone       string
	WheelchairBoarding int
	PlatformCode       string
	ZoneId             string
}

func StopIndex(value string) int {
	indices := []string{"stop_id", "stop_code", "stop_name", "stop_lat", "stop_lon", "location_type", "parent_station", "stop_timezone", "wheelchair_boarding", "platform_code", "zone_id"}
	return IndexOf(value, indices)
}

type Transfer struct {
	FromStopId   string
	ToStopId     string
	FromRouteId  string
	ToRouteId    string
	FromTripId   string
	ToTripId     string
	TransferType int
}

func TransferIndex(value string) int {
	indices := []string{"from_stop_id", "to_stop_id", "from_route_id", "to_route_id", "from_trip_id", "to_trip_id", "transfer_type"}
	return IndexOf(value, indices)
}

type Trip struct {
	RouteId              string
	ServiceId            string
	TripId               string
	RealtimeTripId       string
	TripHeadsign         string
	TripShortName        string
	TripLongName         string
	DirectionId          int
	BlockId              int
	ShapeId              string
	WheelchairAccessible int
	BikesAllowed         int
}

func TripIndex(value string) int {
	indices := []string{"route_id", "service_id", "trip_id", "realtime_trip_id", "trip_headsign", "trip_short_name", "trip_long_name", "direction_id", "block_id", "shape_id", "wheelchair_accessible", "bikes_allowed"}
	return IndexOf(value, indices)
}

type Store struct {
	Agency        []Agency
	CalendarDates []CalendarDate
	Route         []Route
	Shape         []Shape
	StopTime      []StopTime
	Stop          []Stop
	Transfer      []Transfer
	Trip          []Trip
}

func (store *Store) ReadFile(fileType, filePath string) { // we need both fileType and filePath, so we can support having multiple files of one type
	gtfsFile, err := os.Open(filePath)
	gtfsFileLines, err := csv.NewReader(gtfsFile).ReadAll()

	switch fileType {
	case "agency":
		for i, line := range gtfsFileLines {
			if i > 0 { // Skip header line
				agency := Agency{
					Id:       line[AgencyIndex("agency_id")],
					Name:     line[AgencyIndex("agency_name")],
					URL:      line[AgencyIndex("agency_url")],
					Timezone: line[AgencyIndex("agency_timezone")],
					Phone:    line[AgencyIndex("agency_phone")],
				}
				store.Agency = append(store.Agency, agency)
			}
		}
		break
	case "calendar_dates":
		for i, line := range gtfsFileLines {
			if i > 0 { // Skip header line
				calendarDate := CalendarDate{
					ServiceId:     line[CalendarDateIndex("service_id")],
					Date:          line[CalendarDateIndex("date")],
					ExceptionType: ParseInt(line[CalendarDateIndex("exception_type")]),
				}

				store.CalendarDates = append(store.CalendarDates, calendarDate)
			}
		}
		break
	case "routes":
		for i, line := range gtfsFileLines {
			if i > 0 { // Skip header line
				route := Route{
					RouteId:        line[RouteIndex("route_id")],
					AgencyId:       line[RouteIndex("agency_id")],
					ExternalCode:   line[RouteIndex("external_code")],
					RouteShortName: line[RouteIndex("route_short_name")],
					RouteLongName:  line[RouteIndex("route_long_name")],
					RouteDesc:      line[RouteIndex("route_desc")],
					RouteType:      line[RouteIndex("route_type")],
					RouteColor:     line[RouteIndex("route_color")],
					RouteTextColor: line[RouteIndex("route_text_color")],
					RouteURL:       line[RouteIndex("route_url")],
				}

				store.Route = append(store.Route, route)
			}
		}
		break
	case "shapes":
		for i, line := range gtfsFileLines {
			if i > 0 { // Skip header line
				shape := Shape{
					Id:           line[ShapeIndex("shape_id")],
					PTSequence:   ParseInt(line[ShapeIndex("shape_pt_sequence")]),
					Lat:          ParseFloat(line[ShapeIndex("shape_pt_lat")]),
					Lon:          ParseFloat(line[ShapeIndex("shape_pt_lon")]),
					DistTraveled: ParseInt(line[ShapeIndex("shape_dist_traveled")]),
				}

				store.Shape = append(store.Shape, shape)
			}
		}
		break
	case "stop_times":
		for i, line := range gtfsFileLines {
			if i > 0 { // Skip header line
				stopTime := StopTime{
					TripId:            line[StopTimeIndex("trip_id")],
					Sequence:          ParseInt(line[StopTimeIndex("stop_sequence")]),
					StopId:            line[StopTimeIndex("stop_id")],
					StopHeadsign:      line[StopTimeIndex("stop_headsign")],
					ArrivalTime:       line[StopTimeIndex("arrival_time")],
					DepartureTime:     line[StopTimeIndex("departure_time")],
					PickUpType:        ParseInt(line[StopTimeIndex("pickup_type")]),
					DropOffType:       ParseInt(line[StopTimeIndex("drop_off_type")]),
					Timepoint:         ParseInt(line[StopTimeIndex("timepoint")]),
					ShapeDistTraveled: ParseInt(line[StopTimeIndex("shape_dist_traveled")]),
					FareUnitsTraveled: ParseInt(line[StopTimeIndex("fare_units_traveled")]),
				}

				store.StopTime = append(store.StopTime, stopTime)
			}
		}
		break
	case "stops":
		for i, line := range gtfsFileLines {
			if i > 0 { // Skip header line
				stop := Stop{
					Id:                 line[StopIndex("stop_id")],
					Code:               line[StopIndex("stop_code")],
					Name:               line[StopIndex("stop_name")],
					Lat:                ParseFloat(line[StopIndex("stop_lat")]),
					Lon:                ParseFloat(line[StopIndex("stop_lon")]),
					LocationType:       ParseInt(line[StopIndex("location_type")]),
					ParentStation:      line[StopIndex("parent_station")],
					StopTimezone:       line[StopIndex("stop_timezone")],
					WheelchairBoarding: ParseInt(line[StopIndex("wheelchair_boarding")]),
					PlatformCode:       line[StopIndex("platform_code")],
					ZoneId:             line[StopIndex("zone_id")],
				}

				store.Stop = append(store.Stop, stop)
			}
		}
		break
	case "transfers":
		for i, line := range gtfsFileLines {
			if i > 0 { // Skip header line
				transfer := Transfer{
					FromStopId:   line[TransferIndex("from_stop_id")],
					ToStopId:     line[TransferIndex("to_stop_id")],
					FromRouteId:  line[TransferIndex("from_route_id")],
					ToRouteId:    line[TransferIndex("to_route_id")],
					FromTripId:   line[TransferIndex("from_trip_id")],
					ToTripId:     line[TransferIndex("to_trip_id")],
					TransferType: ParseInt(line[TransferIndex("transfer_type")]),
				}

				store.Transfer = append(store.Transfer, transfer)
			}
		}
		break
	case "trips":
		for i, line := range gtfsFileLines {
			if i > 0 { // Skip header line
				trip := Trip{
					RouteId:              line[TripIndex("route_id")],
					ServiceId:            line[TripIndex("service_id")],
					TripId:               line[TripIndex("trip_id")],
					RealtimeTripId:       line[TripIndex("realtime_trip_id")],
					TripHeadsign:         line[TripIndex("trip_headsign")],
					TripShortName:        line[TripIndex("trip_short_name")],
					TripLongName:         line[TripIndex("trip_long_name")],
					DirectionId:          ParseInt(line[TripIndex("direction_id")]),
					BlockId:              ParseInt(line[TripIndex("block_id")]),
					ShapeId:              line[TripIndex("shape_id")],
					WheelchairAccessible: ParseInt(line[TripIndex("wheelchair_accessible")]),
					BikesAllowed:         ParseInt(line[TripIndex("bikes_allowed")]),
				}

				store.Trip = append(store.Trip, trip)
			}
		}
		break

	}

	//calendarDatesFile, err := os.Open("./gtfs-data/calendar_dates.txt")
	//calendarDatesLines, err := csv.NewReader(calendarDatesFile).ReadAll()

	if err != nil {
		log.Fatalln(err)
	}

}

func (store *Store) Export(exportName string) {
	_ = os.Mkdir(exportName, 0755)

	os.Remove(exportName + "/stops.txt")
	stopsFile, _ := os.OpenFile(exportName+"/stops.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	stopsFile.WriteString(CSVRow([]interface{}{"stop_id", "stop_code", "stop_name", "stop_lat", "stop_lon", "location_type", "parent_station", "stop_timezone", "wheelchair_boarding", "platform_code", "zone_id"}))
	for _, stop := range store.Stop {
		var row = []interface{}{}
		row = append(row, stop.Id)
		row = append(row, stop.Code)
		row = append(row, stop.Name)
		row = append(row, stop.Lat)
		row = append(row, stop.Lon)
		row = append(row, stop.LocationType)
		row = append(row, stop.ParentStation)
		row = append(row, stop.StopTimezone)
		row = append(row, stop.WheelchairBoarding)
		row = append(row, stop.PlatformCode)
		row = append(row, stop.ZoneId)

		stopsFile.WriteString(CSVRow(row))
	}

	os.Remove(exportName + "/agency.txt")
	agencyFile, _ := os.OpenFile(exportName+"/agency.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	agencyFile.WriteString(CSVRow([]interface{}{"agency_id", "agency_name", "agency_url", "agency_timezone", "agency_phone"}))
	for _, agency := range store.Agency {
		var row = []interface{}{}
		row = append(row, agency.Id)
		row = append(row, agency.Name)
		row = append(row, agency.URL)
		row = append(row, agency.Timezone)
		row = append(row, agency.Phone)

		agencyFile.WriteString(CSVRow(row))
	}

	os.Remove(exportName + "/calendar_dates.txt")
	calendarFile, _ := os.OpenFile(exportName+"/calendar_dates.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	calendarFile.WriteString(CSVRow([]interface{}{"service_id", "date", "exception_type"}))
	for _, calendarDate := range store.CalendarDates {
		var row = []interface{}{}
		row = append(row, calendarDate.ServiceId)
		row = append(row, calendarDate.Date)
		row = append(row, calendarDate.ExceptionType)

		calendarFile.WriteString(CSVRow(row))
	}

	os.Remove(exportName + "/routes.txt")
	routesFile, _ := os.OpenFile(exportName+"/routes.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	routesFile.WriteString(CSVRow([]interface{}{"route_id", "agency_id", "external_code", "route_short_name", "route_long_name", "route_desc", "route_type", "route_color", "route_text_color", "route_url"}))
	for _, route := range store.Route {
		var row = []interface{}{}
		row = append(row, route.RouteId)
		row = append(row, route.AgencyId)
		row = append(row, route.ExternalCode)
		row = append(row, route.RouteShortName)
		row = append(row, route.RouteLongName)
		row = append(row, route.RouteDesc)
		row = append(row, route.RouteType)
		row = append(row, route.RouteColor)
		row = append(row, route.RouteTextColor)
		row = append(row, route.RouteURL)

		routesFile.WriteString(CSVRow(row))
	}

	os.Remove(exportName + "/shapes.txt")
	shapesFile, _ := os.OpenFile(exportName+"/shapes.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	shapesFile.WriteString(CSVRow([]interface{}{"shape_id", "shape_pt_sequence", "shape_pt_lat", "shape_pt_lon", "shape_dist_traveled"}))
	for _, shape := range store.Shape {
		var row = []interface{}{}
		row = append(row, shape.Id)
		row = append(row, shape.PTSequence)
		row = append(row, shape.Lat)
		row = append(row, shape.Lon)
		row = append(row, shape.DistTraveled)

		shapesFile.WriteString(CSVRow(row))
	}

	os.Remove(exportName + "/stop_times.txt")
	stopTimesFile, _ := os.OpenFile(exportName+"/stop_times.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	stopTimesFile.WriteString(CSVRow([]interface{}{"trip_id", "stop_sequence", "stop_id", "stop_headsign", "arrival_time", "departure_time", "pickup_type", "drop_off_type", "timepoint", "shape_dist_traveled", "fare_units_traveled"}))
	for _, stopTime := range store.StopTime {
		var row = []interface{}{}
		row = append(row, stopTime.TripId)
		row = append(row, stopTime.Sequence)
		row = append(row, stopTime.StopId)
		row = append(row, stopTime.StopHeadsign)
		row = append(row, stopTime.ArrivalTime)
		row = append(row, stopTime.DepartureTime)
		row = append(row, stopTime.PickUpType)
		row = append(row, stopTime.DropOffType)
		row = append(row, stopTime.Timepoint)
		row = append(row, stopTime.ShapeDistTraveled)
		row = append(row, stopTime.FareUnitsTraveled)

		stopTimesFile.WriteString(CSVRow(row))
	}

	os.Remove(exportName + "/transfers.txt")
	transfersFile, _ := os.OpenFile(exportName+"/transfers.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	transfersFile.WriteString(CSVRow([]interface{}{"from_stop_id", "to_stop_id", "from_route_id", "to_route_id", "from_trip_id", "to_trip_id", "transfer_type"}))
	for _, stop := range store.Transfer {
		var row = []interface{}{}
		row = append(row, stop.FromStopId)
		row = append(row, stop.ToStopId)
		row = append(row, stop.FromRouteId)
		row = append(row, stop.ToRouteId)
		row = append(row, stop.FromStopId)
		row = append(row, stop.ToStopId)
		row = append(row, stop.FromTripId)
		row = append(row, stop.ToTripId)
		row = append(row, stop.TransferType)

		transfersFile.WriteString(CSVRow(row))
	}

	os.Remove(exportName + "/trips.txt")
	tripsFile, _ := os.OpenFile(exportName+"/trips.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	tripsFile.WriteString(CSVRow([]interface{}{"route_id", "service_id", "trip_id", "realtime_trip_id", "trip_headsign", "trip_short_name", "trip_long_name", "direction_id", "block_id", "shape_id", "wheelchair_accessible", "bikes_allowed"}))
	for _, trip := range store.Trip {
		var row = []interface{}{}
		row = append(row, trip.RouteId)
		row = append(row, trip.ServiceId)
		row = append(row, trip.TripId)
		row = append(row, trip.RealtimeTripId)
		row = append(row, trip.TripHeadsign)
		row = append(row, trip.TripShortName)
		row = append(row, trip.TripLongName)
		row = append(row, trip.DirectionId)
		row = append(row, trip.BlockId)
		row = append(row, trip.ShapeId)
		row = append(row, trip.WheelchairAccessible)
		row = append(row, trip.BikesAllowed)

		tripsFile.WriteString(CSVRow(row))
	}
}
