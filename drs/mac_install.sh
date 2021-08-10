#!/bin/bash

echo "Downloading drs..."
curl -L "https://github.com/arghyagod-coder/GoLang/releases/latest/download/drs-darwin-amd64" -o drs

echo "Adding drs into PATH..."

mkdir -p ~/.drs;
mv ./drs ~/.drs
echo "PATH=$PATH:~/.drs" >> /etc/environment  
echo "drs installation is completed!"
echo "You need to restart the shell to use drs."