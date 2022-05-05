from time import perf_counter_ns


def bubble_sort(l):
    for sort_pass in range(0, len(l)):
        for i in range(0, len(l) - 1 - sort_pass):
            global comparisons, swap
            comparisons += 1
            if l[i] > l[i + 1]:
                swap += 1
                l[i], l[i + 1] = l[i + 1], l[i]
    return l


if __name__ == "__main__":
    from random import sample
    comparisons = swap = 0
    unsorted = sample(range(0, 10000), 1000)
    start_time = perf_counter_ns()
    print(bubble_sort(unsorted))
    end_time = perf_counter_ns()
    print((end_time - start_time) * 10 ** (-6))
    print(f'Comparison count: {comparisons}\nSwap count: {swap}')
