#!/bin/bash

source ./.env/bin/activate

for i in `seq 0 1`; do
    echo "Starting client $i"
    python client.py --clientID $i &
done

# This will allow you to use CTRL+C to stop all background processes
trap "trap - SIGTERM && kill -- -$$" SIGINT SIGTERM
# Wait for all background processes to complete
wait

python server.py 