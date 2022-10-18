#!/bin/bash

echo starting experiment

./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=101 > results/1.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=1001 > results/2.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=10001 > results/3.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=100001 > results/4.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=1000001 > results/5.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=10000001 > results/6.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=100000001 > results/7.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=1000000001 > results/8.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=10000000001 > results/9.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=100000000001 > results/10.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=1000000000001 > results/11.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=10000000000001 > results/12.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=100000000000001 > results/13.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=1000000000000001 > results/14.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=10000000000000001 > results/15.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=100000000000000001 > results/16.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=1000000000000000001 > results/17.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=10000000000000000001 > results/18.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=100000000000000000001 > results/19.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=1000000000000000000001 > results/20.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=10000000000000000000001 > results/21.txt
sleep 1
./onetenparts -txt=true -png=true -out=s -lat=t -gens=100 -ic=100000000000000000000001 > results/22.txt

echo experiment complete