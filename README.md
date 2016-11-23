# Deployron: A Lightweight Deploy Tool
_Deployron_ is a small and lightweight deployment tool that is preferably used with Docker, but is suitable for most general purpose deployments as well.

## Architecture
![Architecture Image](http://i.imgur.com/zCq1YLQ.png)

## Installation
1. Download the latest release from the release section.
2. Update your `config.yml`
3. Update `config.yml` permissions (otherwise the backend service will not start)
```bash
# Only the owner should be able to access the file (use 644 for read privileges by others)
chmod 600 config.yml

# Change user to root
chown root:root config.yml
```
4. Install the _systemd_ services.
TODO

5. Start and enable the systemd services
```bash
# Start
systemctl start deployron.service
systemctl start deployron_api.service

# Enable (launch when booting)
systemctl enable deployron.service
systemctl enable deployron_api.service
```
