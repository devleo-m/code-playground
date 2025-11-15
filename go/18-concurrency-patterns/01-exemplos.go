package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ============================================
// FAN-IN: Mesclar múltiplos channels
// ============================================

func fanInSelect(ch1, ch2, ch3 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case msg, ok := <-ch1:
				if !ok {
					ch1 = nil // Desabilitar este case quando channel fechar
				} else {
					out <- msg
				}
			case msg, ok := <-ch2:
				if !ok {
					ch2 = nil
				} else {
					out <- msg
				}
			case msg, ok := <-ch3:
				if !ok {
					ch3 = nil
				} else {
					out <- msg
				}
			}
			// Se todos os channels foram fechados, sair
			if ch1 == nil && ch2 == nil && ch3 == nil {
				return
			}
		}
	}()
	return out
}

func fanInDynamic(channels ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan string) {
			defer wg.Done()
			for msg := range c {
				out <- msg
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func producer(name string, delay time.Duration, count int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 1; i <= count; i++ {
			ch <- fmt.Sprintf("%s: mensagem %d", name, i)
			time.Sleep(delay)
		}
	}()
	return ch
}

// ============================================
// FAN-OUT: Distribuir trabalho
// ============================================

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d: processando job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
		results <- job * 2
	}
}

func fanOutExample() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	numWorkers := 3
	var wg sync.WaitGroup

	// Iniciar workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Enviar jobs
	go func() {
		defer close(jobs)
		for j := 1; j <= 9; j++ {
			jobs <- j
		}
	}()

	// Fechar results quando workers terminarem
	go func() {
		wg.Wait()
		close(results)
	}()

	// Coletar resultados
	for result := range results {
		fmt.Printf("Resultado: %d\n", result)
	}
}

// ============================================
// PIPELINE: Encadear estágios
// ============================================

func generateNumbers(count int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= count; i++ {
			out <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			squared := n * n
			fmt.Printf("Estágio 2: %d² = %d\n", n, squared)
			out <- squared
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return out
}

func format(in <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for n := range in {
			formatted := fmt.Sprintf("Resultado: %d", n)
			fmt.Printf("Estágio 3: %s\n", formatted)
			out <- formatted
		}
	}()
	return out
}

func pipelineExample() {
	numbers := generateNumbers(5)
	squared := square(numbers)
	formatted := format(squared)

	for result := range formatted {
		fmt.Printf("Saída final: %s\n", result)
		fmt.Println("------------------------------")
	}
}

// ============================================
// WORKER POOL: Piscina de trabalhadores
// ============================================

type Job struct {
	ID   int
	Data string
}

type WorkerPool struct {
	workers  int
	jobQueue chan Job
	wg       sync.WaitGroup
}

func NewWorkerPool(workers int, queueSize int) *WorkerPool {
	return &WorkerPool{
		workers:  workers,
		jobQueue: make(chan Job, queueSize),
	}
}

func (wp *WorkerPool) Start() {
	for i := 1; i <= wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for job := range wp.jobQueue {
		fmt.Printf("Worker %d: processando job %d (%s)\n", id, job.ID, job.Data)
		time.Sleep(1 * time.Second)
	}
}

func (wp *WorkerPool) Submit(job Job) {
	wp.jobQueue <- job
}

func (wp *WorkerPool) Stop() {
	close(wp.jobQueue)
	wp.wg.Wait()
}

func workerPoolExample() {
	pool := NewWorkerPool(3, 10)
	pool.Start()

	for i := 1; i <= 9; i++ {
		pool.Submit(Job{
			ID:   i,
			Data: fmt.Sprintf("dados do job %d", i),
		})
		fmt.Printf("Job %d submetido\n", i)
	}

	time.Sleep(5 * time.Second)
	pool.Stop()
	fmt.Println("Worker pool encerrado")
}

// ============================================
// WORKER POOL COM CONTEXT
// ============================================

type WorkerPoolWithContext struct {
	workers  int
	taskChan chan Task
	wg       sync.WaitGroup
}

type Task struct {
	ID   int
	Data string
}

func NewWorkerPoolWithContext(workers int) *WorkerPoolWithContext {
	return &WorkerPoolWithContext{
		workers:  workers,
		taskChan: make(chan Task),
	}
}

func (wp *WorkerPoolWithContext) Start(ctx context.Context) {
	for i := 1; i <= wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(ctx, i)
	}
}

func (wp *WorkerPoolWithContext) worker(ctx context.Context, id int) {
	defer wp.wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: cancelado\n", id)
			return
		case task, ok := <-wp.taskChan:
			if !ok {
				fmt.Printf("Worker %d: channel fechado\n", id)
				return
			}
			fmt.Printf("Worker %d: processando task %d\n", id, task.ID)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func (wp *WorkerPoolWithContext) Submit(task Task) {
	wp.taskChan <- task
}

func (wp *WorkerPoolWithContext) Stop() {
	close(wp.taskChan)
	wp.wg.Wait()
}

func workerPoolWithContextExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pool := NewWorkerPoolWithContext(3)
	pool.Start(ctx)

	go func() {
		for i := 1; i <= 10; i++ {
			pool.Submit(Task{ID: i, Data: fmt.Sprintf("data %d", i)})
			time.Sleep(200 * time.Millisecond)
		}
	}()

	<-ctx.Done()
	pool.Stop()
	fmt.Println("Pool encerrado")
}

