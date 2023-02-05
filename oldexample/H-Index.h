#pragma once
#include "common.h"

class Solution {
public:
    int hIndex(vector<int>& citations) {
        map<int, int>mMap;
        vector<int> mVec;
        for (int i = 0; i < citations.size(); i++)
        {
            int x = citations[i];
            if (x == 0) continue;
            mMap[x]++;
            if (mMap[x] == 1)
            {
                mVec.push_back(x);
            }
        }
        int sum = 0;
        int ans = 0;
        if (mVec.size() == 1)
        {
            return min(mVec[0], mMap[mVec[0]]);
        }

        sort(mVec.begin(), mVec.end());
        for (int i = mVec.size() - 1; i >= 0; i--)
        {
            int x = mVec[i];
            sum += mMap[x];
            if (sum <= x && ans < sum)
            {
                ans = sum;
            }
        }
        return ans;
    }
};