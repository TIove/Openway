
# python3 task_1


def A(a: float, b: float, c: float, d: float) -> float:
    return (a + b) / (c + d)


def main():
    try:
        first_val = float(input("Enter a "))
        second_val = float(input("Enter b "))
        third_val = float(input("Enter c "))
        fourth_val = float(input("Enter d "))

        print("Result is " + str(A(first_val, second_val, third_val, fourth_val)))
    except ValueError:
        print("Last value was entered in incorrect format")
        exit(1)
    except ZeroDivisionError:
        print("Divide by zero, values c + d should not be 0")
        exit(2)


if __name__ == '__main__':
    main()
