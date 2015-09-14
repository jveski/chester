#!/bin/bash

###########################
# Install Puppet and Nginx
###########################
sudo apt-get install nginx -y
sudo apt-get install puppet -y

###########################
# Get a random module from the forge
###########################
mkdir -p acceptance/v3/files
curl "https://forgeapi.puppetlabs.com/v3/files/puppetlabs-apache-1.5.0.tar.gz" -o "acceptance/v3/files/puppetlabs-apache-1.5.0.tar.gz"
curl "https://forgeapi.puppetlabs.com/v3/files/puppetlabs-apache-1.6.0.tar.gz" -o "acceptance/v3/files/puppetlabs-apache-1.6.0.tar.gz"


###########################
# Start chester
###########################
go run main.go -modulepath acceptance/v3/files &


###########################
# Start nginx
###########################
/usr/sbin/nginx -p acceptance -c nginx.conf


###########################
# Attempt to install module from the API
###########################
puppet module install puppetlabs/apache --module_repository http://localhost:8080