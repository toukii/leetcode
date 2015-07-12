#include <iostream>

using namespace std;

struct TreeNode{
    int val;
    TreeNode *left, *right;
    TreeNode(int v): val(v), left(NULL), right(NULL){}
    TreeNode(int v, TreeNode *l, TreeNode *r): val(v), left(l), right(r){}
};
#include <stack>
void commonAncestorStack(TreeNode *root, TreeNode *p, stack<TreeNode*> &ps, TreeNode *q, stack<TreeNode*> &qs){
    if(!root) return;
    ps.push(root);
    qs.push(root);
    commonAncestorStack(root->left,p,ps,q,qs);
    commonAncestorStack(root->right,p,ps,q,qs);
    if(ps.top()!=p){
        ps.pop();
    }
    if(qs.top()!=q){
        qs.pop();
    }
}
TreeNode* lowestCommonAncestorStack(TreeNode* root, TreeNode* p, TreeNode* q) {
    stack<TreeNode*> ps,qs;
    commonAncestorStack(root,p,ps,q,qs);
    int psl,qsl;
    psl = ps.size();
    qsl = qs.size();
    if(psl>qsl){
        for(int i=0;i<psl-qsl;i++)
            ps.pop();
    }else{
        for(int i=0;i<qsl-psl;i++)
            qs.pop();
    }
    TreeNode *pi,*qi;
    while(!ps.empty()&&!qs.empty()){
        pi = ps.top();
        qi = qs.top();
        if(pi==qi){
            return pi;
        }
        ps.pop();
        qs.pop();
    }
    return NULL;
}
#include <queue>
#include <deque>
void commonAncestorQueue(TreeNode *root, TreeNode *p, deque<TreeNode*> &ps, TreeNode *q, deque<TreeNode*> &qs){
    if(!root) return;
    ps.push_back(root);
    qs.push_back(root);
    commonAncestorQueue(root->left,p,ps,q,qs);
    commonAncestorQueue(root->right,p,ps,q,qs);
    if(ps[ps.size()-1]!=p){
        ps.pop_back();
    }
    if(qs[qs.size()-1]!=q){
        qs.pop_back();
    }
}
TreeNode* lowestCommonAncestorQueue(TreeNode* root, TreeNode* p, TreeNode* q) {
    deque<TreeNode*> ps,qs;
    commonAncestorQueue(root,p,ps,q,qs);
    TreeNode *pi,*qi,*lowestCommon;
    while(!ps.empty()&&!qs.empty()){
        pi = ps[0];
        qi = qs[0];
        if(pi==qi){
            lowestCommon = pi;
        }else{
            return lowestCommon;
        }
        ps.pop_front();
        qs.pop_front();
    }
    if(lowestCommon) return lowestCommon;
    return NULL;
}
bool travelTree(TreeNode *root, TreeNode *target,int &dep,int &num){
    if(!root) return false;
    //cout<<"dep:"<<dep<<", num:"<<num<<", val:"<<root->val<<endl;
    dep ++;
    if(root==target) return true;
    int tnum = num;
    num = (tnum<<1);
    if(travelTree(root->left,target,dep,num)) return true;
    num = (tnum<<1)+1;
    bool r = travelTree(root->right,target,dep,num);
    if(r) return true;
    dep --;
    return false;
}
TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
    int depp,depq,nump,numq;
    depp = 0;depq = 0;nump = 0;numq = 0;
    travelTree(root,p,depp,nump);
    cout<<"p's depth: "<<depp<<", p's num : "<<nump<<endl;
    travelTree(root,q,depq,numq);
    cout<<"q's depth: "<<depq<<", q's num : "<<numq<<endl;
    depp -=2;
    depq -=2;
    TreeNode *pi=root;
    while(depp>=0&&depq>=0){
        bool pb = nump&(1<<depp);
        bool qb = numq&(1<<depq);
        cout<<"last common:"<<pi->val<<endl;
        cout<<pb<<endl;/**/
        if(pb!=qb){
            return pi;
        }
        if(pb){
            pi = pi->right;
        }else{
            pi = pi->left;
        }
        depp --;
        depq --;
    }
    return pi;
}
int main()
{
    TreeNode *a1 = new TreeNode(3);TreeNode *a2 = new TreeNode(5);
    TreeNode *b1 = new TreeNode(0);TreeNode *b2 = new TreeNode(4,a1,a2);
    TreeNode *b3 = new TreeNode(7);TreeNode *b4 = new TreeNode(9);
    TreeNode *c1 = new TreeNode(2,b1,b2);TreeNode *c2 = new TreeNode(8,b3,b4);
    TreeNode *d1 = new TreeNode(6,c1,c2);
    TreeNode *lowestCommon = lowestCommonAncestor(d1,a1,a2);
    if(lowestCommon){
        cout<<lowestCommon->val<<endl;
    }else{
        cout<<"NULL"<<endl;
    }
    delete a1,a2,b1,b2,b3,b4,c1,c2,d1;
    return 0;
}
