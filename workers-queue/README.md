# Working queue management example

## Disclaimer
This is my implementation of a working queue with a fix number of workers.
My objective was to be able to multithread a set of tasks (the queue)
having a full control of the CPU and memory usage by giving a maximum number
of working threads.

## How to build
go build

## Sample output

```
% ./working-queue
[dispatcher] Send task Un
[dispatcher] Send task Deux
[dispatcher] Send task Trois
[dispatcher] Send task Quatre
[dispatcher] Send task Cinq
[dispatcher] Send task Six
[dispatcher] Send task Sept
[worker 0] Running task Un and sleep 719 seconds
[worker 1] Running task Deux and sleep 106 seconds
[worker 2] Running task Trois and sleep 912 seconds
[worker 1] Running task Quatre and sleep 267 seconds
[dispatcher] Send task Huit
[worker 1] Running task Cinq and sleep 961 seconds
[dispatcher] Send task Neux
[worker 0] Running task Six and sleep 467 seconds
[dispatcher] Send task Dix
[worker 2] Running task Sept and sleep 156 seconds
[worker 2] Running task Huit and sleep 662 seconds
[worker 0] Running task Neux and sleep 757 seconds
[worker 1] Running task Dix and sleep 385 seconds
[worker 2] exit.
[worker 1] exit.
[worker 0] exit.
Un true
Deux true
Trois true
Quatre true
Cinq true
Six true
Sept true
Huit true
Neux true
Dix true
```
