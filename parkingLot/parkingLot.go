package parking_lot

import (
    "sync"
    _ "fmt"
    . "github.com/mohakkataria/gojek_takehome/car"
    "container/heap"
    "fmt"
    "strings"
)

var instance *ParkingLot
var once sync.Once

type ParkingLot struct {
    emptySlots IntHeap // heap of empty slots to optimize to log(n) operations of allocating a smaller empty slot and vacating any given slot
    maxSize int // max size of the parking lot
    isParkingLotCreated bool // check if parking lot has been initialized or not
    regNoSlotMap map[string]int // Map of registration number to the slot
    slotCarMap map[string]Car // Map of Slots to Cars
    colorRegNoMap map[string]map[string]bool // Map of Car Color to registration number Hast set(implemented as another map of string to bool)
}

// colorRegNoMap,regNoSlotMap,slotCarMap is implemented in such a way such that find and remove operations are done in O(1)
// emptySlots in implemented as Heap so that complexity of extract minimum and insert to log(n) on average. Since
// there wasn't much information as to if either of operations of leave and park could be more than the other
// hence a call was to taken to average out the optimality on both the operations

func GetInstance() *ParkingLot {
    once.Do(func() {
        instance = &ParkingLot{}
    })
    return instance
}

func (this *ParkingLot) Initialize(numberOfSlots int) {
    if !numberOfSlots {
        fmt.Println("Parking Lot of Size 0 cannot be created")
        return
    }
    a := make([]int, numberOfSlots)
    for i := range a {
        a[i] = 1 + i
    }
    this.emptySlots = &IntHeap{a}
    heap.Init(this.emptySlots)
    this.slotCarMap = map[int]Car{}
    this.colorRegNoMap = map[string]map[string]bool{}
    this.regNoSlotMap = map[string]int{}
    this.maxSize = numberOfSlots
    this.isParkingLotCreated = true
}

func (this *ParkingLot) mapRegNoToSlot(regNo string, slot int) {
    this.regNoSlotMap[regNo] = slot;
}


func (this *ParkingLot) unmapRegNo(regNo string) {
    delete(this.regNoSlotMap, regNo);
}

func (this *ParkingLot) mapSlotToCar(slot int, car Car) {
    this.slotCarMap[slot] = car;
}


func (this *ParkingLot) unmapSlot(slot int) {
    delete(this.slotCarMap, slot);
}

func (this *ParkingLot) mapColorToRegNo(color string, regNo string) {
    _, exists := this.colorRegNoMap[color]
    if exists {
        this.colorRegNoMap[color][regNo] = true
    } else {
        this.colorRegNoMap[color] = map[string]bool{regNo:true}
    }
}


func (this *ParkingLot) unmapRegNoFromColorMap(color string, regNo string) {
    delete(this.colorRegNoMap[color], regNo);
}


func (this *ParkingLot) Park(car Car) {
    if (!this.isParkingLotCreated) {
        fmt.Println("Parking Lot not created")
        return
    }

    if (this.emptySlots.Len() == 0) {
        fmt.Println("Sorry, parking lot is full");
    } else {
        emptySlot := this.emptySlots.Pop()
        this.mapRegNoToSlot(car.GetRegNo(), emptySlot)
        this.mapSlotToCar(emptySlot, car)
        this.mapColorToRegNo(car.GetColor(), car.GetRegNo())
        fmt.Println("Allocated slot number: " + emptySlot);
    }
}

func (this *ParkingLot) Leave(slot int) {
    if (!this.isParkingLotCreated) {
        fmt.Println("Parking Lot not created")
        return
    }

    if car, exists := this.slotCarMap[slot]; exists {
        this.unmapRegNo(car.GetRegNo())
        this.unmapRegNoFromColorMap(car.GetColor(), car.GetRegNo())
        this.unmapSlot(slot)
        fmt.Println("Slot number" + slot + "is free")
    } else {
        fmt.Println("No car parked on given slot")
    }
}

func (this *ParkingLot) Status() {
    if (!this.isParkingLotCreated) {
        fmt.Println("Parking Lot not created")
        return
    }

    fmt.Println("Slot No.\tRegistration No.\tColour");
    i := 1
    for i <= this.maxSize {
        if car, exists := this.slotCarMap[i]; exists {
           fmt.Println(i + "\t" + car.GetRegNo() + "\t" + car.GetColor())
        }
    }
}

func (this *ParkingLot) GetRegNosForCarsWithColor(color string) []string{
    if (!this.isParkingLotCreated) {
        fmt.Println("Parking Lot not created")
        return
    }

    regNoSlice := []string{}
    if regNoMap, exists := this.colorRegNoMap[color]; exists {
        for regNo, _ := range regNoMap {
            regNoSlice = append(regNoSlice, regNo)
        }
    }

    if len(regNoSlice) > 0 {
        fmt.Println(strings.Join(regNoSlice, ","))
    } else {
        fmt.Println("No cars found with given colour")
    }
    return regNoSlice
}

func (this *ParkingLot) GetSlotNosForCarsWithColor(color string) {
    if (!this.isParkingLotCreated) {
        fmt.Println("Parking Lot not created")
        return
    }

    regNos := this.GetRegNosForCarsWithColor(color)
    slots := []int{}
    for regNo := range regNos {
        if slot, exists := this.regNoSlotMap[regNo]; exists {
            slots = append(slots, slot)
        }
    }

    if len(slots) > 0 {
        fmt.Println(strings.Join(slots, ","))
    } else {
        fmt.Println("No cards found for given colour")
    }

}

func (this *ParkingLot) GetSlotNoForRegNo(regNo string) {
    if (!this.isParkingLotCreated) {
        fmt.Println("Parking Lot not created")
        return
    }

    if slot, exists := this.regNoSlotMap[regNo]; exists {
        fmt.Println(slot)
    } else {
        fmt.Println("Car with this registration no not parked in any slot")
    }
}

