from time import perf_counter_ns


def quicksort(l):
    if not l:
        return l
    partition = len(l) // 2
    left = [i for i in l if i < l[partition]]
    m = [i for i in l if i == l[partition]]
    right = [i for i in l if i > l[partition]]
    return quicksort(left) + m + quicksort(right)


if __name__ == "__main__":
    from random import sample
    comparisons = swap = 0
    unsorted = sample(range(0, 10000), 1000)
    start_time = perf_counter_ns()
    print(quicksort(unsorted))
    end_time = perf_counter_ns()
    print((end_time - start_time) * 10 ** (-6))
    print(f'Comparison count: {comparisons}\nSwap count: {swap}')
