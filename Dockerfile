# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #

ARG BUILDER_IMAGE=golang:1.22-alpine

# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #

FROM ${BUILDER_IMAGE} as development

# Install useful additional tools for development, e.g., curl, make and git:
RUN apk update && apk add --no-cache curl git make

# Ensure staticcheck is executable and in the PATH
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

# Ensure go-critic is executable and in the PATH
RUN go install github.com/go-critic/go-critic/cmd/gocritic@latest

# Set the GOPATH environment variable
ENV $PATH:$(go env GOPATH)/bin

# Install szh shell:
RUN apk add --no-cache zsh

# Install oh-my-zsh:
# Uses "Spaceship" theme with some customization. Uses some bundled plugins and installs some more from github
RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.5/zsh-in-docker.sh)" -- \
    -t https://github.com/denysdovhan/spaceship-prompt \
    -a 'SPACESHIP_PROMPT_ADD_NEWLINE="false"' \
    -a 'SPACESHIP_PROMPT_SEPARATE_LINE="false"' \
    -p git \
    -p ssh-agent \
    -p https://github.com/zsh-users/zsh-autosuggestions \
    -p https://github.com/zsh-users/zsh-completions

# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #