from time import perf_counter_ns


def merge_sort(x):
    if len(x) < 2:
        return x
    result = []
    mid = len(x) // 2
    left = merge_sort(x[:mid])
    right = merge_sort(x[mid:])
    i = j = 0
    while i < len(left) and j < len(right):
        if left[i] > right[j]:
            result += [right[j]]
            j += 1
        else:
            result += [left[i]]
            i += 1
    result += left[i:]
    result += right[j:]
    return result


if __name__ == "__main__":
    from random import sample
    comparisons = swap = 0
    unsorted = sample(range(0, 10000), 1000)
    start_time = perf_counter_ns()
    print(merge_sort(unsorted))
    end_time = perf_counter_ns()
    print((end_time - start_time) * 10 ** (-6))
    print(f'Comparison count: {comparisons}\nSwap count: {swap}')
