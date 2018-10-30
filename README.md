# CRDTexperiments

Implementation of Two Phase and Growth only sets from the paper *A Comprehensive Study of Convergent and Commutative Replicated Data Types* by Shapiro et al.

The two set implementations are used to put together an Add-Remove Partial Order graph for the purposes of experiments with garbage collection in CRDT's. In the case of the Partial order graph GC involves cleaning up tombstones left by removing vertices.

The specifications described by Shapiro et al. are incomplete as they do not provide semantic rules for handling adding and removing edges along with the vertices.
