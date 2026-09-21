# Production deployment

The production stack runs the Go API on `127.0.0.1:8001` and keeps MariaDB
on an internal-only Docker network. GitHub Actions builds and tests each
`master` revision, uploads the immutable image through SSH, and activates it
with `/usr/local/sbin/deploy-rogueserver`.

Required repository Actions secrets:

- `DEPLOY_HOST`
- `DEPLOY_PORT`
- `DEPLOY_USER`
- `DEPLOY_SSH_KEY`
- `DEPLOY_HOST_KEY`

Runtime credentials live only in `/srv/rogueserver/shared/*.env` on the
server. They must not be committed or uploaded to GitHub.

The systemd timer `rogueserver-backup.timer` creates daily compressed SQL
backups in `/srv/rogueserver/backups` and retains 14 days.

