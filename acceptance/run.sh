#!/bin/bash

###########################
# Get a random module from the forge
###########################
mkdir -p acceptance/v3/files
curl "https://forgeapi.puppetlabs.com/v3/files/puppetlabs-apache-1.6.0.tar.gz" -o "acceptance/v3/files/puppetlabs-apache-1.6.0.tar.gz"


###########################
# Start chester
###########################
go run main.go -modulepath acceptance/v3/files &


###########################
# Configure nginx
###########################
cat > "/etc/nginx/conf.d/chester.conf" <<CONFIG
server {
        listen 8081;
        server_name chester;

        root `pwd`/acceptance;

        location /v3/release {
                proxy_pass http://localhost:8080;
        }
}
CONFIG


###########################
# Start nginx
###########################
service nginx restart


###########################
# Attempt to install module from the API
###########################
puppet module install puppetlabs/apache --module_repository http://localhost:8080
