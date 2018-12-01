#! /bin/bash

if [ $# == 0 ]
then
	echo "option error"
elif [ $1 == "status" ]
then
	git status
elif [ $1 == "push" ]
then
	git add .
	git commit -m "$2"
	git push origin master
fi
