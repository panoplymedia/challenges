import sys

min_query = int(sys.argv[1])
max_query = int(sys.argv[2])

# static array of byte data tuples [min, max] from the example in README, missing 344-467
byte_data = [
    [112,149],
    [149,224],
    [224,344],
    [467,515]
]

seq = []

for i, d in enumerate(byte_data):
    min_val = d[0]

    if i != 0 and max_val != min_val:
        seq.append(max_val)
        break

    seq.append(min_val)
    max_val = d[1]

    if i == (len(byte_data) - 1):
        seq.append(max_val)

min_val = min(seq)
max_val = max(seq)

res = (min_query >= min_val and max_query <= max_val)
print('result of query = [{}]'.format(res))