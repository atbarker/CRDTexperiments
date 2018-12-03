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

In the previous blog post we discuss [Conflict-Free Replicated Data Types](https://hal.inria.fr/inria-00609399v1/document) (CRDTs), a class of specialized data structures designed to be replicated across a distributed system. We implemented a few CRDTs as specified by [Shapiro et al.'s _A Comprehensive Study of Convergent and Commutative Replicated Data Types_](https://hal.inria.fr/inria-00555588/document). Some of these implementations run into a perennial problem in distributed systems, distributed garbage collection. In the previous post we discussed multiple issues facing distributed garbage collection such as, high metadata storage costs, fault intolerance, and the need to stronger synchronization that betrays the asynchronous nature of a CRDT. To tackle these problems we look at a series of research avenues that include [pure operation-based CRDTs](https://arxiv.org/abs/1710.04469), [causal trees](https://github.com/gritzko/ctre), and [delta-state CRDTs](https://arxiv.org/pdf/1603.01529.pdf). We then take the Two Phase Set (2P-set) and Add Remove Partial Order (ARPO) implementations from the previous post and discuss methods for extending these to support some form of distributed garbage collection.

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

[Delta-state CRDTs](https://arxiv.org/pdf/1603.01529.pdf) help to avoid the issue of sending the entire state of a data type over a network. Unlike some of the previous work in CRDTs they account for garbage collection in their anti-entropy algorithm designed to enforce causal consistency. In this algorithm each node maintains two maps, one for keeping track of a sequence of state deltas and another for a series of acknowledgements from each neighbor.

Still has considerable metadata overhead (though not with vector clocks)

Relies on something similar to proving causal relationships

When it gets acknowledgements and the delta has been applied then it can garbage collect the delta. No mention of markers or tombstones as of yet in the paper.

### Reducing Vector Clock space overhead

[methods for reducing the space costs of vector clocks](http://www.bailis.org/blog/causality-is-expensive-and-what-to-do-about-it/)

### Synchronized Garbage Collection

[Paxos Commit](https://lamport.azurewebsites.net/video/consensus-on-transaction-commit.pdf) and Two-Phase Commit protocols

## Implementation

## Conclusion


## Implementing Garbage Collection

