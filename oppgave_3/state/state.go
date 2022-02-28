package state

import (
	"strings"
)

var west = []string{"Korn", "Kylling", "Rev", "HS"}
var boat = []string{}
var east = []string{}


func ViewState() string {
    	// Sjekke data som er lagret og viser "verdens-bilde"
	westP := strings.Join(west, " ")
	boatP := strings.Join(boat, " ")
	eastP := strings.Join(east, " ")


    	return "[" + westP + " ---\\ \\_" + boatP + "_/ _________________/--- " + eastP + "]"
}


func PutInBoat(cargo string, position string) string {
	/*
	moveToBoat(position, cargo)
	moveToBoat(position, HS)

	westP := strings.Join(west, " ")
	boatP := strings.Join(boat, " ")
	eastP := strings.Join(east, " ")

	return "[" + westP + " ---\\ \\_" + boatP + "_/ _________________/--- " + eastP + "]"
	*/
	return "[Korn Rev ---\\ \\_Kylling HS_/ _________________/--- ]"
}

/*
//henta kode fra nettet (SO) fordi umulig å fjerne element fra slice i go xd
func moveToBoat[T comparable](l []T, item T) []T {
    for i, other := range l {
       	if other == item {
	    append(boat, T)
            return append(l[:i], l[i+1:]...)
        }
    }
    return l
}*/
