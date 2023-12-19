import argparse
from pathlib import Path
import re
import sys
from typing import Tuple


def min_cubes_rgb(line: str) -> Tuple[int, int, int]:
    """Return the min cubes of each color in a Tuple. Red, Gree, Blue"""
    return (
        max([int(m.group(1)) for m in re.finditer("(\d+) red", line)]),
        max([int(m.group(1)) for m in re.finditer("(\d+) green", line)]),
        max([int(m.group(1)) for m in re.finditer("(\d+) blue", line)]),
    )


def is_game_valid(
    line: str, max_r: int = 12, max_g: int = 13, max_b: int = 14
) -> bool:
    """
    If the game does not violate the color constraints, return True
    else False
    """
    reds, greens, blues = min_cubes_rgb(line)
    return reds <= max_r and greens <= max_g and blues <= max_b


def sum_valid_game_numbers(file) -> int:
    """
    Given a file of games, return the sum of the game numbers that do not
    viloate the color constraints. The games will be a list<int>.
    """
    sum: int = 0
    for line in file.readlines():
        m = re.match("Game (\d+)", line.strip())
        if is_game_valid(line.strip()) and m is not None:
            sum += int(m.group(1))
    return sum


def game_power(line: str) -> int:
    """
    Returns the multiplication of the min number of each cube required.
    """
    reds, greens, blues = min_cubes_rgb(line)
    return reds * greens * blues


def sum_game_power(file) -> int:
    """
    Returns the sum of the game powers, which is the multiplicaiton of min
    cubes of each color.
    """
    return sum([game_power(line.strip()) for line in file.readlines()])


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
    print(sum_valid_game_numbers(args.inputfile))
    args.inputfile.seek(0)  # reset the file handler for pt2
    print(sum_game_power(args.inputfile))
    return


if __name__ == "__main__":
    main(sys.argv[1:])
