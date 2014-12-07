package enigma

import (
	"testing"
)

const (
	Plaintext = "This IS _SAMPLE_ texT. It is used For **TEStinG** - :)"
	Ciphertext = "Nzym FH _FUNVUF_ yfoR. Nf cz wedn Rwu **VUKkojL** - :)"
)

func TestEnigma(t *testing.T) {
	enigma, err := New("main/settings.json")

	if err != nil {
		t.Fatal(err)
	}

	cipher, err := enigma.Encrypt(Plaintext)

	if err != nil {
		t.Fatal(err)
	}

	if cipher != Ciphertext {
		t.Fatal("plaintext -> ciphertext incorrect.")
	}

	enigma, err = New("main/settings.json")

	if err != nil {
		t.Fatal(err)
	}

	plain, err := enigma.Encrypt(cipher)

	if err != nil {
		t.Fatal(err)
	}

	if plain != Plaintext {
		t.Fatal("ciphertext -> plaintext incorrect")
	}
}
