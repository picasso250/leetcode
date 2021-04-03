#include <vector>
#include <map>
#include <algorithm>
#include <iostream>
using namespace std;

class Solution
{
private:
    long long sumVector(vector<int> &A)
    {
        long long sum = 0;
        for (vector<int>::iterator itor = A.begin(); itor != A.end(); ++itor)
        {
            sum += *itor;
        }
        return sum;
    }

public:
    vector<int> fairCandySwap(vector<int> &A, vector<int> &B)
    {
        vector<int> ret = {0, 0};
        long long asum = this->sumVector(A), bsum = this->sumVector(B);
        // cout<<"asum = "<<asum<<" bsum = "<<bsum<<endl;
        int d = (asum - bsum) / 2;
        // cout <<"d "<<d <<endl;

        // build map
        map<int, int> A_;
        for (vector<int>::iterator itor = A.begin(); itor != A.end(); ++itor)
        {
            A_[*itor] = 0;
        }

        // check
        for (vector<int>::iterator itor = B.begin(); itor != B.end(); ++itor)
        {
            int b = *itor;
            int a = b + d;
            if (A_.find(a) != A_.end())
            {
                ret[0] = a;
                ret[1] = b;
                return ret;
            }
        }
        return ret;
    }
};

int main(int argc, char const *argv[])
{
    vector<int> A = {1, 1};
    vector<int> B = {2, 2};
    Solution s;
    vector<int> r = s.fairCandySwap(A, B);
    cout << r[0] << " " << r[1] << endl;
}