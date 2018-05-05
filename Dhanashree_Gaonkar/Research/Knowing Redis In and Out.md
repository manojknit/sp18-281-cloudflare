# Really important know this :
https://martin.kleppmann.com/2015/05/11/please-stop-calling-databases-cp-or-ap.html

# Redis: Basic Definition
Source: https://aphyr.com/posts/283-jepsen-redis
Redis is a fantastic data structure server, typically deployed as a shared heap. It provides fast access to strings, lists, sets, maps, and other structures with a simple text protocol. Since it runs on a single server, and that server is single-threaded, it offers linearizable consistency by default: all operations happen in a single, well-defined order. There’s also support for basic transactions, which are atomic and isolated from one another.

Because of this easy-to-understand consistency model, many users treat Redis as a message queue, lock service, session store, or even their primary database. Redis running on a single server is a CP system, so it is consistent for these purposes.


Redis offers asynchronous primary->secondary replication. A single server is chosen as the primary, which can accept writes. It relays its state changes to secondary servers, which follow along. Asynchronous means that you don’t have to wait for a write to be replicated before the primary returns a response to the client. Writes will eventually arrive on the secondaries, if we wait long enough. In our application, all 5 clients will read from the primary on n1, and n2–n5 will be secondaries.

This is still a CP system, so long as we never read from the secondaries. If you do read from the secondaries, it’s possible to read stale data. That’s just fine for something like a cache! However, if you read data from a secondary, then write it to the primary, you could inadvertently destroy writes which completed but weren’t yet replicated to the secondaries.

What happens if the primary fails? We need to promote one of the secondary servers to a new primary. One option is to use Heartbeat or a STONITH system which keeps a link open between two servers, but if the network partitions we don’t have any way to tell whether the other side is alive or not. If we don’t promote the primary, there could be no active servers. If we do promote the primary, there could be two active servers. We need more nodes.

If one connected component of the network contains a majority (more than N/2) of nodes, we call it a quorum. We’re guaranteed that at most one quorum exists at any point in time–so if a majority of nodes can see each other, they know that they’re the only component in that state. That group of nodes (also termed a “component”) has the authority to promote a new primary.

# Redis Cluster for instance is a system biased towards consistency rather than availability. Redis Sentinel itself is an HA solution with the dogma of consistency and master slave setups.“
