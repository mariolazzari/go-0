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

### Esercizio 6

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Order represents a coffee order
type Order struct {
	ID        int
	CreatedAt time.Time
}

// Get random in range
func getRandomTime(from, to int) time.Duration {
	rand := time.Duration(from + rand.Intn(to))
	return rand * time.Millisecond
}

// barista function handles coffee preparation
func barista(id int, orders <-chan Order, ready chan<- int, errs chan<- error) {
	// Defer recovery to handle eventual panics
	defer func() {
		if r := recover(); r != nil {
			errs <- fmt.Errorf("barista %d recovered from panic: %v", id, r)
		}
	}()

	for order := range orders {
		fmt.Printf("Barista %d started order %d\n", id, order.ID)

		// Simulate preparation time between 1 and 3 seconds
		prepTime := getRandomTime(1000, 3000)
		time.Sleep(prepTime)

		// Simulate a random error (10% chance)
		if rand.Intn(10) == 0 {
			errs <- fmt.Errorf("barista %d: spilled coffee for order %d", id, order.ID)
			continue
		}

		// Send ready notification
		ready <- order.ID
	}
}

// customer function sends an order and waits for completion or timeout
func customer(id int, orders chan<- Order, readyChan <-chan int, errs <-chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	order := Order{ID: id, CreatedAt: time.Now()}
	orders <- order

	// Select with 5 seconds timeout
	select {
	case readyID := <-readyChan:
		if readyID == order.ID {
			fmt.Printf("Customer %d: received coffee %d! Happy.\n", id, readyID)
		}
	case err := <-errs:
		fmt.Printf("Customer %d: sorry, error occurred: %v\n", id, err)
	case <-time.After(5 * time.Second):
		fmt.Printf("Customer %d: order %d took too long. Leaving the shop!\n", id, order.ID)
	}
}

func main() {
	// Channels initialization
	orderChan := make(chan Order, 10)
	readyChan := make(chan int, 10)
	errorChan := make(chan error, 10)

	var wg sync.WaitGroup

	// Start 5 baristas
	for i := 1; i <= 5; i++ {
		go barista(i, orderChan, readyChan, errorChan)
	}

	// Simulation of 10 customers arriving at random intervals
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go customer(i, orderChan, readyChan, errorChan, &wg)

		// Random arrival interval between 500ms and 1.5s
		arrivalInterval := getRandomTime(500, 1500)
		time.Sleep(arrivalInterval)
	}

	// Wait for all customers to finish their experience
	wg.Wait()

	close(orderChan)
	fmt.Println("Cafeteria is closing. Goodbye!")
}
```

## Prograzione di rete

### Server TCP

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func gestConn(conn net.Conn) {
	lettore := bufio.NewReader(conn)
	messaggio, err := lettore.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Messaggio dal client:", messaggio)

	risposta := "Messaggio ricevuto"
	conn.Write([]byte(risposta))

	fmt.Println("Messaggio inviato al client:", risposta)
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go gestConn(conn)

	}
}
```

