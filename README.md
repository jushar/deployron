# Deployron: A Lightweight Deploy Tool
[![Build Status](https://travis-ci.org/Jusonex/deployron.svg?branch=master)](https://travis-ci.org/Jusonex/deployron)

_Deployron_ is a small and lightweight deployment tool that is preferably used with Docker, but is suitable for most general purpose deployments as well.  
It uses a _yaml_ configuration file that holds a single or multiple deploy scripts which can be executed with extended privileges in a secure way.

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

## Configuration
The following snippet is a commented example configuration.
```yml
api:
  ip: 127.0.0.1 # IP the server should listen on (use 0.0.0.0 to listen on all interfaces)
  port: 1337 # Port we're listening on (optional, defaults to 1337)
  unixsocket: "./service_client.sock" # The unix client socket (optional)

service:
  unixsocket: "./service.sock" # The unix backend process server socket (optional)

deployments:
- name: mydeploy1 # name of the deployment entry (you can use any)
  secret: deploy1secret # per deployment secret
  description: "My test deploy service 1" # friendly description (optional)
  user: root # the user who should execute the script below (optional, defaults to root)
  script: # The actual deploy script
  - echo "Hello World from mydeploy1"
  - whoami

- name: mydeploy2
  secret: deploy2secret
  description: "My test deploy service 2"
  user: vm
  script:
  - echo "Hello World from mydeploy2"
  - whoami

```
