#include <iostream>
#include <vector>
#include <map>
#include <string>
#include <bitset>
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

void all_pos(vector<int> cart, vector<int> p)
{
	if (cart.empty())
	{
		for (auto i : p)
		{
			cout<<i<<" ";
		}
		cout<<endl;
		return;
	}
	for (auto i = cart.begin(); i < cart.end(); ++i)
	{
		vector<int> rest = cart;
		cout<<"for-- "<<(*i)<<endl;
		rest.erase(i);
		p.push_back(*i);
		all_pos(rest, p);
		p.pop_back();
	}
}

int main(int argc, char const *argv[])
{
	const int N = 8;
	cout<<(1<<2)<<endl;
	unsigned long long a;
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
	vector<bitset<N>> v;
	for (a=0; a<256; a++)
	{
		bitset<N> bs(a);
		cout<<"for "<<a<<endl;
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
			cout<<"sum1 "<<sum1<<", sum2 "<<sum2<<" ";
			cout<<"ok"<<endl;
			v.push_back(bs);
		}
		else
		{
			continue;
		}
		vector<vector<int>> xmap;
		vector<int> posv;
		assert(xmap.size() == posv.size());
		int min_dist = 200;
		vector<int> min_pos;
		vector<int> cur;
		cout<<"all_pos"<<endl;
		vector<int> tmp;
		all_pos(cart1, tmp);
		// init_posv(cart1, posv);
		for (;;)
		{
			// select(cart1);
			// xmap.push_back(cart1);
			// cur.push_back()
			// posv.push_back()
			int distance = 0;
			int time_;
			int fee;
			int old = 0;
			// print_(xmap, pos);
			// for (int i = 0; i < pos.size(); ++i)
			// {
			// 	int node = xmap[i][pos[i]];
			// 	cout<<"distance between "<<old<<" --> "<<node<<" is "<<dist_map[old][node]<<endl;
			// 	assert(old != node);
			// 	distance += dist_map[old][node];
			// 	old = node;
			// }
			// cout<<"distance "<<distance<<endl;
			if (min_dist > distance)
			{
				min_dist = distance;
				// min_pos = pos;
			}
			// if (!next_(cart1, pos))
			// {
			// 	break;
			// }
		}
		// cout<<"min distance "<<min_dist<<endl;
		// print_(xmap, min_pos);
	}
	// cout<<"total max distance "<<max_dist<<endl;
}
