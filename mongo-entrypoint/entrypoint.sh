#!/usr/bin/env bash
#echo "Creating mongo users..."
#mongosh --authenticationDatabase admin --host localhost -u admin -p admin wiktionary --eval "db.createUser({user: 'user', pwd: 'pass', roles: [{role: 'readWrite', db: 'wiktionary'}]});"
#echo "Mongo users created."
#
#mongosh --host localhost:27017 -u admin -p admin --authenticationDatabase admin
#use admin
#db.createUser({user: "user",pwd: "pass",roles: [{ role: "readWrite", db: "wiktionary"}]})
#use wiktionary

mongosh --host localhost:27017 -u admin -p admin --authenticationDatabase admin < /docker-entrypoint-initdb.d/mongo-script.txt;

echo "Created user for wiktionary db";
