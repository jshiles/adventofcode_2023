import argparse
import logging
import sys
from collections import Counter
from dataclasses import dataclass, field
from pathlib import Path
from typing import List, Tuple


RANKS = {"T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}
RANKS_WC = {"T": 10, "J": 1, "Q": 12, "K": 13, "A": 14}


@dataclass(frozen=True)
class CamelCardsHand:
    hand: list[str]
    bid: int
    wild_card: bool

    @property
    def hand_type(self) -> int:
        """
        Returns a ranking of hand types: Five of a kind (5) ... High Card (0)
        """
        c = Counter(self.hand)

        # if wild card, then assign Jacks to largest repeating value.        
        if self.wild_card and c.get('J', 0) > 0 and c.get('J', 0) != 5:
            first_mc, second_mc = c.most_common(2)[0][0], c.most_common(2)[1][0]
            update_value = second_mc if first_mc == 'J' else first_mc
            c[update_value] += c.get('J', 0)
            c['J'] = 0

        common: List[Tuple(str, int)] = c.most_common()
        if len(common) > 0 and common[0][1] == 5:
            return 6
        elif len(common) > 0 and common[0][1] == 4:
            return 5
        elif len(common) > 1 and common[0][1] == 3 and common[1][1] == 2:
            return 4
        elif len(common) > 0 and common[0][1] == 3:
            return 3
        elif len(common) > 1 and common[0][1] == 2 and common[1][1] == 2:
            return 2
        elif len(common) > 0 and common[0][1] == 2:
            return 1
        return 0

    def __lt__(self, other) -> bool:
        """
        Compares two hands, first by hand_type then element by element
        (unsorted)
        """
        ranks = RANKS_WC if self.wild_card else RANKS
        if self.hand_type == other.hand_type:
            for idx, x in enumerate(self.hand):
                if x != other.hand[idx]:
                    return (int(x) if x.isdigit() else ranks[x]) < (
                        int(other.hand[idx])
                        if other.hand[idx].isdigit()
                        else ranks[other.hand[idx]]
                    )
        return self.hand_type < other.hand_type

    def __repr__(self) -> str:
        return "".join(self.hand)


def read_hands(file, wild_card_logic: bool = False) -> List[CamelCardsHand]:
    """Read hands and bids from a file into a list of Hands"""
    hands: List[CamelCardsHand] = []
    for line in file.readlines():
        hand, bid = line.strip().split()
        hands.append(CamelCardsHand(hand, int(bid), wild_card_logic))
    return hands


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
    parser.add_argument(
        "--wc",
        help="Jacks are wild",
        action='store_true',
    )
    parser.set_defaults(feature=True)
    return parser


def main(argv) -> None:
    parser = init_parser()
    args = parser.parse_args(argv)
    hands = read_hands(args.inputfile, wild_card_logic=args.wc)
    print(sum([(idx+1) * h.bid for idx, h in enumerate(sorted(hands))]))
    return


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    main(sys.argv[1:])
