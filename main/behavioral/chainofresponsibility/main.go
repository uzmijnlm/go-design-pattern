package main

func main() {
	var slot1, slot2, slot3 Slot

	slot1 = &ParamSlot{}
	slot2 = &DegradeSlot{}
	slot3 = &SystemSlot{}

	slot1.setNext(&slot2)
	slot2.setNext(&slot3)

	slot1.check()
}
