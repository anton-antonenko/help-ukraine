#!/bin/bash -x

DESTS=( $( cat sites.txt ) )

TIME="180s"
while :
do
   for DEST in ${DESTS[@]}
   do
       docker run -ti --rm alpine/bombardier -c 1000 -d $TIME -l $DEST
   done
done