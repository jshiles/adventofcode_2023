from adventofcode.almanac import (
    Mapping,
    SeedRange,
    SEED_TO_SOIL,
    SOIL_TO_FERTILIZER,
    FERTILIZER_TO_WATER,
    WATER_TO_LIGHT,
    LIGHT_TO_TEMPERATURE,
    TEMPERATURE_TO_HUMIDITY,
    HUMIDITY_TO_LOCATION,
)
from adventofcode import utils
import logging
import numpy as np
import sys
from typing import List, FrozenSet


@utils.log_output
def source_to_destination(
    source: int, source_dest_map: FrozenSet[Mapping]
) -> int:
    """
    Using our mapping objects, find the mapping that translates source to dest.
    If none exist, then the source and destination are the same.
    """
    for mapping in source_dest_map:
        if mapping.valid_source_number(source):
            return mapping.source_to_destination(source)
    return source


@utils.log_output
def destination_to_source(
    destination: int, source_dest_map: FrozenSet[Mapping]
) -> int:
    """
    Using our mapping objects, find the mapping that translates source to dest.
    If none exist, then the source and destination are the same.
    """
    for mapping in source_dest_map:
        if mapping.valid_destination_number(destination):
            return mapping.destination_to_source(destination)
    return destination


def find_locations(seed: int) -> int:
    """
    Given a seed, find the corresponding location.
    """
    soil = source_to_destination(seed, SEED_TO_SOIL)
    fert = source_to_destination(soil, SOIL_TO_FERTILIZER)
    water = source_to_destination(fert, FERTILIZER_TO_WATER)
    light = source_to_destination(water, WATER_TO_LIGHT)
    temp = source_to_destination(light, LIGHT_TO_TEMPERATURE)
    humidity = source_to_destination(temp, TEMPERATURE_TO_HUMIDITY)
    return source_to_destination(humidity, HUMIDITY_TO_LOCATION)


def find_seed(location: int) -> int:
    """ 
    Given a location, find the corresponding seed.
    """
    humidity = destination_to_source(location, HUMIDITY_TO_LOCATION)
    temp = destination_to_source(humidity, TEMPERATURE_TO_HUMIDITY)
    light = destination_to_source(temp, LIGHT_TO_TEMPERATURE)
    water = destination_to_source(light, WATER_TO_LIGHT)
    fert = destination_to_source(water, FERTILIZER_TO_WATER)
    soil = destination_to_source(fert, SOIL_TO_FERTILIZER)
    return destination_to_source(soil, SEED_TO_SOIL)


def is_starter_seed(seed, seed_ranges: List[SeedRange]) -> bool:
    """
    Is this seed in the set of starter seed ranges? Return bool.
    """
    return any([sr.in_range(seed) for sr in seed_ranges])


def search_ranges(
    start: int, stop: int, seed_ranges: FrozenSet[SeedRange]
) -> List[bool]:
    """
    Given a range of locations, search for the lowest location number that
    maps to a seed in the starter set.
    """
    return [
        is_starter_seed(find_seed(loc), seed_ranges)
        for loc in range(start, stop)
    ]


def main(argv) -> None:
    """ """
    # test ranges.
    # seed_ranges: FrozenSet[SeedRange] = frozenset([
    #     SeedRange(79, 1),
    #     SeedRange(14, 1),
    #     SeedRange(55, 1),
    #     SeedRange(13, 1),
    # ])
    seed_ranges: FrozenSet[SeedRange] = frozenset([
        SeedRange(2019933646, 2719986),
        SeedRange(2982244904, 337763798),
        SeedRange(445440, 255553492),
        SeedRange(1676917594, 196488200),
        SeedRange(3863266382, 36104375),
        SeedRange(1385433279, 178385087),
        SeedRange(2169075746, 171590090),
        SeedRange(572674563, 5944769),
        SeedRange(835041333, 194256900),
        SeedRange(664827176, 42427020),
    ])

    vect_find_seed = np.vectorize(search_ranges)
    start, step = 1, 1000000
    found = False
    while not found:
        logging.info(f"Step: {start} ...")
        result = vect_find_seed(start, start+step, seed_ranges)
        if any(result):
            location = min([idx for idx, e in enumerate(result) if e]) + start
            print(f"Found location: {location}")
            found = True
        start += step

    return


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    main(sys.argv[1:])
