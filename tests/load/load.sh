#!/bin/bash

count=0;
while true; do

	shouldFail=false
	echo "count: $((count % 10))"
	if [ $((count % 10)) -gt 8 ]; then
		echo "Should fail"
		shouldFail=true
	fi

	if [[ "$shouldFail" -eq "false" ]]; then
		echo "POST http://localhost:7000/write?shouldFail=$shouldFail" | vegeta attack -rate=10 -duration=5s | vegeta report
		echo "GET http://localhost:7000/read?shouldFail=$shouldFail" | vegeta attack -rate=10 -duration=5s | vegeta report
	else
# 		echo "POST http://localhost:7000/write?shouldFail=$shouldFail" | vegeta attack -rate=1000 -duration=1s | vegeta report
# 		echo "GET http://localhost:7000/read?shouldFail=$shouldFail" | vegeta attack -rate=1000 -duration=1s | vegeta report
		for i in {1..10}; do
			curl -sS -X GET http://localhost:7000/read?shouldFail=true
			curl -sS -X POST http://localhost:7000/write?shouldFail=true
		done
	fi

	count=$((count+1))
done
