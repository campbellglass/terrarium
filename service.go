package main

import "fmt"

// A service is a process that runs on a machine that clusters can send data to
type Service interface {
  Name() string // placeholder method
}

// An Id is used to uniquely identify a Service
// No two Services should have the same Id
// TODO: Add network ip:port to Id
type Id struct {
  name string
}

// Creates a new Id
func NewId(name string) Id {
  id := Id{
    name: name,
  }
  return id
}

// Gets the name of this Id
func (id *Id) Name() string {
  return id.name
}

// An Announcer is a Service that takes data and Announces it
// Announcing is logging to a file for now
type Announcer struct {
  Id // Unique identifier of this Service
  incoming chan []byte // TODO: create a network handler that feeds bytes from the network into this channel
}

// Creates and returns a pointer to a new Announcer
func NewAnnouncer(id Id) *Announcer {
  announcer := Announcer{
    Id: id,
    incoming: make(chan []byte),
  }
  return &announcer
}

// Runs this Announcer
func (announcer *Announcer) Run() {
  fmt.Printf("Running Global Announcer with name '%s'\n", announcer.Name())
  go func() {
    for {
      announcement := string(<-announcer.incoming)
      toWrite := fmt.Sprintf("%s\n", announcement)
      fmt.Printf(toWrite)  // TODO: change this to a file write
    }
  }()
  return
}
