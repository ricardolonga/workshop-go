package main

type PessoaService interface {
	Save(pessoa *Pessoa) error
}

type PessoaServiceImpl struct {
}

func (psi *PessoaServiceImpl) Save(pessoa *Pessoa) error {
	return nil
}
