#!/bin/bash
#
# Installs DB schema into your MySQL database and
# creates a base user

# Database creation
echo "--------------"
echo "DATABASE SETUP"
echo "--------------"
echo -n "Host (localhost): "
read host
echo -n "Username (root): "
read dbUser
echo -n "Password (root): "
read dbPass

# Default host
if [ -z "$host" ]
then
  host="localhost"
fi

# Default DB user
if [ -z "$dbUser" ]
then
  dbUser="root"
fi

# Default DB Pass
if [ -z "$dbPass" ]
then
  dbPass="root"
fi

db="mysql --host="$host" --user="$dbUser" --password="$dbPass

echo -n "Setting up database... "
$db < sql/schema.sql
echo "OK"
echo

# User creation
echo "-------------"
echo "USER CREATION"
echo "-------------"
echo -n "Name: "
read uName
echo -n "E-mail: "
read uMail
echo -n "Password: "
read -s uPass

hash="$(echo -n "$uPass" | md5 )"
echo -n "Creating user..."
echo -n "INSERT INTO users(name, email, password) VALUES(\""$uName"\",\""$uMail"\",\""$hash"\")" | $db gopherblog
echo "OK"
echo

