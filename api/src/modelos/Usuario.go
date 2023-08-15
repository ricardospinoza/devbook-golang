package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario respresenta um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty`
	Nome     string    `json:nome,omitempty`
	Nick     string    `json:nick,omitempty`
	Email    string    `json:email,omitempty`
	Senha    string    `json:senha,omitempty`
	CriadoEm time.Time `json:criadoEm,omitempty`
}

func (usuario *Usuario) validar() error {

	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório")
	}
	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório")
	}
	if usuario.Email == "" {
		return errors.New("O e-mail é obrigatório")
	}
	if usuario.Senha == "" {
		return errors.New("A senha é obrigatória")
	}
	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
}

// Preparar vai chamar o metodos de validar e formatar usuario recebido
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}
	usuario.formatar()
	return nil
}
