echo "Downloading drs..."
curl -L "https://github.com/arghyagod-coder/GoLang/releases/latest/download/drs-linux-amd64" -o drs

echo "Adding drs into PATH..."

mkdir -p ~/.drs

chmod u+x ./drs

mv ./drs ~/.drs
set -U fish_user_paths ~/.drs/ $fish_user_paths 
echo "drs installation is completed!"
echo "You need to restart the shell to use drs." 