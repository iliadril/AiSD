from time import perf_counter_ns


def heapify(l, end, i):
    left = 2 * i + 1
    right = 2 * (i + 1)
    max_i = i
    global comparisons
    comparisons += 2
    if left < end and l[i] < l[left]:
        max_i = left
    comparisons += 2
    if right < end and l[max_i] < l[right]:
        max_i = right
    if max_i != i:
        global swap
        swap += 1
        l[i], l[max_i] = l[max_i], l[i]
        heapify(l, end, max_i)


def heap_sort(l):
    end = len(l)
    start = end // 2 - 1
    for i in range(start, -1, -1):
        heapify(l, end, i)
    for i in range(end-1, 0, -1):
        global swap
        swap += 1
        l[i], l[0] = l[0], l[i]
        heapify(l, i, 0)
    return l


if __name__ == "__main__":
    from random import sample
    comparisons = swap = 0
    unsorted = sample(range(0, 10000), 1000)
    start_time = perf_counter_ns()
    print(heap_sort(unsorted))
    end_time = perf_counter_ns()
    print((end_time - start_time) * 10 ** (-6))
    print(f'Comparison count: {comparisons}\nSwap count: {swap}')
