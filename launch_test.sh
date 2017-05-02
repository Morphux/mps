#!/usr/bin/env bash

DB="mock3.db"
IP="127.0.0.1"
PORT="6666"
P_CERT="tls/server.crt"
PRIV_CERT="tls/server.key"


./mps -db=./test_database/${DB} -pub ${P_CERT} -priv ${PRIV_CERT} ${IP}:${PORT} &
MPS_PID=$!
sleep 2;
echo ${MPS_PID}
./tests/protocol_tests -a ${IP} -p ${PORT} -d test_database/${DB} -c ${P_CERT} client
RET=$?
kill ${MPS_PID}
exit ${RET};
