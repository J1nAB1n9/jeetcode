#pragma once
#include "common.h"

struct node {
    node()
    {
        x = 0;
        y = 0;
    }
    int x;
    int y;
};

class Solution {
    int SumArea(node a, node b)
    {
        return abs(a.x - b.x) * abs(a.y - b.y);
    }
public:
    map<node, int> kmap;
    bool isRectangleCover(vector<vector<int>>& rectangles) {
        int len = (int)rectangles.size();
        int sumAres = 0;
        for (int i = 0; i < len; i++)
        {
            node node1;
            node node2;

            node1.x = rectangles[i][0];
            node1.y = rectangles[i][1];

            node2.x = rectangles[i][3];
            node2.y = rectangles[i][4];

            if (kmap.count(node1))
            {
                kmap[node1] = 1 + kmap[node1];
            }
            else
            {
                kmap[node1] = 1;
            }

            if (kmap.count(node2))
            {
                kmap[node2] = 1 + kmap[node2];
            }
            else
            {
                kmap[node2] = 1;
            }

            sumAres += SumArea(node1, node2);

            if (kmap[node1] > 4 || kmap[node2] > 4)
            {
                return false;
            }
        }

        node arrNode[4];
        int i = 0;
        for (map<node, int>::iterator itr = kmap.begin(); itr != kmap.end(); ++itr)
        {
            int num = itr->second;
            if (num == 1)
            {
                arrNode[i++] = itr->first;
            }
            else if (num != 4 && num != 2)
            {
                return false;
            }

            if (i > 4) return false;
        }


        if (i != 4) return false;
        if ((SumArea(arrNode[0], arrNode[1]) + SumArea(arrNode[0], arrNode[2]) + SumArea(arrNode[0], arrNode[3])) != sumAres)
            return false;

        return true;
    }
};