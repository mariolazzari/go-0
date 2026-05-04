package utenti

type Utente struct {
	Nome string
	Mail string
}

func NuovoUtente(nome, mail string) *Utente {
	return &Utente{
		Nome: nome,
		Mail: mail,
	}
}
