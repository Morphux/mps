#!/usr/bin/env bash

DB="mock3.db"
IP="127.0.0.1"
PORT="6666"


./mps -db=./test_database/${DB} ${IP}:${PORT} &
MPS_PID=$!
sleep 2;
echo ${MPS_PID}
./tests/protocol_tests -a ${IP} -p ${PORT} -d test_database/${DB} client
RET=$?
kill ${MPS_PID}
exit ${RET};
