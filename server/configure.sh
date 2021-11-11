# set GO111MODULE=off for easy compiling
go env -w GO111MODULE=off

# add current path to GOPATH
echo export GOPATH=$(pwd):$GOPATH >> $HOME/.zshrc
source $HOME/.zshrc 
