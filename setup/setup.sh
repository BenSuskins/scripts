#!/usr/bin/env bash
set +e

echo "Command Line Tools"
xcode-select --install

echo "Git"
brew install git
git config --global alias.st status
git config --global alias.di diff
git config --global alias.co checkout
git config --global alias.ci commit
git config --global alias.br branch
git config --global alias.sta stash
git config --global alias.llog "log --date=local"
git config --global alias.flog "log --pretty=fuller --decorate"
git config --global alias.lg "log --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit --date=relative"
git config --global alias.lol "log --graph --decorate --oneline"
git config --global alias.lola "log --graph --decorate --oneline --all"
git config --global alias.blog "log origin/master... --left-right"
git config --global alias.ds diff --staged
git config --global alias.fixup commit --fixup
git config --global alias.squash commit --squash
git config --global alias.unstage reset HEAD
git config --global alias.rum "rebase master@{u}"

echo "Brew"
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

echo "Python"
brew install python3

echo "Java"
brew install java11

echo "Node"
brew install nvm
nvm install 16
nvm use 16
nvm alias default 16

echo  "Iterm"
brew install --cask iterm2

echo "OhMyZsh"
sh -c "$(wget -O- https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
exec zsh
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k
exec zsh
p10k configure

echo "Docker"
brew install docker
brew install docker-compose

echo "VS Code"
brew install --cask visual-studio-code

echo "Sourcetree"
brew install --cask sourcetree

echo "Checkstyle"
brew install checkstyle

echo "Chrome"
brew install --cask google-chrome

echo "Slack"
brew install --cask slack

echo "Postman"
brew install --cask postman

echo "Jetbrains Toolbox"
brew install --cask jetbrains-toolbox

echo "Github"
brew install gh
brew install --cask github

echo "GCloud"
brew install --cask google-cloud-sdk

echo "Drawio"
brew install --cask drawio

echo "Zoom"
brew install --cask zoom

echo "Powershell"
brew install --cask powershell

echo "Rectangle"
brew install --cask rectangle

echo "Flycut"
brew install --cask flycut

echo "Openshift"
brew install kubectl
brew install helm
brew install tektoncd-cli
brew install openshift-cli

echo "Zscaler"
brew install wget
wget -P /Downloads https://d32a6ru7mhaq0c.cloudfront.net/Zscaler-osx-3.6.0.46-installer.app.zip

echo "Setting up workspace"
mkdir workspace
cd workspace
mkdir lmd
mkdir other
cd lmd
repos=("lmd-client-api2" "lmd-tsp2" "lmd-auth2" "lmd-scheduler2" "lmd-sc" "lmd-am" "lmd-media" "lmd-web-app2" "lmd-ios-app" "lmd-jenkins-files" "lmd-schedule-manager")

for i in "${repos[@]}"; do
    echo "Starting....." $i
  if [ ! -d "$i" ]; then
    git clone git@github.ford.com:LMD/$i.git
    cd $i
    git checkout develop
    # git checkout master
    cd ..
    echo "Finished : " $i
  else
    echo "Already exists : " $i
  fi
done

git clone git@github.ford.com:BSUSKINS/scripts.git

git clone git@github.ford.com:Pro-Tech/software-delivery-toolset.git

brew cleanup

echo "Done"

set -e
