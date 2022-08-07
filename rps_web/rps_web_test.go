package rps_web

import "testing"

func TestPlayRound(t *testing.T) {
	cases := []struct {
		name           string
		expectedResult string
	}{
		{"test for \"draw\"", "It's draw."},
		{"test for \"comp's victory\"", "Computer wins!"},
		{"test for \"player's victory\"", "Human wins!"},
	}
	t.Run(cases[0].name, func(t *testing.T) {
		for i := 0; i < 3; i++ {
			result := PlayRoundConfigurable(i, SpyRPSGenerator, i)
			if cases[0].expectedResult != result.RoundResult {
				t.Errorf("Expected to get %s, but got %s", cases[0].expectedResult, result.RoundResult)
			}
		}
	})
	t.Run(cases[1].name, func(t *testing.T) {
		for i := 0; i < 3; i++ {
			result := PlayRoundConfigurable(i, SpyRPSGenerator, (i+1)%3)
			if cases[1].expectedResult != result.RoundResult {
				t.Errorf("Expected to get %s, but got %s", cases[1].expectedResult, result.RoundResult)
			}
		}
	})
	t.Run(cases[2].name, func(t *testing.T) {
		for i := 0; i < 3; i++ {
			result := PlayRoundConfigurable((i+1)%3, SpyRPSGenerator, i)
			if cases[2].expectedResult != result.RoundResult {
				t.Errorf("Expected to get %s, but got %s", cases[2].expectedResult, result.RoundResult)
			}
		}
	})
}

func TestGetPhrase(t *testing.T) {
	//test if we get right phrases for victory, loss and draw.
	cases := []struct {
		name           string
		winner         int
		expectedPhases []string
	}{
		{"Victory", PLAYERWINS, phrases[PLAYERWINS]},
		{"Loss", COMPUTERWINS, phrases[COMPUTERWINS]},
		{"Draw", DRAW, phrases[DRAW]},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := getPhrase(tc.winner, phrases)
			for _, phr := range phrases[tc.winner] {
				if got == phr {
					return
				}
			}
			t.Errorf("For %s got %s", tc.name, got)
		})
	}
}

func SpyRPSGenerator(n int) int {
	return n
}
