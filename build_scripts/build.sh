CURRENT=`pwd`
BASENAME=`basename "$CURRENT"`

docker build -t "$BASENAME" .