package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "базовый случай", s: "()", want: true},
		{name: "все типы скобок", s: "()[]{}", want: true},
		{name: "вложенные скобки", s: "{[]}", want: true},
		{name: "неправильный порядок", s: "(]", want: false},
		{name: "перепутан порядок закрытия", s: "([)]", want: false},
		{name: "глубокая вложенность", s: "{[()]}", want: true},
		{name: "пустая строка", s: "", want: true},
		{name: "только открывающая", s: "(", want: false},
		{name: "только закрывающая", s: ")", want: false},
		{name: "закрывающая при пустом стеке", s: "){", want: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValid(tc.s)
			if got != tc.want {
				t.Errorf("isValid(%q) = %v, хотели %v", tc.s, got, tc.want)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	s := ""
	for i := 0; i < 1000; i++ {
		s += "({[]})"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsValid(s)
	}
}