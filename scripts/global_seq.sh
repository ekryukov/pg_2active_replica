#!/usr/bin/env bash

DBNAME='svip_db'
DBUSER='svip'
DBSCHEMA='bpc01'


for i in $( psql -U ${DBUSER}  -d ${DBNAME} -t -c "SELECT sequence_name FROM information_schema.sequences where sequence_schema='${DBSCHEMA}'")
do
    echo ${i};
    psql -U svip  -d svip_db -t -c "alter sequence ${i} using bdr"
done;
