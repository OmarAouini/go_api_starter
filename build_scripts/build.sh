CURRENT=`pwd`
BASENAME=`basename "$CURRENT"`
COMMIT_HASH=`git log -1 --format=%H`
docker build -t "$BASENAME:$COMMIT_HASH" .