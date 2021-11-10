# add PRJROOT
export PRJROOT=$(pwd)
echo PROJROOT:
echo $PRJROOT
# add working directory to GOPATH for packge importing
echo GOPATH:
export GOPATH=$PRJROOT:$GOPATH
echo $GOPATH
