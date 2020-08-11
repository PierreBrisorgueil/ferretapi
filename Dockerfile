# build step
FROM golang:latest

# System setup
ENV DEBIAN_FRONTEND=noninteractive \
    TERM=xterm \
    TIMEZONE=UTC

# Install deps + add Chrome Stable + purge all the things
RUN apt-get update && apt-get install -y \
	apt-transport-https \
	ca-certificates \
	curl \
	gnupg \
	--no-install-recommends \
	&& curl -sSL https://dl.google.com/linux/linux_signing_key.pub | apt-key add - \
	&& echo "deb https://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list \
	&& apt-get update && apt-get install -y \
	google-chrome-stable \
	fontconfig \
	fonts-ipafont-gothic \
	fonts-wqy-zenhei \
	fonts-thai-tlwg \
	fonts-kacst \
	fonts-symbola \
	fonts-noto \
	--no-install-recommends \
	&& apt-get purge --auto-remove -y curl gnupg \
	&& rm -rf /var/lib/apt/lists/*

# Create app directory
WORKDIR /app

# args
ARG FERRET_TAG='@latest'

COPY . .
RUN go get github.com/gobs/args
RUN go get github.com/raff/godet
RUN export GO111MODULE=on
RUN go mod init ferret
RUN go mod download github.com/MontFerret/ferret${FERRET_TAG}
RUN go build -o main .

EXPOSE 8080

# Command to run the executable
CMD ["./main"]