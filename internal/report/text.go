// Package report da formato a los resultados de los chequeos.
// Hoy solo imprime texto en la terminal; más adelante puede sumar JSON.
package report

import (
	"fmt"
	"io"
	"text/tabwriter"
	"time"

	"github.com/FGP-bit/domwatch/internal/check"
)

// Text escribe una tabla con los resultados.
func Text(w io.Writer, domain string, results []check.Result) {
	fmt.Fprintf(w, "Dominio: %s\n\n", domain)
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "CHEQUEO\tESTADO\tTIEMPO\tDETALLE")
	for _, r := range results {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", r.Name, r.Status, r.Duration.Round(time.Millisecond), r.Message)
	}
	tw.Flush()
}
