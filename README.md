# Go (Golang) da Zero: Dalle Basi all'Applicazione Pratica

## Introduzione

### Links

[Delve](https://github.com/go-delve/delve)

### Primo programma

```go
package main

import "fmt"

func main(){
	fmt.Println()println("Ciao Mario!")
}
```

```sh
go run main.go
```

### Glossario

- **Array**: una raccolta ordinata di elementi dello stesso tipo.
. **Canale** (Channel): un canale è una struttura di dati utilizzata per la comunicazione tra goroutine in Go.
- **Client HTTP**: in Go, è possibile creare client per effettuare richieste HTTP a server.
- **Concorrenza (Concurrency)**: la concorrenza in Go è la composizione di funzioni indipendentemente eseguibili (goroutine) che comunicano attraverso i canali.
- **Costante (Constant)**: una costante è simile a una variabile, tranne per il fatto che il suo valore non può essere modificato una volta assegnato.
- **Errore (Error)**: un errore è una condizione che indica che qualcosa è andato storto. Gli errori in Go sono rappresentati dall'interfaccia error.
- **Funzione (Function)**: una funzione è un insieme di istruzioni che svolgono un compito specifico. Go supporta funzioni con argomenti e valori di ritorno.
- **Goroutine**: una goroutine è un thread leggero gestito dal runtime di Go.
- **Import**: la parola chiave import viene utilizzata per includere codice da altri pacchetti nel tuo programma.
- **Interfaccia (Interface)**: un'interfaccia è un tipo di dato astratto che definisce un contratto per le funzioni e i metodi che un tipo deve implementare.
- **Interfaccia Vuota**: è un tipo di interfaccia che non ha metodi. Può accettare valori di qualsiasi tipo.
- **Istanza**: un'istanza è un singolo esemplare di un tipo specifico o struttura.
- **Map**: una map è una collezione non ordinata di coppie chiave-valore.
- **Metodo**: un metodo è una funzione associata a un tipo specifico (ad esempio, una struttura).
- **Middleware**: in un'applicazione web, è un componente software che può eseguire operazioni su una richiesta o risposta prima che raggiunga il gestore finale.
- **Pacchetto (Package)**: in Go, i pacchetti sono un modo per raggruppare codice correlato. Ogni programma Go è composto da pacchetti.
- **Panico (Panic)**: un panic è una condizione di errore grave che interrompe l'esecuzione del programma. I panic possono essere causati da errori non gestiti o da condizioni anomale.
- **Puntatore (Pointer)**: un puntatore è una variabile che contiene l'indirizzo di memoria di un'altra variabile.
- **Recupero (Recover)**: il recupero consente di riprendere l'esecuzione di un programma dopo un panico.
- **Riflessione**: in Go, la riflessione consente di esaminare le proprietà dei tipi e delle variabili a runtime.
- **Server HTTP**: in Go, è possibile creare server web per gestire richieste HTTP dai client.
- **Slice**: uno slice è un segmento di un array. Ha una lunghezza variabile, diversamente dagli array.
- **Struttura (Struct)**: in Go, una struttura è un tipo di dato personalizzato utilizzato per raggruppare variabili correlate di diversi tipi.
- **Test (Test)**: i test sono utilizzati per verificare il corretto funzionamento del codice. Go supporta i test unitari e i test integrati.
- **Variabile**: una variabile è un nome assegnato a un'area di memoria per memorizzare valori. Il tipo di valore che può contenere è definito al momento della dichiarazione della variabile.
- **Versione (Version)**: la versione di Go è un numero che identifica la specifica versione del linguaggio.

### Variabili e costanti

- Vaviabile: zona di memoria in cui memorizzare valori, dotata di nome e tipo
- Costante: non possono essere modificate. Non è necessario specificarne il tipo e possono essere correlate usando *iota*

### Tipi

- Primitivi
- Array: sequenza ordinata dello stesso tipo
- Slice: array con dimensione dinamica, utilizzando *append*
- Struttura: raggruppa tipi diversi. Si possono definire metodi su strutture

```go
package main

import "fmt"

type Persona struct {
	Nome  string
	Età   int
	Email string
}

func (p Persona) Saluta() string {
	return "Ciao " + p.Nome
}

func main() {
	p := Persona{
		Nome:  "Mario",
		Età:   51,
		Email: "mario.lazzari@gmail.com",
	}

	saluto := p.Saluta()
	fmt.Println(saluto)
}
```

### Operatori
