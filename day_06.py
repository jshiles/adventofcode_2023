import logging
import sys
from dataclasses import dataclass
from typing import List


@dataclass(frozen=True)
class Race:
    time: int
    record_distance: int

    def beats_record(self, hold_seconds: int) -> bool:
        mm_ps = hold_seconds
        return self.record_distance < (self.time - hold_seconds) * mm_ps


def main(argv) -> None:
    """ """
    races: List[Race] = [
        # Race(7, 9),
        # Race(15, 40),
        # Race(30, 200),
        Race(59, 597),
        Race(79, 1234),
        Race(65, 1032),
        Race(75, 1328),
    ]
    sum = 1
    for race in races:
        sum *= len([
            seconds 
            for seconds in range(1, race.time)
            if race.beats_record(seconds)
        ])
    print(sum)

    # honestly not worth making faster. :/
    long_race = Race(59796575, 597123410321328)
    p2_ways = len([
        seconds 
        for seconds in range(1, long_race.time)
        if long_race.beats_record(seconds)
    ])
    print(p2_ways)

    return


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    main(sys.argv[1:])
