DOMAIN=$1
VERSION=$2
docker tag happycar/$DOMAIN ccr.ccs.tencentyun.com/happycar/$DOMAIN:$VERSION
docker push ccr.ccs.tencentyun.com/happycar/$DOMAIN:$VERSION
