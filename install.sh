
## INSTALL TCP_SERVER
go build -o ./tmp/cache_server ./tcp_server/main.go
DIRECTORY="$HOME/.cache_system"
mkdir -p $DIRECTORY
mv ./tmp/cache_server $DIRECTORY


## INSTALL interactive shell
go build -o ./tmp/cachex ./shell/main.go
mv ./tmp/cachex $DIRECTORY


echo "export PATH=$PATH:$DIRECTORY" >> ~/.bashrc 
echo "export PATH=$PATH:$DIRECTORY" >> ~/.profile 