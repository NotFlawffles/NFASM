package main

type Label struct {
    name string
    address uint16
}

func NewLabel(name string, address uint16) Label {
    return Label{name, address}
}