### Client TCP

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "locahlost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "Ciao, server\n")

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Risposta del server: %s\n", response)
}
```

### Esempio UDP

#### Server UDP

```go
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		length, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Ricevuto messaggio dal client: %s\n", string(buffer[:length]))

		message := []byte("Messaggio ricevuto")
		_, err = conn.WriteToUDP(message, clientAddr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
```

#### CLient UDP

```go
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	message := []byte("Ciao, server!")
	_, err = conn.Write(message)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1024)
	length, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Ricevuto la risposta dal server: %s\n", string(buffer[:length]))
}
```

### Server HTTP

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Primo server HTTP in Go")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Esercizio server HTTP

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ciao Mario!")
}

func main() {
	http.HandleFunc("/", handlerRoot)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Metodi HTTP

I metodi HTTP, noti anche come "verbi HTTP", definiscono le operazioni che possono essere eseguite su una determinata risorsa in una rete HTTP. I principali metodi HTTP includono:

- GET: questo metodo viene utilizzato per richiedere dati da una risorsa specificata. Le richieste GET dovrebbero solo ricevere dati e non avere alcun effetto sulle risorse del server.

- POST: questo metodo viene utilizzato per inviare dati a un server per creare una nuova risorsa. I dati vengono inclusi nel corpo della richiesta.

- PUT: viene utilizzato per aggiornare una risorsa esistente con i dati forniti. A differenza del metodo POST, il metodo PUT è idempotente, il che significa che l'invio ripetuto della stessa richiesta avrà lo stesso effetto di una singola richiesta.

- DELETE: questo metodo viene utilizzato per eliminare una risorsa specificata.

- HEAD: è molto simile al metodo GET, ma richiede solo le intestazioni che avrebbero fatto parte della risposta al metodo GET. Può essere utilizzato per verificare l'esistenza di una risorsa o per verificare le intestazioni di risposta prima di scaricare l'intera risorsa.

- OPTIONS: questo metodo viene utilizzato per descrivere le opzioni di comunicazione per la risorsa di destinazione. Può essere utilizzato per determinare i metodi HTTP supportati da un server o per fare una richiesta "preflight" per determinare se una richiesta di origine incrociata (CORS) sarà accettata.

- PATCH: viene utilizzato per applicare modifiche parziali a una risorsa. A differenza del metodo PUT, che richiede di inviare l'intera entità aggiornata, il metodo PATCH richiede solo di inviare le modifiche specifiche da applicare all'entità.

- CONNECT: questo metodo viene utilizzato da un client per stabilire una rete di tunnel verso un host di destinazione.

- TRACE: questo metodo esegue un test di loopback del messaggio di richiesta lungo il percorso della risorsa di destinazione.

Ciascuno di questi metodi può essere utilizzato per operazioni specifiche su una risorsa HTTP, e capire come e quando utilizzarli può essere essenziale per la creazione di applicazioni web robuste e sicure.

### CLient HTTP

```go
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

	post := `{
"userId": 1,
"id": 1,
"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
"body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
}`

	resp, err = http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", strings.NewReader(post))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
```

### Framework web

- Gorilla mux
- Gin

## Applicazioni Web

### Pacchetti utilizzati

Qui di seguito ti elenco dei pacchetti di Go utili per creare app web (non ho inserito quelli già visti ovvero net, net/http, mux e gin).

#### html/template

Il pacchetto html/template fornisce le funzionalità per creare viste HTML dinamiche. Le viste HTML vengono utilizzate per visualizzare i contenuti di un'applicazione web.

#### database/sql

Il pacchetto database/sql fornisce le funzionalità per interagire con i database. Questo pacchetto consente di accedere ai dati da un database e di aggiornare i dati in un database.

#### gorm

Il pacchetto gorm è un ORM (Object-Relational Mapping) per Go. ORM aiutano a convertire i dati tra oggetti e tabelle di database.

#### crypto/rand

Il pacchetto crypto/rand fornisce le funzionalità per generare numeri casuali. I numeri casuali possono essere utilizzati per una varietà di scopi, tra cui la generazione di token di sicurezza e la crittografia.

#### log

Il pacchetto log fornisce le funzionalità per registrare messaggi di log. I messaggi di log possono essere utilizzati per tracciare l'attività di un'applicazione web e per rilevare eventuali problemi.

#### os

Il pacchetto os fornisce le funzionalità per interagire con il sistema operativo. Questo pacchetto consente di accedere ai file e alle directory dal sistema operativo, e di eseguire comandi del sistema operativo.

#### io

Il pacchetto io fornisce le funzionalità per leggere e scrivere file. Questo pacchetto consente di accedere ai dati da un file e di scrivere dati in un file.

#### time

Il pacchetto time fornisce le funzionalità per gestire il tempo. Questo pacchetto consente di ottenere l'ora corrente, di misurare il tempo di esecuzione di un codice, e di pianificare eventi futuri.

### Sessioni

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour)

	cookie := http.Cookie{
		Name:    "FirstVisit",
		Value:   "1",
		Expires: expiration,
	}

	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "Cookie impostato. Per favore vai alla pagina /check-cookie per verificare il cookie.")
}

func checkCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("FirstVisit")

	if cookie != nil {
		fmt.Fprintf(w, "Bentornato! Hai visitato questo sito %s volta\n", cookie.Value)
	} else {
		fmt.Fprintf(w, "Benvenuto! Non hai mai visitato questo sito prima o il tuo cookie è scaduto.")
	}
}

func main() {
	http.HandleFunc("/set-cookie", setCookie)
	http.HandleFunc("/check-cookie", checkCookie)

	fmt.Println("Server avviato su localhost:8080")

	http.ListenAndServe(":8080", nil)
}
```

