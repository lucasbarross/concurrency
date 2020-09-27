# Exercise 03

## Code Explation

The server basically acts like the linux `cat` command, it sends the received message back and to ensure there were no IO interference the code had no logs.
The RPC implementation used was the gRPC (Google RPC).
The same bash script from exercise 02 was used.

## Results

All data was parsed to ms before any further operation and the data was manipulated using python

### gRPC

| Clients | Mean | Standard Deviation |
| ------- |:----:| ------------------:|
| 1       | 4.39 | 1.60               |
| 2       | 4.64 | 1.59               |
| 3       | 4.49 | 1.51               |
| 4       | 4.16 | 0.71               |
| 5       | 7.64 | 3.24               |
