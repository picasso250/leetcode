#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

struct Interval {
    int start;
    int end;
};

#define BEFORE(v,p) (v    < (p)->start)
#define AFTER(v,p)  ((p)->end < v)

/**
 * return point to the next
 */
struct Interval * copy(struct Interval *res, struct Interval *intervals, int count)
{
    // printf("copy %d\n", count);
    if (count > 0) {
        memcpy((void*)(res), (void*)(intervals), sizeof(struct Interval) * (count));
    }
    return res+count;
}

struct Interval * bsearch_interval(int value, struct Interval * intervals, int size, int *inner)
{
    // printf("find %d from [%d,%d] ==> %d\n", value, intervals[0].start, intervals[0].end, size);
    int low = 0;
    int high = size - 1;
    while (1)
    {
        int mid;
        if ((high - low) <= 1)
        {
            if (high != low)
            {
                if (value <= intervals[low].end)
                {
                    mid = low;
                }
                else if (value >= intervals[high].start)
                {
                    mid = high;
                }
                else
                {
                    mid = low;
                }
            }
            else
            {
                mid = low;
            }
            int start = intervals[mid].start;
            int end = intervals[mid].end;
            if (value < start)
            {
                (*inner) = -1;
                // printf("%d < start:%d, high=%d\n", value, start, mid);
            }
            else if (end < value)
            {
                // printf("end: %d < %d, low=%d\n", end, value, mid);
                (*inner) = 1;
            }
            else
            {
                // printf("%d in [%d,%d]\n", value, start, end);
                (*inner) = 0;
            }
            return intervals + mid;
        }
        else
        {
            mid = (low+high)/2;
        }
        // printf("[%d,%d] => %d\n", low, high, mid);
        // assert(low <= high);
        int start = intervals[mid].start;
        int end = intervals[mid].end;
        if (value < start)
        {
            (*inner) = -1;
            high = mid;
            // printf("%d < start:%d, high=%d\n", value, start, mid);
        }
        else if (end < value)
        {
            // printf("end: %d < %d, low=%d\n", end, value, mid);
            (*inner) = 1;
            low = mid;
        }
        else
        {
            // printf("%d in [%d,%d]\n", value, start, end);
            (*inner) = 0;
            return intervals + mid;
        }
    }
}

