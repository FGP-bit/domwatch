// Comando domwatch: monitor de dominios (DNS, HTTP, SSL, SPF/DKIM/DMARC).
//
// Uso:
//
//	domwatch check <dominio>   ejecuta todos los chequeos una vez
//	domwatch version           muestra la versión
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/FGP-bit/domwatch/internal/check"
	"github.com/FGP-bit/domwatch/internal/report"
)

// version se reemplaza al compilar con:
//
//	go build -ldflags "-X main.version=v0.1.0"
//
// (en noviembre lo hará GoReleaser automáticamente con cada tag).
var version = "dev"

const usage = `domwatch: monitor de dominios

Uso:
  domwatch check <dominio>   ejecuta todos los chequeos una vez
  domwatch version           muestra la versión
  domwatch help              muestra esta ayuda
`

func main() {
	os.Exit(run(os.Args[1:]))
}

// run devuelve el código de salida: 0 = todo bien, 1 = algún chequeo falló, 2 = uso incorrecto.
func run(args []string) int {
	if len(args) == 0 {
		fmt.Fprint(os.Stderr, usage)
		return 2
	}

	switch args[0] {
	case "check":
		if len(args) != 2 {
			fmt.Fprintln(os.Stderr, "uso: domwatch check <dominio>")
			return 2
		}
		return runCheck(args[1])
	case "version":
		fmt.Println("domwatch", version)
		return 0
	case "help", "-h", "--help":
		fmt.Print(usage)
		return 0
	default:
		fmt.Fprintf(os.Stderr, "comando desconocido: %q\n\n%s", args[0], usage)
		return 2
	}
}

func runCheck(domain string) int {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	results := check.Run(ctx, domain)
	report.Text(os.Stdout, domain, results)

	for _, r := range results {
		if r.Status == check.StatusFail {
			return 1
		}
	}
	return 0
}
