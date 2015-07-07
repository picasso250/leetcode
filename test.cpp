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
int main(int argc, char const *argv[])
{
	std::vector<Interval> v = {{1,2},{3,4}};
    bool first = true;
    for (auto it = v.begin(); it != v.end(); ++it)
    {
        *it.
    }
	return 0;
}
