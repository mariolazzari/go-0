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

## Fondamenti

### Funzione

Blocco di codice che può essere riutilizzato

```go
package main

import (
	"errors"
	"fmt"
)

func somma(a, b int) int {
	return a + b
}

func divisione(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divisione per 0")
	}
	return a / b, nil
}

func main() {
	risultato := somma(5, 3)
	fmt.Println(risultato)

	risultatoDiv, err := divisione(5, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(risultatoDiv)
	}

	_, err = divisione(5, 0)
	if err != nil {
		fmt.Println(err)
	}
}
```

### Pacchetto

```sh
go mod init github/mariolazzari/go-0
```

```go
package main

import "github/mariolazzari/go-0/saluta"

func main() {
	saluta.Saluta("Mario")
}
```

### Array e slice

[Go blog](https://go.dev/blog/slices)

### Mappe

```go
package main

import (
	"fmt"
	"sort"
)

var CatalogoLibri map[string]int

func AggiungiLibro(titolo string, pagine int) {
	CatalogoLibri[titolo] = pagine
}

func RimuoviLibro(titolo string) {
	delete(CatalogoLibri, titolo)
}

func AggiornaLibro(titolo string, pagine int) {
	if _, exists := CatalogoLibri[titolo]; exists {
		CatalogoLibri[titolo] = pagine
	}
}

func StampaCatalogo() {
	var chiavi []string
	for k := range CatalogoLibri {
		chiavi = append(chiavi, k)
	}

	sort.Strings(chiavi)
	for _, key := range chiavi {
		fmt.Printf("%s - %d\n", key, CatalogoLibri[key])
	}
}

func main() {
	CatalogoLibri = make(map[string]int)

	CatalogoLibri["Programmazione Go"] = 375
	CatalogoLibri["Go Avanzato"] = 520
	CatalogoLibri["Go per Principianti"] = 200

	AggiungiLibro("Go Concurrency", 300)
	RimuoviLibro("Go per Principianti")
	AggiornaLibro("Go Avanzato", 530)
	StampaCatalogo()

}
```

## Puntatori

### Puntatori e funzioni

```go
package main

import "fmt"

func cambiaValore(ptr *int) {
	*ptr = 100
}

func newInt() *int {
	return new(int)
}

func main() {
	x := 10
	fmt.Println(x)

	cambiaValore(&x)
	fmt.Println(x)

	ptr := newInt()
	*ptr = 3
	fmt.Println(*ptr)

}
```

### Puntatori a strutture

```go
package main

import "fmt"

type Persona struct {
	Nome string
	Età  int
}

type Impiegato struct {
	Persona
	Posizione string
}

func (p *Persona) compleanno() {
	p.Età++
}

func main() {
	p := Persona{Nome: "Mario", Età: 51}

	fmt.Println(p.Nome)
	fmt.Println(p.Età)

	p.compleanno()
	fmt.Println(p.Età)

	e := Impiegato{
		Persona: Persona{
			Nome: "Mario",
			Età:  50,
		},
		Posizione: "Programmatore",
	}

	e.compleanno()
	fmt.Println(e.Nome)
	fmt.Println(e.Età)

	// struttura anonima
	punto := struct{ x, y int }{x: 10, y: 10}
	fmt.Println(punto)
}
```

### Metodi

```go
package main

import (
	"fmt"
	"math"
)

type Cerchio struct {
	raggio float64
}

func (c Cerchio) Area() float64 {
	return math.Pow(c.raggio, 2) * math.Pi
}

func (c Cerchio) Circonferenza() float64 {
	return c.raggio * 2 * math.Pi
}

func main() {
	c := Cerchio{raggio: 10}
	fmt.Printf("Area  cerchio: %.2f\n", c.Area())
	fmt.Printf("Circonferenza: %.2f\n", c.Circonferenza())
}
```

### Interfacce

```go
package main

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

type Quadrato struct {
	Lato float64
}

func (q Quadrato) Area() float64 {
	return math.Pow(q.Lato, 2)
}

func StampaArea(f Forma) {
	fmt.Println("L'area della forma è:", f.Area())
}

func main() {
	q := Quadrato{Lato: 10}
	StampaArea(q)
}
```

### Esercizio 3

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// Definizione di Libro
type Libro struct {
	Titolo string
	Autore string
	Anno   int
}

// Metodo Presentazione per Libro
func (l Libro) Presentazione() string {
	return fmt.Sprintf("Titolo: %s\nAutore: %s\nAnno  : %d\n", l.Titolo, l.Autore, l.Anno)
}

// Interfaccia Biblioteca
type Biblioteca interface {
	AggiungiLibro(libro *Libro)
	PresentaLibri() string
}

// Implementazione concreta
type BibliotecaConcreta struct {
	Libri []*Libro
}

// AggiungiLibro
func (b *BibliotecaConcreta) AggiungiLibro(libro *Libro) {
	b.Libri = append(b.Libri, libro)
}

// PresentaLibri
func (b *BibliotecaConcreta) PresentaLibri() string {
	var presentazioni []string

	for _, libro := range b.Libri {
		presentazioni = append(presentazioni, libro.Presentazione())
	}

	sort.Strings(presentazioni)

	return strings.Join(presentazioni, "\n")
}

// Esempio utilizzo
func main() {
	b := &BibliotecaConcreta{}

	b.AggiungiLibro(&Libro{
		Titolo: "Viaggio al termine della notte",
		Autore: "Celine",
		Anno:   1932,
	})

	b.AggiungiLibro(&Libro{
		Titolo: "1984",
		Autore: "George Orwell",
		Anno:   1949,
	})

	fmt.Println(b.PresentaLibri())
}
```

## File in Go

### Pacchetto ioutil

Il pacchetto **ioutil** che utilizzo negli esempi della sezione è deprecato. Questo significa che, sebbene ancora utilizzabile per compatibilità, ioutil non riceverà più aggiornamenti e potrebbe essere rimosso nelle future versioni di Go.

1. *ioutil.ReadFile* ➔ sostituita da *os.ReadFile*
2. *ioutil.WriteFile* ➔ sostituita da *os.WriteFile*
3. *ioutil.ReadAll* ➔ sostituita da *io.ReadAll*
4. *ioutil.WriteAll* ➔ sostituita da *io.WriteAll*
5. *ioutil.TempFile* ➔ sostituita da *os.CreateTemp*
6. *ioutil.TempDir* ➔ sostituita da *os.MkdirTemp*

Il pacchetto ioutil è stato inizialmente pensato per semplificare la gestione di input e output, ma con l'evoluzione del linguaggio Go, molte di queste funzionalità sono state integrate direttamente in *os* e *io*.

Vi chiedo di utilizzare i nuovi metodi di *os* e *io* al posto di ioutil. In questo modo il vostro codice sarà più allineato agli standard attuali.

```go
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
````

### File di grandi dimensioni

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("largefile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
```

### Scrivere file

```go
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("Ciao, Mario!")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Directory

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// creazione dir
	err := os.Mkdir("example_dir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// leggere contenuto dir
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
```

### Esempio operazioni

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Mkdir("testDir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("testDir/testFile.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString("Ciao, Mario!")
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile("testDir/testFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
```

### Esercizio 4

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Crea una nuova directory chiamata "Compito" e controlla sempre gli errori.
	err := os.Mkdir("Compito", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Crea un nuovo file chiamato "testo.txt" all'interno della directory "Compito".
	file, err := os.Create("Compito/testo.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Scrivi la frase "Buon lavoro con Go!" nel file "testo.txt".
	// Ricorda di chiudere il file dopo aver completato la scrittura.
	_, err = file.WriteString("Buon lavoro con Go!")
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Leggi il contenuto del file "testo.txt" e stampa il contenuto sul terminale.
	data, err := os.ReadFile("Compito/testo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	// Leggi e stampa i nomi di tutti i file presenti nella directory "Compito".
	// Aiutati con ciclo for per iterare sui file.
	files, err := os.ReadDir("Compito")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
```

## Moduli

### Modulo

```sh
go mod init
go get github.com/gin-gonic/gin
go mod tidy 
go run .
```

### Test e build

```sh
go test
go build main.go
go build main.go -o app
```

```go
package main

import "testing"

func Add(a int, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}
}
```

### go.mod e go.sum

- go.mod: informazione pacchetti
- go.sum: somme di controllo crittografiche

```sh
go get
go tidy
```

### Importazioni

```go
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
```

```go
package main

import (
	"fmt"
	"pack/utenti"
)

func main() {
	utente := utenti.NuovoUtente("Mario", "mario,lazzari@gmail.com")
	fmt.Println(utente.Nome)
}
```

## Concorrenza

### Inroduzione alla concorrenza

- Concorrenza: gestione indipendente di più task
- Parallelismo: esecuzione simultanea di più task

#### Modello CSP

- Tony Hoare 1978
- Communicating sequelntial processes
- Comunicazione tramte canali fra processi indipendenti
- Canali trasportano dati e controllo degli stessi

### Goroutines 1

- Goroutine: thread leggero gestito dal runtine di Go
- Thread: più piccola unità di lavoro gestita dal sistema operativo

```go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("Lazzari")
	say("Mario")
}
```

### Goroutines 2

- Done(): fine lavoro
- Wait(): bloccare fino a termine lavoro

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
```

### Mutex

```go
package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (sc *SafeCounter) Inc(key string) {
	sc.mux.Lock()
	// solo una goroutine alla volta può accedere a sc.v
	sc.v[key]++
	sc.mux.Unlock()
}

func (sc *SafeCounter) Value(key string) int {
	sc.mux.Lock()
	defer sc.mux.Unlock()

	// solo una goroutine alla volta può accedere a sc.v
	return sc.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup

	/* 	for range 1000 {
	wg.Add(1)
	go func() {
		c.Inc("myKey")
		wg.Done()
	}()
	}*/

	for range 1000 {
		wg.Go(func() {
			c.Inc("myKey")
		})
	}

	wg.Wait()

	fmt.Println(c.v["myKey"])
}
```

### Esercizio 5

```go
package main

import (
	"fmt"
	"sync"
)

// var counter int = 0
type SafeCounter struct {
	counter int
	mux     sync.Mutex
}

func saluta(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Ciao, %s!\n", name)
}

func main() {
	var wg1, wg2 sync.WaitGroup

	nomi := []string{"Alice", "Bruno", "Carlo", "Davide"}

	for _, nome := range nomi {
		wg1.Add(1)
		go saluta(nome, &wg1)
	}
	wg1.Wait()

	sc := SafeCounter{}
	for range 1000 {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			sc.mux.Lock()
			defer sc.mux.Unlock()
			sc.counter++
		}()
	}

	wg2.Wait()
	fmt.Println("Counter =", sc.counter)
}
```

### Canali

```go
package main

import (
	"fmt"
	"time"
)

func worker(msg string, ch chan string) {
	time.Sleep(time.Second)
	ch <- msg
}

func main() {
	messages := make(chan string)

	go worker("Lavoro completato", messages)
	msg := <-messages
	fmt.Println(msg)
}
```

### Canali bufferizzati

- canale non bufferizzato: può contenere un solo valore
  - blocca la goroutine fino a quando il valore viene ricevuto
- canale bufferizzato: può contenere più valori, fino alla sua capacità massima
  - la goroutine si blocca quando il buffer è pèieno
- close(): chiude un canale in modo da impedire nuovi invii sullo stesso
  - range può ricevere valori fino a quando il canale non viene chiuso
- deadlock: un valore viene inviato al canale ma non ci sono ricezioni dallo stesso

### Seleziona e timeout

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}
```

### Patterns di concorrenza


#### Fan-out fan-in

- Fan-out: suddivide il carico tra diverse goroutine
- Fan-in: raccoglie i risultati da più goroutine in un unico canale

```go
package main

import (
	"fmt"
)

// La funzione calcolo prende un id intero e un canale out in cui inviare i dati.
// Calcola il prodotto di id e i, poi invia il risultato nel canale out.
func calcolo(id int, out chan<- int) {
	for i := range 5 {
		out <- id * i
	}
}

func main() {
	// Creazione di due canali, c1 e c2.
	c1 := make(chan int)
	c2 := make(chan int)

	// Avvio di due goroutine separate. Ognuna esegue la funzione calcolo
	// con diversi valori di id e invia i risultati a un canale separato.
	// Questo è un esempio di "fan-out".
	go calcolo(2, c1)
	go calcolo(3, c2)

	// Creazione di un terzo canale, c.
	c := make(chan int)

	// Avvio di una terza goroutine che riceve i risultati dai canali c1 e c2.
	// Non appena un risultato è disponibile in uno dei due canali,
	// lo invia al canale c. Questo è un esempio di "fan-in".
	go func() {
		for {
			select {
			case v := <-c1:
				c <- v
			case v := <-c2:
				c <- v
			}
		}
	}()

	// Il ciclo for legge e stampa i primi 10 valori dal canale c.
	for range 10 {
		fmt.Println(<-c)
	}

}
```

#### Workers pool

- Coda di task da eseguire
- Più goroutine estraggono dalla coda i task da eseguire

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, lavori <-chan int, risultati chan<- int) {
	// Il ciclo for continua ad estrarre i lavori dal canale "lavori" finché non è vuoto e chiuso.
	for j := range lavori {
		fmt.Println("worker", id, "lavoro iniziato", j)
		// Simula un lavoro lungo un secondo.
		time.Sleep(time.Second)
		fmt.Println("worker", id, "lavoro finito", j)
		// Invia il risultato del lavoro al canale "risultati".
		risultati <- j * 2
	}
}

func main() {
	const numLav = 5
	// Crea un canale per i lavori e un canale per i risultati.
	// Ogni canale ha spazio per "numLav" elementi.
	lavori := make(chan int, numLav)
	risultati := make(chan int, numLav)

	// Avvia tre worker, identificati da ID 1, 2, 3.
	for w := 1; w <= 3; w++ {
		// Ogni goroutine ha il suo ID unico "w" e condivide i canali "lavori" e "risultati".
		go worker(w, lavori, risultati)
	}

	// Inserisce i lavori nel canale "lavori".
	for j := 1; j <= numLav; j++ {
		lavori <- j
	}

	close(lavori)

	for a := 1; a <= numLav; a++ {
		<-risultati
	}
}
```

### Gestione errori

Un errore in una goroutine manda in panic del programma.
Si crea un canale dedicato agli aerori.

### Error channel

```go
package main

import (
	"fmt"
	"time"
)

// funzione che potrebbe fallire
func performTask(i int, errCh chan<- error) {
	// simuliamo un errore per i pari
	if i%2 == 0 {
		errCh <- fmt.Errorf("Errore con l'input %d", i)
		return
	}

	// Se non ci sono errori, simuliamo un operazione con un ritardo
	time.Sleep(2 * time.Second)
	fmt.Printf("Task %d completato con successo\n", i)
}

func main() {
	numTasks := 5
	errCh := make(chan error, numTasks)

	for i := range numTasks {
		go performTask(i, errCh)
	}

	for i := range numTasks {
		select {
		case err := <-errCh:
			if err != nil {
				fmt.Printf("Ricevuto un errore: %s\n", err)
			}
		case <-time.After(5 * time.Second):
			fmt.Printf("Task %d non ha completato in tempo\n", i)
		}
	}
}
```

#### Error group

```go
package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	var urls = []string{
		"http://www.google.com",
		"http://www.yahoo.com",
		"http://www.bing.com",
	}
	for _, url := range urls {
		// Copia l'URL corrente nella variabile locale per evitare problemi di sincronizzazione dei dati.
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Recuperati con successo tutti gli URL.")
	} else {
		fmt.Println("Errore riscontrato:", err)
	}
}
```
