#pragma once
#include "common.h"

class Solution {
public:
    int numberOfArithmeticSlices(vector<int>& nums) {
        int iLen = nums.size();
        if (iLen < 3)
        {
            return 0;
        }
        vector<int> pd;
        int p = 1;
        int Jizhun = nums[1] - nums[0];
        for (int i = 2; i < iLen; i++)
        {
            if (nums[i] - nums[i - 1] == Jizhun)
            {
                p++;
            }
            else if (p)
            {
                pd.push_back(p + 1);
                p = 1;
                Jizhun = nums[i] - nums[i - 1];
            }
        }
        pd.push_back(p + 1);
        int ans = 0;
        for (int i = 0; i < pd.size(); i++)
        {
            if (pd[i] <= 2)
            {
                continue;
            }
            int n = pd[i];
            ans += ((n - 2) * (n - 1) / 2);
        }
        return ans;
    }
};