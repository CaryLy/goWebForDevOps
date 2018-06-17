#! /bin/sh

kill -9 $(pgrep webserver)
cd /data/web/gowork/src/DevOpsAndCloud
git pull https://github.com/CaryLy/goWebForDevOps.git
cd webserver/
./webserver &