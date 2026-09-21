# Production deployment

The production stack runs the Go API on `127.0.0.1:8001`. MariaDB remains on
an internal-only Docker network, while the API also joins a separate edge
network so its loopback-only port can be reached by the host Nginx proxy.
GitHub Actions builds and tests each `master` revision, uploads the immutable
API image through SSH, and activates the release with
`/usr/local/sbin/deploy-rogueserver`.

Normal pushes transfer only the small API image for fast updates. When
bootstrapping a fresh server, manually run the workflow with
`include_database_image` enabled. That bundles the pinned MariaDB image too,
avoiding a runtime dependency on third-party Docker registry mirrors.

Required repository Actions secrets:

- `DEPLOY_HOST`
- `DEPLOY_PORT`
- `DEPLOY_USER`
- `DEPLOY_SSH_KEY`
- `DEPLOY_HOST_KEY`

Runtime credentials live only in `/srv/rogueserver/shared/*.env` on the
server. The daily-run seed key is generated once as
`/srv/rogueserver/shared/secret.key` and mounted read-only. These files must
not be committed or uploaded to GitHub.

The systemd timer `rogueserver-backup.timer` creates daily compressed SQL
backups in `/srv/rogueserver/backups` and retains 14 days.
