import unittest
import task_1

# python3 -m unittest discover


class TestFirstTask(unittest.TestCase):
    def test_main_all_ok(self):
        input_values = [2, 3, 0, 5]
        output = []

        def mock_input(s):
            output.append(s)
            return input_values.pop(0)

        task_1.input = mock_input
        task_1.print = lambda s: output.append(s)

        is_exception_caught = False

        try:
            task_1.main()
        except BaseException:
            is_exception_caught = True

        self.assertFalse(is_exception_caught)
        self.assertEqual(output, ["Enter a", "Enter b", "Enter c", "Enter d", "Result is 1.0",])

    def test_main_incorrect_input_1(self):
        input_values = [':sad_smile', 2, 3, 4]
        output = []

        def mock_input(s):
            output.append(s)
            return input_values.pop(0)

        task_1.input = mock_input
        task_1.print = lambda s: output.append(s)

        is_exception_caught = False

        try:
            task_1.main()
        except BaseException:
            is_exception_caught = True

        assert is_exception_caught

        assert output == ['Enter a', 'Last value was entered in incorrect format']

    def test_main_incorrect_input_2(self):
        input_values = [2, ':sad_smile', 3, 4]
        output = []

        def mock_input(s):
            output.append(s)
            return input_values.pop(0)

        task_1.input = mock_input
        task_1.print = lambda s: output.append(s)

        is_exception_caught = False

        try:
            task_1.main()
        except BaseException:
            is_exception_caught = True

        assert is_exception_caught

        assert output == ['Enter a', 'Enter b', 'Last value was entered in incorrect format']

    def test_main_incorrect_input_3(self):
        input_values = [2, 3, ':sad_smile', 4]
        output = []

        def mock_input(s):
            output.append(s)
            return input_values.pop(0)

        task_1.input = mock_input
        task_1.print = lambda s: output.append(s)

        is_exception_caught = False

        try:
            task_1.main()
        except BaseException:
            is_exception_caught = True

        assert is_exception_caught

        assert output == ['Enter a', 'Enter b', 'Enter c', 'Last value was entered in incorrect format']

    def test_main_incorrect_input_4(self):
        input_values = [2, 3, 4, ':sad_smile']
        output = []

        def mock_input(s):
            output.append(s)
            return input_values.pop(0)

        task_1.input = mock_input
        task_1.print = lambda s: output.append(s)

        is_exception_caught = False

        try:
            task_1.main()
        except BaseException:
            is_exception_caught = True

        assert is_exception_caught

        assert output == ['Enter a', 'Enter b', 'Enter c', 'Enter d', 'Last value was entered in incorrect format']

    def test_main_divide_by_zero_with_zeros_input(self):
        input_values = [0, 0, 0, 0]
        output = []

        def mock_input(s):
            output.append(s)
            return input_values.pop(0)

        task_1.input = mock_input
        task_1.print = lambda s: output.append(s)

        is_exception_caught = False

        try:
            task_1.main()
        except BaseException:
            is_exception_caught = True

        assert is_exception_caught

        assert output == ['Enter a', 'Enter b', 'Enter c', 'Enter d', 'Divide by zero, values c + d should not be 0']

    def test_main_divide_by_zero_with_not_zeros_input(self):
        input_values = [44, 23, -1, 1]
        output = []

        def mock_input(s):
            output.append(s)
            return input_values.pop(0)

        task_1.input = mock_input
        task_1.print = lambda s: output.append(s)

        is_exception_caught = False

        try:
            task_1.main()
        except BaseException:
            is_exception_caught = True

        assert is_exception_caught

        assert output == ['Enter a', 'Enter b', 'Enter c', 'Enter d', 'Divide by zero, values c + d should not be 0']

    def test_A_method_all_ok_res_positive(self):
        self.assertEqual(task_1.A(14, 26, 2, 2), 10)

    def test_A_method_all_ok_res_negative(self):
        self.assertEqual(task_1.A(-20, -10, 1, 1), -15)

    def test_A_method_divide_by_zero(self):
        is_exception_caught = False

        try:
            self.failUnlessRaises(ZeroDivisionError, task_1.A(-20, -10, 2, -2))
        except ZeroDivisionError:
            is_exception_caught = True

        assert is_exception_caught
