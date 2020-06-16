package inventory

import (
	"math/rand"
	"math"
	"time"
	"fmt"
)

type Inventory struct {
	Items map[string][]Item
}

type Item interface {
	Increment()
	Decrement()
	//Name string
	//Quantity int
	//Description string
}

type Weapon struct {
	Name string
	Quantity int
	Type string
	Range string
	DamageDie Die
	Modifier int
}

type Die struct {
	Amount int
	Number int
}

func (w *Weapon) Damage() int {
	roll := 0.0
	// Don't just double the roll
	for i := 0; i < w.DamageDie.Amount; i++ {
		src := rand.NewSource(time.Now().UnixNano())
		r := rand.New(src)
		roll += math.Ceil((r.Float64() * float64(w.DamageDie.Number)))
		roll += float64(w.Modifier)
	}

	return int(roll)
}

func (w *Weapon) ShowDamageDice() string {
	return fmt.Sprintf("%dd%d+%d", w.DamageDie.Amount, w.DamageDie.Number, w.Modifier)
}

func (w *Weapon) Increment() {
	w.Quantity++
}

func (w *Weapon) Decrement() {
	w.Quantity--
}

func ReturnWeapon(name, typ, rang string, quant, mod, die_amount, die_number int) *Weapon {
	damage_die := Die{die_amount, die_number}
	// TODO: Add parameter validation
	weap := Weapon{name, quant, typ, rang, damage_die, mod}
	return &weap
}