// ============================================
// PUB-SUB: Publicar-Assinar
// ============================================

type Message struct {
	Topic   string
	Content string
}

type PubSub struct {
	subscribers map[string][]chan Message
	mutex       sync.RWMutex
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan Message),
	}
}

func (ps *PubSub) Subscribe(topic string) <-chan Message {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	ch := make(chan Message, 10)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)

	return ch
}

func (ps *PubSub) Publish(topic string, content string) {
	ps.mutex.RLock()
	defer ps.mutex.RUnlock()

	msg := Message{Topic: topic, Content: content}

	for _, ch := range ps.subscribers[topic] {
		select {
		case ch <- msg:
		default:
			// Channel cheio, pular (evitar bloqueio)
		}
	}
}

func (ps *PubSub) Close() {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	for _, channels := range ps.subscribers {
		for _, ch := range channels {
			close(ch)
		}
	}
	ps.subscribers = make(map[string][]chan Message)
}

func pubSubExample() {
	pubsub := NewPubSub()

	// Assinante 1: apenas tecnologia
	sub1 := pubsub.Subscribe("tecnologia")
	go func() {
		for msg := range sub1 {
			fmt.Printf("Assinante 1 [%s]: %s\n", msg.Topic, msg.Content)
		}
	}()

	// Assinante 2: tecnologia e esportes
	sub2Tech := pubsub.Subscribe("tecnologia")
	sub2Sports := pubsub.Subscribe("esportes")
	go func() {
		for {
			select {
			case msg := <-sub2Tech:
				fmt.Printf("Assinante 2 [%s]: %s\n", msg.Topic, msg.Content)
			case msg := <-sub2Sports:
				fmt.Printf("Assinante 2 [%s]: %s\n", msg.Topic, msg.Content)
			}
		}
	}()

	// Assinante 3: apenas esportes
	sub3 := pubsub.Subscribe("esportes")
	go func() {
		for msg := range sub3 {
			fmt.Printf("Assinante 3 [%s]: %s\n", msg.Topic, msg.Content)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	pubsub.Publish("tecnologia", "Nova versão do Go lançada!")
	time.Sleep(100 * time.Millisecond)

	pubsub.Publish("esportes", "Brasil vence a copa!")
	time.Sleep(100 * time.Millisecond)

	pubsub.Publish("tecnologia", "Concorrência em Go é incrível!")
	time.Sleep(100 * time.Millisecond)

	pubsub.Publish("esportes", "Jogo de futebol hoje às 20h")
	time.Sleep(1 * time.Second)

	pubsub.Close()
	fmt.Println("Pub-Sub encerrado")
}

// ============================================
// EXEMPLO COMBINADO: Múltiplos padrões
// ============================================

func combinedExample() {
	// 1. Pipeline: gerar -> processar -> formatar
	dataChan := generateData(10)
	processedChan := processData(dataChan)
	formattedChan := formatData(processedChan)

	// 2. Fan-out: distribuir formatação para múltiplos workers
	output1 := workerFunc("Worker 1", formattedChan)
	output2 := workerFunc("Worker 2", formattedChan)
	output3 := workerFunc("Worker 3", formattedChan)

	// 3. Fan-in: mesclar resultados dos workers
	finalOutput := fanInDynamic(output1, output2, output3)

	// Consumir resultado final
	for result := range finalOutput {
		fmt.Println(result)
	}
}

func generateData(count int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= count; i++ {
			out <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return out
}

func processData(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
			time.Sleep(50 * time.Millisecond)
		}
	}()
	return out
}

func formatData(in <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for n := range in {
			out <- fmt.Sprintf("Valor processado: %d", n)
		}
	}()
	return out
}

func workerFunc(name string, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for msg := range in {
			result := fmt.Sprintf("[%s] %s", name, msg)
			out <- result
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return out
}

// ============================================
// MAIN: Executar exemplos
// ============================================

func main() {
	fmt.Println("=== Exemplo 1: Fan-In ===")
	ch1 := producer("Producer 1", 200*time.Millisecond, 3)
	ch2 := producer("Producer 2", 300*time.Millisecond, 3)
	ch3 := producer("Producer 3", 400*time.Millisecond, 3)
	merged := fanInSelect(ch1, ch2, ch3)
	for msg := range merged {
		fmt.Println(msg)
	}
	fmt.Println()

	fmt.Println("=== Exemplo 2: Fan-Out ===")
	fanOutExample()
	fmt.Println()

	fmt.Println("=== Exemplo 3: Pipeline ===")
	pipelineExample()
	fmt.Println()

	fmt.Println("=== Exemplo 4: Worker Pool ===")
	workerPoolExample()
	fmt.Println()

	fmt.Println("=== Exemplo 5: Worker Pool com Context ===")
	workerPoolWithContextExample()
	fmt.Println()

	fmt.Println("=== Exemplo 6: Pub-Sub ===")
	pubSubExample()
	fmt.Println()

	fmt.Println("=== Exemplo 7: Combinado ===")
	combinedExample()
}

