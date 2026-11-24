package check

import "context"

// DNS verifica que el dominio resuelve.
//
// TODO (oct):
//   - Resolver registros A/AAAA con net.DefaultResolver.LookupHost(ctx, domain).
//   - Si no hay respuesta -> StatusFail; si resuelve -> StatusOK con las IPs en Message.
//   - Extra: revisar NS con LookupNS y MX con LookupMX.
func DNS(ctx context.Context, domain string) Result {
	return notImplemented()
}
