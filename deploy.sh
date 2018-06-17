#! /bin/sh

kill -9 $(pgrep webserver)
cd /data/gowork/goWebForDevOps/
git pull https://github.com/CaryLy/goWebForDevOps.git
cd webserver/
./webserver &