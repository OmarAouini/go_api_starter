CURRENT=`pwd`
BASENAME=`basename "$CURRENT"`
COMMIT_HASH=`git log -1 --format=%H`
docker run -it "$BASENAME:$COMMIT_HASH" -p 8080:8080