#!/bin/bash

# check if Docker is installed
if ! [ -x "$(command -v docker)" ]; then
  # get the operating system version
  OS_VERSION=$(lsb_release -cs)

  # install Docker for the correct operating system
  if [ "$OS_VERSION" = "ubuntu" ]; then
    # install Docker for Ubuntu
    sudo apt-get update
    sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
    sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
    sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
    sudo apt-get update
    sudo apt-get install -y docker-ce
  elif [ "$OS_VERSION" = "debian" ]; then
    # install Docker for Debian
    sudo apt-get update
    sudo apt-get install -y apt-transport-https ca-certificates curl gnupg2 software-properties-common
    sudo curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add -
    sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable"
    sudo apt-get update
    sudo apt-get install -y docker-ce
  else
    # unsupported operating system
    echo "Unsupported operating system: $OS_VERSION"
    exit 1
  fi
fi

# build the Docker image for the app
sudo docker build -t go-web-app .

# run the Docker container for the app
sudo docker run -p 8080:8080 go-web-app

