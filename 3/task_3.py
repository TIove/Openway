import math

# python3 task_3


def count_years(interest: float, multiplicity: float, is_composite_interest: bool) -> int:
    if is_composite_interest:
        return math.ceil(math.log(multiplicity, 1 + interest / 100))
    else:
        return math.ceil(multiplicity / (interest / 100))


def main():
    try:
        amount_ = float(input('Enter amount '))
        interest_ = float(input('Enter interest '))
        multiplicity_ = float(input('Enter multiplicity '))
        is_composite_interest_ = input('Is composite interest type (True / False) ')
    except ValueError:
        print("Last entered value was in incorrect format")
        exit(1)

    if is_composite_interest_ == 'True':
        is_composite_interest_ = True
    elif is_composite_interest_ == 'False':
        is_composite_interest_ = False
    else:
        print("You entered incorrect value")

    years_res = count_years(interest_, multiplicity_, is_composite_interest_)
    print(f"After {years_res} you will have more than {amount_ * multiplicity_} !")


if __name__ == '__main__':
    main()
