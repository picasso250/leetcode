#include <iostream>
#include <vector>
#include <map>
#include <string>
#include <bitset>
#include <algorithm>
#include <cmath>
#include <cassert>

using namespace std;

bool prop_weight(float w)
{
	return w > 0 && w <= 35;
}

void print_(vector<vector<int>> &xmap, vector<int> &pos)
{
	cout<<"[";
	for (int i : pos)
	{
		cout<<i<<" ";
	}
	cout<<"] ";
	for (int i = 0; i < xmap.size(); ++i)
	{
		cout<<xmap[i][pos[i]]<<" ";
	}
	cout<<endl;
}

void printv(vector<int> v)
{
	cout<<"[ ";
	for (auto i : v)
	{
		cout<<i<<" ";
	}
	cout<<"]";
}

int dist_map[9][9] = {
	{ 0, 13, 16, 24, 23, 26,  7, 22,  1},
	{13,  0,  4, 12, 13, 16, 16, 10, 12},
	{16,  4,  0,  8,  9, 12, 20,  9, 15},
	{24, 12,  8,  0,  4,  7, 28,  2, 23},
	{23, 13,  9,  4,  0,  4, 26,  2, 22},
	{26, 16, 12,  7,  4,  0, 26,  2, 22},
	{ 7, 16, 20, 28, 26, 29,  0, 27,  6},
	{22, 10,  9,  2,  2,  5, 27,  0, 21},
	{ 1, 12, 15, 23, 22, 25,  6, 21,  0}
};
map<int,int> node_weight = {
	{1,  5},
	{2,  6},
	{3,  8},
	{4, 10},
	{5,  2},
	{6,  4},
	{7,  9},
	{8, 10}
};
int get_dist(vector<int> cart)
{
	int distance = 0;
	int old = 0;
	for (int i = 0; i < cart.size(); ++i)
	{
	     int node = cart[i];
	     // cout<<"distance between "<<old<<" --> "<<node<<" is "<<dist_map[old][node]<<endl;
	     assert(old != node);
	     distance += dist_map[old][node];
	     old = node;
	}
	distance += dist_map[old][0];
	// printv(cart);
	// cout<<" distance "<<distance<<endl;
	return distance;
}
int get_fee(vector<int> cart, int sum)
{
	int total = 0;
	int old = 0;
	for (int i = 0; i < cart.size(); ++i)
	{
	     int node = cart[i];
	     // cout<<"distance between "<<old<<" --> "<<node<<" is "<<dist_map[old][node]<<endl;
	     assert(old != node);
	     total += dist_map[old][node] * 7 * sum;
	     // cout<<"weight: "<<(sum*0.1)<<", ("<<old<<" -> "<<node<<") fee is "<<(dist_map[old][node] * 0.07 * sum)<<endl;
	     sum -= node_weight[node];
	     old = node;
	}
	// printv(cart);
	// cout<<" total "<< total*0.01 <<endl;
	return total;
}
int get_time(vector<int> cart)
{
	double total = 0;
	int old = 0;
	for (int i = 0; i < cart.size(); ++i)
	{
	     int node = cart[i];
	     // cout<<"distance between "<<old<<" --> "<<node<<" is "<<dist_map[old][node]<<endl;
	     assert(old != node);
	     total += 1.0 * dist_map[old][node] / (40.0 / 60) + node_weight[node] * 10;
	     // cout<<"weight: "<<(sum*0.1)<<", ("<<old<<" -> "<<node<<") fee is "<<(dist_map[old][node] * 0.07 * sum)<<endl;
	     old = node;
	}
	total += 1.0 * dist_map[old][0] / (40.0 / 60);
	// printv(cart);
	// cout<<" total "<< total*0.01 <<endl;
	return round(total);
}

int main(int argc, char const *argv[])
{
	const int N = 8;
	cout<<(1<<2)<<endl;
	unsigned long long a;
	int total = 0;
	
	vector<bitset<N>> v;
	int min_dist = 20000;
	int min_fee = 20000;
	int min_time = 20000;
	vector<int> min_dist_v1;
	vector<int> min_dist_v2;
	vector<int> min_fee_v1;
	vector<int> min_fee_v2;
	vector<int> min_time_v1;
	vector<int> min_time_v2;
	for (a=0; a<256; a++)
	{
		bitset<N> bs(a);
		cout<<"## "<<a<<endl;
		int sum1 = 0, sum2 = 0;
		vector<int> cart1;
		vector<int> cart2;
		for (uint i = 0; i < 8; ++i)
		{
			// cout<< (i+1) <<" is on car ";
			if (bs.test(i)) {
				// cout<<1;
				sum1 += node_weight[i+1];
				cart1.push_back(i+1);
			} else {
				// cout<<2;
				sum2 += node_weight[i+1];
				cart2.push_back(i+1);
			}
			// cout<<endl;
		}
		if (prop_weight(sum1) && prop_weight(sum2))
		{
			// cout<<"sum1 "<<sum1<<", sum2 "<<sum2<<" ";
			// cout<<"ok"<<endl;
			v.push_back(bs);
		}
		else
		{
			continue;
		}
		vector<vector<int>> xmap;
		vector<int> posv;
		assert(xmap.size() == posv.size());
		do {
			int distance = 0;
			int time_;
			int fee;
			do {
				total++;
				distance = get_dist(cart1) + get_dist(cart2);
				fee = get_fee(cart1, sum1) + get_fee(cart2, sum2);
				int t1 = get_time(cart1);
				int t2 = get_time(cart2);
				time_ = t1 > t2 ? t1 : t2;
				if (min_dist > distance)
				{
					min_dist = distance;
					min_dist_v1 = cart1;
					min_dist_v2 = cart2;
				}
				if (min_fee > fee)
				{
					min_fee = fee;
					min_fee_v1 = cart1;
					min_fee_v2 = cart2;
				}
				if (min_time > time_)
				{
					min_time = time_;
					min_time_v1 = cart1;
					min_time_v2 = cart2;
				}
			} while (next_permutation(cart2.begin(), cart2.end()));
		} while (next_permutation(cart1.begin(), cart1.end()));
		// cout<<"\tmin distance "<<min_dist<<endl;
		cout<<"\tmin fee "<< (min_fee * 0.01) <<endl;
	}
	cout<<"min distance "<<min_dist<<endl;
	printv(min_dist_v1);
	printv(min_dist_v2);
	cout<<endl;
	cout<<"min time "<< (min_time) <<" minute"<<endl;
	printv(min_time_v1);
	printv(min_time_v2);
	cout<<endl;
	cout<<"min fee "<< (min_fee * 0.01) <<endl;
	printv(min_fee_v1);
	printv(min_fee_v2);
	cout<<endl;
	cout<<total<<endl;
}
