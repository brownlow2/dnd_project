package inventory

import "testing"

func TestReturnWeapon(t *testing.T) {
	weap := ReturnWeapon("Fake Name", "Fake Type", "Melee", 1, 3, 1, 6)
	if weap.Name != "Fake Name" {
		t.Errorf("Name incorrect, got: %v, want: %v.", weap.Name, "Fake Name")
	}

	if weap.Quantity != 1 {
		t.Errorf("Quantity incorrect, got: %v, want: %v.", weap.Quantity, 1)
	}

	if weap.Type != "Fake Type" {
		t.Errorf("Type incorrect, got: %v, want: %v.", weap.Type, "Fake Type")
	}

	if weap.Range != "Melee" {
		t.Errorf("Range incorrect, got: %v, want: %v.", weap.Range, "Melee")
	}

	if weap.DamageDie.Amount != 1 {
		t.Errorf("Die amount incorrect, got: %v, want: %v.", weap.DamageDie.Amount, 1)
	}

	if weap.DamageDie.Number != 6 {
		t.Errorf("Die number incorrect, got: %v, want: %v.", weap.DamageDie.Number, 6)
	}

	if weap.Modifier != 3 {
		t.Errorf("Modifier incorrect, got: %v, want: %v.", weap.Modifier, 3)
	}
}
