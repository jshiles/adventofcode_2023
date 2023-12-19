import argparse
from dataclasses import dataclass
from pathlib import Path
import re
import sys
from typing import Tuple, List
from functools import lru_cache


@dataclass(frozen=True)
class ScratchCard:
    winning_numbers: List[int]
    elfs_numbers: List[int]

    @property
    def matches(self) -> List[int]:
        """Matching numbers."""
        return list(set(self.elfs_numbers) & set(self.winning_numbers))

    @property
    def num_matches(self) -> int:
        """Number of matching numbers."""
        return len(self.matches)

    @property
    def value(self) -> int:
        """Sum the value of the scratchers according the puzzle logic."""
        if len(self.matches) > 0:
            return 2 ** (len(self.matches) - 1)
        return 0


def file_to_scratchcards(file) -> List[ScratchCard]:
    """Transform input into a list of ScratchCards."""
    scratchers: List[ScratchCard] = []
    for line in file.readlines():
        _, winning_str, scratcher_str = re.split(r"[:\|]", line.strip())
        winning_numbers: List[int] = [
            int(m[0]) for m in re.finditer(r"(\d+)", winning_str)
        ]
        elfs_numbers: List[int] = [
            int(m[0]) for m in re.finditer(r"(\d+)", scratcher_str)
        ]
        scratchers.append(ScratchCard(winning_numbers, elfs_numbers))
    return scratchers


def sum_scratchcards_value(scratchers: List[ScratchCard]) -> int:
    """Sum the scratcher values"""
    return sum([sc.value for sc in scratchers])


def recursive_cnt(card_number, x_wins_y: dict) -> int:
    """
    Helper function that recursively caches the count. We could make this
    faster with caching, if we make the x_wins_y immutable.
    """
    if x_wins_y.get(card_number, None) is None:
        return 0
    else:
        return sum(
            [1 + recursive_cnt(scn, x_wins_y) for scn in x_wins_y.get(card_number)]
        )


def count_scratchcards(scratchers: List[ScratchCard]) -> int:
    """
    Build a mapping between card and winning card, then count recursively
    returning the number of cards we end up with.
    """
    # build map
    x_wins_y: dict = {}
    for card_number, sc in enumerate(scratchers):
        if sc.num_matches > 0:
            x_wins_y[card_number] = list(
                range(
                    card_number + 1,
                    min(len(scratchers), card_number + 1 + sc.num_matches),
                )
            )
    # count based on map
    cards: int = 0
    for card_number, sc in enumerate(scratchers):
        cards += 1 + recursive_cnt(card_number, x_wins_y)
    return cards


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
    scratchers = file_to_scratchcards(args.inputfile)
    print(sum_scratchcards_value(scratchers))
    print(count_scratchcards(scratchers))
    return


if __name__ == "__main__":
    main(sys.argv[1:])
