package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuário cadastrado no sistema
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Verifica se campo obrigatório não está vazio
func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}
	if usuario.Senha == "" {
		return errors.New("o senha é obrigatório e não pode estar em branco")
	}

	return nil
}

// Retiras os espaços em brancos do inicio/fim da string
func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}

// Chama métodos validar() e formatar
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}
	usuario.formatar()
	return nil
}
