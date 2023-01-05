package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, errors.New("Number must be positive")
    } else if n == 1 {
        return 0, nil
    } else if n%2 == 0 {
        n, err := CollatzConjecture(n / 2)
        return n + 1, err
    } else {
        n, err := CollatzConjecture(3*n + 1)
        return n + 1, err
    }
}
