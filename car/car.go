package car

type Car struct {
    color string
    regNo string
}

func (this Car) GetRegNo() string{
    return this.regNo
}


func (this Car) GetColor() string{
    return this.color
}