echo "Installing alias 'run'"
echo 'alias run=/home/$USER/scripts/run.sh' >> .bashrc

echo "Setting permissions"
chmod +x ./git
chmod +x ./docker
chmod +x ./setup
chmod +x run.sh