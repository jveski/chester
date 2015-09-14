#!/bin/bash


###########################
# Install Puppet and Nginx
###########################
wget https://apt.puppetlabs.com/puppetlabs-release-precise.deb
sudo dpkg -i puppetlabs-release-precise.deb
sudo apt-get update
sudo apt-get install puppet -y
sudo apt-get install nginx -y


###########################
# Get a random module from the forge
###########################
mkdir -p acceptance/v3/files
curl "https://forgeapi.puppetlabs.com/v3/files/puppetlabs-apache-1.5.0.tar.gz" -o "acceptance/v3/files/puppetlabs-apache-1.5.0.tar.gz"
curl "https://forgeapi.puppetlabs.com/v3/files/puppetlabs-stdlib-4.9.0.tar.gz" -o "acceptance/v3/files/puppetlabs-stdlib-4.9.0.tar.gz"
curl "https://forgeapi.puppetlabs.com/v3/files/puppetlabs-concat-1.2.4.tar.gz" -o "acceptance/v3/files/puppetlabs-concat-1.2.4.tar.gz"
curl "https://forgeapi.puppetlabs.com/v3/files/puppetlabs-apache-1.6.0.tar.gz" -o "acceptance/v3/files/puppetlabs-apache-1.6.0.tar.gz"


###########################
# Start chester
###########################
go run main.go -modulepath acceptance/v3/files &


###########################
# Start nginx
###########################
sudo chmod 777 /var/lib/nginx
/usr/sbin/nginx -v
/usr/sbin/nginx -p acceptance -c nginx.conf &


###########################
# Wait a bit for good measure
###########################
sleep 5


###########################
# Curl the endpoints
###########################
jobs
curl -v -s "http://localhost:8081/v3/release?module=puppetlabs-apache" 1> /dev/null
curl -v -s "http://localhost:8081/v3/files/puppetlabs-apache-1.6.0.tar.gz" 1> /dev/null


###########################
# Attempt to install module from the API
###########################
puppet --version
puppet module install puppetlabs/apache --module_repository http://localhost:8081
