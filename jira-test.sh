#!/bin/sh

USER=`cat config/user`
PASSWD=`cat config/passwd`
URL=`cat config/baseurl`


curl --user $USER:$PASSWD $URL/$1 | prettyjson