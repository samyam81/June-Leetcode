#include <vector>

using namespace std;

class Solution {
private:
    vector<vector<int>> adj, ans;
    
    void DFS(int curr, int root) {
        for (int child : adj[curr]) {
            if (ans[child].empty() || ans[child].back() != root) {
                ans[child].push_back(root);
                DFS(child, root);
            }
        }
    }

public:
    vector<vector<int>> getAncestors(int n, vector<vector<int>>& edges) {
        adj.resize(n);
        ans.resize(n);
        
        for (vector<int>& edge : edges) {
            adj[edge[0]].push_back(edge[1]);
        }

        for (int i = 0; i < n; ++i) {
            DFS(i, i);
        }
        
        return ans;
    }
};
