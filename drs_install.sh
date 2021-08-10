#!/bin/bash

echo "Downloading drs..."
wget "https://github.com/arghyagod-coder/GoLang/releases/latest/download/drs.go-linux-amd64"

echo "Adding drs into PATH..."

mkdir -p $HOME/.drs
mv ./drs.go-linux-amd64 ~/.drs
echo "PATH=$PATH:~/.drs" >> /etc/environment  
echo "drs installation is completed!"
rm -rf ./drs.go-linux-amd64
echo "Deleting useless files"
echo "You need to restart the shell to use drs."