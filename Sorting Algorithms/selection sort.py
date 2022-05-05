from time import perf_counter_ns


def select_sort(l):
    global comparisons, swap
    for i in range(len(l)):
        min_i = i
        for j in range(i+1, len(l)):
            comparisons += 1
            if l[j] < l[min_i]:
                min_i = j
        swap += 1
        l[i], l[min_i] = l[min_i], l[i]
    return l


if __name__ == "__main__":
    from random import sample
    comparisons = swap = 0
    unsorted = sample(range(0, 10000), 1000)
    start_time = perf_counter_ns()
    print(select_sort(unsorted))
    end_time = perf_counter_ns()
    print((end_time - start_time) * 10 ** (-6))
    print(f'Comparison count: {comparisons}\nSwap count: {swap}')
