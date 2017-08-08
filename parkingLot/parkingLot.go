package parkingLot

import (
    "sync"
    _ "fmt"
    . "github.com/mohakkataria/gojek_takehome/car"
    "container/heap"
    "fmt"
    "strings"
    "strconv"
)

var instance *ParkingLot
var once sync.Once

type ParkingLot struct {
    emptySlots IntHeap // heap of empty slots to optimize to log(n) operations of allocating a smaller empty slot and vacating any given slot
    maxSize int // max size of the parking lot
    isParkingLotCreated bool // check if parking lot has been initialized or not
    regNoSlotMap map[string]int // Map of registration number to the slot
    slotCarMap map[int]Car // Map of Slots to Cars
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
    if numberOfSlots <= 0 {
        fmt.Println("Parking Lot of Size <= 0 cannot be created")
        return
    }

    this.emptySlots = IntHeap{}
    i := 1
    for i <= numberOfSlots {
        this.emptySlots = append(this.emptySlots, i)
        i++
    }

    heap.Init(&this.emptySlots)
    this.slotCarMap = map[int]Car{}
    this.colorRegNoMap = map[string]map[string]bool{}
    this.regNoSlotMap = map[string]int{}
    this.maxSize = numberOfSlots
    this.isParkingLotCreated = true

    fmt.Println("Created a parking lot with " + strconv.Itoa(numberOfSlots) + " slots")
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

        if slot := this.GetSlotNoForRegNo(car.GetRegNo(), false); slot == 0 {
            emptySlot := heap.Pop(&this.emptySlots)
            this.mapRegNoToSlot(car.GetRegNo(), emptySlot.(int))
            this.mapSlotToCar(emptySlot.(int), car)
            this.mapColorToRegNo(car.GetColor(), car.GetRegNo())
            fmt.Println("Allocated slot number: " + strconv.Itoa(emptySlot.(int)));
        } else {
            fmt.Println("Car with this registration number already parked at slot: " + strconv.Itoa(slot));
        }

    }
}

func (this *ParkingLot) Leave(slot int) {
    if (!this.isParkingLotCreated) {
        fmt.Println("Parking Lot not created")
        return
    }

    if car, exists := this.slotCarMap[slot]; exists {
        heap.Push(&this.emptySlots, slot)
        this.unmapRegNo(car.GetRegNo())
        this.unmapRegNoFromColorMap(car.GetColor(), car.GetRegNo())
        this.unmapSlot(slot)
        fmt.Println("Slot number " + strconv.Itoa(slot) + " is free")
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
           fmt.Println(strconv.Itoa(i) + "\t" + car.GetRegNo() + "\t" + car.GetColor())
        }
        i++
    }
}

func (this *ParkingLot) GetRegNosForCarsWithColor(color string, print bool) []string{
    if (!this.isParkingLotCreated) {
        if print {
            fmt.Println("Parking Lot not created")
        }
        return []string{}
    }

    regNoSlice := []string{}
    if regNoMap, exists := this.colorRegNoMap[color]; exists {
        for regNo, _ := range regNoMap {
            regNoSlice = append(regNoSlice, regNo)
        }
    }

    if print {
        if len(regNoSlice) > 0 {
            fmt.Println(strings.Join(regNoSlice, ","))
        } else {
            fmt.Println("No cars found with given colour")
        }
    }

    return regNoSlice
}

func (this *ParkingLot) GetSlotNosForCarsWithColor(color string) {
    if (!this.isParkingLotCreated) {
        fmt.Println("Parking Lot not created")
        return
    }

    regNos := this.GetRegNosForCarsWithColor(color, false)
    slots := []string{}
    for _, regNo := range regNos {
        if slot, exists := this.regNoSlotMap[regNo]; exists {
            slots = append(slots, strconv.Itoa(slot))
        }
    }

    if len(slots) > 0 {
        fmt.Println(strings.Join(slots, ","))
    } else {
        fmt.Println("No cars found for given colour")
    }

}

func (this *ParkingLot) GetSlotNoForRegNo(regNo string, print bool) int {
    if (!this.isParkingLotCreated) {
        if print {
            fmt.Println("Parking Lot not created")
        }
        return 0
    }

    if slot, exists := this.regNoSlotMap[regNo]; exists {
        if print {
            fmt.Println(slot)
        }
        return slot
    } else {
        if print {
            fmt.Println("Car with this registration no not parked in any slot")
        }
        return 0
    }
}

