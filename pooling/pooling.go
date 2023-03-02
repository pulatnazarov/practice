package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Config struct {
	MaxConnections int
	ConnString     string
}

type Connection struct {
	value      int32
	connString string
	ctx        context.Context
}

type Pool struct {
	conns []*Connection
	mutex sync.Mutex
}

func NewPool(conf Config, ctx context.Context) *Pool {
	arrOfConnections := make([]*Connection, conf.MaxConnections, conf.MaxConnections)
	for ind := range arrOfConnections {
		arrOfConnections[ind] = &Connection{
			value:      int32(ind),
			connString: conf.ConnString,
			ctx:        ctx,
		}
	}

	pool := &Pool{
		conns: arrOfConnections,
	}

	return pool
}

func (s *Pool) Push(c *Connection) {
	s.conns = append(s.conns, c)
}

func (s *Pool) Pop() *Connection {
	c := s.conns[len(s.conns)-1]
	s.conns = s.conns[:len(s.conns)-1]
	return c
}

func (s *Pool) Top() *Connection {
	return s.conns[len(s.conns)-1]
}

func (s *Pool) Empty() bool {
	return len(s.conns) == 0
}

func (c *Connection) isAlive() bool {
	return c.value > 0
}

func (c *Connection) Connect(connString string) {
	fmt.Println("Connecting to ", connString)
	c.value = 1
}

func (s *Pool) Acquire() *Connection {
	s.mutex.Lock()

	for s.Empty() {
		s.mutex.Unlock()
		time.Sleep(1 * time.Nanosecond)
		s.mutex.Lock()
	}
	conn := s.Top()
	s.Pop()

	if !conn.isAlive() {
		conn.Connect(conn.connString)
	}

	return conn
}

func (s *Pool) Release(conn *Connection) {
	s.Push(conn)
}

func doSomething(dbPool *Pool) {
	for i := 0; i < 100; i++ {
		conn := dbPool.Acquire()

		time.Sleep(100 * time.Millisecond)

		dbPool.Release(conn)
	}
	fmt.Println("Done")
}

func main() {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbPool := NewPool(Config{
		MaxConnections: 20,
		ConnString:     "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	}, ctx)

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		go func() {
			doSomething(dbPool)
			wg.Done()
		}()

		wg.Add(1)
	}

	wg.Wait()

	for _, elem := range dbPool.conns {
		fmt.Println(elem.value)
	}

	end := time.Since(start)
	fmt.Println(end.Seconds())
}

// Stack
//

// 90 connection poolling
//

// max_connections = 100
// buffer = 10

// 1 request
// 1 request
// ... requests

// 10 connections
// 2 requests

// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire

// Done
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire
// Acquire

//

//
