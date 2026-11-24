package check

import "context"

// HTTP verifica que el sitio responde.
//
// TODO (oct):
//   - Crear un http.Client con Timeout (ej. 10 s); nunca uses el cliente sin timeout.
//   - Hacer GET a "https://" + domain con http.NewRequestWithContext.
//   - 2xx/3xx -> StatusOK, 4xx -> StatusWarn, 5xx o error de red -> StatusFail.
//   - Incluir el código de estado en Message.
func HTTP(ctx context.Context, domain string) Result {
	return notImplemented()
}
