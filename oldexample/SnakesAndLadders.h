#pragma once
#include "common.h"

 struct TreeNode {
      int val;
      TreeNode *left;
     TreeNode *right;
      TreeNode(int x) : val(x), left(NULL), right(NULL) {}
  };
 
class Codec {
public:

    void split(const string& s, vector<string>& tokens, char delim = ' ')
    {
        tokens.clear();
        auto string_find_first_not = [s, delim](size_t pos = 0) -> size_t {
            for (size_t i = pos; i < s.size(); i++) {
                if (s[i] != delim) return i;
            }
            return string::npos;
        };
        size_t lastPos = string_find_first_not(0);
        size_t pos = s.find(delim, lastPos);
        while (lastPos != string::npos) {
            tokens.emplace_back(s.substr(lastPos, pos - lastPos));
            lastPos = string_find_first_not(pos);
            pos = s.find(delim, lastPos);
        }
    }


    // Encodes a tree to a single string.
    string serialize(TreeNode* root) {
        string ans;
        vector<string> kVec;
        queue<TreeNode*> kQue;
        kQue.push(root);
        while (!kQue.empty())
        {
            TreeNode* pkNode = kQue.front();
            if (!pkNode)
            {
                kVec.push_back("null");
            }
            else
            {
                kVec.push_back(to_string(pkNode->val));
                kQue.push(pkNode->left);
                kQue.push(pkNode->right);
            }
            kQue.pop();
        }

        for (int i = 0; i < kVec.size(); i++)
        {
            string s = kVec[i];
            if (i == 0)
            {
                ans = s;
            }
            else
            {
                ans = ans + "," + s;
            }
        }
        return ans;
    }

    // Decodes your encoded data to tree.
    TreeNode* deserialize(string data) {
        vector<string> kVec;
        vector<TreeNode*> kAns;
        split(data, kVec, ',');
        
        int Len = (int)kVec.size();
        string s1 = kVec[0];
        if (Len == 0 || s1 =="null")
        {
            return NULL;
        }
        TreeNode* root = new TreeNode(atoi(kVec[0].c_str()));
        kAns.push_back(root);
        for (int i = 1; i < kVec.size(); i++)
        {
            int bCheck = i % 2;
            int iIndex = (bCheck ? ((i - 1) / 2) : (i / 2 - 1));
            TreeNode* pkPar = kAns[iIndex];
            if (kVec[i] == "null")
            {
                kAns.push_back(NULL);
                continue;
            }

            TreeNode* pkNode = new TreeNode(atoi(kVec[i].c_str()));
            if (bCheck)
            {
                pkPar->left = pkNode;
            }
            else
            {
                pkPar->right = pkNode;
            }
            kAns.push_back(pkNode);
        }
        return root;
    }
};

// Your Codec object will be instantiated and called as such:
// Codec codec;
// codec.deserialize(codec.serialize(root));