#include <iostream>
#include <vector>
#include <array>
#include <string>
#include <map>

using namespace std;

 struct Interval {
     int start;
     int end;
     Interval() : start(0), end(0) {}
     Interval(int s, int e) : start(s), end(e) {}
 };
class Solution {
public:
    vector<Interval> insert(vector<Interval>& intervals, Interval newInterval) {
    	if (intervals.size() == 0)
    	{
    		intervals.push_back(newInterval);
    		return intervals;
    	}
    	int first_start = intervals[0].start;
    	if (newInterval.end < first_start)
    	{
    		intervals.insert(intervals.begin(), newInterval);
    		return intervals;
    	}
    	if (newInterval.start < first_start && first_start <= newInterval.end)
    	{
    		intervals.begin()->start = newInterval.start;
    	}
    	int last_end = first_start;
    	for (auto it=intervals.begin(); it != intervals.end(); ++it)
    	{
    		if (last_end <= newInterval.start && newInterval.start < it->start)
    		{
    			it->start = newInterval.start;
    		}
            if (it->start <= newInterval.start && newInterval.start <= it->end) {
            	if (newInterval.end <= it->end)
            	{
            		break;
            	}
            	else
            	{
            		// cout<<"expand "<<it->end<<" => "<<newInterval.end<<endl;
            		int end = it->end = newInterval.end;
	    			if ((it+1) < intervals.end())
	    			{
	    				auto del_begin = intervals.end();
	    				auto del_end = intervals.end();
	            		for (auto itr = it+1; itr < intervals.end();)
	            		{
	            			if (itr->end <= end)
	            			{
	            				if (del_begin == intervals.end())
	            				{
	            					del_begin = itr;
	            				}
	            				del_end = itr;
	            				// intervals.erase(itr);
	            			}
	            			else if (itr->start <= end && end <= itr->end)
	            			{
	            				it->end = itr->end;

	            				if (del_begin == intervals.end())
	            				{
	            					del_begin = itr;
	            				}
	            				del_end = itr;
	            				// intervals.erase(itr);
	            			}
	            			else
	            			{
	            			}
	            				itr++;
	            		}
	            		if (del_begin != intervals.end())
	            		{
	            			intervals.erase(del_begin, del_end+1);
	            		}

	    			}
            		break;
            	}
            }
        }
        auto last = intervals.end();
        last--;
        if (last->end < newInterval.start)
        {
        	intervals.push_back(newInterval);
        }
        return intervals;
    }
};

int main(int argc, char const *argv[])
{
	Solution s;
	vector<Interval> intervals = {{1,3},{6,9}};
	Interval newInterval = {2,5};
	auto v = s.insert(intervals, newInterval);
	cout<<"--------------"<<endl;
	for (auto i : v)
	{
		cout<<"["<<i.start<<","<<i.end<<"]"<<endl;
	}

	intervals = {{1,2},{3,5},{6,7},{8,10},{12,16}};
	newInterval = {4,9};
	v = s.insert(intervals, newInterval);
	cout<<"--------------"<<endl;
	for (auto i : v)
	{
		cout<<"["<<i.start<<","<<i.end<<"]"<<endl;
	}

	intervals = {};
	for (int i = 0; i < 10000; ++i)
	{
		intervals.push_back({i*2+1, i*2+2});
	}
	newInterval = {0,20001};
	v = s.insert(intervals, newInterval);
	cout<<"--------------"<<endl;
	for (auto i : v)
	{
		cout<<"["<<i.start<<","<<i.end<<"]"<<endl;
	}
	return 0;
}
