package check

import "context"

// WarnDays: si al certificado le quedan menos días que esto, el chequeo da WARN.
// Coincide con la alerta de diciembre (SSL con menos de 15 días).
const WarnDays = 15

// TLS revisa el certificado SSL/TLS y cuántos días le quedan.
//
// TODO (oct):
//   - Conectar con tls.Dialer{} a domain + ":443" (usa DialContext con ctx).
//   - Leer conn.ConnectionState().PeerCertificates[0].NotAfter.
//   - Días restantes = time.Until(notAfter).Hours() / 24.
//   - Vencido o error de handshake -> StatusFail; < WarnDays -> StatusWarn; resto -> StatusOK.
func TLS(ctx context.Context, domain string) Result {
	return notImplemented()
}
