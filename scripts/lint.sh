#!/bin/sh

# This script is for maintenance of the projects.
# It executes linting for the different go projects.

CHAPTERS_NB="3 4 5 6 7 8 9"

pwd=$(pwd)

for nb in $CHAPTERS_NB
do
	if [ -d "$pwd/chapter$nb" ]
	then
		echo "enter $pwd/chapter$nb"
		cd $pwd/chapter$nb
        go list -f '{{.Dir}}/...' -m | xargs golangci-lint run
	fi
done
