package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"sync"
)

// @formatter:off

type TrafficManager struct {
	workPiece   int
	calculating bool
	radical     int
	output      *widget.Label
	stop        chan bool
	mu          sync.Mutex
}

func NewTrafficManager(output *widget.Label) *TrafficManager {
	return &TrafficManager{
		workPiece:   0,
		calculating: false,
		radical:     2,
		output:      output,
		stop:        make(chan bool, 1),
	}
}

func (m *TrafficManager) SetWorkPiece(val int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.workPiece = val
}

func (m *TrafficManager) SetRadical(val int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.radical = val
}

func (m *TrafficManager) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.workPiece = 0
	m.calculating = false
}

func (m *TrafficManager) IsCalculating() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.calculating
}

func (m *TrafficManager) SetCalculating(val bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calculating = val
}

func (m *TrafficManager) UpdateOutput(text string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.output.SetText(text)
	fyne.CurrentApp().Driver().CanvasForObject(m.output).Refresh(m.output)
}

func (m *TrafficManager) SetOutput(output *widget.Label) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.output = output
}

func (m *TrafficManager) GetWorkPiece() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.workPiece
}

func (m *TrafficManager) GetRadical() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.radical
}

func (m *TrafficManager) Stop() {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.calculating {
		m.stop <- true
	}
}

func (m *TrafficManager) ShouldStop() bool {
	select {
	case <-m.stop:
		return true
	default:
		return false
	}
}