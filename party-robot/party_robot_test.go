package partyrobot

import "testing"

func TestWelcome(t *testing.T) {
	tests := []struct {
		description string
		name        string
		want        string
	}{
		{
			description: "Greet the guest with a welcoming message",
			name:        "Christiane",
			want:        "Welcome to my party, Christiane!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if got := Welcome(tt.name); got != tt.want {
				t.Errorf("Welcome(%s) = %s, want %s", tt.name, got, tt.want)
			}
		})
	}
}

func TestHappyBirthday(t *testing.T) {
	tests := []struct {
		description string
		name        string
		age         int
		want        string
	}{
		{
			description: "Wish Happy Birthday using the given name and age of the person",
			name:        "Christiane",
			age:         58,
			want:        "Happy birthday Christiane! You are now 58 years old!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if got := HappyBirthday(tt.name, tt.age); got != tt.want {
				t.Errorf("HappyBirthday(%s, %d) = %s, want %s", tt.name, tt.age, got, tt.want)
			}
		})
	}
}

func TestAssignTable(t *testing.T) {
	tests := []struct {
		description   string
		name          string
		direction     string
		tableNumber   int
		distance      float64
		neighbourName string
		want          string
	}{
		{
			description:   "Greet the guest and give them directions to their seat",
			name:          "Christiane",
			direction:     "on the left",
			tableNumber:   27,
			distance:      23.7834298,
			neighbourName: "Frank",
			want:          "Welcome to my party, Christiane!\nYou have been assigned to table 1B. Your table is on the left, exactly 23.8 meters from here.\nYou will be sitting next to Frank.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if got := AssignTable(tt.name, tt.tableNumber, tt.neighbourName, tt.direction, tt.distance); got != tt.want {
				t.Errorf("AssignTable(%s,%d,%s,%s,%f) = %s, want %s", tt.name, tt.tableNumber, tt.neighbourName, tt.direction, tt.distance, got, tt.want)
			}
		})
	}
}