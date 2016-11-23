# Deployron: A Lightweight Deploy Tool
_Deployron_ is a small and lightweight deployment tool that is preferably used with Docker, but is suitable for most general purpose deployments as well.

## Architecture
![Architecture Image](http://i.imgur.com/zCq1YLQ.png)

## Installation
1. Download and extract the latest release from the release section. If you want to use it with _systemd_, make sure you extract the archive to `/var/lib/deployron/`.
2. Update your `config.yml`
3. Update `config.yml` permissions (otherwise the backend service will not start)
  ```bash
  # Only the owner should be able to access the file (use 644 for read privileges by others)
  chmod 600 config.yml

  # Change user to root
  chown root:root config.yml
  ```
4. Install the _systemd_ services (optional)
  ```bash
  # Copy service unit files
  cp systemd/*.service /etc/systemd/system/

  # Create user (hosts the API)
  useradd -d /var/lib/deployron/ -M deployron
  ```

5. Start and enable the systemd services (optional)
  ```bash
  # Start
  systemctl start deployron.service
  systemctl start deployron_api.service

  # Enable (launch when booting)
  systemctl enable deployron.service
  systemctl enable deployron_api.service
  ```
