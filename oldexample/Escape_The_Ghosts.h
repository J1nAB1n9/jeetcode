#pragma once
#include "common.h"
struct node {
    node(int px, int py, int pn)
    {
        x = px;
        y = py;
        n = pn;
    }
    int x;
    int y;
    int n;
};
int posx[4] = { 0,1,0,-1 };
int posy[4] = { 1,0,-1,0 };
class Solution {

private:
    bool checkGhosts(const vector<vector<int>>& ghosts, int x, int y, int n)
    {
        for (int i = 0; i < ghosts.size(); i++)
        {
            int iCheck = abs(ghosts[i][0] - x) + abs(ghosts[i][1] - y);
            if (iCheck <= n)
                return false;
        }

        return true;
    }

    int getMaxAns(const vector<vector<int>>& ghosts, const vector<int>& target)
    {
        int maxn = 0;
        for (int i = 0; i < ghosts.size(); i++)
        {
            int iCheck = abs(ghosts[i][0] - target[0]) + abs(ghosts[i][1] - target[1]);
            if (iCheck > maxn)
                maxn = iCheck;
        }
        return maxn;
    }

    bool bfs(int maxn, const vector<vector<int>>& ghosts, const vector<int>& target)
    {
        int n = 0;
        while (!kQueue.empty())
        {
            kQueue.pop();
        }
        kQueue.push(node(0, 0, 0));
        while (!kQueue.empty())
        {
            node knode = kQueue.front();
            kQueue.pop();
            if (knode.n >= maxn)
            {
                continue;
            }

            if (!checkGhosts(ghosts, knode.x, knode.y, knode.n))
            {
                continue;
            }

            if (knode.x == target[0] && knode.y == target[1])
            {
                return true;
            }

            for (int i = 0; i < 4; i++)
            {
                int px = knode.x + posx[i];
                int py = knode.y + posy[i];
                kQueue.push(node(px, py, knode.n + 1));
            }
        }

        return false;
    }
public:
    bool escapeGhosts(vector<vector<int>>& ghosts, vector<int>& target) {
        int maxn = getMaxAns(ghosts, target);
        if (maxn <= 0) return false;
        return bfs(maxn, ghosts, target);
    }
private:
    queue<node> kQueue;
};
