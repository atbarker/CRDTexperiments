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

In the previous blog post we discuss [Conflict-Free Replicated Data Types](https://hal.inria.fr/inria-00609399v1/document) (CRDTs), a class of specialized data structures designed to be replicated across a distributed system. We implemented a few CRDTs as specified by [Shapiro et al.'s _A Comprehensive Study of Convergent and Commutative Replicated Data Types_](https://hal.inria.fr/inria-00555588/document). Some of these implementations run into a perennial problem in distributed systems, distributed garbage collection. In the previous post we discussed multiple issues facing distributed garbage collection such as, high metadata storage costs, fault intolerance, and the need to stronger synchronization that betrays the asynchronous nature of a CRDT. To tackle these problems we look at a series of research avenues that include Paxos or Two-Phase Commit, [delta-state CRDTs](https://arxiv.org/pdf/1603.01529.pdf), and different methods of reducing metadata overhead for Shapiro et al.'s garbage collection scheme. We then take the Two Phase Set (2P-set) and use it as a basis for exploring these different CRDT garbage collection possibilities. We do not consider the Add-Remove Partial Order (ARPO) implementation from the previous blog post as the 2P-set encompasses the same challenges with far fewer lines of code. 

## Delta State CRDTs

Describe how Deltas are propogated across nodes

[Delta-state CRDTs](https://arxiv.org/pdf/1603.01529.pdf) help to avoid the issue of sending the entire state of a data type over a network. The potentially large message sizes involved with classical state-based CRDTs result in them only being practical for small objects such as counters. A solution to this problem is only to transmit a delta state that encompasses only the changes made to a replica. These deltas can utilize the same join operations to apply updates to a local state with the additional ability to join multiple deltas into a group. The authors present a few anti-entropy algorithms that are used to ensure convergence and either eventual or causal consistency of the delta CRDT. In the causal consistency algorithm each node maintains two maps, one for keeping track of a sequence of deltas and another for a series of acknowledgements from each neighbor. Recall that the garbage collection scheme presented by Shapiro et al. involves maintaining a set of vector clocks for each replica. With which one can infer causal relationships and determine whether or not an update is "stable" at which point the tombstones associated with it can be deleted. A delta state CRDT instead stores a sequence of deltas and a map of acknowledgements that contains the latest version of a delta acknowledged by neighbors. 

The previous blog post showed that the space complexity for storing the metadata necessary for $N$ nodes is $O(N^2)$. In the case of a delta CRDT we see that the metadata cost at each node for $|A|$ neighbors and $|D|$ stored deltas is $O(A + D)$. So instead of the state scaling quadratically with more replicas we see the state grow linearly depending on how many neighbors each node is keeping track of and how many deltas have been sent. Even though the delta CRDT metadata is stated by the authors to be for garbage collection of deltas cached on each node, they can also be used to kill two birds with one stone. Providing both a means to perform garbage collection on deltas and tombstones due to the fact they both rely on determining causal relationships. When a delta that contains an operation that created a tombstone is garbage collected then we can assume that the tombstone is also garbage collected. This means that the programmer must maintain links between which tombstone applies to which delta but that can be done with a simple pointer.

Since the causally consistent delta CRDT metadata allows us to garbage collect tombstones we can now compare it to the $O(N^2)$ overhead of Shapiro et al.'s approach. If we assume that the set of neighbors is all $N$ nodes we see that the complexity is $O(N+D)$. Therefore if we keep the number of deltas in check and run local garbage collection regularly the metadata overhead is considerably decreased in comparison to the previous approach. Though the approximately $O(N)$ overhead may still prove costly when operating at scale.

Therefore we can see that delta-CRDTs, although a bit more difficult to implement, can solve the problem of exploding state through tombstones.

### Synchronized Garbage Collection

[Paxos Commit](https://lamport.azurewebsites.net/video/consensus-on-transaction-commit.pdf) and Two-Phase Commit protocols

### Space-Saving Optimizations

If a CRDT is not expected to have a long lifespan it is likely sufficient to simply perform some common space saving optimizations in order to minimize the effect of state explosion. As discussed in the previous blog post a tombstone can actually be rather small. In the case of a 2P-set or a ARPO the tombstone set can be represented as a bitmap with each bit corresponding to an element in a set. A bit is set when the corresponding element is deleted. Using this trick the storage space needed for tombstones become negligible. Although this does not eliminate the problem of garbage collection in its entirety as the set can still contain deleted elements after all replicas have marked them as deleted. Therefore this sort of optimization is best applied in conjunction with another garbage collection scheme.

Most programmers would consider a metadata overhead of $O(N^2)$ to be unacceptable. As shown with delta-CRDTs it is possible to abandon vector clocks and still perform distributed garbage collection. While delta-CRDTs are clearly a better solution there are other ways of achieving the same goal of $O(N)$ metadata overhead. Recall that in Shapiro et al.'s scheme the key for garbage collecting tombstones is to prove causal relationships. Peter Bailis among others have described multiple [methods for reducing the space costs of vector clocks](http://www.bailis.org/blog/causality-is-expensive-and-what-to-do-about-it/). Two of the described methods involve decreasing availability and eliminating "happens-before" (causal) relations therefore necessitating synchronization. The third option of explicitly specifying relevant relationships only applies in certain cases (a message replying to another). This leaves us with restricting the number of replicas participating. Either putting a total upper bound on the number of participants or by only requiring a replica to store information about its immediate "neighbors" (likely spatial neighbors to keep latency down). 

## Conclusion

