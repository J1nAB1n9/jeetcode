#include "common.h"
#define iMaxLevel 8
class Node {
public:
    Node() {}
    Node(int _Level, int _iValue);
    ~Node();

    bool bHasPush;
    int iValue;
    Node** pkNext;
};
Node::Node(int _Level, int _iValue) {
    bHasPush = true;
    iValue = _iValue;
    pkNext = new Node*[_Level + 1];

    memset(pkNext,0,sizeof(Node*) * (_Level + 1));
}
Node::~Node() {
    delete[] pkNext;
}
class Skiplist {
public:
    Skiplist() {
        iNowLevel = 0;
        iValueNums = 0;

        pkHeader = new Node(0, iMaxLevel);
    }

    ~Skiplist() {
        delete pkHeader;
    }

    bool search(int target) {
        Node* _p = pkHeader;
        for (int i = iNowLevel; i >= 0; i--) {
            while (_p->pkNext[i] && _p->pkNext[i]->iValue < target)
            {
                _p = _p->pkNext[i];
            }
        }

        _p = _p->pkNext[0];
        if (_p && _p->iValue == target) {
            return true;
        }

        return false;
    }

    void add(int num) {
        Node* current = pkHeader;

        Node* update[iMaxLevel + 1];
        for (int i = iNowLevel; i >= 0; i--) {
            while (current->pkNext[i] && current->pkNext[i]->iValue < num)
            {
                current = current->pkNext[i];
            }
            update[i] = current;
        }

        current = current->pkNext[0];
        if (current && current->iValue == num) {
            return;
        }

        if (current == NULL || current->iValue != num) {
            int iRandom = randlvl();
            if (iRandom > iNowLevel) {
                for (int i = iRandom; i > iNowLevel; i--) {
                    update[i] = pkHeader;
                }

                iNowLevel = iRandom;
            }

            Node* _node = new Node(iMaxLevel,num);
            for (int i = 0; i <= iRandom; i++) {
                _node->pkNext[i] = update[i]->pkNext[i];
                update[i]->pkNext[i] = _node;
            }

            iValueNums++;
        }
    }

    bool erase(int num) {
        Node* current = pkHeader;
        Node* update[iMaxLevel + 1];
        memset(update, 0, sizeof(Node*) * (iMaxLevel + 1));

        for (int i = iNowLevel; i >= 0; i--) {
            while (current->pkNext[i] != NULL && current->pkNext[i]->iValue < num) {
                current = current->pkNext[i];
            }
            update[i] = current;
        }

        current = current->pkNext[0];
        if (current != NULL && current->iValue == num) {
            for (int i = 0; i <= iNowLevel; i++) {
                if (update[i]->pkNext[i] != current)
                    break;

                update[i]->pkNext[i] = current->pkNext[i];
            }

            while (iNowLevel > 0 && pkHeader->pkNext[iNowLevel] == 0) {
                iNowLevel--;
            }

            return true;
        }

        return false;
    }

    int randlvl(){
        return rand() % iMaxLevel + 1;
    }

public:
    int iNowLevel;
    Node* pkHeader;
    int iValueNums;
};

/**
 * Your Skiplist object will be instantiated and called as such:
 * Skiplist* obj = new Skiplist();
 * bool param_1 = obj->search(target);
 * obj->add(num);
 * bool param_3 = obj->erase(num);
 */