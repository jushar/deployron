# Docker Autodeploy
This small application provides a HTTP rest API for Dockerhub/-registry web hooks to automatically deploy Docker containers.

## Installation
Since we don't want to launch the HTTP server as root, we use the _setuid_ flag to run a deploy script as root
(which is required for `docker` commands to connect to the Docker daemon).
However, due to security reasons set setuid flag is disabled for shell scripts in the Linux kernel.
Therefore, we have to compile a small helper program that securely executes some scripts. This requires us to manually setup permissions correctly.

1.) Compile setuid helper:
```shell
gcc -o setuidhelper setuidhelper.c
```

2.) Update permissions
```shell
# Deploy script permissions (replace deploy.sh perhaps)
chown root:root ./deploy.sh
chmod 700 ./deploy.sh

# Folder permissions to prevent replacing
chown root:root .
chmod 700 .
```

3.) Launch the API
```shell
go run service.go
```