from itertools import pairwise
import time

a = [0] * 1000000

start_time = time.time()
for i,k in zip(a[0::2], a[1::2]):
    i + k

print("--- %s seconds ---" % (time.time() - start_time))
