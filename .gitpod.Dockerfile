FROM gitpod/workspace-full

# Install neovim
RUN brew install neovim

# Install Golang v1.18
RUN go install golang.org/dl/go1.18beta2@latest
RUN go1.18beta2 download

# Install delve for debug
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Ensure it's now go1.18 using in the PATH
RUN export PATH="$(go1.18beta2 env GOROOT)/bin:${PATH}"

