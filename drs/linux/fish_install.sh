echo "Downloading drs..."
wget "https://github.com/arghyagod-coder/GoLang/releases/download/1.0.0/drs-linux-amd64" 

echo "Adding drs into PATH..."

mkdir -p ~/.drs

mv ./drs-linux-amd64 ~/.drs/drs
chmod +x ~/.drs/drs
set -U fish_user_paths ~/.drs/ $fish_user_paths 
echo "drs installation is completed!"
echo "You need to restart the shell to use drs." 