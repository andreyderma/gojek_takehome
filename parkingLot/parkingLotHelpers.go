package parkingLot

import . "github.com/mohakkataria/gojek_takehome/car"

// Helper function to add to HashMap, mapping of RegNo to Slot
func mapRegNoToSlot(regNo string, slot int) {
    this := GetInstance()
    this = instance
    this.regNoSlotMap[regNo] = slot;
}

// Helper function to remove from HashMap, mapping of RegNo to Slot
func unmapRegNo(regNo string) {
    this := GetInstance()
    delete(this.regNoSlotMap, regNo);
}

// Helper function to add to HashMap, mapping of slot to Car
func mapSlotToCar(slot int, car Car) {
    this := GetInstance()
    this.slotCarMap[slot] = car;
}

// Helper function to remove from HashMap, mapping of slot to Car
func unmapSlot(slot int) {
    this := GetInstance()
    delete(this.slotCarMap, slot);
}

// Helper function to add to HashSet at given color key in the HashMap
func mapColorToRegNo(color string, regNo string) {
    this := GetInstance()
    _, exists := this.colorRegNoMap[color]
    if exists {
        this.colorRegNoMap[color][regNo] = true
    } else {
        this.colorRegNoMap[color] = map[string]bool{regNo:true}
    }
}

// Helper function to remove from HashSet at given color key in the HashMap
func unmapRegNoFromColorMap(color string, regNo string) {
    this := GetInstance()
    delete(this.colorRegNoMap[color], regNo);
}
