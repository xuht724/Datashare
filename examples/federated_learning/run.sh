#!/bin/bash

source ./.env/bin/activate

for i in `seq 0 3`; do
    for i in `seq 0 1`; do
        python client.py --clientID $i
    done
done

for i in `seq 0 3`; do
    python server.py 
    for i in `seq 0 1`; do
        python client.py --clientID $i
    done
done