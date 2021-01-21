#!/bin/bash
echo " ---------- Downloading pitop ---------- "
wget https://github.com/PierreKieffer/pitop/raw/master/bin/32/pitop

chmod +x pitop
sudo mv pitop /usr/local/bin
echo " ---------- pitop is installed ---------- "
echo " ---------- usage ---------- "
echo "    start cmd : pitop "
echo "    press q to exit "
echo " --------------------------- "



