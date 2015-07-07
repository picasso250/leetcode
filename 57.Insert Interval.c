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
    struct Interval *res = malloc(sizeof(newInterval) * (intervalsSize+1));
    int copy_to = -1;
    struct Interval *rp;
    for (int i = 0; i < intervalsSize; ++i)
    {
        struct Interval *p = intervals + i;
        // printf("for [%d,%d]\n", p->start, p->end);
    	if (BEFORE(newInterval.start, p))
    	{
            // printf("[%d,%d].start before [%d,%d]\n", newInterval.start, newInterval.end, p->start, p->end);
            // 1. ensure before has copied and do
            copy_to = i-1;
            rp = copy(res, intervals, copy_to+1);
    		if (BEFORE(newInterval.end, p))
    		{
    			// 1.1 insert and copy all after
			    rp = copy(rp, &newInterval, 1);
                int n = intervalsSize - i; // i included
			    copy(rp, intervals+i, n);
                (*returnSize) = intervalsSize+1;
			    return res;
    		}
            else if (AFTER(newInterval.end, p))
            {
                // 1.3 delete and check after
                for (int j = 0; j < intervalsSize; ++j)
                {
                    if (AFTER(newInterval.end, intervals+j))
                    {
                        // 1.3.3 go on
                        // printf("1.3.3 go on with [%d,%d]\n", (intervals+j)->start, (intervals+j)->end);
                    }
                    else if (BEFORE(newInterval.end, intervals+j))
                    {
                        // 1.3.1 delete before, add me, copy after
                        copy(rp, &newInterval, 1);
                        int n = intervalsSize-j; // j included
                        copy(rp+1, intervals+j, n);
                        (*returnSize) = i+1+n;
                        return res;
                    }
                    else
                    {
                        // 1.3.2 delete before, join me, copy after
                        newInterval.end = (intervals+j)->end;
                        rp = copy(rp, &newInterval, 1);
                        int n = intervalsSize-j-1;
                        copy(rp, intervals+j, n);
                        (*returnSize) = i+1+n;
                        return res;
                    }
                }
                // after
                copy(rp, &newInterval, 1);
                (*returnSize) = i+1;
                return res;
            }
            else
            {
                // 1.2 join and copy all after
                rp->start = newInterval.start;
                rp->end = p->end;
                int n = intervalsSize - i - 1;
                copy(rp+1, intervals + i + 1, );
                (*returnSize) = intervalsSize;
                return res;
            }
    	}
    	else if (AFTER(newInterval.start, p))
    	{
    		copy_to = i;
    	}
    	else// if (IN(newInterval.start, p))
    	{
            // 2 ensure before copied, and do
            // printf("start in [%d,%d]\n", p->start, p->end);
            rp = copy(res, intervals, copy_to+1);
    		newInterval.start = p->start;
            for (int j = i; j < intervalsSize; ++j)
            {
                if (BEFORE(newInterval.end, intervals+j))
                {
                    // 2.3 ignore all below
                    // printf("2.3 ignore all below\n", (intervals+j)->start, (intervals+j)->end);
                    rp = copy(rp, &newInterval, 1);
                    int n = intervalsSize - j; // j included
                    // printf("%d-%d=%d j included\n", intervalsSize, j, n);
                    copy(rp, intervals+j, n);
                    (*returnSize) = i+1+n;
                    return res;
                }
                if (AFTER(newInterval.end, intervals+j))
                {
                    // 2.2 check more, ie. do nothing
                    // printf("2.2 check more, ie. do nothing. skip [%d,%d]\n", (intervals+j)->start, (intervals+j)->end);
                }
                else
                {
                    // 2.1 ignore, ie. copy all
                    // but with consideration of that it could be entered
                    // by last loop, so we just need to copy the rest
                    // ignore means ignore below interval(s)
                    // printf("2.1 end in [%d,%d]\n", (intervals+j)->start, (intervals+j)->end);
                    newInterval.end = (intervals+j)->end;
                    rp = copy(rp, &newInterval, 1);
                    int n = intervalsSize - j - 1; // j not included
                    copy(rp, intervals+j+1, n);
                    (*returnSize) = i+1+n;
                    return res;
                }
            }
            // after
            rp = copy(rp, &newInterval, 1);
            (*returnSize) = i+1;
            return res;
    	}
    }
    rp = copy(res, intervals, intervalsSize);
    rp = copy(rp, &newInterval, 1);
    (*returnSize) = intervalsSize + 1;
    return res;
}

int main(int argc, char const *argv[])
{
    int returnSize;
    struct Interval newInterval = {2,5};
    struct Interval intervals[] = {{1,3},{6,9}};
    struct Interval * res = insert(intervals, sizeof(intervals)/sizeof(struct Interval), newInterval, &returnSize);
    printf("size %d\n", returnSize);
    for (int i = 0; i < returnSize; ++i)
    {
        printf("[%d,%d]\n", res[i].start, res[i].end);
    }

    struct Interval newInterval2 = {4,9};
    struct Interval intervals2[] = {{1,2},{3,5},{6,7},{8,10},{12,16}};
    res = insert(intervals2, sizeof(intervals2)/sizeof(struct Interval), newInterval2, &returnSize);
    printf("size %d\n", returnSize);
    for (int i = 0; i < returnSize; ++i)
    {
        printf("[%d,%d]\n", res[i].start, res[i].end);
    }

    struct Interval newInterval3 = {6,6};
    struct Interval intervals3[] = {{3,5},{12,15}};
    res = insert(intervals3, sizeof(intervals3)/sizeof(struct Interval), newInterval3, &returnSize);
    printf("size %d\n", returnSize);
    for (int i = 0; i < returnSize; ++i)
    {
        printf("[%d,%d]\n", res[i].start, res[i].end);
    }

    struct Interval newInterval4 = {0,20001};
    int n = 10000;
    struct Interval *intervals4 = malloc(n*sizeof(struct Interval));
    for (int i = 0; i < n; ++i)
    {
        intervals4[i].start = i*2+1;
        intervals4[i].end = i*2+2;
    }
    res = insert(intervals4, n, newInterval4, &returnSize);
    printf("#### size %d\n", returnSize);
    for (int i = 0; i < returnSize; ++i)
    {
        printf("[%d,%d]\n", res[i].start, res[i].end);
    }

    struct Interval intervals6[] = {{1,5}};
    struct Interval newInterval5 = {6,8};
    res = insert(intervals6, sizeof(intervals6)/sizeof(struct Interval), newInterval5, &returnSize);
    printf("### size %d\n", returnSize);
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

    struct Interval intervals_9[] = {{1,4},{10,12},{13,14},{16,16},{19,20},{21,24},{33,33},{36,39},{44,46},{48,50}};
    struct Interval newInterval_9 = {5,13};
    res = insert(intervals_9, sizeof(intervals_9)/sizeof(struct Interval), newInterval_9, &returnSize);
    printf("8. ### size %d\n", returnSize);
    for (int i = 0; i < returnSize; ++i)
    {
        printf("[%d,%d]\n", res[i].start, res[i].end);
    }
    return 0;
}
