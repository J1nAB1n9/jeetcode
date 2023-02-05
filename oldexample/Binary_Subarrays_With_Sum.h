#pragma once
#include "common.h"
class Solution {
public:
    int numSubarraysWithSum(vector<int>& nums, int goal) {
        vector<int> intAns;
        int queSum = 0, n = nums.size(), ans = 0;
        for (int i = 0; i < n; i++)
        {
            queSum += nums[i];
            intAns.push_back(queSum);
        }

        int i = 0, j = 1;
        if (n == 1)
        {
            if (nums[0] == goal)
            {
                return 1;
            }
            else
            {
                return 0;
            }
        }

        for (; i < n; i++)
        {
            j = i;
            if (i != 0) {
                while (j < n && (intAns[j] - intAns[i - 1]) <= goal)
                {
                    if ((intAns[j] - intAns[i - 1]) == goal)
                        ans++;
                    j++;
                }
            }
            else {
                while (j < n && intAns[j] <= goal)
                {
                    if (intAns[j] == goal)
                        ans++;
                    j++;
                }
            }
        }

        return ans;
    }
};