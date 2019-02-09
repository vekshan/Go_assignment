package main

import (
	"errors"
	"fmt"
	"strings"
)

type Trip struct {
	destination string
	weight float32
	deadline int
}

type Vehicle struct{
	vehicle string
	name string
	destination string
	speed float32
	capacity float32
	load float32
}
type Truck struct {
	Vehicle

}

type Pickup struct{
	Vehicle
	isPrivate bool
}

type TrainCar struct{
	Vehicle
	railway string
}

func NewTruck() Truck {
	return Truck{
		Vehicle: Vehicle{
			vehicle:     "Truck",
			name:        "Truck",
			destination: "",
			speed:       40,
			capacity:    10,
			load:        0,
		},
	}
}

func NewPickUp() Pickup {
	return Pickup{
		Vehicle: Vehicle{
			vehicle:     "Pickup",
			name:        "Pickup",
			destination: "",
			speed:       60,
			capacity:    2,
			load:        0,
		},
		isPrivate: true,
	}
}

func NewTrainCar() TrainCar {
	return TrainCar{
		Vehicle: Vehicle{
			vehicle:     "TrainCar",
			name:        "TrainCar",
			destination: "",
			speed:       30,
			capacity:    30,
			load:        0,
		},
		railway: "CNR",
	}
}

type Transporter interface {
		addLoad(trip Trip)(err error)
		print()
}

func NewTorontoTrip(weight float32, deadline int) (trip *Trip) {
	return &Trip{
		destination: "Toronto",
		weight: weight,
		deadline: deadline,
	}
}

func NewMontrealTrip(weight float32, deadline int) (trip *Trip) {
	return &Trip{
		destination: "Montreal",
		weight: weight,
		deadline: deadline,
	}
}
func (vehicle *Vehicle) addLoad(trip Trip) (err error){
	if vehicle.destination == "" {
		vehicle.destination = trip.destination
	}else if vehicle.destination != trip.destination{
		return errors.New("Other destination")
	}

	if vehicle.capacity - vehicle.load < trip.weight {
		return errors.New("Out of capacity")
	} else {
		vehicle.load += trip.weight
	}

	if trip.destination == "Montreal" && (200/vehicle.speed) > float32(trip.deadline) {
		return errors.New("Insufficient time")
	}

	if trip.destination == "Toronto" && (400/vehicle.speed) > float32(trip.deadline) {
		return errors.New("Insufficient time")
	}

	return nil
}


func (truck *Truck) addLoad(trip Trip) (err error){
	return truck.Vehicle.addLoad(trip)
}

func (pickup *Pickup) addLoad(trip Trip) (err error){
	return pickup.Vehicle.addLoad(trip)
}

func (trainCar *TrainCar) addLoad(trip Trip) (err error){
	return trainCar.Vehicle.addLoad(trip)
}

func (truck *Truck) print(){
	fmt.Printf("%s to %s with %f tons\n", truck.name, truck.destination, truck.load);

}

func (pickup *Pickup) print(){
	fmt.Printf("%s to %s with %f tons (Private: %t) \n", pickup.name, pickup.destination, pickup.load, pickup.isPrivate);

}

func (trainCar *TrainCar) print(){
	fmt.Printf("%s to %s with %f tons (%s)\n", trainCar.name, trainCar.destination, trainCar.load, trainCar.railway);

}

func main() {

	var transporters [6] Transporter

	truckA := NewTruck()
	truckB := NewTruck()

	pickupA := NewPickUp()
	pickupB := NewPickUp()
	pickupC := NewPickUp()

	trainCarA := NewTrainCar()

	truckA.name = "Truck A"
	truckB.name = "Truck B"

	pickupA.name = "Pickup A"
	pickupB.name = "Pickup B"
	pickupC.name = "Pickup C"

	trainCarA.name = "TrainCar A"

	transporters[0] = &truckA
	transporters[1] = &truckB
	transporters[2] = &pickupA
	transporters[3] = &pickupB
	transporters[4] = &pickupC
	transporters[5] = &trainCarA

	//var flag bool

	var trips []Trip

	for  {

		var destination string
		var weight float32
		var deadline int

		fmt.Print("Destination: (t)oronto, (m)ontreal, else exit? ")
		fmt.Scanln(&destination)

		destination = destination[:1]
		if strings.ToLower(destination) != "m" &&  strings.ToLower(destination) != "t" {
			break
		}


		fmt.Print("Weight: ")
		fmt.Scanln(&weight)

		fmt.Print("Deadline (in hours): ")
		fmt.Scanln(&deadline)



		var trip Trip
		if strings.ToLower(destination) == "m" {
			trip = *NewMontrealTrip(weight, deadline)
		} else if strings.ToLower(destination) == "t" {
			trip = *NewTorontoTrip(weight, deadline)
		}

		for _,transporter := range transporters{
			err:= Transporter.addLoad(transporter, trip)
			if err != nil {
				fmt.Println(err)
			}else{
				trips = append(trips, trip)
				break
			}
		}
	}

	fmt.Println("Not going to TO or Montreal, bye!")
	fmt.Println("Trips: ",trips)
	fmt.Println("Vehicles: ")
	truckA.print()
	truckB.print()
	pickupA.print()
	pickupB.print()
	pickupC.print()
	trainCarA.print()
}





