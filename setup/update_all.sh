#!/bin/zsh

echo "Updating Brew"
brew update

brew upgrade

brew cleanup

echo "Updating OhMyZsh"
omz update