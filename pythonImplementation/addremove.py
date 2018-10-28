from abc import ABCMeta, abstractmethod, abstractproperty
from copy import deepcopy
import base64
import random


def random_client_id():
    """Returns a random client identifier"""
    return 'py_%s' % base64.b64encode(str(random.randint(1, 0x40000000)))


class StateCRDT(object):
    __metaclass__ = ABCMeta

    #
    # Abstract methods
    #

    @abstractmethod
    def __init__(self):
        pass

    @abstractproperty
    def value(self):
        """Returns the expected value generated from the payload"""
        pass

    @abstractproperty
    def payload(self):
        """This is a deepcopy-able version of the CRDT's payload.
        If the CRDT is going to be serialized to storage, this is the
        data that should be stored.
        """
        pass

    @classmethod
    @abstractmethod
    def merge(cls, X, Y):
        """Merge two replicas of this CRDT"""
        pass

    #
    # Built-in methods
    #

    def __repr__(self):
        return "<%s %s>" % (self.__class__, self.value)

    def clone(self):
        """Create a copy of this CRDT instance"""
        return self.__class__.from_payload(deepcopy(self.payload))

    @classmethod
    def from_payload(cls, payload, *args, **kwargs):
        """Create a new instance of this CRDT using a payload.  This
        is useful for creating an instance using a deserialized value
        from a datastore."""
        new = cls(*args, **kwargs)
        new.payload = payload
        return new

class SetStateCRDT(StateCRDT, MutableSet):

    def __contains__(self, element):
        return self.value.__contains__(element)

    def __iter__(self):
        return self.value.__iter__()

    def __len__(self):
        return self.value.__len__()





class AddRemovePartialOrder(StateCRDT):

    def __init__(self):
        self.V = TwoPSet()
        self.E = GSet()

    def queryLookup(self, v):
        isPresent = False
        return isPresent

    #depends on the query lookup function
    def queryBefore(self, u, v):
        if queryLookup(self, u) and queryLookup(self, v):
            isBefore = False 
            return isBefore
        else:
            return False

    def updateAtSource(self, u, v, w):
        pass

    def updateDownstream(self, u, v, w):
        pass

    def removeAtSource(self, v):
        pass

    def removeDownstream(self, v):
        pass
