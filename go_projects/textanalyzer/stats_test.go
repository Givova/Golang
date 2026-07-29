package main

import "testing"

func TestCountLine(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected int
	}{
		{"Zero count", "", 0},
		{"One count", "my one string", 1},
		{"Several count", "ariva viva \n sasat", 2},
		{"String with n in end", "string \n", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountLines(tt.text)
			if result != tt.expected {
				t.Errorf("CountLine(%s) = %d, ожидали %d",
					tt.text, result, tt.expected)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected int
	}{
		{"Zero count", "", 0},
		{"One count", "one", 1},
		{"Several count", "ariva viva sasat", 3},
		{"String with n in end", "string    sds", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWords(tt.text)
			if result != tt.expected {
				t.Errorf("CountWords(%s) = %d, ожидали %d",
					tt.text, result, tt.expected)
			}
		})
	}
}

func TestCountChars(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected int
	}{
		{"Zero count", "", 0},
		{"Kirillitsa", "тевирп крош!", 12},
		{"ASCII", "we", 2},
		{"Emoji", "Go🚀", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountChars(tt.text)
			if result != tt.expected {
				t.Errorf("CountChars(%s) = %d, ожидали %d",
					tt.text, result, tt.expected)
			}
		})
	}
}


func TestWordFrequency(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected map[string]int
	}{
		{"Base", "Top  gun", map[string]int{"top": 1, "gun":1}},
		{"Repeat", "my one my", map[string]int{"my": 2, "one":1}},
		{"Diferent registry", "Viva viva ViVa ar",map[string]int{"viva": 3, "ar":1} },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WordFrequency(tt.text)
			if len(result) != len(tt.expected){
				t.Fatalf("WordFrequency(%q): получили %d ключей, хотели %d",
                    tt.text, len(result), len(tt.expected))
			}
			for k, v := range tt.expected{
				if result[k] != v{
					 t.Errorf("WordFrequency(%q)[%q] = %d, хотели %d",
                        tt.text, k, result[k], v)
				}
			}
		})
	}
}

// func TestTopWords(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input     map[string]int
// 		count 	int
// 		expected []WordCount
// 	}{
// 		{"Base", map[string]int{"top": 1, "gun":1}, 2, []WordCount{"top": 1, "gun":1}},
// 		{"Base", map[string]int{"top": 1, "gun":3}, 10, []WordCount{"gun": 3, "top":1}},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result := TopWords(tt.text)
// 			if len(result) != len(tt.expected){
// 				t.Fatalf("WordFrequency(%q): получили %d ключей, хотели %d",
//                     tt.text, len(result), len(tt.expected))
// 			}
// 			for k, v := range tt.expected{
// 				if result[k] != v{
// 					 t.Errorf("WordFrequency(%q)[%q] = %d, хотели %d",
//                         tt.text, k, result[k], v)
// 				}
// 			}
// 		})
// 	}
// }

func TestTopWords(t *testing.T) {
    freq := map[string]int{"go": 5, "is": 3, "great": 1}
 
    t.Run("top 2", func(t *testing.T) {
        result := TopWords(freq, 2)
        if len(result) != 2 {
            t.Fatalf("хотели 2 элемента, получили %d", len(result))
        }
        if result[0].Word != "go" || result[0].Count != 5 {
            t.Errorf("первый элемент: %v, хотели {go 5}", result[0])
        }
        if result[1].Word != "is" || result[1].Count != 3 {
            t.Errorf("второй элемент: %v, хотели {is 3}", result[1])
        }
    })
 
    t.Run("n больше количества слов", func(t *testing.T) {
        result := TopWords(freq, 100)
        if len(result) != 3 {
            t.Errorf("хотели 3 элемента, получили %d", len(result))
        }
    })
}



