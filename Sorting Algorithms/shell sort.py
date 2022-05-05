from time import perf_counter_ns


def shell_sort(l):
    gap = len(l) // 2
    global comparisons, swap
    while gap > 0:
        # literally insertion
        for i in range(gap, len(l)):
            val = l[i]
            j = i
            while j >= gap and l[j - gap] > val:
                comparisons += 1
                l[j] = l[j - gap]
                j -= gap
            l[j] = val
        gap //= 2
    return l


if __name__ == "__main__":
    from random import sample
    comparisons = swap = 0
    unsorted = sample(range(0, 10000), 1000)
    start_time = perf_counter_ns()
    print(shell_sort(unsorted))
    end_time = perf_counter_ns()
    print((end_time - start_time) * 10 ** (-6))
    print(f'Comparison count: {comparisons}\nSwap count: {swap}')
