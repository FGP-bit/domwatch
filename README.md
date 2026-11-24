# domwatch

Monitor de dominios escrito en Go. Revisa **DNS**, **estado HTTP**, **vencimiento del certificado SSL/TLS** y la autenticación de correo (**SPF, DKIM, DMARC**).

Lo uso como proyecto eje para aprender infraestructura: Docker, CI/CD, Prometheus, Grafana, Alertmanager, AWS y Terraform.

## Uso

```bash
go run ./cmd/domwatch check example.com
```

Salida esperada cuando los chequeos estén implementados:

```text
Dominio: example.com

CHEQUEO  ESTADO  TIEMPO  DETALLE
dns      OK      12ms    93.184.215.14
http     OK      210ms   200 OK
tls      OK      95ms    vence en 84 días
email    WARN    20ms    falta DMARC
```

Código de salida: `0` todo bien · `1` algún chequeo falló · `2` uso incorrecto.

## Estructura

```text
domwatch/
├── cmd/domwatch/main.go   # entrada del CLI: subcomandos check, version
├── internal/
│   ├── check/             # un archivo por chequeo, todos con la misma firma
│   │   ├── check.go       # tipos Result/Status, lista de chequeos y Run()
│   │   ├── dns.go
│   │   ├── http.go
│   │   ├── tls.go
│   │   └── email.go       # SPF / DKIM / DMARC
│   └── report/text.go     # imprime los resultados como tabla
├── Makefile               # make run / build / test
└── go.mod
```

Lo que se irá sumando, según el roadmap:

| Mes | Se agrega |
|---|---|
| Oct | `Dockerfile` (multi-stage + distroless, imagen < 20 MB) |
| Nov | `.github/workflows/`, `.goreleaser.yaml`, `internal/metrics/`, subcomando `serve` con `/metrics` |
| Dic | `deploy/` con `prometheus/`, `grafana/`, `alertmanager/` y `docker-compose.yml` |
| Ene | `infra/terraform/` (state remoto en S3, módulos) |

## Roadmap

- [ ] `domwatch check` con los 4 chequeos implementados
- [ ] `docker run domwatch check dominio.com` con imagen < 20 MB
- [ ] CI con tests y lint; cada tag publica binarios e imagen en GHCR
- [ ] Daemon con `/metrics` para Prometheus
- [ ] Dashboard en Grafana con HTTPS y alerta cuando a un SSL le quedan < 15 días
- [ ] Infraestructura en AWS con Terraform; un push a `main` despliega solo

## Desarrollo

```bash
make run DOMAIN=example.com   # ejecutar
make build                    # binario en ./bin/domwatch
make test                     # tests
```
