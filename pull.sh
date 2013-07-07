#!/bin/bash

./fmt.sh

git add .
git commit -a -m "do some"
https_proxy=localhost:8087 git push origin master