/**
 * Return an array of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
struct Interval* insert(struct Interval* intervals, int intervalsSize, struct Interval newInterval, int* returnSize) {
    // assert(newInterval.start <= newInterval.end);
    // ni.start can equal ni.end
    if (intervalsSize == 0)
    {
    	struct Interval *i = malloc(sizeof(newInterval));
    	*returnSize = 1;
        copy(i, &newInterval, 1);
        return i;
    }
    // find the left using bsearch
    int left = 0;
    int right = 0;
    int inner;
    struct Interval *p = bsearch_interval(newInterval.start, intervals, intervalsSize, &inner);
    if (inner == 0)
    {
        newInterval.start = p->start;
        left = p - intervals;
    }
    else if (inner == 1)
    {
        left = p - intervals + 1;
        if (p != (intervals+intervalsSize))
        {
            p++;
        }
        else
        {
            left = p - intervals;
        }
    }
    else
    {
        left = p - intervals;
    }
    // printf("left %d\n", left);
    int n = intervalsSize - left;
    p = bsearch_interval(newInterval.end, p, n, &inner);
    // printf("find %d in [%d], inner %d\n", newInterval.end, p - intervals, inner);
    if (inner == 0)
    {
        newInterval.end = p->end;
        p++;
    }
    else if (inner == -1)
    {

    }
    else
    {
        p++;
    }
    right = intervals + intervalsSize - (p);
    // printf("right %d\n", right);
    (*returnSize) = left + right + 1;
    struct Interval *res = malloc(sizeof(newInterval) * (*returnSize));
    struct Interval *rp = res;
    rp = copy(rp, intervals, left);
    rp->start = newInterval.start;
    rp->end = newInterval.end;
    rp = copy(rp+1, p, right);
    return res;
}

int main(int argc, char const *argv[])
{
    int returnSize;
    struct Interval newInterval = {2,5};
    struct Interval intervals[] = {{1,3},{6,9}};
    struct Interval * res = insert(intervals, sizeof(intervals)/sizeof(struct Interval), newInterval, &returnSize);
    printf("1. ### size %d\n", returnSize);
    for (int i = 0; i < returnSize; ++i)
    {
        printf("[%d,%d]\n", res[i].start, res[i].end);
    }

    struct Interval newInterval2 = {4,9};
    struct Interval intervals2[] = {{1,2},{3,5},{6,7},{8,10},{12,16}};
    res = insert(intervals2, sizeof(intervals2)/sizeof(struct Interval), newInterval2, &returnSize);
    printf("2. ### size %d\n", returnSize);
    for (int i = 0; i < returnSize; ++i)
    {
        printf("[%d,%d]\n", res[i].start, res[i].end);
    }

    struct Interval newInterval3 = {6,6};
    struct Interval intervals3[] = {{3,5},{12,15}};
    res = insert(intervals3, sizeof(intervals3)/sizeof(struct Interval), newInterval3, &returnSize);
    printf("3. ### size %d\n", returnSize);
    for (int i = 0; i < returnSize; ++i)
    {
        printf("[%d,%d]\n", res[i].start, res[i].end);
    }

    // struct Interval newInterval4 = {0,20001};
    // int n = 10000;
    // struct Interval *intervals4 = malloc(n*sizeof(struct Interval));
    // for (int i = 0; i < n; ++i)
    // {
    //     intervals4[i].start = i*2+1;
    //     intervals4[i].end = i*2+2;
    // }
    // res = insert(intervals4, n, newInterval4, &returnSize);
    // printf("#### size %d\n", returnSize);
    // for (int i = 0; i < returnSize; ++i)
    // {
    //     printf("[%d,%d]\n", res[i].start, res[i].end);
    // }

    struct Interval intervals6[] = {{1,5}};
    struct Interval newInterval5 = {6,8};
    res = insert(intervals6, sizeof(intervals6)/sizeof(struct Interval), newInterval5, &returnSize);
    printf("5. ### size %d\n", returnSize);
    for (int i = 0; i < returnSize; ++i)
    {
        printf("[%d,%d]\n", res[i].start, res[i].end);
    }

    struct Interval intervals7[] = {{1,4},{9,12},{19,22}};
    struct Interval newInterval7 = {7,13};
    res = insert(intervals7, sizeof(intervals7)/sizeof(struct Interval), newInterval7, &returnSize);
    printf("### size %d\n", returnSize);
    for (int i = 0; i < returnSize; ++i)
    {
        printf("[%d,%d]\n", res[i].start, res[i].end);
    }

    struct Interval intervals_8[] = {{2,7},{8,8},{10,10},{12,13},{16,19}};
    struct Interval newInterval_8 = {9,17};
    res = insert(intervals_8, sizeof(intervals_8)/sizeof(struct Interval), newInterval_8, &returnSize);
    printf("8. ### size %d\n", returnSize);
    for (int i = 0; i < returnSize; ++i)
    {
        printf("[%d,%d]\n", res[i].start, res[i].end);
    }

    // struct Interval intervals_9[] = {{1,4},{10,12},{13,14},{16,16},{19,20},{21,24},{33,33},{36,39},{44,46},{48,50}};
    // struct Interval newInterval_9 = {5,13};
    // res = insert(intervals_9, sizeof(intervals_9)/sizeof(struct Interval), newInterval_9, &returnSize);
    // printf("9. ### size %d\n", returnSize);
    // for (int i = 0; i < returnSize; ++i)
    // {
    //     printf("[%d,%d]\n", res[i].start, res[i].end);
    // }

    struct Interval intervals_10[] = {{1,5}};
    struct Interval newInterval_10 = {2,3};
    res = insert(intervals_10, sizeof(intervals_10)/sizeof(struct Interval), newInterval_10, &returnSize);
    printf("10. ### size %d\n", returnSize);
    for (int i = 0; i < returnSize; ++i)
    {
        printf("[%d,%d]\n", res[i].start, res[i].end);
    }

    return 0;
}
