#include <vector>
#include <algorithm>

using namespace std;

class Solution {
    static bool comp(vector<int>& a, vector<int>& b) {
        return a[0] > b[0];
    }

    int findparent(int node, vector<int>& parent) {
        if (node == parent[node]) 
            return node;
        
        return parent[node] = findparent(parent[node], parent); 
    }

    bool unionn(int u, int v, vector<int>& parent, vector<int>& rank) {
        u = findparent(u, parent);
        v = findparent(v, parent);

        if (u != v) {
            if (rank[u] < rank[v]) 
                parent[u] = v;
             else if (rank[u] > rank[v]) 
                parent[v] = u;
             else {
                parent[v] = u;
                rank[u]++;
            }
            return true;
        }
        return false;
    }

public:
    int maxNumEdgesToRemove(int n, vector<vector<int>>& edges) {
        vector<int> parentAlice(n + 1);
        vector<int> parentBob(n + 1);
        vector<int> aliceRank(n + 1, 1);
        vector<int> bobRank(n + 1, 1);

        sort(edges.begin(), edges.end(), comp);

        for (int i = 1; i <= n; ++i) {
            parentAlice[i] = i;
            parentBob[i] = i;
        }

        int removeEdges = 0;
        int mergeAlice = 1;
        int mergeBob = 1;

        for (auto& it : edges) {
            int type = it[0];
            int u = it[1];
            int v = it[2];

            if (type == 3) {
                bool tempAlice = unionn(u, v, parentAlice, aliceRank);
                bool tempBob = unionn(u, v, parentBob, bobRank);

                if (tempAlice) 
                    mergeAlice++;
                
                if (tempBob) 
                    mergeBob++;
                
                if (!tempAlice && !tempBob) 
                    removeEdges++;
                
            } else if (type == 1) {
                bool tempAlice = unionn(u, v, parentAlice, aliceRank);
                if (tempAlice) 
                    mergeAlice++;
                 else 
                    removeEdges++;
                
            } else if (type == 2) {
                bool tempBob = unionn(u, v, parentBob, bobRank);
                if (tempBob) 
                    mergeBob++;
                 else 
                    removeEdges++;
                
            }
        }

        if (mergeAlice != n || mergeBob != n) 
            return -1;
        
        return removeEdges;
    }
};
