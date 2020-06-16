package character

type Character struct {
	Overview
	Skills
	Stats
}

type Overview struct {
	Name string
	Race string
	Class string
	Level int
}

type Skills struct {
	Strength Skill
	Dexterity Skill
	Constitution Skill
	Intelligence Skill
	Wisdom Skill
	Charisma Skill
}

type Skill struct {
	Name string
	Value int
	Modifier int
	Skills []string
	SkillValues map[string]int
	SavingThrow int
}

type Stats struct {
	HP int
	AC int
	Initiative int
	Speed int
	ProfBonus int
	Inspiration int
	HitDice int
}

type Abilities struct {

}

func NewCharacter() *Character {
	o := Overview{Name: "Uldar", Race: "Elf", Class: "class", Level: 5}
	skills := createBaseSkills()
	c := Character{Overview: o, Skills: skills}
	return &c
}

func createBaseSkills() Skills {
	str_sks := []string{"Athletics"}
	str_vals := map[string]int{"Athletics": 1}
	str := Skill{Name: "Strength", Value: 1, Modifier: 1, Skills: str_sks, SkillValues: str_vals, SavingThrow: 1}
	dex_sks := []string{"Acrobatics", "Sleight of Hand", "Stealth"}
	dex_vals := map[string]int{"Acrobatics": 1, "Sleight of Hand": 1, "Stealth": 1}
	dex := Skill{Name: "Dexterity", Value: 1, Modifier: 1, Skills: dex_sks, SkillValues: dex_vals, SavingThrow: 1}
	con_sks := []string{}
	con_vals := map[string]int{}
	con := Skill{Name: "Consitution", Value: 1, Modifier: 1, Skills: con_sks, SkillValues: con_vals, SavingThrow: 1}
	int_sks := []string{"Arcana", "History", "Investigation", "Nature", "Religion"}
	int_vals := map[string]int{"Arcana": 1, "History": 1, "Investigation": 1, "Nature": 1, "Religion": 1}
	inte := Skill{Name: "Intelligence", Value: 1, Modifier: 1, Skills: int_sks, SkillValues: int_vals, SavingThrow: 1}
	wis_sks := []string{"Animal Handling", "Insight", "Medicine", "Perception", "Survival"}
	wis_vals := map[string]int{"Animal Handling": 1, "Insight": 1, "Medicine": 1, "Perception": 1, "Survival": 1}
	wis := Skill{Name: "Wisdom", Value: 1, Modifier: 1, Skills: wis_sks, SkillValues: wis_vals, SavingThrow: 1}
	cha_sks := []string{"Deception", "Intimidation", "Performance", "Persuasion"}
	cha_vals := map[string]int{"Deception": 1, "Intimidation": 1, "Performance": 1, "Persuasion": 1}
	cha := Skill{Name: "Charisma", Value: 1, Modifier: 1, Skills: cha_sks, SkillValues: cha_vals, SavingThrow: 1}

	skills := Skills{Strength: str, Dexterity: dex, Constitution: con, Intelligence: inte, Wisdom: wis, Charisma: cha}
	return skills
}
