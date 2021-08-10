#!/bin/bash

echo "Downloading drs..."
curl -L "https://github.com/arghyagod-coder/GoLang/releases/latest/download/drs-linux-amd64" -o drs

echo "Adding drs into PATH..."

mkdir -p ~/.drs

chmod u+x ./drs

mv ./drs ~/.drs
echo "export PATH=$PATH:~/.drs" >> ~/.zshrc
echo "drs installation is completed!"
echo "You need to restart the shell to use drs."