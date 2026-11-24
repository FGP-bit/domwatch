// Package check contiene los chequeos que domwatch ejecuta sobre un dominio.
//
// Cada chequeo es una función con la firma de CheckFunc: recibe un dominio y
// devuelve un Result. Así el CLI (`domwatch check`) y, más adelante, el daemon
// con /metrics (`domwatch serve`) usan exactamente la misma lógica.
package check

import (
	"context"
	"time"
)

// Status es el resultado de un chequeo.
type Status string

const (
	StatusOK   Status = "OK"   // todo bien
	StatusWarn Status = "WARN" // funciona, pero requiere atención (ej. SSL vence pronto)
	StatusFail Status = "FAIL" // roto
	StatusSkip Status = "SKIP" // no se ejecutó (ej. todavía no implementado)
)

// Result es lo que devuelve cada chequeo.
type Result struct {
	Name     string        // "dns", "http", "tls", "email"
	Status   Status        // OK, WARN, FAIL, SKIP
	Message  string        // explicación legible para humanos
	Duration time.Duration // cuánto tardó el chequeo
}

// CheckFunc es la firma que comparten todos los chequeos.
type CheckFunc func(ctx context.Context, domain string) Result

// Check asocia un nombre con su función.
type Check struct {
	Name string
	Fn   CheckFunc
}

// All lista los chequeos en el orden en que se ejecutan.
// Para agregar uno nuevo: crea su archivo en este paquete y súmalo aquí.
var All = []Check{
	{"dns", DNS},
	{"http", HTTP},
	{"tls", TLS},
	{"email", Email},
}

// Run ejecuta todos los chequeos sobre un dominio y mide cuánto tarda cada uno.
// (Más adelante: correrlos en paralelo con goroutines.)
func Run(ctx context.Context, domain string) []Result {
	results := make([]Result, 0, len(All))
	for _, c := range All {
		start := time.Now()
		r := c.Fn(ctx, domain)
		r.Name = c.Name
		r.Duration = time.Since(start)
		results = append(results, r)
	}
	return results
}

// notImplemented es un atajo para los chequeos que aún no están escritos.
func notImplemented() Result {
	return Result{Status: StatusSkip, Message: "todavía no implementado"}
}
