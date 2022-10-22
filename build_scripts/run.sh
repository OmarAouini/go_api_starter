CURRENT=`pwd`
BASENAME=`basename "$CURRENT"`

docker run -it "$BASENAME" -p 8080:8080