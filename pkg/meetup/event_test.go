package meetup

import "testing"

func TestNewEvent(t *testing.T) {
	e, _ := NewEvent("https://www.meetup.com/pt-BR/Golang-Campinas/events/272398725/")

	if e.Group != "Golang-Campinas" {
		t.Error("unexpected meetup group")
	}
	if e.ID != "272398725" {
		t.Error("unexpected meetup event ID")
	}
}

func TestInvalidNewEvent(t *testing.T) {
	_, err := NewEvent("https://globoesporte.globo.com/blogs/meia-encarnada/post/2020/02/13/corinthians-eliminado-pelo-guarani-desilusao-mas-nao-vergonha.ghtml")

	if err == nil {
		t.Error("expected error")
	}
}