### Gestione input

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHanlder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	nome := r.FormValue("nome")
	email := r.FormValue("email")

	// validazione
	if nome == "" || email == "" {
		fmt.Fprintf(w, "Nome e email sono obbligatori")
	}

	fmt.Fprintf(w, "Nome: %s\nEmail: %s", nome, email)
}

func main() {
	http.HandleFunc("/", formHanlder)
	http.ListenAndServe(":8080", nil)
}
```

### Database

#### Driver 

```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// select
	rows, err := db.Query("Select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	}

	// insert
	stmt, err := db.Prepare("insert into users(name) values(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Mario")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Last ID:", lastId)
}
```

#### Gorm

```sh
go get -u gorm.io/driver/sqlite
```

```go
package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "Mario"})
}
```

### Transazioni

```go
package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec("insert into users(name) values(?)", "Mario")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
```

### Web CRUD

```go
package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("non riesco a collegarmi con il database")
	}

	db.AutoMigrate(&User{})
}

func main() {
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		db.Create(&User{Name: name})
		fmt.Fprintf(w, "Utente creato")
	})

	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		var users []User
		db.Find(&users)

		fmt.Fprintf(w, "Lista utente:\n")
		for _, user := range users {
			fmt.Fprintf(w, user.Name, "\n")
		}
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		name := r.URL.Query().Get("name")

		var user User
		db.First(&user, id)
		user.Name = name
		db.Save(&user)

		fmt.Fprintf(w, "Utente aggiornato")
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		db.Delete(&User{}, id)

		fmt.Fprintf(w, "Utente cancellato")
	})

	http.ListenAndServe(":8080", nil)
}
```

### Template HTML

```go
package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title  string
	Header string
	Body   string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := &Page{
			Title:  "Benvenuto",
			Header: "Ciao, mondo!",
			Body:   "Benvenuto nella mia applicazione!",
		}

		t, _ := template.ParseFiles("layout.html", "content.html")

		t.ExecuteTemplate(w, "content", p)
	})

	http.ListenAndServe(":8080", nil)
}
```

```html
{{define "content"}}
<h1>{{.Header}}</h1>
<p>{{.Body}}</p>
{{end}}
```

```html
<!doctype html>
<html>
    <head>
        <title>{{.Title}}</title>
    </head>
    <body>
        {{template "content" .}}
    </body>
</html>
```

### Autenticazione

#### Basic 

```go
package main

import (
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok || username != "admin" || password != "password" {
		w.Header().Set("WWW-Authenticate", `Basic realm="Inserire utente e password"`)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Non sei autorizzato")
	}
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.ListenAndServe(":8080", nil)
}
```

#### O-Auth

## Progetto gestione utenti

### Setup

```sh
go mod init user_app
```

```go
package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WIP")
	})
	http.ListenAndServe(":8080", nil)
}
```

### Connessione database

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

var db *sql.DB
var sessionMap = make(map[string]int)

func init() {
	var err error
	connectionString := "user=poastgres password=password dbname=users sslmode=disable"
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connesso a Postgres")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WIP")
	})
	http.ListenAndServe(":8080", nil)
}
```

### Login

```go
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

var db *sql.DB
var sessionMap = make(map[string]int)

func init() {
	var err error
	connectionString := "user=poastgres password=password dbname=users sslmode=disable"
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connesso a Postgres")
}

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			t, _ := template.ParseFiles("login.html")
			t.Execute(w, nil)
		}
		fmt.Fprintf(w, "WIP")

		name := r.FormValue("name")
		password := r.FormValue("password")

		var user User
		err := db.QueryRow("select id, name, email from users where name = $1 and password = $2", name, password).Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			http.Error(w, "Credenziali non valide", 400)
		}

		sessionID := uuid.New().String()
		cookie := &http.Cookie{
			Name:  "session_id",
			Value: sessionID,
		}

		http.SetCookie(w, cookie)
		sessionMap[sessionID] = user.ID
		http.Redirect(w, r, "/user", http.StatusSeeOther)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WIP")
	})
	http.ListenAndServe(":8080", nil)
}
```

### Registrazione

```go
```
