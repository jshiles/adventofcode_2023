import argparse
from pathlib import Path
import re
import sys
from typing import Tuple, Optional


TEXT_DIGITS = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]


def extract_first_digit(line: str) -> Optional[int]:
    "Given a line (str), extract the first number return it."
    val = next((i for i in line if i.isdigit()), None)
    if val is not None:
        return int(val)
    return


def extract_first_and_last_digit(line: str) -> Tuple[int, int]:
    """
    Given a line(str), extract the first and last integer and return them as
    a tuple.
    """
    numbers: Tuple(int, int) = []
    regex = "(?=[0-9]|one|two|three|four|five|six|seven|eight|nine)"
    start_idxs = [m.start() for m in re.finditer(regex, line)]
    for idx in start_idxs:
        if line[idx].isdigit():
            numbers.append([idx, int(line[idx])])
        else:
            numbers = numbers + [
                (idx, TEXT_DIGITS.index(x) + 1)
                for x in TEXT_DIGITS
                if line[idx:].startswith(x)
            ]

    return (numbers[0][1], numbers[-1][1])


def parse_calibration_coordinates(file) -> int:
    """
    Parse file into lines and extract first and last digits, return sum over
    all lines.
    """
    sum: int = 0
    for line in file.readlines():
        a, b = extract_first_and_last_digit(line.strip())
        sum += a * 10 + b
    return sum


def is_valid_file(parser, arg):
    if not Path(arg).exists:
        parser.error("The file %s does not exist!" % arg)
    else:
        return open(arg, "r")


def init_parser():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "--inputfile",
        help="location of the input file",
        metavar="FILE",
        type=lambda x: is_valid_file(parser, x),
    )
    parser.set_defaults(feature=True)
    return parser


def main(argv) -> None:
    """ """
    parser = init_parser()
    args = parser.parse_args(argv)
    sum = parse_calibration_coordinates(args.inputfile)
    print(sum)
    return


if __name__ == "__main__":
    main(sys.argv[1:])
