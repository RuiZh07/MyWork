#!/bin/bash

echo "Build WaCave"
sudo docker compose build

echo "Run WaCave network"
sudo docker compose up

sudo docker compose down