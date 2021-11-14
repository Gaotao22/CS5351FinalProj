# set GO111MODULE=off for easy compiling
go env -w GO111MODULE=off

# add current path to GOPATH
# for zsh
if [ ! -d "$HOME/.zshrc"]; then 
echo export GOPATH=$(pwd):$GOPATH >> $HOME/.zshrc
source $HOME/.zshrc
fi

# for bash
if [ ! -d "$HOME/.bashrc"]; then 
echo export GOPATH=$(pwd):$GOPATH >> $HOME/.bashrc
source $HOME/.bashrc 
fi