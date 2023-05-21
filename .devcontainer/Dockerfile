FROM golang:1.20.4-bullseye

RUN apt update

# Add a non-root user.
ARG DEVCONTAINER_UID
RUN if test -z "$DEVCONTAINER_UID"; then >&2 echo "DEVCONTAINER_UID not specified, exiting"; exit 1; fi
RUN useradd --uid $DEVCONTAINER_UID --user-group --create-home vscode --shell /bin/bash

# Grant sudo privileges to the user.
RUN apt install -y sudo
RUN usermod --append --groups sudo vscode
RUN echo "vscode ALL=(ALL) NOPASSWD:ALL" | tee --append /etc/sudoers.d/vscode

USER vscode
WORKDIR /workspace
