package check

import "context"

// Email revisa la autenticación de correo del dominio: SPF, DMARC y DKIM.
//
// TODO (oct):
//   - SPF: net.DefaultResolver.LookupTXT(ctx, domain) y buscar un TXT que empiece con "v=spf1".
//   - DMARC: LookupTXT de "_dmarc." + domain y buscar "v=DMARC1".
//   - DKIM: necesita un selector (ej. "default", "google", "selector1"):
//     LookupTXT de selector + "._domainkey." + domain y buscar "v=DKIM1".
//     Puedes recibir el selector como flag del CLI más adelante.
//   - Falta SPF o DMARC -> StatusWarn; ninguno -> StatusFail; todos -> StatusOK.
func Email(ctx context.Context, domain string) Result {
	return notImplemented()
}
