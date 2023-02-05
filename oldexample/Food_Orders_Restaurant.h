#pragma once
#include "common.h"
struct user {
    user() {
        in = false;
    }
    bool in;
    string table;
    map<string, int> food;
};

class Solution {
public:
    set<string> kAllFood;
    vector<user> kVuser;
    Solution()
    {
        kVuser.resize(510);
    }

    vector<vector<string>> displayTable(vector<vector<string>>& orders) {
        int max = 0;
        for (int i = 0; i < (int)orders.size(); i++)
        {
            vector<string> kUser = orders[i];
            int iNum = atoi(kUser[1].c_str());
            if (iNum > max)
            {
                max = iNum;
            }
            kVuser[iNum].in = true;
            kVuser[iNum].table = kUser[1];
            kVuser[iNum].food[kUser[2]] = kVuser[iNum].food[kUser[2]] + 1;
            kAllFood.insert(kUser[2]);
        }

        vector<vector<string>> ans;
        vector<string> lie1;
        lie1.push_back("Table");
        for (set<string>::iterator iter = kAllFood.begin(); iter != kAllFood.end(); ++iter)
        {
            string str = *iter;
            lie1.push_back(str);
        }
        ans.push_back(lie1);

        for (int i = 1; i <= max; i++)
        {
            vector<string> lie;
            if (!kVuser[i].in)
            {
                continue;
            }
            lie.push_back(kVuser[i].table);

            for (set<string>::iterator iter = kAllFood.begin(); iter != kAllFood.end(); ++iter)
            {
                string str = *iter;
                int n = kVuser[i].food[str];
                lie.push_back(to_string(n));
            }
            ans.push_back(lie);
        }

        return ans;
    }
};