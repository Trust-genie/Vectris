import numpy as np
import random


def createMatrix(size):
    matrix = np.random.rand(size, size)
    return matrix

def testMatrixcreate(benchmark):
    size =10000
    result = benchmark(createMatrix, size)
    benchmark.pedantic(createMatrix, args=(size,), iterations=10, rounds=30)