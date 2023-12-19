import argparse
from dataclasses import dataclass
from pathlib import Path
import re
import sys
from typing import Tuple, List
from functools import lru_cache


@dataclass(frozen=True)
class Coord:
    row: int
    col: int

    def __repr__(self) -> str:
        return f"({self.row}, {self.col})"


@dataclass(frozen=True)
class Symbol:
    coord: Coord
    value: str


@dataclass(frozen=True)
class Number:
    row: int
    col_start: int
    col_end: int
    value: int

    def __repr__(self) -> str:
        return f"[{self.row}, {self.col_start}-{self.col_end}] -> {self.value}"

    @property
    @lru_cache
    def adjacent_coordinates(self) -> List[Coord]:
        """
        Computes the list of coordinates adjacent to our number as a property
        and the result is cached.
        """
        coords: List[Coord] = []
        coords = coords + [
            Coord(self.row - 1, x) for x in range(self.col_start - 1, self.col_end + 2)
        ]
        coords.append(Coord(self.row, self.col_start - 1))
        coords.append(Coord(self.row, self.col_end + 1))
        coords = coords + [
            Coord(self.row + 1, x) for x in range(self.col_start - 1, self.col_end + 2)
        ]
        return coords


def file_to_list(file) -> Tuple[List[Number], List[Symbol]]:
    """Parse the input file into 'Numbers' and 'Symbols' and return two lists"""
    row: int = 0
    numbers: List[Number] = []
    symbols: List[Coord] = []

    for line in file.readlines():
        line = line.strip()
        numbers = numbers + [
            Number(row, m.start(), m.end() - 1, int(m.string[m.start():m.end()]))
            for m in re.finditer(r"(\d+)", line)
        ]
        symbols = symbols + [
            Symbol(Coord(row, m.start()), m.string[m.start()])
            for m in re.finditer("([^0-9.]{1})", line)
        ]
        row += 1

    return numbers, symbols


def adjacent_to_symbol(num: Number, coords: List[Coord]) -> bool:
    """Return True if the Number (num) is adjacent to a Symbol otherwise False"""
    return any([adj in coords for adj in num.adjacent_coordinates])


def sum_machine_parts(file) -> int:
    """
    A Number is a machine part if it is adjacent to a Symbol. This will sum
    the Numbers which are machine parts, and return that value.
    """
    numbers, symbols = file_to_list(file)
    coords = [sym.coord for sym in symbols]
    return sum([mp.value for mp in numbers if adjacent_to_symbol(mp, coords)])


def find_gears(numbers: List[Number], symbols: List[Symbol]) -> List[Tuple[Symbol, int]]:
    """
    A Symbol is a GEAR if it has the value '*' and is adjacent to EXACTLY two
    numbers. A GEAR RATIO is the multiplication of those two adjacent numbers.
    This returns a list of Gears and their Gear Ratios.
    """
    gears: List[Tuple[Symbol, int]] = []
    for sym in symbols:
        if sym.value != "*":  # gears must be a "*"
            continue
        adj_nums = [
            num.value 
            for num in numbers 
            if sym.coord in num.adjacent_coordinates
        ]
        if len(adj_nums) == 2:  # gears must be adjacent to exactly 2!
            gears.append((sym, adj_nums[0] * adj_nums[1]))
    return gears


def sum_gear_ratios(file) -> int:
    """Sums the Gear Ratios"""
    numbers, symbols = file_to_list(file)
    gears: List[Tuple[Symbol, int]] = find_gears(numbers, symbols)
    return sum([g[1] for g in gears])


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

    # part 1
    print(sum_machine_parts(args.inputfile))

    args.inputfile.seek(0)  # reset for part 2

    # part 2
    print(sum_gear_ratios(args.inputfile))
    return


if __name__ == "__main__":
    main(sys.argv[1:])
