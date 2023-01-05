package collatzconjecture

import (
    "errors"
)

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, errors.New("cannot be a negative number")
    }

    var result int
    for n > 1 {
        if n%2 == 0 {
            n = n / 2
        } else {
            n = 3*n + 1
        }
        result++
    }
    return result, nil
}
