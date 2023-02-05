#pragma once
#include "common.h"
class Solution {
    // vector<vector<int>> beginBoard = {{1,2,3},{4,5,6}};
    int ed[50000];
    int pd[50000];
    queue<vector<vector<int>>> bfs;
    vector<vector<int>> BeginInt;
    int Noop;
    int arrX[4] = {1,0,-1,0};
    int arrY[4] = {0,1,0,-1};
public:
    Solution()
    {
        vector<int> col1 = { 1,2,3 };
        vector<int> col2 = { 4,5,0 };
        BeginInt.push_back(col1);
        BeginInt.push_back(col2);
        memset(ed, 0, sizeof(ed));
        memset(pd, 0, sizeof(pd));
        Noop = Get_Base6Int(BeginInt);

        bfs.push(BeginInt);
        while (!bfs.empty())
        {
            vector<vector<int>>& board = bfs.front();
  
            int param = Get_Base6Int(board);
            if (ed[param])
            {
                bfs.pop();
                continue;
            }
            ed[param] = 1;
            int x = -1, y = -1;
            Decimal_Base6(board,x,y);
            //if (x < 0 || x >= 2 || y < 0 || y >= 3)
            //{
            //    continue;
            //}
            for (int i = 0; i < 4; i++)
            {
                int nx = x + arrX[i];
                int ny = y + arrY[i];
                if(nx < 0 || nx >= 2|| ny < 0 || ny >= 3)
                {
                    continue;
                }
                vector<vector<int>> kNow = board;
                kNow[nx][ny] = kNow[nx][ny] + kNow[x][y];
                kNow[x][y] = kNow[nx][ny] - kNow[x][y];
                kNow[nx][ny] = kNow[nx][ny] - kNow[x][y];
                int NParam = Get_Base6Int(kNow);
                if (ed[NParam]) continue;
                
                pd[NParam] = pd[param] + 1;
                bfs.push(kNow);
            }
            bfs.pop();
        }
    }

    int Get_Base6Int(vector<vector<int>>& board)
    {
        int ans = 0;
        for (int i = 0; i < 2; i++)
        {
            for (int j = 0; j < 3; j++)
            {
                int value = board[i][j];
                ans = ans * 6 + value;
            }
        }
        return ans;
    }

    void Decimal_Base6(vector<vector<int>>& board,int& x,int& y)
    {
        for (int i = 0; i < 2; i++)
        {
            for (int j = 0; j < 3; j++)
            {
                if (board[i][j] == 0)
                {
                    x = i;
                    y = j;
                    break;
                }
            }
        }
    }

    int slidingPuzzle(vector<vector<int>>& board) {
        //return Get_Base6Int(board);
        int query = Get_Base6Int(board);
        if(ed[query])
            return pd[query];

        return -1;
    }
};