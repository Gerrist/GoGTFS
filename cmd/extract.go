package cmd

import (
	"github.com/Gerrist/gtfs-cli/GTFS"
	"github.com/Gerrist/gtfs-cli/util"
	"github.com/spf13/cobra"
	"log"
)

var filterAgency string
var inputDir string
var outputDir string

func init() {
	versionCmd.PersistentFlags().StringVarP(&filterAgency, "agency", "a", "", "agency to extract data from")
	versionCmd.PersistentFlags().StringVarP(&inputDir, "input", "i", "", "Input GTFS directory")
	versionCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "", "Directory where output is stored")
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract agency from GTFS",
	Long:  `Extract agency from GTFS`,
	Run: func(cmd *cobra.Command, args []string) {
		if filterAgency == "" {
			log.Panicln("agency flag can't be empty (example: -agency=CXX)")
		}
		if inputDir == "" {
			log.Panicln("input flag can't be empty (example: -input=gtfs-data)")
		}
		if outputDir == "" {
			log.Panicln("output flag can't be empty (example: -input=niag-gtfs)")
		}

		if !util.DirectoryExists(inputDir) {
			log.Panicln("Input directory does not exists")
		} else {
			gtfs := GTFS.Store{}

			log.Println("[Import]", "Importing agency.txt")
			gtfs.ReadFile("agency", "./" + inputDir + "/agency.txt")

			log.Println("[Import]", "Importing calendar_dates.txt")
			gtfs.ReadFile("calendar_dates", "./" + inputDir + "/calendar_dates.txt")

			log.Println("[Import]", "Importing routes.txt")
			gtfs.ReadFile("routes", "./" + inputDir + "/routes.txt")

			log.Println("[Import]", "Importing shapes.txt")
			gtfs.ReadFile("shapes", "./" + inputDir + "/shapes.txt")

			log.Println("[Import]", "Importing stop_times.txt")
			gtfs.ReadFile("stop_times", "./" + inputDir + "/stop_times.txt")

			log.Println("[Import]", "Importing stops.txt")
			gtfs.ReadFile("stops", "./" + inputDir + "/stops.txt")

			log.Println("[Import]", "Importing transfers.txt")
			gtfs.ReadFile("transfers", "./" + inputDir + "/transfers.txt")

			log.Println("[Import]", "Importing trips.txt")
			gtfs.ReadFile("trips", "./" + inputDir + "/trips.txt")

			log.Println("[Filter]", "Filtering GTFS with", filterAgency, "data")

			newGtfs := GTFS.Store{}

			for _, agency := range gtfs.Agency {
				if agency.Id == filterAgency {
					newGtfs.Agency = append(newGtfs.Agency, agency)
				}
			}

			routeIds := make([]string, 0)
			for _, route := range gtfs.Route {
				if route.AgencyId == filterAgency {
					newGtfs.Route = append(newGtfs.Route, route)
					routeIds = append(routeIds, route.RouteId)
				}
			}

			tripIds := make([]string, 0)
			shapeIds := make([]string, 0)
			serviceIds := make([]string, 0)
			for _, trip := range gtfs.Trip {
				if util.IndexOf(trip.RouteId, routeIds) > -1 {
					newGtfs.Trip = append(newGtfs.Trip, trip)
					tripIds = append(tripIds, trip.TripId)
					serviceIds = append(serviceIds, trip.ServiceId)
					shapeIds = append(shapeIds, trip.ShapeId)
				}
			}

			for _, calendarDate := range gtfs.CalendarDates {
				if util.IndexOf(calendarDate.ServiceId, serviceIds) > -1 {
					newGtfs.CalendarDates = append(newGtfs.CalendarDates, calendarDate)
				}
			}

			stopIds := make([]string, 0)
			for _, stopTime := range gtfs.StopTime {
				if util.IndexOf(stopTime.TripId, tripIds) > -1 {
					newGtfs.StopTime = append(newGtfs.StopTime, stopTime)
					stopIds = append(stopIds, stopTime.StopId)
				}
			}

			for _, stop := range gtfs.Stop {
				if util.IndexOf(stop.Id, stopIds) > -1 {
					newGtfs.Stop = append(newGtfs.Stop, stop)
				}
			}

			for _, shape := range gtfs.Shape {
				if util.IndexOf(shape.Id, shapeIds) > -1 {
					newGtfs.Shape = append(newGtfs.Shape, shape)
				}
			}

			log.Println("[Export]", "Exporting new GTFS with", filterAgency, "data to", outputDir)

			newGtfs.Export(outputDir)
		}

	},
}
