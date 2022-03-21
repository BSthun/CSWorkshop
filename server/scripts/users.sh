#!/bin/bash

BASEDIR=$(dirname "$0")

# Preserve the original field separator
ORIGINAL_IFS=$IFS

# Change the field separator to the new line character
IFS=$'\n'

# Read the lines with crednetials from the CSV file
read -d '' -ra users < $BASEDIR/users.csv

# Change the file separtor to a comma
IFS=','

# Iterate over the users
for user in "${users[@]}"
do
   # Split the users into two separate variables
   read iuser ipassword <<< "$user"

   # Create the user
   useradd -m $iuser

   # Set the password
   echo "$ipasswd"$'\n'"$ipasswd" | passwd $iuser &>/dev/null

   # Create user dir
   mkdir -p /var/www/comps/$iuser &>/dev/null

   # Set user dir permission
   chown -R $iuser:$iuser /var/www/comps/$iuser
   chmod -R 744 /var/www/comps/$iuser

   echo "Initialized user '$iuser'"
done

# Return the original field separator
IFS=$ORIGINAL_IFS