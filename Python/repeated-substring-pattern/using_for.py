import time

my_list = [1,2,3,4,5,6,7]
nb_items = 2
for i in range(0, len(my_list), nb_items):
    x = my_list[i:i+nb_items]
    print(i, x)

a = [0] * 1000000
start_time = time.time()
for i in range(0, len(a), 2):
    sum(a[i:i+nb_items])

print("--- %s seconds ---" % (time.time() - start_time))
