# GoDis Key-Value Store

A key-value store implementation in Go, supporting multiple data structures and advanced features.

## Quick Start

Start a local server:
```bash
./spawn_redis_server.sh
```

This starts a Redis server on port **6379** by default.

Start a replica server:
```bash
./spawn_redis_server.sh --port 3000 --replicaof localhost 6379
```

Test the server:
```bash
$ redis-cli ping
PONG
```

---

## Features

### Basic Commands

#### PING
Check server connectivity.
```bash
PING
# Returns: PONG
```

#### ECHO
Echo the given string.
```bash
ECHO "Hello World"
# Returns: "Hello World"
```

---

### String Operations

#### SET
Set key to hold the string value with optional expiration.
```bash
SET mykey "value"
SET session "data" PX 5000  # Expire in 5000ms
```

#### GET
Get the value of a key.
```bash
GET mykey
# Returns: "value"
```

#### INCR
Increment the integer value of a key by one.
```bash
INCR counter
# Returns: 1
```

---

### List Operations

Lists are ordered collections of strings. Elements can be added/removed from both ends.

#### LPUSH
Insert elements at the head of a list.
```bash
LPUSH mylist "element1" "element2"
# Returns: 2
```

#### RPUSH
Append elements to the tail of a list.
```bash
RPUSH mylist "element3" "element4"
# Returns: 4
```

#### LPOP
Remove and return the first element of a list.
```bash
LPOP mylist
# Returns: "element2"
```

#### RPOP
Remove and return the last element of a list.
```bash
RPOP mylist
# Returns: "element4"
```

#### LLEN
Get the length of a list.
```bash
LLEN mylist
# Returns: 2
```

#### LRANGE
Get a range of elements from a list.
```bash
LRANGE mylist 0 -1  # Get all elements
# Returns: ["element1", "element3"]
```

#### BLPOP
Blocking version of LPOP. Waits for an element to become available.
```bash
BLPOP mylist 5  # Wait up to 5 seconds
# Returns: ["mylist", "element1"]
```

---

### Sorted Set Operations

Sorted sets store unique members with associated scores, automatically sorted by score.

#### ZADD
Add members with scores to a sorted set.
```bash
ZADD leaderboard 100 "player1" 200 "player2"
# Returns: 2
```

#### ZCARD
Get the number of members in a sorted set.
```bash
ZCARD leaderboard
# Returns: 2
```

#### ZRANK
Get the rank (index) of a member in a sorted set.
```bash
ZRANK leaderboard "player1"
# Returns: 0
```

#### ZSCORE
Get the score of a member in a sorted set.
```bash
ZSCORE leaderboard "player1"
# Returns: "100"
```

#### ZRANGE
Get members in a score range.
```bash
ZRANGE leaderboard 0 -1 WITHSCORES
# Returns: ["player1", "100", "player2", "200"]
```

#### ZREM
Remove members from a sorted set.
```bash
ZREM leaderboard "player1"
# Returns: 1
```

---

### Stream Operations

Streams are append-only log data structures for event streaming.

#### XADD
Append a new entry to a stream.
```bash
XADD mystream * field1 value1 field2 value2
# Returns: "1640000000000-0"
```

#### XRANGE
Query a range of entries in a stream.
```bash
XRANGE mystream - +
# Returns all entries
```

#### XREAD
Read entries from one or more streams.
```bash
XREAD STREAMS mystream 0
# Returns new entries since ID 0
```

---

### Pub/Sub

Publish/Subscribe messaging pattern for real-time communication.

#### PUBLISH
Post a message to a channel.
```bash
PUBLISH news "Breaking news!"
# Returns: 1 (number of subscribers)
```

#### SUBSCRIBE
Subscribe to channels.
```bash
SUBSCRIBE news sports
# Receives messages from news and sports channels
```

#### UNSUBSCRIBE
Unsubscribe from channels.
```bash
UNSUBSCRIBE news
# Stops receiving messages from news channel
```

---

### Transactions

Execute multiple commands atomically.

#### MULTI
Mark the start of a transaction block.
```bash
MULTI
```

#### EXEC
Execute all queued commands in the transaction.
```bash
EXEC
# Executes all commands queued since MULTI
```

#### DISCARD
Discard all commands in the transaction queue.
```bash
DISCARD
# Cancels the transaction
```

---

### Replication

Master-replica replication for high availability and read scaling.

#### INFO
Get server information including replication status.
```bash
INFO replication
# Returns: role:master, connected_slaves:1, ...
```

#### REPLCONF
Internal command used by replicas to configure replication.
```bash
REPLCONF listening-port 6380
```

#### PSYNC
Internal command used by replicas to synchronize with master.
```bash
PSYNC ? -1
# Initiates full synchronization
```

---

## Architecture

- **RESP Protocol**: Redis Serialization Protocol for client-server communication
- **Skip List**: Efficient sorted set implementation with O(log n) operations
- **Pub/Sub**: In-memory message broker with channel subscriptions
- **Replication**: Asynchronous master-replica data synchronization
- **Transactions**: ACID-compliant transaction support with command queuing

---

## Testing

Run all tests:
```bash
go test ./app/tests/...
```

Run specific test:
```bash
go test ./app/tests/zadd_test.go -v
```

---

## Implementation Details

- Written in Go for performance and concurrency
- Custom RESP protocol parser
- Skip list data structure for sorted sets
- Non-blocking pub/sub with goroutines
- Polling-based blocking operations (BLPOP)
- Thread-safe operations with mutex locks
