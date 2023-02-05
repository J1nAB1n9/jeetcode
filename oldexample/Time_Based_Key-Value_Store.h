#pragma once
#include "common.h"
struct Node
{
    Node(string val, int timestamp) :val(val), timestamp(timestamp), next(NULL) {}

    string val;
    int timestamp;
    Node* next;
};
typedef map<string, Node*> MapNode;

class TimeMap {
public:
    /** Initialize your data structure here. */
    MapNode key2Node;
    TimeMap() {
        key2Node.clear();
    }

    void push(Node* root, string val, int timestamp_prev)
    {
        Node* ptl = root;
        while (ptl->next)
        {
            Node* next = ptl->next;
            if (timestamp_prev > next->timestamp)
            {
                Node* nNode = new Node(val, timestamp_prev);
                ptl->next = nNode;
                nNode->next = next;
                return;
            }
            else
            {
                ptl = next;
            }
        }

        Node* nNode = new Node(val, timestamp_prev);
        ptl->next = nNode;
        return;
    }

    string getV(Node* root, int timestamp_prev)
    {
        Node* ptl = root;
        while (ptl)
        {  
            if (timestamp_prev < ptl->timestamp || !ptl->timestamp)
            {
                ptl = ptl->next;
            }
            else
            {
                return ptl->val;
            }
        }

        return "";
    }

    void set(string key, string value, int timestamp) {
        MapNode::iterator itr = key2Node.find(key);
        if (itr == key2Node.end())
        {
            Node* head = new Node("", 0);
            push(head, value, timestamp);
            key2Node[key] = head;
        }
        else
        {
            Node* head = itr->second;
            push(head, value, timestamp);
        }
    }

    string get(string key, int timestamp) {
        MapNode::iterator itr = key2Node.find(key);
        if (itr == key2Node.end())
        {
            return "";
        }
        else
        {
            Node* head = itr->second;
            return getV(head, timestamp);
        }
    }
};

/**
 * Your TimeMap object will be instantiated and called as such:
 * TimeMap* obj = new TimeMap();
 * obj->set(key,value,timestamp);
 * string param_2 = obj->get(key,timestamp);
 */