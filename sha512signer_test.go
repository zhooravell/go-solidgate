package solidgate

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSha512SineGenerator_Generate(t *testing.T) {
	s := NewSha512Signer("unicorn", []byte("20c20ee3-4173-4daa-87e5-dbcce8c7949d"))

	p := make(map[string]interface{})
	p["amount"] = 105
	p["currency"] = "EUR"

	jsonBytes, err := json.Marshal(p)

	if err != nil {
		t.Error(err)
	}

	sing, err := s.Sine(jsonBytes)

	if err != nil {
		t.Error(err)
	}

	expected := "YjZkZTc2Yzk2ZjNlYzUzNjFmM2JiYjljNTJkNGVhYTZjYmNlOGJiZDVkNjljMTQ1MWFmZmNhNjBkNzQ0YjkyY2NjYzljM2E0NWEwNzE3NWEwNWFiMzgyOWNkZTk0ZDA2MzAyNTJmYzkzMWM2NThiYmYxZDFhZDUzMGRjOTgwNjE="
	if expected != sing {
		t.Fail()
	}
}

func BenchmarkSha512SineGenerator(b *testing.B) {
	s := NewSha512Signer("unicorn", []byte("20c20ee3-4173-4daa-87e5-dbcce8c7949d"))

	for n := 0; n < b.N; n++ {
		_, err := s.Sine([]byte(`{"amount":105,"currency":"EUR"}`))

		if err != nil {
			log.Panicln(err)
		}
	}
}
