---
title: Implementing a Garbage-Collected Graph CRDT (Part 2 of 2)
author: Austen Barker
layout: single
classes: wide
---

<script type="text/javascript"
src="http://cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML,http://composition.al/javascripts/MathJaxLocal.js">
</script>

by Austen Barker &middot; edited by  and Lindsey Kuper

## Background & Recap

[Conflict-Free Replicated Data Types](https://hal.inria.fr/inria-00609399v1/document) (CRDTs) are a class of specialized data structures designed to be replicated across a distributed system while providing eventual consistency and high availability. CRDTs can be modified concurrently without coordination while providing a means to reconcile conflicts between replicas.

While CRDTs are a promising solution to the problem of building an eventually consistent distributed system, numerous practical implementation challenges remain. To deal with issues that arise when processing concurrent operations such as, for example, conflicting additions and removals of elements in a set, many CRDT specifications rely on the use of _tombstones_, which are markers to represent deleted items. These tombstones can accumulate over time and necessitate the use of a garbage collection system in order to avoid unacceptably costly growth of underlying data structures. These garbage collection systems can prove to be difficult to implement in practice. This series of two posts will chronicle my exploration of garbage collection in the context of CRDTs, and my attempts to implement a non-trivial garbage-collected CRDT based on the specifications in [Shapiro et al.'s _A Comprehensive Study of Convergent and Commutative Replicated Data Types_](https://hal.inria.fr/inria-00555588/document).

Shapiro et al. present two styles of specifying CRDTs: _state-based_ and _operation-based_. The difference comes from how the replicas propagate updates to one another. In the state-based model, replicas transmit their entire local state to other replicas, which then reconcile inconsistencies through a commutative, associative, and idempotent _merge_ operation. As seen later in this blog post, a merge operation can often be represented by a union between two sets.

Operation-based, or op-based, CRDTs transmit their state by sending only the update operations performed to other replicas, so each operation is individually replayed on the recipient replica. In this model, the operations must be commutative but not necessarily idempotent. Op-based CRDTs are more particular about the messaging protocol between replicas, but require less bandwidth than state-based CRDTs, which must transmit the entire local state instead of small operations. State-based CRDTs, on the other hand, provide an associative merge operation.

Shapiro et al. give specifications for a variety of CRDT data structures, including sets, counters, registers, and graphs. This blog post is primarily concerned with the implementation of sets and graphs. The two simplest sets specified are the Grow-only Set (G-Set) and the so-called Two-Phase Set (2P-Set). The G-Set is a set of elements that grows monotonically, with no removal operation. 2P-Sets, on the other hand, support removing items from the set.  In the case of a state-based 2P-Set, conflicts between add and remove operations during a merge necessitate some record of which elements have been removed from the set; an additional G-Set, sometimes called the _tombstone set_, maintains markers or tombstones denoting removed elements.

Shapiro et al. then use 2P-Sets and G-Sets to represent sets of vertices and edges in a directed graph. Their Montonic Directed Acyclic Graph (Monotonic DAG) CRDT specification is simply two G-Sets, one for the vertices and one for the edges. In this data structure, there is no operation for removing vertices, and its contents are monotonically increasing.

Finally, Shapiro et al. introduce the Add-Remove Partial Order (ARPO) graph CRDT as a solution to the mess that arises when one attempts to include vertex removals in their Monotonic DAG specification. They define the ARPO using a 2P-Set to represent vertices and a G-Set to represent edges.

## The need for garbage collection

An example of a situation in which garbage collection is useful is when an update to an ARPO is applied and considered stable; at that point, one can discard the set of removed vertices.  For any CRDT that maintains tombstones, such as the state-based 2P-Set and the ARPO, the tombstones might pile up and cause unnecessary bloat. The difficulty with implementing garbage collection is that it will often require synchronization. Shapiro et al. present two challenges related to garbage collection: _stability_ and _commitment_.

Stability refers to whether an update has been received by all replicas.  The purpose of tombstones is to help resolve conflicts between concurrent operations by having a record of removed elements. Eventually, a tombstone is no longer required when all concurrent updates have been “delivered” and an update can be considered stable. The paper applies a modified form of [Wuu and Bernstein’s stability algorithm](https://dl.acm.org/citation.cfm?id=806750), which requires each replica to maintain a set of all the other replicas and for there to be a mechanism to detect when a replica crashes. The algorithm uses vector clocks to determine concurrency of updates.

Commitment issues arise when one needs to perform an operation with a need for synchronization, such as removing tombstones from a 2P-Set or resetting all the replicas of an object to their initial values. Shapiro et al.'s conclusion is to require atomic agreement between all replicas concerning the application of the desired operation.

## Research Avenues

### Pure Operation Based CRDTs

[pure operation-based CRDTs](https://arxiv.org/abs/1710.04469)

### Causal Trees

[causal trees](https://github.com/gritzko/ctre)

### Delta State CRDTs

[delta-state CRDTs](https://arxiv.org/pdf/1603.01529.pdf)

### Reducing Vector Clock space overhead

[methods for reducing the space costs of vector clocks](http://www.bailis.org/blog/causality-is-expensive-and-what-to-do-about-it/)

### Synchronized Garbage Collection

[Paxos Commit](https://lamport.azurewebsites.net/video/consensus-on-transaction-commit.pdf) and Two-Phase Commit protocols

## Conclusion


## Implementing Garbage Collection

Having implemented an ARPO-like CRDT from the original specification, my next step was to investigate garbage collection. Implementing garbage collection for a CRDT like this one is a challenge.  First, establishing the stability of an update as described in the paper assumes that the set of all replicas is known and that they do not crash permanently. Thus the implementation must include a way to detect crashed replicas (in practice, using a timeout) and a way to communicate the failure of a replica reliably to all other replicas.

Another issue is the metadata storage requirements for implementing garbage collection.  Assuming causal delivery of updates requires the use of vector clocks or some similar mechanism to establish causality. Shapiro et al. specifically mention using vector clocks for determining stability in section 4.1 of their paper.   As the definition of stability depends on causality, one can use the same vector clocks to establish both. However, the paper's scheme for determining stability requires each replica to store a copy of the last received vector clock from every other known replica. Therefore, the space complexity required to store the vector clocks locally for $N$ replicas is $O(N^2)$, and total space consumption across the whole set of replicas, $O(N^3)$ --- considerably worse than the usual $O(N)$ necessary to store a single vector clock at each replica for tracking causal relationships, and enough to make programmers uneasy. 

When adding the class of commitment problems to the already mounting pile of dilemmas, the programmer loses hope for the availability and performance of their system. The solutions discussed by Shapiro et al. include the [Paxos Commit](https://lamport.azurewebsites.net/video/consensus-on-transaction-commit.pdf) and Two-Phase Commit protocols, which add considerably to the complexity of the implementation along with sacrificing availability. Shapiro et al. suggest performing operations requiring strong synchronization during periods when network partitions are rare; it may also help to limit such operations to when the availability of a system is not paramount. For example, one could run a garbage collection job during a scheduled server maintenance window.

To sum up, distributed garbage collection requires confronting some of the hardest problems in distributed systems.  Perhaps the easiest solution to the unbounded growth of a CRDT via tombstones is to use the [ostrich algorithm](https://en.wikipedia.org/wiki/Ostrich_algorithm), or to avoid CRDTs that use tombstones entirely.

