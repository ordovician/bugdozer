#!/bin/sh
USER=`cat config/user`
PASSWD=`cat config/passwd`
URL=`cat config/baseurl`

curl --user $USER:$PASSWD "$URL/search?jql=assignee=$USER&fields=id,key,description,summary" | prettyjson