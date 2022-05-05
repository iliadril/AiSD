from time import perf_counter_ns


def insertion_sort(l):
    global comparisons, swap
    for i in range(1, len(l)):
        marker = l[i]
        stepper = i

        while stepper > 0 and l[stepper - 1] > marker:
            comparisons += 1
            l[stepper] = l[stepper - 1]
            stepper -= 1
        l[stepper] = marker
    return l


if __name__ == "__main__":
    from random import sample
    comparisons = swap = 0
    unsorted = sample(range(0, 10000), 1000)
    start_time = perf_counter_ns()
    print(insertion_sort(unsorted))
    end_time = perf_counter_ns()
    print((end_time - start_time) * 10 ** (-6))
    print(f'Comparison count: {comparisons}\nSwap count: {swap}')
