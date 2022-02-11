FROM gitpod/workspace-full:latest

# Install neovim
RUN sudo apt-get install software-properties-common
RUN sudo add-apt-repository ppa:neovim-ppa/stable
RUN sudo apt-get update
RUN sudo apt-get install -y neovim 

# Install Golang v1.18
RUN sudo mkdir /workspace/go && sudo chown gitpod:gitpod -R /workspace/go
RUN go install golang.org/dl/go1.18beta2@latest
RUN go1.18beta2 download

# Install delve for debug
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Ensure it's now go1.18 using in the PATH
ENV PATH="/home/gitpod/sdk/go1.18beta2/bin:${PATH}"
ENV GOROOT="/home/gitpod/sdk/go1.18beta2"